package processor

import (
	"encoding/json"

	"github.com/xtech-cloud/omo-mod-net/protocol"
)

var jsonHandlers = make(map[string]func(*protocol.Request, *protocol.Response), 0)

func ProcessJson(_bytes []byte) ([]byte, error) {
	req := &protocol.Request{}
	rsp := &protocol.Response{}

	err := json.Unmarshal(_bytes, req)
	if nil != err {
		rsp.Head.ErrCode = -2
		rsp.Head.ErrString = err.Error()
		rsp.Body = &protocol.EmptyBlock{}
		return jsonToBytes(rsp)
	}

	rsp.Head.Msg = req.Head.Msg
	rsp.Head.Session = req.Head.Session

	if _, ok := jsonHandlers[req.Head.Msg]; !ok {
		rsp.Head.ErrCode = -1
		rsp.Head.ErrString = "handler not found"
		rsp.Body = &protocol.EmptyBlock{}
		return jsonToBytes(rsp)
	}

	jsonHandlers[req.Head.Msg](req, rsp)
	return jsonToBytes(rsp)
}

func BindJsonHandler(_msg string, _handler func(*protocol.Request, *protocol.Response)) {
	jsonHandlers[_msg] = _handler
}

func jsonToBytes(_json *protocol.Response) ([]byte, error) {
	return json.Marshal(_json)
}
