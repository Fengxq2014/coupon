package db

import (
	"github.com/Fengxq2014/coupon/common/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	DB, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=feng sslmode=disable")
	if err != nil{
		log.Fatalf("connect db error:%s", err)
	}
	defer DB.Close()
}
