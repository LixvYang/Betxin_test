package router

import (
	"betxin/api/v1/administrator"
	"betxin/api/v1/category"
	"betxin/api/v1/oauth"
	"betxin/utils"
	"betxin/utils/jwt"

	"github.com/gin-gonic/gin"
)

// func createMyRender() multitemplate.Renderer {
// 	p := multitemplate.NewRenderer()
// 	p.AddFromFiles("/", "dist/index.html")
// 	return p
// }

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.LoadHTMLFiles("dist/index.html", "dist/welcome.html")

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(jwt.JwtToken())
	{
		// 分类模块
		auth.GET("/category/:id", category.GetCategoryInfo)
		auth.POST("/category/add", category.CreateCatrgory)
		auth.PUT("/category/:id", category.UpdateCategory)
		auth.DELETE("/category/:id", category.DeleteCategory)
		auth.GET("/category/list", category.ListCategories)

		// 接收用户订单
	}

	router := r.Group("api/v1")
	{
		// 登录控制模块
		router.POST("login", administrator.Login)
		router.GET("/oauth/redirect", oauth.MixinOauth)

		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", "flysnow_org")
		})
		r.GET("/welcome", func(c *gin.Context) {
			c.HTML(200, "welcome.html", "flysnow_org")
		})

	}

	_ = r.Run(utils.HttpPort)
}
