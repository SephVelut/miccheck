package miccheckmock

import testifyMock "github.com/stretchr/testify/mock"

type mock struct {
	testifyMock testifyMock.Mock
}

func (m *mock) On(method string) *testifyMock.Call {
	return m.testifyMock.On(method)
}

func (m *mock) AssertExpectations(t testifyMock.TestingT) bool {
	return m.testifyMock.AssertExpectations(t)
}

func (m *mock) AssertNotCalled(t testifyMock.TestingT, method string) bool {
	return m.testifyMock.AssertNotCalled(t, method)
}
