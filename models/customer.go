package models

import (
	"github.com/Fengxq2014/coupon/db"
)

type Customer struct {
	Phone string `gorm:"primary_key"`
	Name  string
}

type CustomerModel struct{}

var CustomerTableName = "customer"

func (c CustomerModel) One(phone string) (customer Customer, err error) {
	err = db.GetDB().Table(CustomerTableName).Where("phone = ?", phone).Take(&customer).Error
	return
}
