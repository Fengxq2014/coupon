package controllers

import (
	"errors"
	"github.com/Fengxq2014/coupon/common/cache"
	"github.com/Fengxq2014/coupon/common/random"
	"github.com/Fengxq2014/coupon/models"
	"github.com/gin-gonic/gin"
	"github.com/qichengzx/qcloudsms_go"
	"os"
	"strconv"
	"time"
)

type CommonController struct{}

func (ctrl CommonController) SendSMS(c *gin.Context) {
	//resultOk(c, nil)
	//return
	opt := qcloudsms.NewOptions(os.Getenv("SMS_APPID"), os.Getenv("SMS_APPKEY"), os.Getenv("SMS_SIGN"))
	i, err := strconv.Atoi(os.Getenv("SMS_TPLID"))
	if err != nil {
		resultFail(c, err)
		return
	}
	phone := c.Param("phone")
	var customer models.CustomerModel
	_, err = customer.GetCustomer(phone)
	if err != nil {
		resultFail(c, "您还不是特约用户")
		return
	}
	code := strconv.Itoa(random.RandRangeNum(1000, 9999))
	var client = qcloudsms.NewClient(opt)
	b, err := client.SendSMSSingle(qcloudsms.SMSSingleReq{
		Tel: qcloudsms.SMSTel{
			Nationcode: "86",
			Mobile:     phone,
		},
		TplID:  i,
		Params: []string{code, "5分钟"},
		Sig:    os.Getenv("sign"),
	})
	if !b {
		resultFail(c, err)
	} else {
		cache.GetCache().Set(phone, code, 5*time.Minute)
		resultOk(c, nil)
	}
}

func (ctrl CommonController) CheckSMS(c *gin.Context) {
	v, err := cache.GetCache().Get(c.Param("phone"))
	if err != nil {
		resultFail(c, err)
		return
	}
	if v != c.Param("code") {
		resultFail(c, errors.New("验证码错误"))
		return
	}
	resultOk(c, nil)
}
