//Package controllers 接收请求并处理数据返回
/**
* @Author: Kanesong
* @Date: 2019/4/9 3:17 PM
 */
package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Hello example
func (i *ExampleController) Hello(c *gin.Context) {
	returnData := "This is an example"
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "Welcome.",
		"data":      returnData,
		"timestamp": time.Now().Unix(),
	})
	return
}

//RequestCaptchaResult 获取容器信息，根据POST或GET请求，查询数据库返回数据
func (i *CaptchaResult) RequestCaptchaResult(c *gin.Context) {
	var ticket string
	var randstr string
	if c.Request.Method == "GET" {
		log.Println(c.Params)
		ticket = c.DefaultQuery("ticket", "")
		randstr = c.DefaultQuery("randstr", "")
	} else {
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		body := buf[0:n]
		var playLoad PlayLoad
		err := json.Unmarshal(body, &playLoad)
		if err != nil {
			log.Println("params analysis failed")
		} else {
			ticket = playLoad.Ticket
			randstr = playLoad.Randstr

		}
	}
	code, err := GetCaptchaResult(ticket, randstr)
	msg := "success"
	if code == 0 {
		msg = "请求失败"
	} else if code == 1 {
		msg = "验证成功"
	} else {
		msg = err.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"message":     msg,
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
		"CaptchaCode": code,
	})
	return
}
