package controller

import (
	"chat_server/app/controller/protocol"
	"github.com/golang/protobuf/proto"
	"github.com/wonderivan/logger"
)

type Controller struct {
	Handles   map[string]func(flag int32 , data []byte)
}

func NewController() *Controller {
	return &Controller{
		Handles:make(map[string]func(flag int32 , data []byte)),
	}
}

func (c *Controller) Dispense(data []byte) {
	req , err := UnpackReq(data)
	if err != nil {
		logger.Error("request api unmarshal failed : %s", err)
		return
	}

	handle ,ok := c.Handles[req.GetReqName()]
	if !ok {
		logger.Error("reqName not registry handle...")
		return
	}
	handle(req.GetFlag(),req.GetData())
}

func (c *Controller) PackMessage(flag int32 , ackName string , pb proto.Message) (bytes []byte , err error) {
	return PackMessage(&app.ApiAck{Flag:proto.Int32(flag)} , ackName , pb)
}

func (c *Controller) PackCstMessage(flag int32 , cstName string , pb proto.Message) (bytes []byte , err error) {
	return PackCstMessage(&app.ApiReq{Flag:proto.Int32(flag)} , cstName , pb)
}