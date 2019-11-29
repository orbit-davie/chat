package processor

import (
	"chat_server/chat/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/wonderivan/logger"
)

func(r *Room) EnterRoomCst(flag int32 , data []byte) {
	cst := &protocol.EnterRoomCst{}
	if err := r.Controller.Unmarshal(data , cst) ; err != nil {
		logger.Error("EnterRoomCst unmarshal failed...")
		return
	}

	rep := &protocol.EnterRoomCstResponse{
		PlayerId:cst.PlayerId,
		RoomName:cst.RoomName,
		RoomId:proto.String(r.Model.Id),
	}
	if r.Model.Name != cst.GetRoomName() {
		rep.Result = proto.Bool(false)
	}else {
		r.Model.Enter(cst.GetPlayerId())
		rep.Result = proto.Bool(true)
	}

	r.Controller.CstButler(cst.GetPlayerId() , flag ,"EnterRoomCstResponse" , rep)
}

func(r *Room) LeaveRoomCst(flag int32 , data []byte) {
	cst := &protocol.LeaveRoomCst{}
	if err := r.Controller.Unmarshal(data , cst) ; err != nil {
		logger.Error("LeaveRoomCst unmarshal failed...")
		return
	}

	rep := &protocol.LeaveRoomCstResponse{
		PlayerId:cst.PlayerId,
		RoomName:cst.RoomName,
	}

	if !r.Model.Check(cst.GetPlayerId()) {
		rep.Result = proto.Bool(false)
	}else {
		rep.Result = proto.Bool(true)
		r.Model.Leave(cst.GetPlayerId())
	}
	r.Controller.CstButler(cst.GetPlayerId() , flag ,"LeaveRoomCstResponse" , rep)
}

func(r *Room) ChatCst(flag int32 , data []byte) {
	cst := &protocol.ChatCst{}
	if err := r.Controller.Unmarshal(data , cst) ; err != nil {
		logger.Error("ChatCst unmarshal failed...")
		return
	}
	if !r.Model.Check(cst.GetPlayerId()) {
		cstRep := &protocol.ChatCstResponse{
			PlayerId:cst.PlayerId,
			Result:proto.Bool(false),
		}
		r.Controller.CstButler(cst.GetPlayerId(), flag,"ChatCstResponse",cstRep)
		return
	}

	rep := &protocol.ChatAck{
		Data:proto.String(cst.GetData()),
	}
	for _ , id := range r.Model.All() {
		r.Controller.Reply(id,flag,"ChatAck" ,rep)
	}
}