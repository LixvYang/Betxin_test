package model

import (
	"betxin/utils/errmsg"
	"time"

	"github.com/gofrs/uuid"
	"github.com/shopspring/decimal"
)

type UserToTopic struct {
	Id            int             `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	UserId        string          `gorm:"type:varchar(50);not null;index" json:"user_id"`
	TopicUuid     uuid.UUID       `gorm:"varchar(36);" json:"topic_uuid"`
	YesRatioPrice decimal.Decimal `gorm:"type:decimal(10,10);not null;" json:"yes_ratio_price"`
	NoRatioPrice  decimal.Decimal `gorm:"type:decimal(10,10);not null;" json:"no_ratio_price"`

	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at"`
}

func CreateUserToTopic(data *UserToTopic) int {
	if err := db.Model(&UserToTopic{}).Create(data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func UpdateUserToTopic(data *UserToTopic) int {
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
	maps["YesRatioPrice"] = data.YesRatioPrice
	maps["NoRatioPrice"] = data.NoRatioPrice
	if err := db.Model(&UserToTopic{}).Where("user_id = ?", data.UserId).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	if err := tx.Commit().Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeleteUserToTopic(id int) int {
	if err := db.Delete(&User{}, id).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func ListUserToTopics(pageSize int, pageNum int) ([]UserToTopic, int64, int) {
	if pageSize == 0 {
		pageSize = 50
	}

	userToTopics := make([]UserToTopic, 0)
	var count int64

	if err := db.Model(&Topic{}).Count(&count).Error; err != nil {
		return userToTopics, count, errmsg.ERROR
	}
	if err := db.Where("").Offset(pageSize).Limit((pageNum - 1) * pageSize).Order("id desc").Find(&userToTopics).Error; err != nil {
		return userToTopics, count, errmsg.ERROR
	}

	return userToTopics, count, errmsg.SUCCSE
}
