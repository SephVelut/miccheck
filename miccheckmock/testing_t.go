package miccheckmock

type testingT interface {
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	FailNow()
}
