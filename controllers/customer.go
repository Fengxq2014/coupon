package controllers

import (
	"github.com/Fengxq2014/coupon/common/cache"
	"github.com/Fengxq2014/coupon/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

type CustomerController struct{}

var customerModel models.CustomerModel

func (ctrl CustomerController) Get(c *gin.Context) {
	phone := c.Param("phone")
	customer, err := customerModel.One(phone)
	if err == nil {
		resultOk(c, customer)
		return
	} else {
		resultFail(c, err)
	}
}

func (ctrl CustomerController) Login(c *gin.Context) {
	var model models.CustomerModel
	if err := c.ShouldBindJSON(&model); err != nil {
		e := err.(validator.ValidationErrors)
		for _, err := range e {
			resultFail(c, err.Field+"不合法")
			return
		}
	}
	v, err := cache.GetCache().Get(model.Phone)
	if err != nil {
		resultFail(c, "请先发送验证码")
		return
	}
	if v != model.Code {
		resultFail(c, "验证码错误")
		return
	}
	customer, err := customerModel.One(model.Phone)
	if err == nil {
		resultOk(c, customer)
		return
	} else {
		resultFail(c, "不是本系统用户")
		return
	}
	resultOk(c, nil)
}

func (ctrl CustomerController) GetCopList(c *gin.Context) {
	var coscopModel models.CusCopModel
	phone := c.Query("phone")
	coscopModel.Phone = phone
	ccs, err := coscopModel.GetList()
	if err != nil {
		resultFail(c, err.Error())
		return
	}
	resultOk(c, ccs)
}

func (ctrl CustomerController) GetCop(c *gin.Context) {
	var model models.CusCopModel
	if err := c.ShouldBindJSON(&model); err != nil {
		e := err.(validator.ValidationErrors)
		for _, err := range e {
			resultFail(c, err.Field+"不合法")
			return
		}
	}
	v, err := cache.GetCache().Get(model.Phone)
	if err != nil {
		resultFail(c, "请先发送验证码")
		return
	}
	code := c.Query("code")
	if v != code {
		resultFail(c, "验证码错误")
		return
	}
	couponModel := models.CouponModel{}
	_, err = couponModel.Get(model.CopID)
	if err != nil {
		resultFail(c, "无效的优惠券")
		return
	}
	copModel := models.CusCopModel{}
	_, err = copModel.Get(model.CopID)
	if err == nil {
		resultFail(c, "您已经领取该优惠券")
		return
	}
	var insert models.CusCopModel
	insert.CopID = model.CopID
	insert.Phone = model.Phone
	err = insert.Insert()
	if err != nil {
		resultFail(c, "领取失败")
		return
	}
	resultOk(c, nil)
}
