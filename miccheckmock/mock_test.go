package miccheckmock

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type someMock struct {
	mock
}

func (s *someMock) SomeMethod() {
	s.testifyMock.Called()
}

func TestSetsMethodExpectation(t *testing.T) {
	Convey("Given I expect SomeMethod to be called", t, func() {
		s := new(someMock)
		s.On("SomeMethod")

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

		Convey("When I specify a contract expectation for SomeMethod", func() {

			Convey("And I call SomeMethod", func() {

				Convey("Then it will write the expectation", nil)

			})

			Convey("And I dont call SomeMethod", func() {

				Convey("Then it will not write the expectation", nil)

			})

		})

	})
}
