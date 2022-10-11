package oauth

import (
	"betxin/model"
	"betxin/service"
	"betxin/utils"
	"betxin/utils/errmsg"
	"fmt"
	"log"
	"net/http"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func MixinOauth(c *gin.Context) {
	var code = c.Query("code")
	var pathUrl string = "/"

	if c.Query("state") != "" {
		pathUrl = c.Query("state")
	}
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
	// checked := model.CheckUserAuthorization(userinfo.UserID)
	// if checked == errmsg.SUCCSE {
	// 	userauth := &model.UserAuthorization{
	// 		UserId:      userinfo.UserID,
	// 		Provider:    "mixin",
	// 		AccessToken: access_token,
	// 	}
	// 	_ = model.CreateUserAuthorization(userauth)
	// } else {
	// 	// 如果用户已经认证过了 那就更新assess_token
	// 	_ = model.UpdateUserAuthorization(userinfo.UserID, access_token)
	// }

	user := model.User{
		AvatarUrl: userinfo.AvatarURL,
		FullName:  userinfo.FullName,
		MixinId:   userinfo.IdentityNumber,
		IdentityNumber:    userinfo.IdentityNumber,
		MixinUuid: userinfo.UserID,
		SessionId: userinfo.SessionID,
	}
	// 如果用户不存在
	if checked := model.CheckUser(userinfo.UserID); checked != errmsg.SUCCSE {
		if coded := model.CreateUser(&user); coded != errmsg.SUCCSE {
			log.Fatalf("error!!!")
		}

		sessionToken := uuid.NewV4().String()
		session := sessions.Default(c)
		session.Set("userId", user.IdentityNumber)
		session.Set("token", sessionToken)
		session.Save()
	} else {
		//用户存在 就更新数据
		if coded := model.UpdateUser(userinfo.UserID, &user); coded != errmsg.SUCCSE {
			log.Fatalf("更新失败")
		}

		session := sessions.Default(c)
		sessionToken := session.Get("token")
		session.Set("userId", user.IdentityNumber)
		if sessionToken == nil {
			sessionToken = uuid.NewV4().String()
			session.Set("userId", user.IdentityNumber)
			session.Set("token", sessionToken)
			session.Save()
		}
	}

	// v1.SendResponse(c, errmsg.SUCCSE, user)

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprint("http://localhost:8080", pathUrl))
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
