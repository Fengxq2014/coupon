package tests

import (
	"github.com/Fengxq2014/coupon/db"
	"github.com/Fengxq2014/coupon/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"reflect"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()
}

func TestCustomer(t *testing.T) {
	var model models.CustomerModel
	phone := "13333333333"
	customer, err := model.GetCustomer(phone)
	if err != nil {
		t.Log(gorm.IsRecordNotFoundError(err))
		t.Log(reflect.TypeOf(err))
		t.Error(err)
	}
	t.Log(customer.Name)
}
