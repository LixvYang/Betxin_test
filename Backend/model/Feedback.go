package model

import (
	"betxin/utils/errmsg"
	"time"
)

type FeedBack struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId    string    `gorm:"type:varchar(36);not null;index" json:"user_id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	CreatedAt time.Time `gorm:"type:datetime(3); not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(3);not null;" json:"updated_at"`
}

func CreateFeedBack(data *FeedBack) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func ListFeedBack() ([]FeedBack, int, int) {
	var feedback []FeedBack
	var err error
	var total int64

	err = db.Model(&feedback).Count(&total).Error
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return feedback, int(total), errmsg.SUCCSE
}

// 根据user_id获取messages
func ListFeedBackByUserId(user_id string) ([]FeedBack, int, int) {
	var message []FeedBack
	var err error
	var total int64

	err = db.Where("user_id = ?", user_id).Error
	db.Model(&message).Count(&total)
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return message, int(total), errmsg.SUCCSE
}

func DeleteFeedBackByMessageId(id string) int {
	if err := db.Where("id = ?", id).Delete(&MixinMessage{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
