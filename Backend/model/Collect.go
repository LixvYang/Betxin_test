package model

import (
	"betxin/utils/errmsg"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Collect struct {
	gorm.Model
	UserId  string    `gorm:"type:varchar(50);not null;index:user_collect_topic" json:"user_id"`
	TopicId uuid.UUID `gorm:"index:user_collect_topic;type:varchar(36) not null;" json:"topic_id"`
}

//Create Collect
func CreateCollect(data *Collect) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// list Collect
func ListCollects(offset int, limit int) ([]Collect, int, int) {
	var collects []Collect
	var count int64

	if err := db.Model(&Collect{}).Count(&count).Error; err != nil {
		return collects, int(count), errmsg.ERROR
	}
	if err := db.Where("").Offset(offset).Limit(limit).Order("id desc").Find(&collects).Error; err != nil {
		return collects, int(count), errmsg.ERROR
	}

	return collects, int(count), errmsg.SUCCSE
}

// 根据标签user_id获取收藏数据.
func GetCollectByUserId(user_id string) ([]Collect, int, int) {
	var collects []Collect
	var total int64
	// select * from collect from user_id = user_id
	if err := db.Select("user_id = ?", user_id).Model(&collects).Count(&total).Error; err != nil {
		return nil, int(total), errmsg.ERROR
	}
	if err := db.Select("user_id = ?", user_id).Find(&collects).Error; err != nil {
		return nil, int(total), errmsg.ERROR
	}
	return collects, int(total), errmsg.SUCCSE
}

// Delete collect by id
func DeleteCollect(user_id string, topic_id uuid.UUID) int {
	if err := db.Where("user_id = ? and topic_id = ?", user_id, topic_id).Delete(&Collect{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
