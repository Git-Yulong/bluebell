/**
  @author: ZYL
  @date:
  @note
*/
package utils

// This file is auto-generated, don't edit it. Thanks.

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

//func main() {
//	err := _main(tea.StringSlice(os.Args[1:]))
//	if err != nil {
//		panic(err)
//	}
//}

// SendMsg 发送验验证码
// tel 手机号 string类型
// code 验证码 string类型
func IsMobile(mobile string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)

}

func SendMsg(tel string, code string) (error, string) {
	client, err := CreateClient(tea.String("LTAI5tARRzaXyK9QcxkNXYBy"), tea.String("f38T0a5SUsITB18cFQKHJLKkJs7llz"))
	if err != nil {
		zap.L().Error("初始化阿里云客户端失败")
		return err, "failed"
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("阿里云短信测试"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(tel),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	// 复制代码运行请自行打印 API 的返回值
	_, err = client.SendSms(sendSmsRequest)
	if err != nil {
		return err, "failed"
	}

	return nil, "success"
}

// Code 随机生成一个6位数的验证码
func TelCheckCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code) //转字符串返回
	return res
}
