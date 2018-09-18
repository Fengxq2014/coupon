package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qichengzx/qcloudsms_go"
	"os"
)

type CommonController struct{}

func (ctrl CommonController) SendSMS(c *gin.Context) {
	opt := qcloudsms.NewOptions(os.Getenv("appid"), os.Getenv("appkey"), os.Getenv("sign"))
	var client = qcloudsms.NewClient(opt)
	client.SetDebug(true)
	client.SendSMSSingle(qcloudsms.SMSSingleReq{
		Tel: qcloudsms.SMSTel{
			Nationcode: "86",
			Mobile:     c.Param("phone"),
		},
		Type:   0,
		Sign:   "",
		TplID:  0,
		Params: nil,
		Msg:    "",
		Sig:    os.Getenv("sign"),
	})
}
