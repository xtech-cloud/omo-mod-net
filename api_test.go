package net

import (
	"testing"

	"github.com/xtech-cloud/omo-mod-net/service"
)

func Test_RunBroker(_t *testing.T) {
	service.RunBroker("test", "tcp://*.9999", "tcp://*.9990")
}
