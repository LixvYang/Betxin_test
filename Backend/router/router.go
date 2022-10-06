package router

import (
	"betxin/api/v1/administrator"
	"betxin/api/v1/category"
	"betxin/api/v1/oauth"
	"betxin/api/v1/topic"
	"betxin/utils"
	"betxin/utils/cors"
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
	
	gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLFiles("dist/index.html", "dist/welcome.html")
	r.Use(cors.Cors())
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
		auth.POST("/category/list", category.ListCategories)

		// 接收用户订单

		// 管理员
		// administrator.CreateAdministratorLixin()
		auth.POST("/administrator/add", administrator.CreateAdministrator)

	}

	router := r.Group("api/v1")
	{
		// 登录控制模块
		router.POST("/login", administrator.Login)
		router.GET("/oauth/redirect", oauth.MixinOauth)
		// 管理用户

		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", "flysnow_org")
		})
		r.GET("/welcome", func(c *gin.Context) {
			c.HTML(200, "welcome.html", "flysnow_org")
		})

		//话题
		router.POST("/topic/create", topic.CreateTopic)
		router.POST("/topic/cid/:cid", topic.GetTopicByCid)
		router.POST("/topic/list", topic.ListTopics)
		router.GET("/topic/:id", topic.GetTopicInfoById)
		// 用户
		// router.POST("/user/add", user)

		
	}

	_ = r.Run(utils.HttpPort)
}
