package controllers

import (
	"github.com/Fengxq2014/coupon/models"
	"github.com/gin-gonic/gin"
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
