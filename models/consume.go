package models

import (
	"github.com/Fengxq2014/coupon/db"
	"github.com/satori/go.uuid"
)

type Consume struct {
	ID       string `gorm:"column:id;primary_key"`
	CuscopID string `gorm:"column:cuscopid"`
	CusID    string `gorm:"column:cusid"`
	MchID    string `gorm:"column:mchid"`
	Amount   int    `gorm:"default:0"`
}

type ConsumeModel struct {
	CuscopID string `json:"cuscopid" binding:"required"`
	CusID    string `json:"cusid" binding:"required"`
	MchID    string `json:"mchid" binding:"required"`
}

var ConsumeTableName = "consume"

func (con ConsumeModel) Consume(amount int) error {
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
		ID:       uuid.NewV4().String(),
		CuscopID: con.CuscopID,
		CusID:    con.CusID,
		MchID:    con.MchID,
		Amount:   amount,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Table(CosCopTabelName).Update("status", "1").Where("id", con.CuscopID).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
