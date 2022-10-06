package model

import (
	"fmt"
	"log"

	"betxin/utils"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	redisClient *redis.Client
	db          *gorm.DB
)

func InitDb() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：Warn
		Logger: logger.Default.LogMode(logger.Warn),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		log.Panic("连接数据库失败,请检查参数:", err)
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
	ping, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(ping)
	// 如果存在表则删除（删除时会忽略、删除外键约束)
	// db.Migrator().DropTable(&User{})
	// db.Migrator().DropTable(&Category{})
	// db.Migrator().DropTable(&Topic{})
	// db.Migrator().DropTable(&Collect{})
	// db.Migrator().DropTable(&Bonuse{})
	// db.Migrator().DropTable(&Currency{})
	// db.Migrator().DropTable(&MixinMessage{})
	// db.Migrator().DropTable(&SwapOrder{})
	// db.Migrator().DropTable(&MixinNetworkSnapshot{})
	// db.Migrator().DropTable(&UserAuthorization{})
	// db.Migrator().DropTable(&MixinOrder{})
	// db.Migrator().DropTable(&UserToTopic{})
	// db.Migrator().DropTable(&Administrator{})

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	db.AutoMigrate(
		// &User{},
		&Category{},
		&Topic{},
		// &Collect{},
		// &Bonuse{},
		// &Currency{},
		// &MixinMessage{},
		// &SwapOrder{},
		// &MixinNetworkSnapshot{},
		// &UserAuthorization{},
		// &MixinOrder{},
		// &UserToTopic{},
		// &Administrator{},
	)
	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	// SetMaxOpenCons 设置数据库的最大连接数量。
	// SetConnMaxLifetiment 设置连接的最大可复用时间
	sqlDB.SetMaxIdleConns(1000)
	sqlDB.SetMaxOpenConns(100000)
	sqlDB.SetConnMaxLifetime(-1)
}
