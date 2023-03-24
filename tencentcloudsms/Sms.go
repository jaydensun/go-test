package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func netOk() bool {
	resp, err := http.Get("https://www.baidu.com")
	if err == nil {
		return resp.StatusCode == 200
	}
	return false
}

func main() {
	var phone, secretId, secretKey, signName, smsSdkAppId, templateId string
	flag.StringVar(&phone, "phone", "", "phone number")
	flag.StringVar(&secretId, "secretId", "", "secretId")
	flag.StringVar(&secretKey, "secretKey", "", "secretKey")
	flag.StringVar(&signName, "signName", "", "signName")
	flag.StringVar(&smsSdkAppId, "smsSdkAppId", "", "smsSdkAppId")
	flag.StringVar(&templateId, "templateId", "", "templateId")
	flag.Parse()

	if len(phone) == 0 {
		panic("phone list is empty!!!")
	}

	fmt.Printf("%s, %s, %s, %s, %s, %s\n", phone, secretId, secretKey, signName, smsSdkAppId, templateId)
	phones := strings.Split(phone, ",")
	fmt.Println(phones)
	if len(phones) == 0 {
		panic("phone list is empty!!!")
	}

	for {
		if netOk() {
			fmt.Println("网络连通: ", nowDate())
			sendSms(phones, secretId, secretKey, signName, smsSdkAppId, templateId)
		} else {
			fmt.Println("网络不连通: ", nowDate())
		}
		time.Sleep(time.Minute * 60)
	}

	//sendSms()
}

func sendSms(phones []string, secretId string, secretKey string, signName string, smsSdkAppId string, templateId string) {
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(secretId, secretKey)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.ap-guangzhou.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(phones)
	request.SignName = common.StringPtr(signName)
	request.SmsSdkAppId = common.StringPtr(smsSdkAppId)
	request.TemplateId = common.StringPtr(templateId)
	request.TemplateParamSet = common.StringPtrs([]string{"远程网络已经连通", "请确认是否正常(" + nowDate() + ")。"})

	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	_, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	// fmt.Printf("%s", response.ToJsonString())
}

func nowDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
