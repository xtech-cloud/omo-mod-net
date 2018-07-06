package service

import (
	zmq "github.com/pebbe/zmq4"
)

func RunWorker(_name string, _proc string, _processor func([]byte) ([]byte, error)) {
	if nil == _processor {
		panic("need processor")
	}

	responder, err := zmq.NewSocket(zmq.REP)
	if nil != err {
		panic(err)
	}
	defer responder.Close()
	responder.Connect(_proc)

	for {
		//  Wait for next request from client
		req, err := responder.RecvBytes(0)
		if nil != err {
			continue
		}

		rsp, err := _processor(req)
		if nil != err {
			rsp = make([]byte, 0)
			rsp[0] = 255
			responder.SendBytes(rsp, 0)
			continue
		}
		responder.SendBytes(rsp, 0)
	}
}
