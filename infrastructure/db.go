package infrastructure

import (
	"module/config"
	"module/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() {
	_db, err := gorm.Open(postgres.Open(config.GetDsn()), &gorm.Config{})
	if err != nil {
		util.HandleError(err)
	}
	db = _db
	err = db.AutoMigrate()
}
