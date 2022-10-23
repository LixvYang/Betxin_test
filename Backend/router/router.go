package router

import (
	"betxin/api/sd"
	"betxin/api/v1/administrator"
	"betxin/api/v1/bonuse"
	"betxin/api/v1/category"
	"betxin/api/v1/collect"
	"betxin/api/v1/comment"
	"betxin/api/v1/currency"
	"betxin/api/v1/feedback"
	"betxin/api/v1/message"
	"betxin/api/v1/mixinorder"
	"betxin/api/v1/mixpayorder"
	"betxin/api/v1/oauth"
	"betxin/api/v1/praisecomment"
	"betxin/api/v1/sendback"
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

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()

	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)
	r.HTMLRender = createMyRender()
	r.Use(logger.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())
	r.Use(gin.Logger())

	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "./web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	session.EnableCookileSession(r)
	r.POST("/api/v1/backend/login", administrator.Login)
	r.GET("/oauth/redirect", oauth.MixinOauth)
	r.POST("/api/v1/topic/list", topic.ListTopics)
	r.POST("/api/v1/topic/search", topic.GetTopicByTitle)
	r.POST("/api/v1/topic/:cid", topic.GetTopicByCid)
	r.GET("/api/v1/topic/:tid", topic.GetTopicInfoById)
	r.POST("/api/v1/category/list", category.ListCategories)
	r.POST("/api/v1/currency/list", currency.ListCurrencies)
	r.POST("/api/v1/feedback/add", feedback.CreateFeedback)
	r.POST("/api/v1/usertotopic/check", usertotopic.CheckUserToTopic)
	r.POST("/api/v1/usertotopic/:id", usertotopic.GetUserToTopic)

	r.POST("/mixpayorder/:traceid", mixpayorder.GetMixpayOrder)


	// administrator.CreateAdministratorME()

	auth := r.Group("api/v1")
	auth.Use(jwt.JwtToken())
	{
		//管理员
		auth.POST("/backend/administrator/add", administrator.CreateAdministrator)
		auth.DELETE("/backend/administrator/:id", administrator.DeleteAdministrator)
		auth.GET("/backend/administrator/:id", administrator.GetAdministratorInfo)
		auth.POST("/backend/administrator/list", administrator.ListAdministrators)
		auth.PUT("/backend/administrator/:id", administrator.UpdateAdministrator)

		// bonuse 奖金
		// auth.POST("/bounuse/add", bonuse.CreateBonuse)
		auth.DELETE("/backend/bonuse/:id", bonuse.DeleteBonuse)
		auth.GET("/backendbonuse/:trace_id", bonuse.GetBonuseByTraceId)
		// auth.GET("bonuse/:id", bonuse.GetBonuseById)
		auth.POST("/backend/bonuse/list", bonuse.ListBonuses)
		auth.PUT("/backend/bonuse/:id", bonuse.UpdateBonuse)

		// 分类模块
		auth.GET("/backend/category/:id", category.GetCategoryInfo)
		auth.POST("/backend/category/add", category.CreateCatrgory)
		auth.PUT("/backend/category/:id", category.UpdateCategory)
		auth.DELETE("/backend/category/:id", category.DeleteCategory)
		auth.POST("/backend/category/list", category.ListCategories)

		// 收藏
		auth.POST("/backend/collect/list", collect.ListCollects)

		// 加密货币
		auth.POST("/backend/currency/list", currency.ListCurrencies)

		// Mixin信息
		auth.POST("/backend/message/add", message.CreateMessage)
		auth.POST("/backend/message/:id", message.DeleteCollect)
		auth.GET("/backend/message/:id", message.GetMessage)
		auth.POST("/backend/message/list", message.ListMessages)
		auth.PUT("/backend/message/:id", message.UpdateMessage)

		// MixinOrder 接收用户的币
		auth.POST("/backend/mixinorder/add", mixinorder.CreateMixinOrder)
		auth.DELETE("/backend/mixinorder/:traceId", mixinorder.DeleteMixinOrder)
		auth.GET("/backend/mixinorder/:traceId", mixinorder.GetMixinOrderById)
		auth.POST("/backend/mixinorder/list", mixinorder.ListMixinOrder)
		auth.PUT("/backend/mixinorder/:traceId", mixinorder.UpdateMixinOrder)

		// snapshot 反馈给用户的钱
		auth.POST("/backend/snapshot/add", snapshot.CreateMixinNetworkSnapshot)
		auth.POST("/backend/snapshot/:traceId", snapshot.DeleteSnapshot)
		auth.GET("/backend/snapshot/:traceId", snapshot.GetMixinNetworkSnapshot)
		auth.POST("/backend/snapshot/list", snapshot.ListMixinNetworkSnapshots)
		auth.PUT("/backend/snapshot/:traceId", snapshot.UpdateMixinNetworkSnapshot)

		// swaporder 管理从4swap的转账金钱
		auth.POST("/backend/swaporder/add", swaporder.CreateSwapOrder)
		auth.DELETE("/backend/swaporder/:traceId", swaporder.DeleteSwapOrder)
		auth.GET("/backend/swaporder/:traceId", swaporder.GetSwapOrder)
		auth.POST("/backend/swaporder/list", swaporder.ListSwapOrderNoLimit)
		auth.PUT("/backend/swaporder/:traceId", swaporder.UpdateMessage)

		// topic 管理话题
		auth.POST("/backend/topic/add", topic.CreateTopic)
		auth.DELETE("/backend/topic/:id", topic.DeleteTopic)
		auth.POST("/backend/topic/stop", topic.StopTopic)
		auth.POST("/backend/topic/list", topic.ListTopics)
		auth.PUT("/backend/topic/:id", topic.UpdateTopic)

		// upload   上传文件
		auth.POST("/backend/file", upload.Upload)

		// user 用户管理
		auth.POST("/backend/user/add", user.CreateUser)
		auth.DELETE("/backend/user/delete", user.DeleteUser)
		auth.GET("/backend/user/:userId", user.GetUserInfoByUserId)
		// auth.GET("/user/:fullName", user.GetUserInfoByUserFullName)
		auth.POST("/backend/user/list", user.ListUser)
		auth.POST("/backend/user/:userId", user.UpdateUser)

		// usertotopic 用户买的话题
		auth.POST("/backend/usertotopic/add", usertotopic.CreateUserToTopic)
		auth.DELETE("/backend/usertotopic/delete", usertotopic.DeleteUserToTopic)
		auth.POST("/backend/usertotopic/list", usertotopic.ListUserToTopics)
		auth.POST("/backend/usertotopic/:topicId", usertotopic.ListUserToTopicsByTopicId)
		auth.PUT("/backend/usertotopic/update", usertotopic.UpdateUserToTopic)

		// feedback
		auth.POST("/backend/feedback/list", feedback.ListFeedBack)
		auth.DELETE("/backend/feedback/:id", feedback.DeleteFeedBack)

		// sendback
		auth.POST("/backend/sendback/list", sendback.ListSendBack)
		auth.DELETE("/backend/sendback/:id", sendback.DeleteSendback)
		auth.POST("/backend/sendback/add", feedback.CreateFeedback)

		// mixpay order
		auth.POST("/backend/mixpayorder/list", mixpayorder.ListMixpayOrder)

		// praise comment
		auth.POST("/backend/praisecomment/list", praisecomment.ListPraiseComment)

		// comment
		auth.POST("/backend/comment/list", comment.ListComment)


		auth.GET("/backend/health", sd.HealthCheck)
		auth.GET("/backend/disk", sd.DiskCheck)
		auth.GET("/backend/cpu", sd.CPUCheck)
		auth.GET("/backend/ram", sd.RAMCheck)
	}

	router := r.Group("api/v1")
	router.Use(session.AuthMiddleware())
	{
		router.POST("/user/info", user.GetUserInfoByUserId)
		router.POST("/usertotopic/list", usertotopic.ListUserToTopicsByUserIdNoLimit)
		router.POST("/usertotopic/add", usertotopic.CreateUserToTopic)

		router.POST("/collect/list", collect.GetCollectByUserId)
		router.POST("/collect/add", collect.CreateCollect)
		router.POST("/collect/check", collect.CheckCollect)
		router.POST("/collect/delete", collect.DeleteCollect)

		router.POST("/comment/:tid", comment.ListCommentByTid)
		router.POST("/comment/add", comment.CreateComment)
		router.GET("/comment/:id", comment.GetCommentById)

		router.POST("/praisecomment/delete", praisecomment.DeletePraiseComment)
		router.POST("/praisecomment/add", praisecomment.CreatePraiseComment)
		router.POST("/praisecomment/check", praisecomment.CheckPraise)

		router.POST("/mixpayorder/add", mixpayorder.CreateMixinpayOrder)
		router.POST("/mixpayorder/update", mixpayorder.UpdateMixpayOrder)
	}

	_ = r.Run(utils.HttpPort)
}
