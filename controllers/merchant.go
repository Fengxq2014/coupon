package controllers

import (
	"github.com/Fengxq2014/coupon/common/log"
	"github.com/Fengxq2014/coupon/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
	"time"
)

type MerchantController struct{}

func (ctrl MerchantController) Consume(c *gin.Context) {
	var model models.ConsumeModel
	if err := c.ShouldBindJSON(&model); err != nil {
		e := err.(validator.ValidationErrors)
		for _, err := range e {
			log.Info(err.Field)
			resultFail(c, err.Field+"不合法")
			return
		}
	}

	copModel := models.CusCopModel{}
	cc, err := copModel.Get(model.CopID)
	if err != nil {
		log.Errorf("cuscop查询错误：%s", err)
		resultFail(c, "券码错误")
		return
	}
	if cc.Status != 0 {
		resultFail(c, "优惠券状态不能核销")
		return
	}

	err = model.Consume(20, cc.Phone)
	if err != nil {
		resultFail(c, err)
	} else {
		resultOk(c, nil)
	}
}

func (ctrl MerchantController) List(c *gin.Context) {
	model := models.ConsumeModel{MchPhone: c.Query("phone")}
	start, err := time.Parse("2006-01-02", c.Query("start"))
	if err != nil {
		log.Info(err.Error())
		resultFail(c, "时间有误")
		return
	}
	end, err := time.Parse("2006-01-02", c.Query("end"))
	if err != nil {
		resultFail(c, "时间有误")
		return
	}
	var result []models.DailySumModel
	for !start.After(end) {
		sum, err := model.DailySum(start)
		if err != nil {
			resultFail(c, err.Error())
			return
		}
		result = append(result, sum)
		start = start.Add(24*time.Hour)
	}
	//cons, err := model.List(start, end)
	if err != nil {
		resultFail(c, err.Error())
		return
	}
	resultOk(c, result)
}
