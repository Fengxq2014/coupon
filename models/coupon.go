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
	StartTime time.Time `gorm:"column:start_time"`
	ExpTime   time.Time `gorm:"column:exp_time"`
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

func (c CouponModel) GetByCusCopID(id string) (coupon Coupon, err error) {
	err = db.GetDB().Table(couponTabelName).Select("coupon.id, title, amount, start_time, exp_time").Joins("left join cus_cop on copid = coupon.id").Where("cus_cop.id = ?", id).Take(&coupon).Error
	return
}
