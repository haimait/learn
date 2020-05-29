package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"oneNet/lib/oneNet"
)

func main() {
	DeviceAdd()
}

//Test Pass
func DeviceAdd() {
	on := oneNet.NewOneNet("IKiCAxceXjJqpXwoTlBTfivq8ibwISuY8LUTlpL3DG8=")
	device := make(map[string]interface{})
	//device["title"] = "my test device"
	//device["private"] = true
	deviceKey := on.GetApiKey()

	device["name"] = "test555555"
	device["desc"] = "some descriptiontest00002test0000277777"
	device["key"] = deviceKey
	fmt.Println("device:11111", device)

	//on.SetApiKey("PvJ6UuQDb+c9j8SY3gsqUFIjoEuXUHlqOzudgQuRN2s=")
	//ret, s := on.DeviceAdd2(device)
	ret := on.DeviceAdd2(device)
	//ret, s := on.Device(599079417)
	fmt.Println("ret:", ret)
	//fmt.Println("s:", *s)



	//if ret == true {
	//	t.Log(ret)
	//	t.Log(*s)
	//} else {
	//	t.Error(ret)
	//	if s != nil {
	//		t.Error(*s)
	//	} else {
	//		t.Error(on.GetErrorNo())
	//		t.Error(on.GetError())
	//
	//	}
	//}
	//on.SetApiKey(deviceKey)

}

//func GetToken(){
//	version := "2018-10-31"
//	id:=1234
//	res := fmt.Sprintf("products/%d",id)  // 通过产品ID访问产品API
//	//# 用户自定义token过期时间
//	et := strconv.Itoa(int(time.Now().Add(time.Millisecond *3600*12).Unix()))
//	//# 签名方法，支持md5、sha1、sha256
//	method := "sha1"
//	//# 对access_key进行decode
//	accessKey := "PvJ6UuQDb+c9j8SY3gsqUFIjoEuXUHlqOzudgQuRN2s="
//	key := base64DecodeStr(accessKey)
//	//# 计算sign
//	org := et + "\n" + method + "\n" + res + "\n" + version
//	//sign_b = hmac.new(key=key, msg=org.encode(), digestmod=method)
//	bStr,_:=json.Marshal(org)
//	sign_b = hmacSha1(string(bStr),key)
//	sign = base64.b64encode(sign_b.digest()).decode()
//	sign = base64EncodeStr(sign_b.digest()).decode()
//
//	//# value 部分进行url编码，method/res/version值较为简单无需编码
//	sign = quote(sign, safe='')
//	res = quote(res, safe='')
//
//	//# token参数拼接
//	token = 'version=%s&res=%s&et=%s&method=%s&sign=%s' % (version, res, et, method, sign)
//	return token
//}
func hmacSha1(data string, secret string) string {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//base解码
func base64DecodeStr(src string) string {
	a, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "error"
	}
	return string(a)
}

//base编码
func base64EncodeStr(src string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(src)))
}
