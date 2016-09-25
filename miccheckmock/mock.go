package miccheckmock

import (
	"regexp"
	"runtime"
	"strings"

	testifyMock "github.com/stretchr/testify/mock"
)

type mock struct {
	testifyMock    testifyMock.Mock
	contractWriter contractWriterMediator
	expectations   map[string][]map[string]interface{}
}

func (m *mock) Called() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("Couldn't get the caller information")
	}

	functionName := m.lastFunctionNameFromRuntime(pc)

	if functionName != "" && m.expectations[functionName] != nil {
		m.contractWriter.ExpectationFullfilled(m.expectations[functionName])
	}
}

func (m *mock) On(method string, request, response []map[string]interface{}) *call {
	if request != nil {
		m.expectations = map[string][]map[string]interface{}{method: request}

		if response != nil {
			m.contractWriter.ExpectationPromised(request, response)
		} else {
			m.contractWriter.ExpectationPromised(request, []map[string]interface{}{})
		}
	}

	call := &call{}
	tCall := m.testifyMock.On(method)
	call.setCall(tCall)

	return call
}

func (m *mock) AssertExpectations(t testingT) bool {
	return m.testifyMock.AssertExpectations(t)
}

func (m *mock) AssertNotCalled(t testingT, method string) bool {
	return m.testifyMock.AssertNotCalled(t, method)
}

func (m *mock) SetContractWriter(contractWriter contractWriterMediator) {
	m.contractWriter = contractWriter
}

func (m *mock) lastFunctionNameFromRuntime(pc uintptr) string {
	functionPath := runtime.FuncForPC(pc).Name()

	re := regexp.MustCompile("\\.pN\\d+_")
	if re.MatchString(functionPath) {
		functionPath = re.Split(functionPath, -1)[0]
	}

	parts := strings.Split(functionPath, ".")
	functionName := parts[len(parts)-1]

	return functionName
}
