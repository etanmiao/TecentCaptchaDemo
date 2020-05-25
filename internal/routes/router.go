/**
* @Author: Kanesong
* @Date: 2019/4/5 9:16 PM
 */
package routes

import (
	"git.code.oa.com/etanmiao/captchaDemo/internal/controllers"
	"github.com/gin-gonic/gin"
)

var indexCtl = new(controllers.IndexController)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	{
		indexCtl := new(controllers.IndexController)
		router.GET("/", indexCtl.Welcome)
		router.NoRoute(indexCtl.Handle404)
	}

	// new api
	exampleRouter := router.Group("/example")
	exampleCtl := new(controllers.ExampleController)
	{
		exampleRouter.GET("/hello", exampleCtl.Hello)
	}

	//container api
	captchaRouter := router.Group("/captcha")
	captchaCtl := new(controllers.CaptchaResult)
	{
		captchaRouter.GET("/authresult", captchaCtl.RequestCaptchaResult)
	}
	return router
}
