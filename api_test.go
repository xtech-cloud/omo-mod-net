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
	processor.BindJsonHandler("ping", handlePing)

	reporter, _ := service.NewReporter()
	go reporter.Run(":18999", processor.ProcessJson)

	broker, _ := service.NewBroker()
	go broker.Run("tcp://*:3000", "tcp://*:9990")

	worker, _ := service.NewWorker()
	worker.Run("tcp://127.0.0.1:9990", processor.ProcessJson)
}
