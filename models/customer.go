package models

import (
	"github.com/Fengxq2014/coupon/db"
)

type Customer struct {
	Phone string `gorm:"primary_key"`
	Name  string
	Type  int
}

type CustomerModel struct {
	Phone string `binding:"required"`
	Code  string `binding:"required"`
}

var CustomerTableName = "customer"

func (c CustomerModel) One(phone string) (customer Customer, err error) {
	err = db.GetDB().Table(CustomerTableName).Where("phone = ?", phone).Take(&customer).Error
	return
}

func (c CustomerModel) GetCustomer(phone string) (customer Customer, err error) {
	err = db.GetDB().Table(CustomerTableName).Where("phone = ?", phone).Take(&customer).Error
	return
}
