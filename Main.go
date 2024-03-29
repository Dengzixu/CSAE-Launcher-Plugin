//go:generate goversioninfo
package main

import (
	"CSAE-Launcher-Plugin/src/common/Init"
	"CSAE-Launcher-Plugin/src/common/Logs"
	"CSAE-Launcher-Plugin/src/common/global"
	"CSAE-Launcher-Plugin/src/runner"
	"flag"
	"github.com/judwhite/go-svc/svc"
	"os"
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
		Logs.G.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	// 初始化日志
	Init.Logger()

	// 第一次启动初始化
	Init.First()

	// 检查更新
	Init.CheckUpdate()

	// 进入主程序代码段
	help := flag.Bool("h", true, "获取帮助")

	chooseFile := flag.Bool("c", false, "选择 CSAE 程序路径")
	setFile := flag.String("C", "", "设置 CSAE `程序路径`")

	runGame := flag.Bool("l", false, "运行离线游戏")

	portable := flag.Bool("p", false, "以 便携模式 运行")
	server := flag.Bool("s", false, "以 服务模式 运行")

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
	case *portable:
		global.IsService = false
		runner.Portable()
	case *server:
		global.IsService = env.IsWindowsService()
		if !env.IsWindowsService() {
			Logs.G.Error("This mode only run as Windows Service")
			os.Exit(2)
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
		Logs.G.Infow("Init Windows Service...")
		<-p.quit
		Logs.G.Infow("Quit signal received...")
		p.wg.Done()
	}()

	return nil
}

func (p *program) Stop() error {
	// The Stop method is invoked by stopping the Windows service, or by pressing Ctrl+C on the console.
	// This method may block, but it's a good idea to finish quickly or your process may be killed by
	// Windows during a shutdown/reboot. As a general rule you shouldn't rely on graceful shutdown.
	Logs.G.Infow("Stopping...")
	close(p.quit)
	p.wg.Wait()
	Logs.G.Infow("Stopped.")
	return nil
}
