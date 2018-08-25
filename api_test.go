package net

import (
	"fmt"
	"net"
	"testing"

	"github.com/xtech-cloud/omo-mod-net/processor"
	"github.com/xtech-cloud/omo-mod-net/protocol"
	"github.com/xtech-cloud/omo-mod-net/service"
)

func handleReporterPing(_req *protocol.Request, _rsp *protocol.Response, _sender interface{}) {
	_rsp.Head.Msg = "pong"
	_rsp.Body = &protocol.EmptyBlock{}
	sender := _sender.(*net.UDPAddr)
	fmt.Println(sender.String())
}

func handleWorkerPing(_req *protocol.Request, _rsp *protocol.Response, _sender interface{}) {
	_rsp.Head.Msg = "pong"
	_rsp.Body = &protocol.EmptyBlock{}
}

func Test_RunBroker(_t *testing.T) {
	reporterProcessor := processor.NewProcessor()
	reporterProcessor.BindJsonHandler("ping", handleReporterPing)

	reporter, _ := service.NewReporter()
	go reporter.Run(":18999", reporterProcessor.ProcessJson)

	broker, _ := service.NewBroker()
	go broker.Run("tcp://*:3000", "tcp://*:9990")

	workerProcessor := processor.NewProcessor()
	workerProcessor.BindJsonHandler("ping", handleWorkerPing)
	worker, _ := service.NewWorker()
	worker.Run("tcp://127.0.0.1:9990", workerProcessor.ProcessJson)
}
