package processor

import (
	"chat_server/app/processor"
	"chat_server/chat/models/room"
	"chat_server/chat/session"
)

type Room struct {
	Stopped 	bool
	Model     	*room.Room
	Controller  *HandlesManage
	Processor 	*processor.Processor
}

func NewRoom(playerId , serverId ,roomName string) *Room {
	r := &Room{
		Model:room.NewRoom(roomName),
		Controller:NewHandlesManage(&session.RoomSession{
			PlayerId:playerId,
			ServerId:serverId,
		}),
		Processor:processor.NewProcessor("room"),
	}
	r.setRoute()
	return r
}

func (v *Room) Cast(pattern string , bytes []byte) {
	v.Processor.Receive(pattern, bytes)
}

func (v *Room) Run() {
	v.Processor.HandleLoop(func(pattern string, message []byte) {
		if v.Stopped {
			return
		}
		switch pattern {
		case "cast":
			v.Controller.Dispense(message)
		case "stop":
			v.Processor.Stop()
			v.Stopped = true
		}
	})
}

func (r *Room) setRoute() {
	route := r.Controller.Route
	route("EnterRoomCst" , r.EnterRoomCst)
	route("LeaveRoomCst" , r.LeaveRoomCst)
	route("ChatCst" , r.ChatCst)

}