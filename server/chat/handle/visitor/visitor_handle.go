package visitor

import (
	"chat_server/chat/protocol"
	"github.com/golang/protobuf/proto"
)

func (h *HandlesManage) Chat(flag int32 , data []byte) {
	req := &protocol.ChatReq{}
	if err := h.Unmarshal(data , req) ; err != nil {
		h.ReplyErr(h.Session.PlayerId , flag , "ChatAck" , err)
		return
	}

	cst := &protocol.ChatCst{
		Data:req.Data,
		PlayerId:proto.String(h.Session.PlayerId),
	}
	h.CastButler(h.Session.PlayerId,flag,"ChatCst",cst)
}

func (h *HandlesManage) CreateRoom(flag int32 , data []byte) {
	req := &protocol.CreateRoomReq{}
	if err := h.Unmarshal(data , req) ; err != nil {
		h.ReplyErr(h.Session.PlayerId , flag , "CreateRoomAck" , err)
		return
	}

	cst := &protocol.CreateRoomCst{
		PlayerId:proto.String(h.Session.PlayerId),
		RoomName:req.RoomName,
	}
	h.CastButler(h.Session.PlayerId,flag,"CreateRoomCst" ,cst)
}

func (h HandlesManage) EnterRoom(flag int32 , data []byte) {
	req := &protocol.EnterRoomReq{}
	if err := h.Unmarshal(data , req) ; err != nil {
		h.ReplyErr(h.Session.PlayerId , flag , "CreateRoomAck" , err)
		return
	}

	cst := &protocol.EnterRoomCst{
		PlayerId:proto.String(h.Session.PlayerId),
		RoomName:req.RoomName,
	}
	h.CastButler(h.Session.PlayerId,flag,"CreateRoomCst" ,cst)
}