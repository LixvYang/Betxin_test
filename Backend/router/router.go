package router

import (
	"betxin/api/v1/administrator"
	"betxin/api/v1/bonuse"
	"betxin/api/v1/category"
	"betxin/api/v1/collect"
	"betxin/api/v1/message"
	"betxin/api/v1/mixinorder"
	"betxin/api/v1/oauth"
	"betxin/api/v1/snapshot"
	"betxin/api/v1/swaporder"
	"betxin/api/v1/topic"
	"betxin/api/v1/upload"
	"betxin/api/v1/user"
	"betxin/api/v1/usertotopic"
	"betxin/utils"
	"betxin/utils/cors"
	"betxin/utils/jwt"
	"betxin/utils/logger"
	"betxin/utils/session"

	"github.com/gin-contrib/sessions"
	redisStore "github.com/gin-contrib/sessions/redis"
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
	r.Use(gin.Logger())
	r.LoadHTMLFiles("dist/index.html", "dist/welcome.html")

	auth := r.Group("api/v1")
	auth.Use(jwt.JwtToken())
	{
		//管理员
		auth.POST("/administrator/add", administrator.CreateAdministrator)
		auth.DELETE("/administrator/:id", administrator.DeleteAdministrator)
		auth.GET("/administrator/:id", administrator.GetAdministratorInfo)
		auth.POST("/administrator/list", administrator.ListAdministrators)
		auth.PUT("/administrator/:id", administrator.UpdateAdministrator)

		// bonuse 奖金
		auth.POST("/bounuse/add", bonuse.CreateBonuse)
		auth.DELETE("/bonuse/:id", bonuse.DeleteBonuse)
		auth.GET("bonuse/:trace_id", bonuse.GetBonuseByTraceId)
		// auth.GET("bonuse/:id", bonuse.GetBonuseById)
		auth.POST("bonuse/list", bonuse.ListBonuses)
		auth.PUT("bonuse/:id", bonuse.UpdateBonuse)

		// 分类模块
		auth.GET("/category/:id", category.GetCategoryInfo)
		auth.POST("/category/add", category.CreateCatrgory)
		auth.PUT("/category/:id", category.UpdateCategory)
		auth.DELETE("/category/:id", category.DeleteCategory)
		auth.POST("/category/list", category.ListCategories)

		// 收藏
		auth.POST("/collect/add", collect.CreateCollect)
		auth.DELETE("/collect/delete", collect.DeleteCollect)
		auth.GET("/collect/:userId", collect.GetCollect)
		auth.POST("/collect/list", collect.ListCollects)

		// Mixin信息
		auth.POST("/message/add", message.CreateMessage)
		auth.POST("/message/:id", message.DeleteCollect)
		auth.GET("/message/:id", message.GetMessage)
		auth.POST("/message/list", message.ListMessages)
		auth.PUT("/message/:id", message.UpdateMessage)

		// MixinOrder 接收用户的币
		auth.POST("/mixinorder/add", mixinorder.CreateMixinOrder)
		auth.DELETE("/mixinorder/:traceId", mixinorder.DeleteMixinOrder)
		auth.GET("/mixinorder/:traceId", mixinorder.GetMixinOrderById)
		auth.POST("/mixinorder/list", mixinorder.ListMixinOrder)
		auth.PUT("/mixinorder/:traceId", mixinorder.UpdateMixinOrder)

		// snapshot 反馈给用户的钱
		auth.POST("/snapshot/add", snapshot.CreateMixinNetworkSnapshot)
		auth.POST("/snapshot/:traceId", snapshot.DeleteSnapshot)
		auth.GET("/snapshot/:traceId", snapshot.GetMixinNetworkSnapshot)
		auth.POST("/snapshot/list", snapshot.ListMixinNetworkSnapshots)
		auth.PUT("/snapshot/:traceId", snapshot.UpdateMixinNetworkSnapshot)

		// swaporder 管理从4swap的转账金钱
		auth.POST("/swaporder/add", swaporder.CreateSwapOrder)
		auth.DELETE("/swaporder/:traceId", swaporder.DeleteSwapOrder)
		auth.GET("/swaporder/:traceId", swaporder.GetSwapOrder)
		auth.POST("/swaporder/list", swaporder.ListSwapOrder)
		auth.PUT("/swaporder/:traceId", swaporder.UpdateMessage)

		// topic 管理话题
		auth.POST("/topic/add", topic.CreateTopic)
		auth.DELETE("/topic/:tid", topic.DeleteTopic)
		auth.GET("/topic/:tid", topic.GetTopicInfoById)
		auth.POST("/topic/:cid", topic.GetTopicByCid)
		auth.POST("/topic/list", topic.ListTopics)
		auth.POST("topic/stop", topic.StopTopic)

		// upload   上传文件
		auth.POST("/file", upload.Upload)

		// user 用户管理
		auth.POST("/user/add", user.CreateUser)
		auth.DELETE("/user/delete", user.DeleteUser)
		auth.GET("/user/:userId", user.GetUserInfoByUserId)
		// auth.GET("/user/:fullName", user.GetUserInfoByUserFullName)
		auth.POST("/user/list", user.ListUser)
		auth.POST("/user/:userId", user.UpdateUser)

		// usertotopic 用户买的话题
		auth.POST("/usertotopic/add", usertotopic.CreateUserToTopic)
		auth.DELETE("/usertotopic/delete", usertotopic.DeleteUserToTopic)
		auth.POST("/usertotopic/list", usertotopic.ListUserToTopics)
		auth.POST("/usertotopic/:userId", usertotopic.ListUserToTopicsByUserId)
		// auth.POST("/usertotopic/:topicId", usertotopic.ListUserToTopicsByTopicId)
		auth.PUT("/usertotopic/update", usertotopic.UpdateUserToTopic)
	}

	store, _ := redisStore.NewStore(10, "tcp", "localhost:6379", "123456", []byte("secret"))
	r.Use(sessions.Sessions("betxin_api", store))
	r.GET("/oauth/redirect", oauth.MixinOauth)

	router := r.Group("api/v1")
	router.Use(session.AuthMiddleware())
	{
		// 登录控制模块
		router.POST("/login", administrator.Login)

		// 管理用户
		// r.GET("/", func(c *gin.Context) {
		// 	c.HTML(200, "index.html", "flysnow_org")
		// })
		// r.GET("/welcome", func(c *gin.Context) {
		// 	c.HTML(200, "welcome.html", "flysnow_org")
		// })

		// 管理奖金

		//种类

		//收藏

		//

		//话题
		// usertotopic 用户买的话题
		// router.POST("/usertotopic/add", usertotopic.CreateUserToTopic)
		// router.DELETE("/usertotopic/delete", usertotopic.DeleteUserToTopic)
		// router.POST("/usertotopic/list", usertotopic.ListUserToTopics)
		// router.POST("/usertotopic/:userId", usertotopic.ListUserToTopicsByUserId)
		// // auth.POST("/usertotopic/:topicId", usertotopic.ListUserToTopicsByTopicId)
		// router.PUT("/usertotopic/update", usertotopic.UpdateUserToTopic)
		// 用户
		router.POST("/user/info", user.GetUserInfoByUserId)

		// router.POST("/user/add", user)
		// router.POST("/file", upload.Upload)
	}

	_ = r.Run(utils.HttpPort)
}
