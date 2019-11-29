package session

import "github.com/gorilla/websocket"

type Session struct {
	PlayerId 	string
	ServerId 	string
	Conn 		*websocket.Conn

	CacheResponseCount 	int32
}

// visitor actor模型 会话
type ControllerSession struct {
	Flag 		int32
	PlayerId 	string
	ServerId 	string
}

// room actor模型 会话
type RoomSession struct {
	Id			string
	PlayerId 	string
	ServerId 	string
}