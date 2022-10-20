package controller

import (
	"ginapp/param"
	"ginapp/service"
	"ginapp/tool"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendCode", mc.sendSmsCode)
	engine.POST("/api/login_sms", mc.smsLogin)
}

func (mc *MemberController) sendSmsCode(c *gin.Context) {
	//发送验证码

	phone, exist := c.GetQuery("phone")
	if !exist {
		tool.Failed(c, "解析失败")
		//c.JSON(200, map[string]interface{}{
		//	"code": 0,
		//	"msg":  "参数解析失败",
		//})
		return
	}

	ms := service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend {
		tool.Success(c, "发送成功")
		//c.JSON(200, map[string]interface{}{
		//	"code": 1,
		//	"msg":  "发送成功",
		//})
		return
	}
	tool.Success(c, "发送失败")
	//c.JSON(200, map[string]interface{}{
	//	"code": 0,
	//	"msg":  "发送失败",
	//})
}

func (mc *MemberController) smsLogin(ctx *gin.Context) {
	//解析结构体参数
	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(ctx.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Success(ctx, "参数解析失败")
		return
	}
	//完成登陆手机+验证码
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)

	if member != nil {
		tool.Success(ctx, member)
		return
	}

	tool.Failed(ctx, "登陆失败")
}
