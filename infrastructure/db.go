package infrastructure

import (
	"module/config"
	"module/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() error {
	_db, err := gorm.Open(postgres.Open(config.GetDsn()), &gorm.Config{})
	if err != nil {
		return err
	}
	db = _db
	err = db.AutoMigrate(&domain.User{}, &domain.Address{})
	if err != nil {
		return err
	}
	return nil
}

func GetDb() *gorm.DB {
	return db
}
