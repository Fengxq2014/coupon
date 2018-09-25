package models

import (
	"github.com/Fengxq2014/coupon/db"
	"github.com/satori/go.uuid"
)

type CosCop struct {
	ID     string
	CopID  string
	Phone  string
	Status int
}
type CusCopModel struct {
	CopID string
	Phone string
}

var CosCopTabelName = "cus_cop"

func (c CusCopModel) Insert() error {
	return db.GetDB().Create(&CosCop{
		ID:     uuid.NewV4().String(),
		CopID:  c.CopID,
		Phone:  c.Phone,
		Status: 0,
	}).Error
}

func (c CusCopModel) Get(id string) (cc CosCop, err error) {
	err = db.GetDB().Table(CosCopTabelName).Where("id = ?", id).Take(&cc).Error
	return
}
