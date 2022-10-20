package service

import (
	"fmt"
	"ginapp/dao"
	"ginapp/model"
	"ginapp/param"
	"ginapp/tool"
	"math/rand"
	"time"
)

//短信 发送的服务
type MemberService struct {
}

//func (ms *MemberService) GetUserInfo(userId string) *model.Member {
//	id, err := strconv.Atoi(userId)
//	if err != nil {
//		return nil
//	}
//	memberDao := dao.MemberDao{tool.DbEngine}
//	return memberDao.QueryMemberById(id)
//}

// 发送短信验证码
func (ms *MemberService) SendCode(phone string) bool {

	//1.产生一个验证码
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	fmt.Printf("当前手机号%v 发送的短信验证码是:%v\n", phone, code)

	//2.调用阿里云sdk 完成发送
	// config := tool.GetConfig().Sms
	// client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	return false
	// }

	// request := dysmsapi.CreateSendSmsRequest()
	// request.Scheme = "https"
	// request.SignName = config.SignName
	// request.TemplateCode = config.TemplateCode
	// request.PhoneNumbers = phone
	// par, err := json.Marshal(map[string]interface{}{
	// 	"code": code,
	// })
	// request.TemplateParam = string(par)

	// response, err := client.SendSms(request)
	// fmt.Println(response)
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	return false
	// }

	//3.接收返回结果，并判断发送状态
	//短信验证码发送成功
	// if response.Code == "OK" {
	// 	return true
	// }
	smsCode := model.SmsCode{Phone: phone, Code: code, BizId: "121212", CreateTime: time.Now().Unix()}
	//return true
	memberDao := dao.MemberDao{tool.DbEngine}
	fmt.Println("操作开始", smsCode)
	result := memberDao.InsertCode(smsCode)
	fmt.Println("操作成功")
	return result > 0

}

//用户手机号+验证码的登录
func (ms *MemberService) SmsLogin(loginparam param.SmsLoginParam) *model.Member {
	//1.获取到手机号和验证码
	// fmt.Printf("传递的参数是%v\n", loginparam)
	//2.验证手机号+验证码是否正确
	md := dao.MemberDao{tool.DbEngine}
	sms := md.ValidateSmsCode(loginparam.Phone, loginparam.Code)
	if sms.Id == 0 {
		return nil
	}

	//3、根据手机号member表中查询记录
	member := md.QueryByPhone(loginparam.Phone)
	if member.Id != 0 {
		return member
	}

	//4.新创建一个member记录，并保存
	user := model.Member{}
	user.UserName = loginparam.Phone
	user.Mobile = loginparam.Phone
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)

	return &user

}
