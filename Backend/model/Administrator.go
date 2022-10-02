package model

import (
	"betxin/utils/errmsg"

	"gorm.io/gorm"
)

type Administrator struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
}

// CreateAdministrator 新增管理员
func CreateAdministrator(data *Administrator) int {
	//data.Password = ScryptPw(data.Password)
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// Delete 管理员
func DeleteAdministrator(id int) int {
	var user User
	if err := db.Where("id = ? ", id).Delete(&user).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (Administrator, int) {
	var user Administrator

	db.Where("username = ? AND password = ?", username, password).First(&user)
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}

	return user, errmsg.SUCCSE
}
