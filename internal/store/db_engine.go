package store

import (
	"mgp_example/config"
	"mgp_example/internal/store/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	defaultDB *gorm.DB
)

func GetDefaultDB() *gorm.DB {
	return defaultDB
}

func initDefaultDB() {
	var err error
	defaultDB, err = gorm.Open(mysql.Open(config.GetConf().MySQL.Dsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		TranslateError:                           true,
		PrepareStmt:                              true,
		QueryFields:                              true,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := defaultDB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if err = defaultDB.AutoMigrate(
		new(model.Audit),
	); err != nil {
		panic(err)
	}

	if config.GetConf().LogLevel == "debug" {
		defaultDB = defaultDB.Debug()
	}
}

func init() {
	initDefaultDB()
}
