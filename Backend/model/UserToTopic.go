package model

import (
	"betxin/utils/errmsg"
	"time"

	"github.com/shopspring/decimal"
)

type UserToTopic struct {
	Id            int             `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	TopicUuid     string          `gorm:"type:varchar(36);not null;index:useid_topicid_index" json:"tid"`
	UserId        string          `gorm:"type:varchar(50);not null;index:useid_topicid_index" json:"user_id"`
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

func DeleteUserToTopic(userId, TopicUuid string) int {
	if err := db.Where("user_id = ? AND topic_uuid = ?", userId, TopicUuid).Delete(&User{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func ListUserToTopicsByUserId(userId string, offset, limit int) ([]UserToTopic, int, int) {
	var userToTopics []UserToTopic
	var count int64

	if err := db.Model(&userToTopics).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return userToTopics, 0, errmsg.ERROR
	}

	if err := db.Model(&userToTopics).Where("user_id = ?").Limit(limit).Offset(offset).Order("created_at DESC").Find(userToTopics).Error; err != nil {
		return userToTopics, 0, errmsg.ERROR
	}

	return userToTopics, int(count), errmsg.SUCCSE
}

func ListUserToTopicsByTopicId(topicId string, offset, limit int) ([]UserToTopic, int, int) {
	var userToTopics []UserToTopic
	var count int64

	if err := db.Model(&userToTopics).Where("topic_uuid = ?", topicId).Count(&count).Error; err != nil {
		return userToTopics, 0, errmsg.ERROR
	}

	if err := db.Model(&userToTopics).Where("topic_uuid = ?").Limit(limit).Offset(offset).Order("created_at DESC").Find(userToTopics).Error; err != nil {
		return userToTopics, 0, errmsg.ERROR
	}

	return userToTopics, int(count), errmsg.SUCCSE
}

func ListUserToTopics(offset, limit int) ([]UserToTopic, int, int) {
	var userToTopics []UserToTopic
	var count int64

	if err := db.Model(&Topic{}).Count(&count).Error; err != nil {
		return userToTopics, 0, errmsg.ERROR
	}

	if err := db.Limit(limit).Offset(offset).Order("created_at DESC").Find(userToTopics).Error; err != nil {
		return userToTopics, 0, errmsg.ERROR
	}

	return userToTopics, int(count), errmsg.SUCCSE
}
