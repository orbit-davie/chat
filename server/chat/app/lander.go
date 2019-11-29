package main

import (
	"chat_server/app/lander"
	"chat_server/app/mongo"
	"github.com/wonderivan/logger"
)

type LanderServer struct {
	ServerId 	string
	Url 		string	//url
	Database 	string //db name
}

func (l *LanderServer) OnInit() {
	lander.SetLander(mongo.NewMongoManger(mongo.Config{ServerId:l.ServerId},mongo.MongoCfg{Url:l.Url,Database:l.Database}))
}

func (l *LanderServer) Run() {
	lander.GetLander().Run()
}

func (l *LanderServer) OnDestroy() {
	lander.GetLander().StopSafe()
	logger.Info("lander stopped ...")
}