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
		log.Println("Get userInfo fail!!!")
		c.Redirect(http.StatusPermanentRedirect, fmt.Sprint(utils.IP+utils.HttpPort+pathUrl))
	}

	user := model.User{
		AvatarUrl:      userinfo.AvatarURL,
		FullName:       userinfo.FullName,
		MixinId:        userinfo.IdentityNumber,
		IdentityNumber: userinfo.IdentityNumber,
		MixinUuid:      userinfo.UserID,
		SessionId:      userinfo.SessionID,
	}

	// 如果用户不存在
	if checked := model.CheckUser(userinfo.UserID); checked != errmsg.SUCCSE {
		if coded := model.CreateUser(&user); coded != errmsg.SUCCSE {
			log.Println("Get userInfo fail!!!")
		}

		sessionToken := uuid.NewV4().String()
		session := sessions.Default(c)
		session.Set("userId", user.MixinUuid)
		session.Set("token", sessionToken)
		session.Save()
	} else {
		//用户存在 就更新数据
		if coded := model.UpdateUser(userinfo.UserID, &user); coded != errmsg.SUCCSE {
			log.Println("Update userInfo fail!!!")
		}

		session := sessions.Default(c)
		session.Clear()
		sessionToken := uuid.NewV4().String()
		session.Set("userId", user.MixinUuid)
		session.Set("token", sessionToken)
		session.Save()
	}

	c.Redirect(http.StatusPermanentRedirect, fmt.Sprint(utils.IP+utils.HttpPort+pathUrl))
}