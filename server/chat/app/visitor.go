package main

import (
	"chat_server/app/processor"
	visitor_handle"chat_server/chat/handle/visitor"
	"chat_server/chat/session"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"time"
)

type Visitor struct {
	Stopped 	bool
	Session 	*session.Session		//websocket session
	Controller  *visitor_handle.HandlesManage
	Processor 	*processor.Processor
}

func NewVisitor(playerId , serverId string , sess *session.Session) *Visitor {
	return &Visitor{
		Session:sess,
		Controller:visitor_handle.NewHandlesManage(&session.ControllerSession{
			PlayerId:playerId,
			ServerId:serverId,
		}),
		Processor:processor.NewProcessor("visitor"),
	}
}

func (v *Visitor) Cast(pattern string , bytes []byte) {
	v.Processor.Receive(pattern, bytes)
}

const DeadTime = 10 * time.Second

func (v *Visitor) receiveLoop() {
	for  {
		v.Session.Conn.SetReadDeadline(time.Now().Add(DeadTime))
		messageType, message , err := v.Session.Conn.ReadMessage()
		if len(message) == 0 || err == io.EOF || err != nil{
			break
		}
		if messageType != websocket.BinaryMessage {
			fmt.Errorf("%s: message type `%v` invalid", "connection", messageType)
			continue
		}

		v.Controller.Dispense(message)
	}
}

func (v *Visitor) responseLoop() {
	v.Processor.HandleLoop(func(pattern string, message []byte) {
		if v.Stopped {
			return
		}
		switch pattern {
		case "reply":
			if err := v.Session.Conn.WriteMessage(websocket.BinaryMessage , message) ; err != nil {

			}
		case "stop":
			v.Processor.Stop()
			v.Session.Conn.Close()
			v.Stopped = true
		}
	})
}