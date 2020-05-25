//Package controllers 接收请求并处理数据返回
/**
* @Author: shanizeng
* @Date: 2019/4/17 4:45 PM
 */
package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	// 导入数据库
	"encoding/json"

	"git.code.oa.com/etanmiao/captchaDemo/configs"
)

//SignatureGeneration 生成签名函数
func SignatureGeneration(ticket string, randstr string) (*PostData, error) {
	var PostParam PostData
	PostParam.Ticket = ticket
	PostParam.Randstr = randstr
	Cfg := configs.LoadDBConfig("secret")
	PostParam.Action = "DescribeCaptchaResult"
	PostParam.Version = "2019-07-22"
	PostParam.CaptchaType = 9
	PostParam.AppSecretKey = Cfg.Key("APPSECRETKEY").MustString("")
	PostParam.CaptchaAppID = Cfg.Key("APPID").MustInt()

	//公共参数
	_ = Cfg.Key("SECRETKEY")
	PostParam.Host = "captcha.tencentcloudapi.com"
	PostParam.XTCAction = "DescribeCaptchaResult"
	PostParam.Service = "captcha"
	PostParam.XTCRegion = "ap-guangzhou"
	PostParam.XTCTimestamp = time.Now().Unix()
	PostParam.XTCVersion = "2019-07-22"

	//step1 build canonical request string
	PostParam.HTTPRequestMethod = "POST"
	canonicalURI := "/"
	// 默认POST请求的话为空
	canonicalQueryString := ""
	canonicalHeaders := "content-type:application/json\n" + "host:" + PostParam.Host + "\n"
	signedHeaders := "content-type;host"
	fmt.Println(ticket)
	var playLoadIns PlayLoad
	playLoadIns = PlayLoad{CaptchaType: 9, Ticket: ticket, UserIP: "1.1.1.1", Randstr: randstr,
		CaptchaAppID: PostParam.CaptchaAppID, AppSecretKey: PostParam.AppSecretKey}

	playload, err := json.Marshal(playLoadIns)
	if err != nil {
		log.Fatalf("playload json marshaling failed:%s", err)
	}
	fmt.Println(string(playload))
	hashedRequestPayload := sha256hex(string(playload))
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		PostParam.HTTPRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	fmt.Println(canonicalRequest)
	fmt.Println("canonicalRequest")

	// step 2: build string to sign
	algorithm := "TC3-HMAC-SHA256"
	date := time.Unix(PostParam.XTCTimestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, PostParam.Service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		PostParam.XTCTimestamp,
		credentialScope,
		hashedCanonicalRequest)
	fmt.Println(string2sign)

	// step 3: sign string
	secretKey := Cfg.Key("SECRETKEY").MustString("")
	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(PostParam.Service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	fmt.Println(signature)

	// step 4: build authorization
	secretID := Cfg.Key("SECRETID").MustString("")
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretID,
		credentialScope,
		signedHeaders,
		signature)
	curl := fmt.Sprintf(`curl -X POST https://%s\
		-H "Authorization: %s"\
		 -H "Content-Type: application/json"\
		 -H "Host: %s" -H "X-TC-Action: %s"\
		 -H "X-TC-Timestamp: %d"\
		 -H "X-TC-Version: %s"\
		 -H "X-TC-Region: %s"\
		 -d '%s'`, PostParam.Host, authorization, PostParam.Host, PostParam.Action, PostParam.
		XTCTimestamp, PostParam.Version, PostParam.XTCRegion, playload)
	fmt.Println(curl)
	PostParam.Authorization = authorization
	PostParam.PlayLoad = string(playload)
	return &PostParam, nil
}

//GetCaptchaResult 获取签名结果，返回给前端
func GetCaptchaResult(ticket string, randstr string) (int, error) {
	postParam, err := SignatureGeneration(ticket, randstr)
	if err != nil {
		log.Println("签名生成失败")
		return 0, nil
	}
	client := &http.Client{Transport: &http.Transport{Proxy: SetProxy}}
	req, err := http.NewRequest("POST", "https://"+postParam.Host, strings.NewReader(postParam.PlayLoad))
	if err != nil {
		log.Panicln("req build error")
		return 0, nil
	}
	req.Header.Set("X-TC-Action", postParam.Action)
	req.Header.Set("X-TC-Timestamp", strconv.FormatInt(postParam.XTCTimestamp, 10))
	req.Header.Set("X-TC-Version", postParam.Version)
	req.Header.Set("X-TC-Region", postParam.XTCRegion)
	req.Header.Set("Authorization", postParam.Authorization)
	req.Header.Set("Content-Type", "application/json")
	log.Println(req.Header["Authorization"][0])

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return 0, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln("接口访问失败")
		return 0, nil
	}
	fmt.Println(string(body))
	var resultResponse DescribeCaptchaResultResponse
	err = json.Unmarshal(body, &resultResponse)
	if err != nil {
		log.Println("解析返回结果失败")
		return 0, nil
	}
	err = errors.New(resultResponse.Response.CaptchaMsg)
	log.Println(resultResponse.Response.CaptchaCode, err.Error())
	return resultResponse.Response.CaptchaCode, err
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}
