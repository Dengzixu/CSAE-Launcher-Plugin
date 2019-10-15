//go:generate goversioninfo
package main

import (
	"CSAELauncherPlugin/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/judwhite/go-svc/svc"
	"github.com/lxn/walk"
	"log"
	"net/http"
	"sync"
)

const VERSION = "0.0.1"
const CHANNEL = "alpha"

type program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

func main() {
	prg := &program{}

	// Call svc.Run to start your program/service.
	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {

	fmt.Println("注意: 此窗口非常重要，请不要关闭。")

	if !env.IsWindowsService() {
		fmt.Println("提示: 如果您 *不是* 网吧用户，建议您通过安装包进行安装。")
	}
	go RunServer()
	return nil
}

func (p *program) Start() error {
	// The Start method must not block, or Windows may assume your service failed
	// to start. Launch a Goroutine here to do something interesting/blocking.
	p.quit = make(chan struct{})

	p.wg.Add(1)
	go func() {
		log.Println("Starting...")
		<-p.quit
		log.Println("Quit signal received...")
		p.wg.Done()
	}()

	return nil
}

func (p *program) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.
	log.Println("Stopping...")
	close(p.quit)
	p.wg.Wait()
	log.Println("Stopped.")
	return nil
}

func RunServer() {
	router := gin.Default()

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

	err := router.Run("127.0.0.1:23232")

	if err != nil {
		walk.MsgBox(nil, "CSAE Launcher Plugin", "错误: 初始化失败, 请检查是否同时运行了个本程序，如果无法解决，请联系开发人员。", walk.MsgBoxIconError)
		fmt.Println("错误: 初始化失败, 请检查是否同时运行了多个本程序，如果无法解决，请联系开发人员。")
	}
}
