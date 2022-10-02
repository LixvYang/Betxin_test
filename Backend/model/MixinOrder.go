// 记录mixin转账到机器人的交易  接收用户转账
package model

import (
	"betxin/utils/errmsg"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type MixinOrder struct {
	Type       string          `gorm:"type:varchar(20)" json:"type"`
	SnapshotId string          `gorm:"type:varchar(50)" json:"snapshot_id"`
	AssetId    string          `gorm:"type:varchar(50);not null;index" json:"asset_id"`
	Amount     decimal.Decimal `gorm:"type:decimal(16, 8)" json:"amount"`
	TraceId    string          `gorm:"type:varchar(50);not null;unique" json:"trace_id"`
	Memo       string          `gorm:"type:varchar(255);" json:"memo"`
	CreatedAt  time.Time       `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:datetime(3)" json:"updated_at"`
}

func CreateMixinOrder(data *MixinOrder) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeleteMixinOrder(trace_id int) int {
	if err := db.Delete(&User{}, trace_id).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetMixinOrderById(trace_id string) (MixinOrder, int) {
	var mixinOrder MixinOrder
	if err := db.First(&mixinOrder, trace_id).Error; err != nil {
		return MixinOrder{}, errmsg.ERROR
	}
	return mixinOrder, errmsg.SUCCSE
}

func ListMixinOrder(pageSize, pageNum int) ([]MixinOrder, int) {
	var mixinOrder []MixinOrder
	var total int64
	db.Model(&mixinOrder).Count(&total)
	err := db.Find(&mixinOrder).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return mixinOrder, int(total)
}
