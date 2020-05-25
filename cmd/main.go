/**
* @Author: Kanesong
* @Date: 2019/4/5 9:15 PM
 */
package main //import git.code.oa.com/etanmiao/captchaDemo

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"git.code.oa.com/etanmiao/captchaDemo/configs"
	"git.code.oa.com/etanmiao/captchaDemo/internal/routes"
	"github.com/gin-gonic/gin"
)

var (
	v           bool
	versionPath = "configs/VERSION"
	logPath     = "logs/gin.log"
)

func init() {
	flag.BoolVar(&v, "v", false, "show version")
	flag.Usage = usage
}

func usage() {
	_, err := os.Stat(versionPath)
	if err != nil {
		versionPath = "../configs/VERSION"
	}
	version, err := ioutil.ReadFile("configs/VERSION")
	if err != nil {
		log.Println(err)
	}
	log.Print("Version:", string(version))
}

func main() {
	// 查看版本号
	flag.Parse()
	if v {
		flag.Usage()
		return
	}
	// init router
	router := routes.SetupRouter()

	// init config
	serverConfig := configs.LoadServerConfig()
	gin.SetMode(serverConfig.RunMode)

	// init log
	_, err := os.Stat(logPath)
	if err != nil {
		logPath = "../logs/gin.log"
	}
	f, _ := os.Create(logPath)
	gin.DefaultWriter = io.MultiWriter(f)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverConfig.HTTPPort),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
