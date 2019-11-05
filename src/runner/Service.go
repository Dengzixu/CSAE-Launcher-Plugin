package runner

import (
	"CSAE-Launcher-Plugin/src/common/utils"
	"CSAE-Launcher-Plugin/src/controller"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lxn/walk"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
	"net/http"
	"os"
)

const VERSION = "0.0.5"
const CHANNEL = "alpha"

func Service() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://hlds.zixutech.cn"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "UPDATE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: false,
		ExposeHeaders:    []string{"*"},
		MaxAge:           60 * 60 * 24 * 7, // 七天
	}))

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	router.GET("/about", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"version": VERSION,
			"channel": CHANNEL,
		})
	})

	router.POST("/launch", controller.LaunchController)

	router.GET("/choose", controller.ChooseFileController)

	//err := router.Run("127.0.0.1:23232")

	router.Use(tlsHandler())
	err := router.RunTLS("127.0.0.1:23232", utils.GetSSLDir()+"\\certificate.crt", utils.GetSSLDir()+"\\key.pem")

	if err != nil {
		fmt.Println(err)
		walk.MsgBox(nil, "CSAE Launcher Plugin", "错误: 初始化失败, 请检查是否运行了多个程序，如果无法解决，请联系开发人员。", walk.MsgBoxIconError)
		log.WithField("component", "Web").Error("初始化失败: 无法创建服务")
		os.Exit(2)
	}
}

func tlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "127.0.0.1:23232",
		})
		// If there was an error, do not continue.
		if err := secureMiddleware.Process(c.Writer, c.Request); err != nil {
			return
		}
		c.Next()
	}
}
