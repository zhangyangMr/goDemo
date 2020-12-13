package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func main() {
	accessKeyId := "accessKeyId"
	accessSecret := "accessSecret"
	phoneNum := "phoneNum"
	//signName := "阿里云短信测试专用"
	signName := "阿里云短信测试专用"
	templateCode := "SMS_138530360"
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)
	if err != nil {
		fmt.Print(err.Error())
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNum
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = "{\"code\":" + "777777" + "}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
