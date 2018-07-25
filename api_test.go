package net

import (
	"testing"

	"github.com/xtech-cloud/omo-mod-net/processor"
	"github.com/xtech-cloud/omo-mod-net/protocol"
	"github.com/xtech-cloud/omo-mod-net/service"
)

func handlePing(_req *protocol.Request, _rsp *protocol.Response) {
	_rsp.Head.Msg = "pong"
	_rsp.Body = &protocol.EmptyBlock{}
}

func Test_RunBroker(_t *testing.T) {
	go service.RunBroker("test", "tcp://*:3000", "tcp://*:9990")

	processor.BindJsonHandler("ping", handlePing)
	service.RunWorker("test", "tcp://127.0.0.1:9990", processor.ProcessJson)
}
