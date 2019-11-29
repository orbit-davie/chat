package main

import (
	"chat_server/chat/cache"
	"chat_server/chat/processor"
	"chat_server/chat/session"
	"github.com/gorilla/websocket"
	"github.com/wonderivan/logger"
	"net/http"
	"runtime/debug"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatServer struct {
	Conf 		Config
	Address 	string

	wg 			*sync.WaitGroup
}

func (c *ChatServer) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	playerId := r.Header.Get("PlayerId")
	if playerId == "" {
		return
	}

	conn ,err := upgrader.Upgrade(w,r, nil )
	if err != nil {
		logger.Error("websocket handle failed : %s" ,err)
		return
	}

	session := &session.Session{
		PlayerId:playerId,
		ServerId:c.Conf.Server.Id,
		Conn:conn,
		CacheResponseCount: int32(0),
	}
	v := NewVisitor(playerId,c.Conf.Server.Id , session)
	b := processor.NewButler(playerId ,c.Conf.Server.Id)
	//重复登录
	if old := cache.Find(playerId) ; old != nil {
		logger.Error("playerId is exist : %s",playerId)
		old.Cast("stop",nil)
		cache.UnRegistry(playerId)
	}

	cache.Registry(playerId,v)
	processor.Registry(playerId , b)

	go c.RunVisitorSafe(v)
	go c.RunButlerSafe(playerId , b)
}

func (c *ChatServer) RunVisitorSafe(v *Visitor) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("panic: %s", r)
			logger.Error("stack: %s", debug.Stack())
		}
	}()
	c.wg.Add(1)
	go v.receiveLoop()
	v.responseLoop()

	cache.UnRegistry(v.Session.PlayerId)

	c.wg.Done()
}

func (c *ChatServer) RunButlerSafe(playerId string ,b *processor.Butler) {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("panic: %s", r)
			logger.Error("stack: %s", debug.Stack())
		}
	}()

	b.Run()
	processor.UnRegistry(playerId)
}

func (c *ChatServer) OnInit() {
	c.wg =new(sync.WaitGroup)

	http.HandleFunc("ws" , c.handleWebsocket)
}

func (c *ChatServer) Run() {
	defer func() {
		if r := recover(); r != nil {
			logger.Error("panic: %s", r)
			logger.Error("stack: %s", debug.Stack())
		}
	}()

	err := http.ListenAndServe(c.Address , nil)
	if err != nil {
		logger.Error("http listen failed : %s",err)
	}
}

func (c *ChatServer) OnDestroy() {
	for _ , v := range cache.AllVisitors() {
		v.Cast("stop" , nil)
	}

	for _ , v := range processor.AllButlers() {
		v.Cast("stop" , nil)
	}

	for _  ,v := range processor.RoomsAll() {
		v.Cast("stop" , nil)
	}
	c.wg.Wait()
}