package main

import (
	"chat_server/app"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/wonderivan/logger"
	"os"
	"os/signal"
	"syscall"
)

type Flags struct {
	ConfigPath string
	ServerName string
	ServerId   string
}

var (
	flags         Flags
	config        Config
)

//解析参数
func parseFlag() {
	flag.StringVar(&flags.ConfigPath, "config", "/Users/wangbowei/chat_server/doc/config/config.toml", "-config <config>")
	flag.StringVar(&flags.ServerId, "server-id", "", "-server-id <server-id>")
	flag.Parse()
}

func parseConfig(){
	if _, err := toml.DecodeFile(flags.ConfigPath,&config) ; err != nil {
		logger.Error("parse config failed : %s",err)
	}
}

func init() {
	parseFlag()
	parseConfig()
}

func AppConfig() (conf app.Config) {
	conf.Name = config.Server.Name
	conf.App.Environment = config.App.Environment
	return
}

func main() {
	app.Init(AppConfig())
	app.AddServer(&ChatServer{
		Conf:config,
		Address:config.Address.Port,
	})
	//DB
	app.AddServer(&LanderServer{
		ServerId:config.Server.Id,
		Url:config.Mongo.Url,
		Database:config.Mongo.Database,
	})

	go app.Start()
	StopAndWait()
}

func StopAndWait(){
	c := make(chan os.Signal , 1)
	signal.Notify(c , os.Interrupt ,os.Kill , syscall.SIGHUP ,syscall.SIGTERM)
	go func() {
		<- c
		close(app.App().Signal.Stop)
	}()

	<- app.App().Signal.Stopped
}