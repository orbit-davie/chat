package processor

import (
	"chat_server/app/controller"
	"chat_server/chat/cache"
	"chat_server/chat/models/visitor"
	"chat_server/chat/protocol"
	"chat_server/chat/session"
	"github.com/golang/protobuf/proto"
	"github.com/wonderivan/logger"
)

type HandlesManage struct {
	Controller 	*controller.Controller
	Session 	*session.RoomSession
}

func NewHandlesManage(session *session.RoomSession) *HandlesManage {
	return &HandlesManage{
		Controller:controller.NewController(),
		Session:session,
	}
}

func(h *HandlesManage) Dispense(data []byte) {
	h.Controller.Dispense(data)
}

func(h *HandlesManage) Route(handleName string , handle func(flag int32 , data []byte)){
	if _ , ok := h.Controller.Handles[handleName] ; ok {
		logger.Error("handle name has registered...")
		return
	}
	h.Controller.Handles[handleName] = handle
}

func(h *HandlesManage) reply(playerId string , flag int32 , ackName string , pb proto.Message) {
	data , err := h.Controller.PackMessage(flag, ackName , pb)
	if err != nil {
		logger.Error("pack ack message failed : %s",err)
		return
	}

	visitor := cache.Find(playerId)
	if visitor == nil {
		logger.Error("can not found visitor")
		return
	}
	visitor.Cast("reply",data)
}

func(h *HandlesManage) Reply(playerId string , flag int32 , ackName string , pb proto.Message) {
	if playerId == "" {
		return
	}

	h.reply(playerId , flag , ackName , pb)
}

func(h *HandlesManage) ReplyErr(playerId string , flag int32 , ackName string , err error) {
	pb := &protocol.ErrAck{
		AckName:proto.String(ackName),
		Err:proto.String(err.Error()),
	}
	data , err := h.Controller.PackMessage(flag, "ErrorAck" , pb)
	if err != nil {
		logger.Error("pack ack message failed : %s",err)
		return
	}

	visitor := cache.Find(playerId)
	if visitor == nil {
		logger.Error("can not found visitor")
		return
	}
	visitor.Cast("reply",data)
}

func(h *HandlesManage) FindDo(playerId string , handle func(visitor *visitor.Visitor)) {
	if v := visitor.Find(playerId) ; v != nil {
		handle(v)
	}
}

func(h *HandlesManage) Unmarshal(data []byte , pb proto.Message) error {
	if err := proto.Unmarshal(data , pb) ; err != nil {
		return err
	}
	return nil
}

func(h *HandlesManage) Marshal(pb proto.Message) ([]byte , error) {
	return proto.Marshal(pb)
}

func(h *HandlesManage) CstButler(playerId string , flag int32 , cstName string , pb proto.Message) {
	data , err := h.Controller.PackMessage(flag, cstName , pb)
	if err != nil {
		logger.Error("pack ack message failed : %s",err)
		return
	}

	b := FindButler(playerId)
	if b == nil {
		logger.Error("can not find butler")
		return
	}

	b.Cast("cst",data)
}

func(h *HandlesManage) CstRoom(roomName string , flag int32 , cstName string , pb proto.Message) {
	data , err := h.Controller.PackMessage(flag, cstName , pb)
	if err != nil {
		logger.Error("pack ack message failed : %s",err)
		return
	}

	r := FindRoom(roomName)
	if r == nil {
		logger.Error("can not find room")
		return
	}

	r.Cast("cst",data)
}