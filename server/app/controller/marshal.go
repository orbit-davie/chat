package controller

import (
	"chat_server/app/controller/protocol"
	"fmt"
	"github.com/golang/protobuf/proto"
)

// 将 ApiAck 编译成数据流
func PackMessage(ack *app.ApiAck , ackName string ,pb proto.Message) ( []byte , error ){
	data , err := proto.Marshal(pb)
	if err != nil {
		fmt.Errorf("marshal pb failed %s", err)
		return nil,err
	}

	message ,err := PackAckApi(ack , ackName ,data)
	if err != nil {
		fmt.Errorf("pack apiAck failed %s",err)
		return nil ,err
	}

	return message, nil
}

// 将 ApiReq 编译成数据流 , 用于cast到line中
func PackCstMessage(ack *app.ApiReq , cstName string ,pb proto.Message) ( []byte , error ){
	data , err := proto.Marshal(pb)
	if err != nil {
		fmt.Errorf("marshal pb failed %s", err)
		return nil,err
	}

	message ,err := PackReqApi(ack , cstName ,data)
	if err != nil {
		fmt.Errorf("pack apiAck failed %s",err)
		return nil ,err
	}

	return message, nil
}


func UnpackReq(message []byte) (*app.ApiReq, error){
	req := new(app.ApiReq)
	err := proto.Unmarshal(message,req)
	return req, err
}

func PackAckApi(ack *app.ApiAck ,ackName string , data []byte ) (message []byte , err error) {
	if ack == nil {
		ack = new(app.ApiAck)
	}
	ack.AckName = proto.String(ackName)
	ack.Data = data
	message , err = proto.Marshal(ack)
	return
}

func PackReqApi(req *app.ApiReq ,cstName string , data []byte) (message []byte ,err error ){
	if req == nil {
		req = new(app.ApiReq)
	}
	req.ReqName = proto.String(cstName)
	req.Data =data
	message ,err = proto.Marshal(req)
	return
}