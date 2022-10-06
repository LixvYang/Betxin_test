package router

import (
	"betxin/api/v1/administrator"
	"betxin/api/v1/category"
	"betxin/api/v1/oauth"
	"betxin/api/v1/topic"
	"betxin/api/v1/upload"
	"betxin/utils"
	"betxin/utils/cors"
	"betxin/utils/jwt"
	"betxin/utils/logger"

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

	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)
	r.Use(logger.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())
	r.LoadHTMLFiles("dist/index.html", "dist/welcome.html")
	
	auth := r.Group("api/v1")
	auth.Use(jwt.JwtToken())
	{
		//管理员
		auth.POST("/administrator/add", administrator.CreateAdministrator)

		// bonuse 奖金

		// 分类模块
		auth.GET("/category/:id", category.GetCategoryInfo)
		auth.POST("/category/add", category.CreateCatrgory)
		auth.PUT("/category/:id", category.UpdateCategory)
		auth.DELETE("/category/:id", category.DeleteCategory)
		auth.POST("/category/list", category.ListCategories)
		// 收藏

		// Mixin信息

		// MixinOrder 接收用户的币

		// snapshot 反馈给用户的钱

		// swaporder 管理从4swap的转账金钱

		// topic 管理话题

		// upload   上传文件


		// user 用户管理

		// 人\userauth 用户登录

		// usertotopic 用户买的话题

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

		// 管理奖金

		//种类

		//收藏

		//

		//话题
		router.POST("/topic/create", topic.CreateTopic)
		router.POST("/topic/cid/:cid", topic.GetTopicByCid)
		router.POST("/topic/list", topic.ListTopics)
		router.GET("/topic/:id", topic.GetTopicInfoById)
		// 用户
		// router.POST("/user/add", user)
		router.POST("/file", upload.Upload)
	}

	_ = r.Run(utils.HttpPort)
}
