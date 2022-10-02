package oauth

import (
	"log"
	"betxin/model"
	"betxin/service"
	"betxin/utils"
	"betxin/utils/errmsg"
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-gonic/gin"
)

func MixinOauth(c *gin.Context) {
	var code = c.Query("code")
	access_token, _, err := mixin.AuthorizeToken(c, utils.ClientId, utils.AppSecret, code, "")
	if err != nil {
		log.Printf("AuthorizeToken: %v", err)
		return
	}

	userinfo, err := service.GetUserInfo(access_token)
	if err != nil {
		log.Fatalf("获取用户信息失败!!!")
	}
	// 如果用户没有认证过
	checked := model.CheckUserAuthorization(userinfo.UserID)
	if checked == errmsg.SUCCSE {
		userauth := &model.UserAuthorization{
			UserId:      userinfo.UserID,
			Provider:    "mixin",
			AccessToken: access_token,
		}
		_ = model.CreateUserAuthorization(userauth)
	} else {
		// 如果用户已经认证过了 那就更新assess_token
		_ = model.UpdateUserAuthorization(userinfo.UserID, access_token)
	}

	// 如果用户不存在
	if checked := model.CheckUser(userinfo.UserID); checked == errmsg.SUCCSE {
		user := model.User{
			AvatarUrl: userinfo.AvatarURL,
			FullName:  userinfo.FullName,
			MixinId:   userinfo.IdentityNumber,
			UserId:    userinfo.UserID,
		}
		if cod := model.CreateUser(&user); cod == errmsg.ERROR {
			log.Fatalf("error!!!")
		}
	}
	//用户存在 就更新数据

	c.Redirect(http.StatusPermanentRedirect, "/welcome")
}

// func MixinOauth(c *gin.Context) {
// 	var code = c.Query("code")
// 	key := mixin.GenerateEd25519Key()
// 	store, err := mixin.AuthorizeEd25519(c, utils.ClientId, utils.AppSecret, code, "", key)
// 	authorizationID = store.AuthID
// 	if err != nil {
// 		log.Printf("AuthorizeEd25519: %v", err)
// 		return
// 	}

// 	client, err := mixin.NewFromOauthKeystore(store)
// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	user, err := client.UserMe(c)
// 	if err != nil {
// 		log.Printf("UserMe: %v", err)
// 		return
// 	}
// 	log.Println("user", user.UserID)

// 	// snapshots, _ := client.ReadSnapshots(c, "", time.Now(), "", 10)
// 	// for _, snapshot := range snapshots {
// 	// 	fmt.Println("snapshot.TraceID" + snapshot.TraceID)
// 	// 	fmt.Println("snapshot.Receiver" + snapshot.Receiver)
// 	// 	fmt.Println("snapshot.OpponentID" + snapshot.OpponentID)
// 	// 	fmt.Println("snapshot.Sender" + snapshot.Sender)
// 	// }

// 	snapshots, _ := client.ReadSnapshotsWithOptions(c, time.Now(),  10, mixin.ReadSnapshotsOptions{})
// 	for _, snapshot := range snapshots {
// 		fmt.Println("snapshot.TraceID" + snapshot.TraceID)
// 		fmt.Println("snapshot.Receiver" + snapshot.Receiver)
// 		fmt.Println("snapshot.OpponentID" + snapshot.OpponentID)
// 		fmt.Println("snapshot.Sender" + snapshot.Sender)
// 	}

// }
