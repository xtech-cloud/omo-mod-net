package service

import (
	zmq "github.com/pebbe/zmq4"
)

type _Publisher struct {
	msgChan chan string
	socket  *zmq.Socket
}

var publishers = make(map[string]*_Publisher)

func RunPublisher(_name string, _proc string) {
	socket, err := zmq.NewSocket(zmq.PUB)
	if nil != err {
		panic(err)
	}
	defer socket.Close()

	publisher := &_Publisher{}
	publisher.socket = socket
	publisher.socket.Bind(_proc)
	publishers[_name] = publisher

	for {
		select {
		case msg := <-publisher.msgChan:
			publisher.socket.Send(msg, 0)
		}
	}

}

func Publish(_name string, _msg string) {
	if _, ok := publishers[_name]; !ok {
		return
	}
	publishers[_name].msgChan <- _msg
}