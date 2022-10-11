package model

import (
	"betxin/utils/errmsg"
	"errors"
	"time"

	"gorm.io/gorm"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Topic struct {
	Tid           string          `gorm:"type:varchar(36);index;" json:"tid"`
	Cid           int             `gorm:"type:int;not null" json:"cid"`
	Category      Category        `gorm:"foreignKey:Cid" json:"category"`
	Title         string          `gorm:"type:varchar(50);not null;index:title_intro_topic_index" json:"title"`
	Intro         string          `gorm:"type:varchar(255);not null;index:title_intro_topic_index" json:"intro"`
	CollectCount  int             `gorm:"type:int;default 0" json:"collect_count"`
	YesRatio      decimal.Decimal `gorm:"type:decimal(5,2);default 0.00;" json:"yes_ratio"`
	NoRatio       decimal.Decimal `gorm:"type:decimal(5,2);default 0.00" json:"no_ratio"`
	YesRatioPrice decimal.Decimal `gorm:"type:decimal(16,8);default 0" json:"yes_ratio_price"`
	NoRatioPrice  decimal.Decimal `gorm:"type:decimal(16,8);default 0" json:"no_ratio_price"`
	TotalPrice    decimal.Decimal `gorm:"type:decimal(32,8);default 0;" json:"total_price"`
	ReadCount     int             `gorm:"type:int;not null;default:0" json:"read_count"`
	ImgUrl        string          `gorm:"varchar(255);" json:"img_url"`
	IsStop        int             `gorm:"type:int;default 0;" json:"is_stop"`

	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at"`
}

func (t *Topic) BeforeCreate(tx *gorm.DB) error {
	t.Tid = uuid.NewV4().String()
	return nil
}

func (t *Topic) BeforeUpdate(tx *gorm.DB) error {
	if t.IsStop == 1 {
		return errors.New("话题已经停止")
	}
	return nil
}

func CheckTopic(title string) int {
	var topic Topic
	db.Select("tid").Where("title = ?", title).First(&topic)
	if topic.Intro != "" {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 将某个话题停止
func StopTopic(uuid string) int {
	if err := db.Model(&Topic{}).Where("tid = ?", uuid).Update("is_stop", 1).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetTopicTotalPrice(tid string) (int, int) {
	var totalPrice int
	if err := db.Model(&Topic{}).Where("tid = ?", tid).First(&totalPrice).Error; err != nil {
		return totalPrice, errmsg.ERROR
	}
	return totalPrice, errmsg.SUCCSE
}

// GetCateArt 查询分类下的所有话题
func GetTopicByCid(cid int, limit int, offset int) ([]Topic, int, int) {
	var topicList []Topic
	var total int64
	err := db.Preload("Category").Limit(limit).Offset(offset).Where("cid =?", cid).Find(&topicList).Error
	db.Model(&topicList).Where("cid =?", cid).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}

	return topicList, int(total), errmsg.SUCCSE
}

// 根据uuid获取话题数据.  查询单个话题
func GetTopicById(uuid string) (Topic, int) {
	var topic Topic
	err := db.Where("tid = ?", uuid).Preload("Category").Joins("Category").First((&topic)).Error
	db.Model(&topic).Where("tid = ?", uuid).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return topic, errmsg.ERROR
	}
	return topic, errmsg.SUCCSE
}

// 创建新话题
func CreateTopic(data *Topic) int {
	if err := db.Create(data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 根据tid删除标签
func DeleteTopic(tid string) int {
	if err := db.Where("tid = ?", tid).Delete(&Topic{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func UpdateTopic(uuid string, data *Topic) int {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return errmsg.ERROR
	}

	// 锁住指定 id 记录
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Last(&Topic{}, uuid).Error; err != nil {
		tx.Rollback()
		return errmsg.ERROR
	}

	var maps = make(map[string]interface{})
	maps["cid"] = data.Cid
	maps["Intro"] = data.Intro
	maps["title"] = data.Title
	maps["CollectCount"] = data.CollectCount
	maps["YesRatio"] = data.YesRatio
	maps["NoRatio"] = data.NoRatio
	maps["YesRatioPrice"] = data.YesRatioPrice
	maps["NoRatioPrice"] = data.NoRatioPrice
	maps["TotalPrice"] = data.TotalPrice
	maps["ImgUrl"] = data.ImgUrl

	var topic Topic
	if err := db.Where("uuid = ?", uuid).Model(&Topic{}).First(&topic).Error; err != nil || topic.IsStop == 1 {
		tx.Rollback()
		return errmsg.ERROR
	}

	if err := db.Model(&Topic{}).Where("tid = ?", uuid).Updates(maps).Error; err != nil {
		return errmsg.ERROR
	}
	if err := tx.Commit().Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 更新话题某一个价钱
// func UpdateTopicPrice

// 更新话题的总价钱
func UpdateTopicTotalPrice(tid string, selectWin string, plusPrice decimal.Decimal) int {
	// selectWin yes_ratio, no_ratio
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return errmsg.ERROR
	}

	// 锁住指定 id 记录
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Last(&Topic{}, tid).Error; err != nil {
		tx.Rollback()
		return errmsg.ERROR
	}

	db.Model(&Topic{}).Where("tid = ?", tid).Update("total_price", gorm.Expr("total_price + ?", plusPrice))

	if selectWin == "yes_win" {
		db.Model(&Topic{}).Where("tid = ?", tid).Update("yes_ratio", gorm.Expr("(? + yes_ratio_prict)/total_price", plusPrice)).Update("yes_ratio_price", gorm.Expr("yes_ratio_price + ?", plusPrice))
	} else {
		db.Model(&Topic{}).Where("tid = ?", tid).Update("no_ratio", gorm.Expr("(? + no_ratio_prict)/total_price", plusPrice)).Update("no_ratio_price", gorm.Expr("no_ratio_price + ?", plusPrice))
	}

	if err := tx.Commit().Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetArt 查询话题列表
func ListTopics(offset int, limit int) ([]Topic, int, int) {
	var topicList []Topic
	var err error
	var total int64
	err = db.Select("tid, cid, title, intro, collect_count, yes_ratio, no_ratio, yes_ratio_price, no_ratio_price, total_price, read_count,img_url, Category.category_name, Category.id, created_at, updated_at").
		Limit(limit).Offset(offset).Order("Created_At DESC").Joins("Category").Find(&topicList).Error
	// 单独计数
	db.Model(&topicList).Count(&total)
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return topicList, int(total), errmsg.SUCCSE
}

// 搜索标题
func SearchTopic(offset int, limit int, query interface{}, args ...interface{}) ([]Topic, int, int) {
	var topicList []Topic
	var err error
	var total int64
	err = db.Select("tid, cid, title, intro, collect_count, yes_ratio, no_ratio, yes_ratio_price, no_ratio_price, total_price, read_count,img_url, Category.category_name, Category.id, created_at, updated_at").
		Order("Created_At DESC").Joins("Category").Where(query, args...).Limit(limit).Offset(offset).Find(&topicList).Count(&total).Error
	//单独计数
	// db.Model(&topicList).Count(&total)
	if err != nil {
		return nil, int(total), errmsg.ERROR
	}
	return topicList, int(total), errmsg.SUCCSE
}

// // 搜索标题
// func SearchTopic(title, content, intro string, offset int, limit int) ([]Topic, int, int) {
// 	var topicList []Topic
// 	var err error
// 	var total int64
// 	err = db.Select("tid, cid, title, intro,  content, collect_count, yes_ratio, no_ratio, yes_ratio_price, no_ratio_price, total_price, read_count, Category.category_name, Category.id, created_at, updated_at").
// 		Order("Created_At DESC").Joins("Category").Where("title LIKE ? OR content = ? OR intro = ?", "%"+title+"%", "%"+content+"%", "%"+intro+"%").Limit(limit).Offset(offset).Find(&topicList).Error
// 	//单独计数
// 	db.Model(&topicList).Where("title LIKE ? OR content = ? OR intro = ?", "%"+title+"%", "%"+content+"%", "%"+intro+"%").Count(&total)
// 	if err != nil {
// 		return nil, int(total), errmsg.ERROR
// 	}
// 	return topicList, int(total), errmsg.SUCCSE
// }
