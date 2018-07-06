package service

import (
	"strings"

	zmq "github.com/pebbe/zmq4"
)

func RunSubscriber(_name string, _proc string, _filter string, _processor func(string)) {
	if nil == _processor {
		panic("need processor")
	}
	subscriber, err := zmq.NewSocket(zmq.SUB)
	if nil != err {
		panic(err)
	}
	defer subscriber.Close()
	subscriber.Connect(_proc)

	subscriber.SetSubscribe(_filter)

	for {
		msg, _ := subscriber.Recv(0)

		if msgs := strings.Fields(msg); len(msgs) > 1 {
			_processor(msgs[1])
		}
	}
}
