package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"time"
	"weshierNext/config"
	"weshierNext/model"
	"weshierNext/pkg/logger"
	"weshierNext/router"
	"weshierNext/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func pingServer(port string) error {
	for i := 0; i < viper.GetInt("maxPingCount"); i++ {
		addr := "http://127.0.0.1"
		res, err := http.Get(addr + port + "/api/sd/health")
		if err == nil && res.StatusCode == http.StatusOK {
			return nil
		}
		logger.Logger.Info("waiting for the router, retry in 1 second")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}

func main() {
	var cfg string
	flag.StringVar(&cfg, "c", "", "配置文件路径")
	flag.Parse()
	// 初始化配置
	if err := config.Init(cfg); err != nil {
		panic(err)
	}
	// 初始化数据库
	model.DB.Init()
	defer model.DB.Close()

	gin.SetMode(viper.GetString("mode"))
	g := gin.New()
	publicPrefixer := viper.GetString("publicPrefix")
	if publicPrefixer == "" {
		publicPrefixer = "/public"
	}
	g.Static(publicPrefixer, "/public")
	mw := []gin.HandlerFunc{
		middleware.Logging(),
		middleware.RequestID(),
	}
	// 加载路由
	router.Load(g, mw...)
	port := fmt.Sprintf(":%s", viper.GetString("port"))
	go func() {
		if err := pingServer(port); err != nil {
			logger.Logger.Fatal("the router has no response, or it might took to long to start up.", zap.String("error", err.Error()))
		}
		logger.Logger.Info("the router has been deployed successfuly")
	}()
	logger.Logger.Info(http.ListenAndServe(port, g).Error())
}
