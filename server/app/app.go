package app

import (
	"github.com/wonderivan/logger"
	"runtime/debug"
)

const pkgName = "app"

type Server interface {
	OnInit()
	Run()
	OnDestroy()
}

type AppServer struct {
	Conf 		Config
	Signal  	*Signal

	Servers 		[] Server
	InitServers 	[] Server
}

var app = new(AppServer)

func App() *AppServer {
	return app
}

func Init(conf Config){
	app.Conf = conf
	app.Signal = NewSignal()
}

func AddServer(server Server)  {
	app.Servers = append(app.Servers , server)
}

func Start(){
	//退出处理
	defer func() {
		if r := recover(); r != nil {
			logger.Error("%s: app will collapse for reason: %s", pkgName, r)
			logger.Error("%s: stack: %s", pkgName, debug.Stack())
		}
		for _, server := range app.InitServers {
			destroySafe(server)
		}
		logger.Info("all server stop successfully ...")
		//发送信号，所有服务都以结束。通知主线可以正常退出
		close(app.Signal.Stopped)
	}()

	for _ , s := range app.Servers{
		s.OnInit()
		app.InitServers = append(app.InitServers , s)
		go s.Run()
	}
	//等待结束服务控制器的信号，进行退出处理
	<- app.Signal.Stop
}

func destroySafe(server Server){
	defer func() {
		if r := recover(); r != nil {
			logger.Error("%s: destroy server failed: %s", pkgName, r)
		}
	}()
	server.OnDestroy()
}