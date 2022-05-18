package util

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)

func SendCheckCodeMessage(mobile, checkcode string) error {

	url := "https://u.smsyun.cc/sms-partner/access/b031k3/sendsms"
	post := "{\"clientid\":\"" + "b031k3" +
		"\",\"password\":\"" + "25d55ad283aa400af464c76d713c07ad" +
		"\",\"mobile\":\"" + mobile +
		"\",\"smstype\":\"" + "4" +
		"\",\"content\":\"" + "【Filer】您的验证码是" + checkcode + "，如非本人操作，请忽略此条短信。" +
		"\"}"

	var jsonStr = []byte(post)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return errors.New("SendCheckCodeMessage err 001")
	}
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("SendCheckCodeMessage err 002")
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(gjson.Get(string(body), "data.0.msg"))
	if gjson.Get(string(body), "data.0.msg").String() != "成功" {
		return errors.New("SendCheckCodeMessage err 003")
	}
	return nil
}

///**
// * 使用AK&SK初始化账号Client
// * @param accessKeyId
// * @param accessKeySecret
// * @return Client
// * @throws Exception
// */
//func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
//	config := &openapi.Config{
//		// 您的AccessKey ID
//		AccessKeyId: accessKeyId,
//		// 您的AccessKey Secret
//		AccessKeySecret: accessKeySecret,
//	}
//	// 访问的域名
//	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
//	_result = &dysmsapi20170525.Client{}
//	_result, _err = dysmsapi20170525.NewClient(config)
//	return _result, _err
//}
//
//func _main (args []*string) (_err error) {
//	client, _err := CreateClient(tea.String("accessKeyId"), tea.String("accessKeySecret"))
//	if _err != nil {
//		return _err
//	}
//
//	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{}
//	// 复制代码运行请自行打印 API 的返回值
//	_, _err = client.SendSms(sendSmsRequest)
//	if _err != nil {
//		return _err
//	}
//	return _err
//}
