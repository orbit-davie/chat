package processor

import (
	"chat_server/chat/models/butler"
	"chat_server/chat/protocol"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/wonderivan/logger"
)

func(h *Butler) CreateRoomCst(flag int32 , data []byte) {
	if !(h.Model.RoomId == "" && h.Model.Status == butler.OUT_ROOM) {
		h.Controller.ReplyErr(h.Model.PlayerId , flag , "CreateRoomAck" , errors.New("您现在在房间中，请先退出房间"))
		return
	}
	cst := &protocol.CreateRoomCst{}
	if err := h.Controller.Unmarshal(data , cst) ; err != nil {
		h.Controller.ReplyErr(h.Model.PlayerId , flag , "CreateRoomAck" , errors.New("unmarshal failed"))
		return
	}

	room := NewRoom(cst.GetPlayerId() , h.Controller.Session.ServerId ,cst.GetRoomName())
	go room.Run()
	RoomsRegistry(room.Model.Name, room)

	h.Model.EnterRoom(room.Model.Id,room.Model.Name)
	cstRoom := &protocol.EnterRoomCst{
		RoomName:proto.String(room.Model.Name),
		PlayerId:proto.String(cst.GetPlayerId()),
	}

	data , err := h.Controller.Controller.PackCstMessage(flag, "EnterRoomCst" , cstRoom)
	if err != nil {
		logger.Error("pack ack message failed : %s",err)
		return
	}

	room.Cast("cast",data)
}

func(h *Butler) EnterRoomCst(flag int32 , data []byte) {
	if !(h.Model.RoomId == "" && h.Model.Status == butler.OUT_ROOM) {
		h.Controller.ReplyErr(h.Model.PlayerId , flag , "EnterRoomAck" , errors.New("您现在在房间中，请先退出房间"))
		return
	}

	cst := &protocol.EnterRoomCst{}
	if err := h.Controller.Unmarshal(data , cst) ; err != nil {
		h.Controller.ReplyErr(h.Model.PlayerId , flag , "EnterRoomAck" , errors.New("unmarshal failed"))
		return
	}

	cstRoom := &protocol.EnterRoomCst{
		RoomName:proto.String(cst.GetRoomName()),
		PlayerId:proto.String(cst.GetPlayerId()),
	}

	h.Controller.CstRoom(cst.GetRoomName() , flag , "EnterRoomCst" , cstRoom)
}

func(h *Butler) ChatCst(flag int32 , data []byte) {
	if h.Model.RoomId == "" || h.Model.Status != butler.IN_ROOM {
		h.Controller.ReplyErr(h.Model.PlayerId , flag , "ChatAck" , errors.New("您现在在房间中，请先退出房间"))
		return
	}

	cst := &protocol.ChatCst{}
	if err := h.Controller.Unmarshal(data , cst) ; err != nil {
		h.Controller.ReplyErr(h.Model.PlayerId , flag , "ChatAck" , errors.New("unmarshal failed"))
		return
	}

	cstRoom := &protocol.ChatCst{
		RoomName:cst.RoomName,
		Data:cst.Data,
		PlayerId:cst.PlayerId,
	}
	h.Controller.CstRoom(cst.GetRoomName() , flag , "ChatCst" , cstRoom)
}

func(h *Butler) EnterRoomCstResponse(flag int32 , data []byte) {
	cstResponse := &protocol.EnterRoomCstResponse{}
	if err := h.Controller.Unmarshal(data , cstResponse) ; err != nil {
		h.Controller.ReplyErr(h.Model.PlayerId , flag , "EnterRoomAck" , errors.New("unmarshal failed"))
		return
	}

	if cstResponse.GetResult() {
		h.Model.EnterRoom(cstResponse.GetRoomId() , cstResponse.GetRoomName())
	}

	ack := &protocol.EnterRoomAck{
		RoomId:cstResponse.RoomId,
		RoomName:cstResponse.RoomName,
		Result:cstResponse.Result,
	}
	h.Controller.Reply(h.Model.PlayerId , flag , "EnterRoomAck" , ack)
}