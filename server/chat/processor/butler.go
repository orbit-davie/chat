package processor

import (
	"chat_server/app/processor"
	"chat_server/chat/models/butler"
	"chat_server/chat/session"
)

type Butler struct {
	Stopped 	bool
	Model 		*butler.Butler
	Controller  *HandlesManage
	Processor 	*processor.Processor
}

func NewButler(playerId , serverId string) *Butler {
	b := &Butler{
		Model:butler.NewButler(playerId),
		Controller:NewHandlesManage(&session.RoomSession{
			PlayerId:playerId,
			ServerId:serverId,
		}),
		Processor:processor.NewProcessor("butler"),
	}

	b.setRoute()
	return b
}

func (b *Butler) Cast(pattern string , bytes []byte) {
	b.Processor.Receive(pattern, bytes)
}

func (b *Butler) Run() {
	b.Processor.HandleLoop(func(pattern string, message []byte) {
		if b.Stopped {
			return
		}
		switch pattern {
		case "cast":
			b.Controller.Dispense(message)
		case "stop":
			b.Processor.Stop()
			b.Stopped = true
		}
	})
}

func (b *Butler) setRoute() {
	route := b.Controller.Route
	route("CreateRoomCst" , b.CreateRoomCst)
	route("EnterRoomCst" , b.EnterRoomCst)
	route("ChatCst" , b.ChatCst)

	route("EnterRoomCstResponse" , b.EnterRoomCstResponse)
}