package main

// import (
// 	"betxin/utils"
// 	"betxin/utils/errmsg"
// 	"context"
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"

// 	"github.com/fox-one/mixin-sdk-go"
// 	"github.com/jasonlvhit/gocron"
// 	"github.com/shopspring/decimal"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// 	"gorm.io/gorm/schema"
// )

// var db *gorm.DB

// type Currency struct {
// 	AssetId  string          `gorm:"uniqueIndex;type:varchar(255);not null" json:"asset_id"`
// 	PriceUsd decimal.Decimal `gorm:"type:decimal(20,10); not null" json:"price_usd"`
// 	PriceBtc decimal.Decimal `gorm:"type:decimal(20,10); not null" json:"price_btc"`
// 	ChainId  string          `gorm:"type:varchar(255); not null" json:"chain_id"`
// 	IconUrl  string          `gorm:"type:varchar(255)" json:"icon_url"`
// 	Symbol   string          `gorm:"type:varchar(255)" json:"symbol"`

// 	CreatedAt time.Time `gorm:"type:datetime(3)" json:"created_at"`
// 	UpdatedAt time.Time `gorm:"type:datetime(3)" json:"updated_at"`
// }

// var AllCurrency = [...]string{
// 	utils.PUSD,
// 	utils.BTC,
// 	utils.BOX,
// 	utils.XIN,
// 	utils.ETH,
// 	utils.MOB,
// 	utils.USDC,
// 	utils.USDT,
// 	utils.EOS,
// 	utils.SOL,
// 	utils.UNI,
// 	utils.DOGE,
// 	utils.RUM,
// 	utils.DOT,
// 	utils.WOO,
// 	utils.ZEC,
// 	utils.LTC,
// 	utils.SHIB,
// 	utils.BCH,
// 	utils.MANA,
// 	utils.FIL,
// 	utils.BNB,
// 	utils.XRP,
// 	utils.SC,
// 	utils.MATIC,
// 	utils.ETC,
// 	utils.XMR,
// 	utils.DCR,
// 	utils.TRX,
// 	utils.ATOM,
// 	utils.CKB,
// 	utils.LINK,
// 	utils.GTC,
// 	utils.HNS,
// 	utils.DASH,
// 	utils.XLM,
// }

// func updateRedisCurrency(ctx context.Context) {
// 	var wg sync.WaitGroup

// 	for _, currency := range AllCurrency {
// 		wg.Add(1)
// 		go func(currency string) {
// 			defer wg.Done()
// 			asset, err := mixin.ReadNetworkAsset(ctx, currency)
// 			if err != nil {
// 				return
// 			}
// 			fmt.Println(asset)
// 			currencies := &Currency{
// 				AssetId:  asset.AssetID,
// 				PriceUsd: asset.PriceUSD,
// 				PriceBtc: asset.PriceBTC,
// 				ChainId:  asset.ChainID,
// 				IconUrl:  asset.IconURL,
// 				Symbol:   asset.Symbol,
// 			}
// 			// 有值
// 			if CheckCurrency(asset.AssetID) != errmsg.SUCCSE {
// 				UpdateCurrency(currencies)
// 			} else {
// 				CreateCurrency(currencies)
// 			}
// 		}(currency)
// 	}
// 	wg.Wait()
// }

// func DailyCurrency(ctx context.Context) {
// 	gocron.Every(1).Minute().Do(updateRedisCurrency, ctx)
// }

// func InitDb() {
// 	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		utils.DbUser,
// 		utils.DbPassWord,
// 		utils.DbHost,
// 		utils.DbPort,
// 		utils.DbName,
// 	)

// 	var err error
// 	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
// 		// gorm日志模式：Warn
// 		Logger: logger.Default.LogMode(logger.Warn),
// 		// 外键约束
// 		DisableForeignKeyConstraintWhenMigrating: true,
// 		// 禁用默认事务（提高运行速度）
// 		SkipDefaultTransaction: true,
// 		NamingStrategy: schema.NamingStrategy{
// 			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
// 			SingularTable: true,
// 		},
// 	})
// 	if err != nil {
// 		log.Panic("连接数据库失败,请检查参数:", err)
// 	}
// 	db.AutoMigrate(
// 		&Currency{},
// 	)

// 	sqlDB, _ := db.DB()
// 	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
// 	// SetMaxOpenCons 设置数据库的最大连接数量。
// 	// SetConnMaxLifetiment 设置连接的最大可复用时间
// 	sqlDB.SetMaxIdleConns(1000)
// 	sqlDB.SetMaxOpenConns(5000)
// 	sqlDB.SetConnMaxLifetime(time.Hour / 2)
// }

// func CheckCurrency(assetId string) int {
// 	var currency Currency
// 	if err := db.Model(&Currency{}).Where("asset_id = ?", assetId).Find(&currency).Error; err != nil || currency.AssetId != "" {
// 		// 有值
// 		return errmsg.ERROR
// 	}
// 	// 无值
// 	return errmsg.SUCCSE
// }

// func CreateCurrency(data *Currency) int {
// 	if err := db.Model(&Currency{}).Create(&data).Error; err != nil {
// 		return errmsg.ERROR
// 	}
// 	return errmsg.SUCCSE
// }

// func UpdateCurrency(data *Currency) int {
// 	tx := db.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()
// 	if err := tx.Error; err != nil {
// 		return errmsg.ERROR
// 	}

// 	// 锁住指定 id 的 User 记录
// 	if err := tx.Set("gorm:query_option", "FOR UPDATE").Model(&Currency{}).Where("asset_id = ?", data.AssetId).Error; err != nil {
// 		tx.Rollback()
// 		return errmsg.ERROR
// 	}
// 	var maps = make(map[string]interface{})
// 	maps["PriceUsd"] = data.PriceUsd
// 	maps["PriceBtc"] = data.PriceBtc
// 	maps["ChainId"] = data.ChainId
// 	maps["Symbol"] = data.Symbol
// 	if err := db.Model(&Currency{}).Where("asset_id = ?", data.AssetId).Updates(maps).Error; err != nil {
// 		return errmsg.ERROR
// 	}
// 	if err := tx.Commit().Error; err != nil {
// 		return errmsg.ERROR
// 	}
// 	return errmsg.SUCCSE
// }

// func main() {
// 	ctx := context.Background()
// 	InitDb()
// 	// updateRedisCurrency(ctx)
// 	// DailyCurrency(ctx)
	
// }
