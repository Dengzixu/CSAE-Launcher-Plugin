//go:generate goversioninfo
package main

import (
	"CSAELauncherPlugin/common/Init"
	"CSAELauncherPlugin/common/global"
	"CSAELauncherPlugin/runner"
	"flag"
	"fmt"
	"github.com/judwhite/go-svc/svc"
	log "github.com/sirupsen/logrus"
	"sync"
)

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
	// 初始化日志
	Init.Logger()

	// 第一次启动初始化
	Init.First()

	// 检查更新
	Init.CheckUpdate()

	// 初始化配置文件
	Init.Config()

	// 进入主程序代码段
	help := flag.Bool("h", true, "获取帮助")

	chooseFile := flag.Bool("c", false, "选择 CSAE 程序路径")
	setFile := flag.String("C", "", "设置 CSAE `程序路径`")

	runGame := flag.Bool("l", false, "运行离线游戏")
	server := flag.Bool("s", false, "运行服务")

	flag.Parse()

	flag.Usage = runner.Default

	switch {
	case *chooseFile:
		runner.ChooseFile()
		break
	case "" != *setFile:
		runner.UnSupport()
		break
	case *runGame:
		runner.UnSupport()
		break
	case *server:
		fmt.Println("注意: 此窗口非常重要，请不要关闭。")
		global.IsService = env.IsWindowsService()
		if !env.IsWindowsService() {
			log.Warn("提示: 如果您 不是 网吧用户，建议您通过安装包进行安装。")
		}
		go runner.Service()
		break
	case *help:
		runner.Default()
	}

	return nil
}

func (p *program) Start() error {
	// The Start method must not block, or Windows may assume your service failed
	// to start. Launch a Goroutine here to do something interesting/blocking.
	p.quit = make(chan struct{})

	p.wg.Add(1)
	go func() {
		log.WithField("component", "Service Control").Info("正在启动服务...")
		<-p.quit
		log.WithField("component", "Service Control").Info("收到退出信号...")
		p.wg.Done()
	}()

	return nil
}

func (p *program) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.
	log.WithField("component", "Service Control").Info("正在停止服务...")
	close(p.quit)
	p.wg.Wait()
	log.WithField("component", "Service Control").Info("服务停止.")
	return nil
}
