// 结束topic返还给用户的钱
package model

import (
	"betxin/utils/errmsg"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Bonuse struct {
	Id          int             `gorm:"type:int;primaryKey;autoIncrement" json:"id"`
	UserId      int             `gorm:"type:int;not null;index;" json:"user_id"`
	Title       string          `gorm:"type:varchar(50);not null;" json:"title"`
	Description string          `gorm:"type:varchar(200);not null;" json:"description"`
	AssetId     string          `gorm:"type:varchar(50);not null;" json:"asset_id"`
	Amount     decimal.Decimal `gorm:"type:decimal(16, 8)" json:"amount"`
	Memo        string          `gorm:"type:varchar(255);" json:"memo"`
	TraceId     string          `gorm:"type:varchar(50);not null;uniqueIndex;" json:"trace_id"`
	CreatedAt   time.Time       `gorm:"type:datetime(3); not null" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"type:datetime(3); not null" json:"updated_at"`
}

func CheckBonuse(trace_id string) int {
	var bonuse Bonuse
	db.Select("id").Where("trace_id = ?", trace_id).Last(&bonuse)
	if bonuse.Id > 0 {
		return errmsg.ERROR_BONUSE_EXIST
	}
	return errmsg.SUCCSE
}

func CreateBonuse(data *Bonuse) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func ListBonuses(pageSize int, pageNum int) ([]Bonuse, int) {
	var bonuse []Bonuse
	var total int64
	db.Model(&bonuse).Count(&total)
	err := db.Find(&bonuse).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return bonuse, int(total)
}

func DeleteBonuse(trace_id string) int {
	if err := db.Where("trace_id = ?", trace_id).Delete(&Bonuse{}).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetBonuseByTraceId(trace_id string) (*Bonuse, int) {
	var bonuse *Bonuse
	if err := db.Where("trace_id = ?", trace_id).Last(&bonuse).Error; err != nil {
		return nil, errmsg.ERROR
	}
	return bonuse, errmsg.SUCCSE
}

func GetBonusesByUserId(user_id int) ([]*Bonuse, int) {
	var bonuse []*Bonuse
	var total int64
	db.Model(&bonuse).Count(&total)
	if err := db.Find(&bonuse).Where("user_id = ?", user_id).Error; err != nil {
		return nil, 0
	}
	return bonuse, int(total)
}
