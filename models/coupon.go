package models

import (
	"github.com/Fengxq2014/coupon/db"
	"time"
)

type Coupon struct {
	ID        string
	Title     string
	Code      string
	Amount    int
	StartTime time.Time
	ExpTime   time.Time
	Content   string
	Remarks   string
}

type CouponModel struct {
	ID string
}

var couponTabelName = "coupon"

func (c CouponModel) Get(id string) (coupon Coupon, err error) {
	err = db.GetDB().Table(couponTabelName).Where("id = ?", id).Take(&coupon).Error
	return
}
