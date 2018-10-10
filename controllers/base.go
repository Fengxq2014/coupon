package controllers

import (
	"github.com/Fengxq2014/coupon/common/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func result(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "msg": msg})
}

func resultOk(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": "成功"})
}

func resultOkMsg(ctx *gin.Context, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": msg})
}

func resultOkData(ctx *gin.Context, data interface{}, msg string) {
	if msg == "" {
		msg = "成功"
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": data, "msg": msg})
}

func resultFail(ctx *gin.Context, err interface{}) {
	log.Info(err)
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusBadRequest, "data": nil, "msg": err})
}
