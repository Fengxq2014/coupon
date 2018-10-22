package models

import (
	"github.com/Fengxq2014/coupon/db"
	"strconv"
	"time"
)

type CosCop struct {
	CopID  string `gorm:"column:copid"`
	Phone  string
	Status int
	ID     int `gorm:"column:id"`
}
type CusCopModel struct {
	CopID string `json:"cop_id" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}

type CouponItem struct {
	Code            string `json:"code" gorm:"column:code"`
	Denominations   int    `json:"denominations" gorm:"column:amount"`
	OriginCondition int    `json:"originCondition" gorm:"-"`
	Value           int    `json:"value" gorm:"column:amount"`
	Name            string `json:"name" gorm:"column:title"`
	StartAt         Time   `json:"startAt" gorm:"column:start_time"`
	EndAt           Time   `json:"endAt" gorm:"column:exp_time"`
}

var CosCopTabelName = "cus_cop"

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	*t = Time(tm)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (c CusCopModel) Insert(id int) error {
	return db.GetDB().Table(CosCopTabelName).Create(&CosCop{
		CopID:  c.CopID,
		Phone:  c.Phone,
		Status: 0,
		ID:     id,
	}).Error
}

func (c CusCopModel) Get(id string) (cc CosCop, err error) {
	err = db.GetDB().Table(CosCopTabelName).Where("id = ?", id).Take(&cc).Error
	return
}

func (c CusCopModel) GetList() (ccs []CouponItem, err error) {
	err = db.GetDB().Table(CosCopTabelName).Select("cus_cop.id as code, title, amount, start_time, exp_time").Joins("left join coupon on copid = coupon.id").Where("phone=?", c.Phone).Scan(&ccs).Error
	for k := range ccs {
		ccs[k].Value = ccs[k].Denominations
	}
	return
}

func (c CusCopModel) GetByIDAndPhone(status int) (cc CosCop, err error) {
	err = db.GetDB().Table(CosCopTabelName).Where("copid = ? AND phone = ? AND status = ?", c.CopID, c.Phone, status).Take(&cc).Error
	return
}

func (c CusCopModel) IsNotUseID(id int) bool {
	var count int
	db.GetDB().Table(CosCopTabelName).Where("id = ? AND status = 0", id).Count(&count)
	return count <= 0
}

func (c CusCopModel) HasCouponCode(code int) bool {
	return !db.GetDB().Table(CosCopTabelName).Select("count(1)").Joins("left join coupon on copid = coupon.id").Where("cus_cop.id=?", code).RecordNotFound()
}
