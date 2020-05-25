package controllers

import (
	// gin框架
	_ "github.com/gin-gonic/gin"
)

// 接口信息

//ExampleController 接口示例
type ExampleController struct{}

//CaptchaResult container控制器，接受GET和POST请求
type CaptchaResult struct{}

//IndexController 根目录控制器
type IndexController struct{}

//PostData DescribeCaptchaResult接口Post参数
type PostData struct {
	Action       string `json:"Action"`
	Version      string `json:"Version"`
	CaptchaType  int    `json:"CaptchaType"`
	Ticket       string `json:"Ticket"`
	UserIP       string `json:"UserIp"`
	Randstr      string `json:"Randstr"`
	CaptchaAppID int    `json:"CaptchaAppId"`
	AppSecretKey string `json:"AppSecretKey"`

	//公共参数
	Host              string `json:"captcha.tencentcloudapi.com"`
	Service           string `json:"captcha"`
	HTTPRequestMethod string `json:"POST"`
	XTCAction         string `json:"X-TC-Action"`
	XTCRegion         string `json:"X-TC-Region"`
	XTCTimestamp      int64  `json:"X-TC-Timestamp"`
	XTCVersion        string `json:"X-TC-Version"`
	Authorization     string `json:"Authorization"`
	XTCToken          string `json:"X-TC-Token"`
	PlayLoad          string `json:"PlayLoad"`
}

//PlayLoad 负荷内容
type PlayLoad struct {
	CaptchaType  int    `json:"CaptchaType"`
	Ticket       string `json:"Ticket"`
	UserIP       string `json:"UserIp"`
	Randstr      string `json:"Randstr"`
	CaptchaAppID int    `json:"CaptchaAppId"`
	AppSecretKey string `json:"AppSecretKey"`
}

//DescribeCaptchaResultResponse DescribeCaptchaResult接口的返回数据
type DescribeCaptchaResultResponse struct {
	Response struct {
		CaptchaCode int    `json:"CaptchaCode"`
		CaptchaMsg  string `json:"CaptchaMsg"`
	} `json:"Response"`
	RetCode int `json:"retcode"`
}
