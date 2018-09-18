package controllers

import (
	"github.com/Fengxq2014/coupon/common/cache"
	"github.com/gin-gonic/gin"
	"github.com/qichengzx/qcloudsms_go"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type CommonController struct{}

func (ctrl CommonController) SendSMS(c *gin.Context) {
	opt := qcloudsms.NewOptions(os.Getenv("SMS_APPID"), os.Getenv("SMS_APPKEY"), os.Getenv("SMS_SIGN"))
	i, err := strconv.Atoi(os.Getenv("SMS_TPLID"))
	if err != nil {
		resultFail(c, err)
		return
	}
	code := strconv.Itoa(rand.Intn(1001))
	phone := c.Param("phone")
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
