package model

import (
	"betxin/utils/errmsg"
	"time"
)

type UserAuthorization struct {
	UserId      string    `gorm:"type:varchar(50);not null;index" json:"user_id"`
	Provider    string    `gorm:"type:varchar(50);comment:third party auth provider" json:"provider"`
	AccessToken string    `gorm:"not null" json:"access_token"`
	Raw         string    `gorm:"type:varchar(255);" json:"third pary user info"`
	CreatedAt   time.Time `gorm:"type:datetime(3);" json:"created_at"`
	UpdatedAt   time.Time `gorm:"type:datetime(3);" json:"updated_at"`
	PublicKey   string    `gorm:"type:varchar(255);" json:"public_key"`
}

func CheckUserAuthorization(user_id string) int {
	var userauth UserAuthorization
	db.Select("user_id").Where("user_id = ?", user_id).First(&userauth)
	if userauth.UserId != "" {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

func CreateUserAuthorization(data *UserAuthorization) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func UpdateUserAuthorization(user_id string, access_token string) int {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return errmsg.ERROR
	}

	// 锁住指定 id 的 User 记录
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Last(&UserAuthorization{}, user_id).Error; err != nil {
		tx.Rollback()
		return errmsg.ERROR
	}

	if err := db.Model(&UserAuthorization{}).Where("user_id = ?", user_id).Update("access_token", access_token).Error; err != nil {
		return errmsg.ERROR
	}
	if err := tx.Commit().Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeleteUserAuthorization(user_id string) int {
	if err := db.Delete(&User{}, user_id).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetUserAuthorization(user_id string) (string, int) {
	var userAuthorization *UserAuthorization
	if err := db.First(&userAuthorization, user_id).Error; err != nil {
		return "", errmsg.ERROR
	}
	return userAuthorization.AccessToken, errmsg.SUCCSE
}
