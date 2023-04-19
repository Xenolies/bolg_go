package main

/**
* @Author: Xenolies
* @Date: 2023/4/16 12:10
* @Version: 1.0
 */

import (
	"bolg_go/config"
	"bolg_go/controller"
	"bolg_go/database"
	"bolg_go/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	database.InitMySQL(config.GlobalConfig.Admin)
	router := gin.Default()

	// 读取页面问价目录
	router.LoadHTMLGlob("./web/*")

	// 访问错误内容知己转到NotFound
	router.NoRoute(controller.NotFound)

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", 0)
	})
	groupUser := router.Group("/user")
	groupUser.POST("/register", controller.Register)
	groupUser.POST("/login", controller.Login)

	groupAdmin := router.Group("/admin")
	groupAdmin.POST("/login", controller.Login)

	groupAdmin.Use(middleware.AuthMiddleware(true))
	//groupAdmin.POST("/login", controller.AdminLogin)
	groupAdmin.GET("/users", controller.GetAllUsers)
	groupAdmin.GET("/usersfind", controller.FindUsers)
	groupAdmin.GET("/users/:userId", controller.GetOneUser)
	groupAdmin.PUT("/users", controller.CreateOneUser)
	groupAdmin.POST("/users/:userId", controller.UpdateOneUser)
	groupAdmin.DELETE("/users/:userId", controller.DeleteOneUser)

	groupArticles := router.Group("/articles")
	groupArticles.POST("/postarticles", controller.PostArticles)

	groupFiles := router.Group("/files")
	groupFiles.POST("/uploadonefile", controller.UploadOneFile)
	groupFiles.POST("/uploadfiles", controller.UploadFiles)

	//监听服务器端口
	router.Run(":8080")
}
