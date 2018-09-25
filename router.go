package main

import (
	"github.com/Fengxq2014/coupon/common/log"
	"github.com/Fengxq2014/coupon/controllers"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
)

func initRouter() *gin.Engine {
	if gin.Mode() == "release" {
		pwd, _ := os.Getwd()
		s := filepath.Join(pwd, "log", "server.log")
		file, err := os.OpenFile(s, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Error("open log file error")
			os.Exit(1)
		}
		gin.DefaultErrorWriter = io.MultiWriter(file, os.Stdout)
	}

	r := gin.New()
	r.Use(loggerWithWriter(gin.DefaultWriter), gin.Recovery())

	v1 := r.Group("/v1")
	{
		customer := new(controllers.CustomerController)
		common := new(controllers.CommonController)
		mch := new(controllers.MerchantController)
		v1.GET("/customer/:phone", customer.Get)
		v1.GET("/common/sms/check", common.CheckSMS)
		v1.GET("/common/sms/send/:phone", common.SendSMS)
		v1.POST("/mch/consume", mch.Consume)
	}

	return r
}
