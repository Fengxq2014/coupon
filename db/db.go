package db

import (
	"github.com/Fengxq2014/coupon/common/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=coupon password=feng sslmode=disable")
	if err != nil {
		log.Fatalf("Connect db error:%s", err)
	}
	db.LogMode(true)
}

func Close() error {
	return db.Close()
}

func GetDB() *gorm.DB {
	return db
}
