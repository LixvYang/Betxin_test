// 从4swap交换出来tx的结构保存到数据库
package model

import (
	"betxin/utils/errmsg"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SwapOrder struct {
	Type       string          `gorm:"type:varchar(20);" json:"type"`
	SnapshotId string          `gorm:"type:varchar(50)" json:"snapshot_id,omitempty"`
	AssetID    string          `gorm:"type:varchar(50)" json:"asset_id"`
	Amount     decimal.Decimal `gorm:"type:decimal(16, 8)" json:"amount"`
	TraceId    string          `gorm:"type:varchar(50);not null" json:"trace_id"`
	Memo       string          `gorm:"type:varchar(255)" json:"memo"`
	State      string          `gorm:"type:varchar(20)" json:"state"`
	CreatedAt  time.Time       `gorm:"type:datetime(3)" json:"created_at"`
}

func CreateSwapOrder(data *SwapOrder) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func DeleteSwapOrder(trace_id string) int {
	if err := db.Delete(&SwapOrder{}, trace_id).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetSwapOrder(trace_id string) (SwapOrder, int) {
	var swapOrder SwapOrder
	if err := db.First(&swapOrder, trace_id).Error; err != nil {
		return SwapOrder{}, errmsg.ERROR
	}
	return swapOrder, errmsg.SUCCSE
}

func ListSwapOrders(pageSize int, pageNum int) ([]*SwapOrder, int) {
	var swapOrder []*SwapOrder
	var total int64
	db.Model(&swapOrder).Count(&total)
	err := db.Find(&swapOrder).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return swapOrder, int(total)
}
