package miccheckmock

type contractWriterMediator interface {
	ExpectationPromised([]map[string]interface{})
	ExpectationFullfilled([]map[string]interface{})
}
