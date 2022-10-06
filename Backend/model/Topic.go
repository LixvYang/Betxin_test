package model

import (
	"betxin/utils/errmsg"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"

	"github.com/shopspring/decimal"
)

type Topic struct {
	Uuid          uuid.UUID       `gorm:"type:varchar(36);index;" json:"uuid"`
	Cid           int             `gorm:"type:int;not null" json:"cid"`
	Category      Category        `gorm:"foreignKey:Cid" json:"category"`
	Title         string          `gorm:"type:varchar(50);not null;index:title_intro_content_topic_index" json:"title"`
	Intro         string          `gorm:"type:varchar(50);not null;index:title_intro_content_topic_index" json:"intro"`
	Content       string          `gorm:"type:varchar(50);not null;index:title_intro_content_topic_index" json:"content"`
	CollectCount  int             `gorm:"type:int;default 0" json:"collect_count"`
	YesRatio      decimal.Decimal `gorm:"type:decimal(4,2);default 0.00;" json:"yes_ratio"`
	NoRatio       decimal.Decimal `gorm:"type:decimal(4,2);default 0.00" json:"no_ratio"`
	YesRatioPrice decimal.Decimal `gorm:"type:decimal(16,8);default 0" json:"yes_ratio_ratio"`
	NoRatioPrice  decimal.Decimal `gorm:"type:decimal(16,8);default 0" json:"no_ratio_ratio"`
	TotalPrice    decimal.Decimal `gorm:"type:decimal(32,8);default 0;" json:"total_price"`
	ReadCount     int             `gorm:"type:int;not null;default:0" json:"read_count"`

	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at"`
}

func (t *Topic) BeforeCreate(tx *gorm.DB) error {
	t.Uuid = uuid.NewV4()
	return nil
}

func CheckTopic(title string) int {
	var topic Topic
	db.Select("uuid").Where("title = ?", title).First(&topic)
	if topic.Intro != "" {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
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
func GetTopicById(uuid uuid.UUID) (Topic, int) {
	var topic Topic
	err := db.Where("uuid = ?", uuid).Preload("Category").First((&topic)).Error
	db.Model(&topic).Where("uuid = ?", uuid).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return topic, errmsg.ERROR
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

// GetArt 查询文章列表
func ListTopics(offset int, limit int) ([]Topic, int, int) {
	var topicList []Topic
	var err error
	var total int64
	err = db.Select("topic.uuid, cid, title, intro,  content, collect_count, yes_ratio, no_ratio, yes_ratio_price, no_ratio_price, total_price, read_count, Category.category_name, Category.id, created_at, updated_at").
		Limit(limit).Offset(offset).Order("Created_At DESC").Joins("Category").Find(&topicList).Error
	// 单独计数
	db.Model(&topicList).Count(&total)
	if err != nil {
		return nil, 0, errmsg.ERROR
	}
	return topicList, int(total), errmsg.SUCCSE
}

// 搜索标题
func SearchTopic(title, content, intro string, offset int, limit int) ([]Topic, int, int) {
	var topicList []Topic
	var err error
	var total int64
	err = db.Select("topic.uuid, cid, title, intro,  content, collect_count, yes_ratio, no_ratio, yes_ratio_price, no_ratio_price, total_price, read_count, Category.category_name, Category.id, created_at, updated_at").
		Order("Created_At DESC").Joins("Category").Where("title LIKE ? OR content = ? OR intro = ?", "%"+title+"%", "%"+content+"%", "%"+intro+"%").Limit(limit).Offset(offset).Find(&topicList).Error
	//单独计数
	db.Model(&topicList).Where("title LIKE ? OR content = ? OR intro = ?", "%"+title+"%", "%"+content+"%", "%"+intro+"%").Count(&total)
	if err != nil {
		return nil, int(total), errmsg.ERROR
	}
	return topicList, int(total), errmsg.SUCCSE
}
