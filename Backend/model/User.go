package model

import (
	"betxin/utils/errmsg"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        int       `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	UserId    string    `gorm:"type:varchar(50);not null;index" json:"user_id"`
	MixinUuid uuid.UUID `gorm:"index;" json:"mixin_uuid"`
	FullName  string    `gorm:"type:varchar(50);not null" json:"full_name"`
	AvatarUrl string    `gorm:"type:varchar(255);not null" json:"avatar_url"`
	MixinId   string    `gorm:"type:varchar(50);not null;index;" json:"mixin_id"`
	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("创建之间")
	u.MixinUuid = uuid.NewV4()
	fmt.Println(u.MixinUuid)
	return nil
}

// CheckUser 查询用户是否存在
func CheckUser(user_id string) int {
	var user User
	db.Select("id").Where("user_id = ?", user_id).First(&user)
	if user.Id != 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// Create user
func CreateUser(data *User) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//
func GetUserById(id int) (User, int) {
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return User{}, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

//Delete user
func DeleteUser(id int) int {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// EditUser 编辑用户信息
func EditUser(id string, data *User) int {
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
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Last(&User{}, data.UserId).Error; err != nil {
		tx.Rollback()
		return errmsg.ERROR
	}

	var maps = make(map[string]interface{})
	maps["username"] = data.FullName
	if err := db.Model(&User{}).Where("id = ? ", id).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	if err := tx.Commit().Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetUserByName gets an user by the username.
func GetUserByName(full_name string) (*User, int) {
	var user *User
	if err := db.Where("full_name = ?", full_name).First(&user).Error; err != nil {
		return nil, errmsg.ERROR
	}
	return user, errmsg.SUCCSE
}

//List users
func ListUser(pageSize int, pageNum int) ([]*User, int64, int) {
	if pageSize == 0 {
		pageSize = 50
	}

	users := make([]*User, 0)
	var count int64
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return users, count, errmsg.ERROR
	}

	if err := db.Where("").Offset(pageSize).Limit((pageNum - 1) * pageSize).Order("id desc").Find(&users).Error; err != nil {
		return users, count, errmsg.ERROR
	}
	return users, count, errmsg.SUCCSE
}
