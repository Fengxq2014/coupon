package controllers

import (
	"github.com/Fengxq2014/coupon/common/log"
	"github.com/Fengxq2014/coupon/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
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
	cc, err := copModel.Get(model.CuscopID)
	if err != nil {
		log.Errorf("cuscop查询错误：%s", err)
		resultFail(c, "未领取优惠券")
		return
	}
	if cc.Status != 0 {
		resultFail(c, "优惠券状态不能核销")
		return
	}

	err = model.Consume(20)
	if err != nil {
		resultFail(c, err)
	} else {
		resultOk(c, nil)
	}
}
