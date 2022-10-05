package model

import (
	"betxin/utils/errmsg"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/shopspring/decimal"
)

type Topic struct {
	Uuid          uuid.UUID       `gorm:"type:varchar(36);" json:"uuid"`
	CategoryName  string          `gorm:"type:varchar(20)" json:"category_name"`
	Title         string          `gorm:"type:varchar(50);not null;" json:"title"`
	Intro         string          `gorm:"type:varchar(50);not null;" json:"intro"`
	Content       string          `gorm:"type:varchar(50);not null;" json:"content"`
	CollectCount  int             `gorm:"type:int;default 0" json:"collect_count"`
	YesRatio      decimal.Decimal `gorm:"type:decimal(4,2);default 0;" json:"yes_ratio"`
	NoRatio       decimal.Decimal `gorm:"type:decimal(4,2);default 0" json:"no_ratio"`
	YesRatioPrice decimal.Decimal `gorm:"type:decimal(16,8);default 0" json:"yes_ratio_ratio"`
	NoRatioPrice  decimal.Decimal `gorm:"type:decimal(16,8);default 0" json:"no_ratio_ratio"`
	TotalPrice    decimal.Decimal `gorm:"type:decimal(32,8);default 0;" json:"total_price"`
	CreatedAt     time.Time       `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt     time.Time       `gorm:"type:datetime(3)" json:"updated_at"`
}

func (t *Topic) BeforeCreate(tx *gorm.DB) error {
	t.Uuid = uuid.NewV4()
	fmt.Println(t.Uuid)
	return nil
}

// 根据标签id获取标签数据.
func GetTopicById(id int) (Topic, int) {
	var topic Topic
	if err := db.Find(&topic, id).Error; err != nil {
		return Topic{}, errmsg.ERROR
	}
	return topic, errmsg.SUCCSE
}

// 创建新标签
func CreateTopic(data *Topic) int {
	if err := db.Create(data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 根据标签id删除标签
func DeleteTopic(id int) int {
	if err := db.Where("id = ?", id).Delete(&Topic{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func UpdateTopic(data *Topic) int {
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
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Last(&Topic{}, data.Uuid).Error; err != nil {
		tx.Rollback()
		return errmsg.ERROR
	}

	var maps = make(map[string]interface{})
	// maps["CategoryId"] = data.CategoryName
	// maps["Title"] = data.Title
	maps["Intro"] = data.Intro
	maps["Content"] = data.Content
	maps["CollectCount"] = data.CollectCount
	maps["YesRatio"] = data.YesRatio
	maps["NoRatio"] = data.NoRatio
	maps["YesRatioPrice"] = data.YesRatioPrice
	maps["NoRatioPrice"] = data.NoRatioPrice
	maps["TotalPrice"] = data.TotalPrice
	if err := db.Model(&UserToTopic{}).Where("Id = ?", data.Uuid).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	if err := tx.Commit().Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//list Topics
func ListTopics(pageSize int, pageNum int) ([]Topic, int64, int) {
	if pageSize == 0 {
		pageSize = 50
	}

	Topics := make([]Topic, 0)
	var count int64

	if err := db.Model(&Topic{}).Count(&count).Error; err != nil {
		return Topics, count, errmsg.ERROR
	}
	if err := db.Where("").Offset(pageSize).Limit((pageNum - 1) * pageSize).Order("id desc").Find(&Topics).Error; err != nil {
		return Topics, count, errmsg.ERROR
	}

	return Topics, count, errmsg.SUCCSE
}
