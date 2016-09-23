package miccheck

type builder struct {
	assembledExpectation *expectation
}

func (b *builder) assemble(expects map[string]interface{}) {
	var exp *expectation
	if b.assembledExpectation != nil {
		exp = &expectation{data: expects, nextExpector: b.assembledExpectation}
	} else {
		exp = &expectation{data: expects}
	}

	b.assembledExpectation = exp
}

func (b *builder) build() *expectation {
	return b.assembledExpectation
}
