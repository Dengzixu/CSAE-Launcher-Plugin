package runner

import (
	"CSAE-Launcher-Plugin/src/common/Logs"
	"CSAE-Launcher-Plugin/src/common/utils"
	"CSAE-Launcher-Plugin/src/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go.uber.org/zap"
	"net/http"
	"os"
)

const VERSION = "0.0.11"
const CHANNEL = "beta"

func Service() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://hlds.zixutech.cn"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
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

	//err := router.Run("127.0.0.1:23232")

	router.Use(tlsHandler())
	err := router.RunTLS("127.0.0.1:23232", utils.GetSSLDir()+"\\certificate.crt", utils.GetSSLDir()+"\\key.pem")

	if err != nil {
		Logs.G.Error("Init web server failed", zap.Error(err))
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
