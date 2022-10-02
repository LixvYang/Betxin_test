package model

import (
	"betxin/utils/errmsg"

	"gorm.io/gorm"
)

type Collect struct {
	gorm.Model
	UserId  int    `gorm:"type:int;not null;index:user_collect_topic;not null;" json:"user_id"`
	TopicId string `gorm:"index:user_collect_topic;type:varchar(50) not null;" json:"topic_id"`
}

//Create Collect
func CreateCollect(data *Collect) int {
	if err := db.Model(&Collect{}).Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// list Collect
func ListCollects(pageSize int, pageNum int) ([]Collect, int64, int) {
	if pageSize == 0 {
		pageSize = 50
	}

	collects := make([]Collect, 0)
	var count int64

	if err := db.Model(&Collect{}).Count(&count).Error; err != nil {
		return collects, count, errmsg.ERROR
	}
	if err := db.Where("").Offset(pageSize).Limit((pageNum - 1) * pageSize).Order("id desc").Find(&collects).Error; err != nil {
		return collects, count, errmsg.ERROR
	}

	return collects, count, errmsg.SUCCSE
}

// 根据标签user_id获取收藏数据.
func GetCollectById(user_id int) (*[]Collect, int) {
	var collects *[]Collect
	// select * from collect from user_id = user_id
	if err := db.Select("user_id = ?", user_id).Find(&collects).Error; err != nil {
		return nil, errmsg.ERROR
	}
	return collects, errmsg.SUCCSE
}

// Delete collect by id
func DeleteCollect(user_id string, topic_id int) int {
	if err := db.Where("user_id = ? and topic_id = ?", user_id, topic_id).Delete(&Collect{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
