package util

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	imgUrl          = "https://investors.oss-cn-beijing.aliyuncs.com"
	endpoint        = "http://oss-cn-beijing.aliyuncs.com"
	accessKeyID     = "LTAI5tLvK7ZSuCXHDKQAJZWp"
	accessKeySecret = "XFX63hmUufFXKLUqydk3vJaT9p4oqc"
	bucketName      = "investors"
	fileFolder      = "test"
)

//oss添加文件
func setFileOss(base64str string) string {
	base64str = strings.Replace(base64str, "data:image/png;base64,", "", 1)
	byte1, err := base64.StdEncoding.DecodeString(base64str) //成图片文件并把文件写入到buffe
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		return ""
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return ""
	}
	t := time.Now().Nanosecond()
	s2 := rand.NewSource(int64(t))
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(10000000000))
	x1 := strconv.Itoa(t)
	x2 := strconv.Itoa(r2.Intn(10000000))
	fmt.Println(x1 + x2)
	// 指定访问权限为公共读，缺省为继承bucket的权限。
	//objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	// 上传字符串。
	err = bucket.PutObject(fileFolder+"/"+x1+x2+".png", bytes.NewReader(byte1))
	if err != nil {
		return ""
	}
	fmt.Println(imgUrl + "/" + fileFolder + "/" + x1 + x2 + ".png")
	return imgUrl + "/" + fileFolder + "/" + x1 + x2 + ".png"
}
