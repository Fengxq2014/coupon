package main

import (
	"fmt"
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
	var noLogPath []string
	noLogPath = append(noLogPath, "/front/")
	filepath.Walk("./front", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		noLogPath = append(noLogPath, "/" + path)
		return nil
	})
	r.Use(loggerWithWriter(gin.DefaultWriter, noLogPath...), gin.Recovery(), CORSMiddleware())
	r.Static("/front", "./front")

	v1 := r.Group("/v1")
	{
		customer := new(controllers.CustomerController)
		common := new(controllers.CommonController)
		mch := new(controllers.MerchantController)
		//v1.GET("/customer/:phone", customer.Get)
		v1.GET("/common/sms/check", common.CheckSMS)
		v1.GET("/common/sms/send/:phone", common.SendSMS)
		v1.POST("/mch/consume", mch.Consume)
		v1.GET("/mch/list", mch.List)
		v1.POST("/customer/login", customer.Login)
		v1.GET("/customer/coplist", customer.GetCopList)
	}

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
