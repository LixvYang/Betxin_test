// 记录 topic结束  转账给用户的表
package model

import (
	"betxin/utils/errmsg"
	"time"

	"github.com/shopspring/decimal"
)

type MixinNetworkSnapshot struct {
	SnapshotId string          `gorm:"type:varchar(50)" json:"snapshot_id"`
	TraceId    string          `gorm:"type:varchar(50);not null;" json:"trace_id"`
	UserId     string          `gorm:"type:varchar(50);not null" json:"user_id,omitempty"`
	AssetId    string          `gorm:"type:varchar(50);not null;index" json:"asset_id"`
	OpponentID string          `gorm:"type:varchar(50)" json:"opponent_id,omitempty"`
	Amount     decimal.Decimal `gorm:"type:decimal(16, 8)" json:"amount"`
	
	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at"`
}

func CreateMixinNetworkSnapshot(data *MixinNetworkSnapshot) int {
	if err := db.Create(&data).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.ERROR
}

func DeleteMixinNetworkSnapshot(trace_id string) int {
	if err := db.Delete(&MixinNetworkSnapshot{}, trace_id).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetMixinNetworkSnapshot(trace_id string) int {
	var mixinNetworkSnapshot *MixinNetworkSnapshot
	if err := db.First(mixinNetworkSnapshot, trace_id).Error; err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
