package models

import (
	"github.com/Fengxq2014/coupon/db"
	"github.com/satori/go.uuid"
	"time"
)

type Consume struct {
	ID        string    `gorm:"column:id;primary_key"`
	CuscopID  string    `gorm:"column:cuscopid"`
	CustPhone string    `gorm:"column:cust_phone"`
	MchPhone  string    `gorm:"column:mch_phone"`
	Amount    int       `gorm:"default:0"`
	SuccessAt time.Time `gorm:"column:success_at"`
}

type ConsumeModel struct {
	CopID    string `json:"code" binding:"required"`
	MchPhone string `json:"mchid" binding:"required"`
}

type DailySumModel struct {
	Date string `json:"date"`
	Sum int `json:"sum"`
}

var ConsumeTableName = "consume"

func (con ConsumeModel) Consume(amount int, custPhone string) error {
	tx := db.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Table(ConsumeTableName).Create(&Consume{
		ID:        uuid.NewV4().String(),
		CuscopID:  con.CopID,
		CustPhone: custPhone,
		MchPhone:  con.MchPhone,
		Amount:    amount,
		SuccessAt: time.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table(CosCopTabelName).Where("copid = ?", con.CopID).Update("status", "1").Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (con ConsumeModel) List(start, end time.Time) (cons []Consume, err error) {
	err = db.GetDB().Table(ConsumeTableName).Where("mch_phone = ? AND success_at >= ? AND success_at <=?", con.MchPhone, start, end).Scan(&cons).Error
	for k := range cons {
		cons[k].SuccessAt = cons[k].SuccessAt.Add(-8*time.Hour)
	}
	return
}

func (con ConsumeModel) DailySum(date time.Time) (dailySumModel DailySumModel, err error) {
	dailySumModel.Date = date.Format("2006-01-02")
	err = db.GetDB().Table(ConsumeTableName).Select("sum(amount)").Where("mch_phone = ? AND success_at >= ? AND success_at <=?", con.MchPhone, date, date.Add(23*time.Hour).Add(59*time.Minute).Add(59*time.Second)).First(&dailySumModel).Error
	return
}
