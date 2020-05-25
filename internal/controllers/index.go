//Package controllers 接收请求并处理数据返回
/**
* @Author: Kanesong
* @Date: 2019/4/5 9:33 PM
 */
package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Welcome example
func (i *IndexController) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "Welcome.",
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})

	return
}

//Handle404 404错误处理器
func (i *IndexController) Handle404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":      404,
		"msg":       "Page is Not Found.",
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})
	return
}
