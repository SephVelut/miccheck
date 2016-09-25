package miccheckmock

import testify "github.com/stretchr/testify/mock"

type call struct {
	tCall *testify.Call
}

func (c *call) setCall(tCall *testify.Call) {
	c.tCall = tCall
}

func (c *call) andReturn(args ...interface{}) {
	c.tCall.Return(args...)
}
