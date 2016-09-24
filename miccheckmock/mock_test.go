package miccheckmock

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	testify "github.com/stretchr/testify/mock"
)

type ContractWriterMock struct {
	testify.Mock
}

func (c *ContractWriterMock) ExpectationPromised(expectation []map[string]interface{}) {
	c.Called(expectation)
}

func (c *ContractWriterMock) ExpectationFullfilled(expectation []map[string]interface{}) {
	c.Called(expectation)
}

type someMock struct {
	mock
}

func (s *someMock) SomeMethod() {
	s.Called()

	s.testifyMock.Called()
}

func TestSetsMethodExpectation(t *testing.T) {
	Convey("Given I expect SomeMethod to be called", t, func() {
		s := new(someMock)
		s.On("SomeMethod", nil)

		Convey("When I call SomeMethod", func() {
			s.SomeMethod()

			Convey("Then it will pass", func() {
				s.AssertExpectations(t)
			})

		})

		Convey("When I dont call SomeMethod", func() {
			s = new(someMock)
			Convey("Then it will fail", func() {
				s.AssertNotCalled(t, "SomeMethod")
			})

		})

		Convey("When I specify an expectation for SomeMethod", func() {
			contractWriter := &ContractWriterMock{}
			s := &someMock{}
			s.SetContractWriter(contractWriter)

			Convey("Then it will notify the mediator of the promised expectation", func() {
				expectation := []map[string]interface{}{map[string]interface{}{"key": "value"}}
				contractWriter.On("ExpectationPromised", expectation).Once()

				s.On("SomeMethod", expectation)

				contractWriter.AssertExpectations(t)
			})

			Convey("When I call SomeMethod once", func() {
				expectation := []map[string]interface{}{map[string]interface{}{"key": "value"}}
				contractWriter.On("ExpectationFullfilled", expectation).Once()

				contractWriter.On("ExpectationPromised", expectation)
				s.On("SomeMethod", expectation)
				s.SomeMethod()

				Convey("Then it will notify the mediator of the fullfilled expectation", func() {
					contractWriter.AssertExpectations(t)
				})
			})

			Convey("When I call SomeMethod multiple times", func() {
				expectation := []map[string]interface{}{map[string]interface{}{"key": "value"}}
				contractWriter.On("ExpectationFullfilled", expectation).Times(4)

				Convey("Then it will notify the mediator of the fullfilled expectation each time", func() {
					contractWriter.On("ExpectationPromised", expectation)
					s.On("SomeMethod", expectation)
					s.SomeMethod()
					s.SomeMethod()
					s.SomeMethod()
					s.SomeMethod()

					contractWriter.AssertExpectations(t)
				})

			})

			Convey("When I dont call SomeMethod", func() {

				Convey("Then it will not write the expectation", func() {
					contractWriter.AssertNotCalled(t, "ExpectationFullfilled")
				})

			})

		})

	})
}
