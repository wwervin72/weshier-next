package model

import (
	"fmt"
	"time"
	"weshierNext/pkg/logger"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	// mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB global database variable
var DB *DataBase

// DataBase db define
type DataBase struct {
	Self      *gorm.DB
	RedisPool *redis.Pool
}

// Init Init global DB
func (db *DataBase) Init() {
	DB = &DataBase{
		Self:      GetSelfDB(),
		RedisPool: GetRedisPool(),
	}
	DB.Self.AutoMigrate(&UserModel{}, &UserAuth{})
	InsertAdminUser()
}

// GetSelfDB self DB
func GetSelfDB() *gorm.DB {
	return InitSelfDb()
}

// GetRedisPool init redis connection
func GetRedisPool() *redis.Pool {
	return &redis.Pool{
		// 最大空闲数
		MaxIdle: 8,
		// 最大活跃数，0 表示无限制
		MaxActive: 0,
		// 最大空闲时间，超过这个时间后，空闲的链接将会被关闭
		IdleTimeout: time.Millisecond * 1000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", viper.GetString("redis.addr"))
			if err != nil {
				logger.Logger.DPanic("redis connect failed", zap.String("error", err.Error()))
				return nil, err
			}
			logger.Logger.Info("redis connect success")
			return conn, nil
		},
	}
}

// InitSelfDb init self db
func InitSelfDb() *gorm.DB {
	return OpenDB(viper.GetString("mysql.username"), viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"), viper.GetString("mysql.dbname"))
}

// OpenDB connect Database
func OpenDB(username, password, addr, dbname string) *gorm.DB {
	conf := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s", username, password, addr, dbname, true, "Local")
	db, err := gorm.Open("mysql", conf)
	if err != nil {
		logger.Logger.DPanic(fmt.Sprintf("mysql connect failed. Database name: %s", dbname), zap.String("error", err.Error()))
		return nil
	}
	logger.Logger.Info("mysql connect success")
	setupDb(db)
	return db
}

func setupDb(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	// 用于设值最大打开的连接数，默认值为0表示不限制，
	// 设置最大的连接数，可以避免并发太高导致连接 mysql 出现 too many connections 的错误
	db.DB().SetMaxOpenConns(1000)
	// 用于设置闲置的连接数，
	// 设置闲置的连接数则，则当开启的一个连接使用完成后，可以放在池里等候下一次使用
	db.DB().SetMaxIdleConns(10)
}

// Close close database
func (db *DataBase) Close() {
	db.Self.Close()
}
