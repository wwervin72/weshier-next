package router

import (
	"fmt"
	"weshierNext/handler/article"
	"weshierNext/handler/page"
	"weshierNext/handler/sd"
	"weshierNext/handler/user"
	"weshierNext/router/middleware"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Load 加载路由
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	g.NoRoute(page.NotFound)
	// 配置 swagger
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", viper.GetString("port")))
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	g.GET("/", page.Home)

	api := g.Group("/api")
	{
		svcd := api.Group("/sd")
		{
			svcd.GET("/health", sd.HealthCheck)
			svcd.GET("/disk", sd.DiskCheck)
			svcd.GET("/cpu", sd.CPUCheck)
			svcd.GET("/ram", sd.RAMCheck)
		}
		userGroup := api.Group("/user")
		{
			userGroup.GET("", user.QueryUserInfo)
			// github login
			userGroup.GET("/auth/github/callback", user.GithubLogin)
			userGroup.POST("/login", user.Login)
			userGroup.GET("/logout", middleware.LoginRequired, user.Logout)
			userGroup.POST("/register", user.Register)
		}
		articleGroup := api.Group("/article")
		{
			articleGroup.POST("", middleware.LoginRequired, article.Create)
			articleGroup.GET("/detail/:articleId", article.QueryArticleDetailByID)
			articleGroup.GET("/list", article.QueryArticleList)
		}
	}
	return g
}
