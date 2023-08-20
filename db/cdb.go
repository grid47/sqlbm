package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CDB = gorm.DB

func NewCDB(dsn string, models ...interface{}) (*CDB, error) {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(models...)
	if err != nil {
		return nil, err
	}

	return db, nil
}
