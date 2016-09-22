package miccheck

import "reflect"

type expectation struct {
	data         map[string]interface{}
	nextExpector *expectation
}

func (e *expectation) validate(expectation *expectation) []map[string]interface{} {
	var existingMatches = []map[string]interface{}{}
	if !e.isTerminal() {
		existingMatches = e.nextExpector.validate(expectation)
	}

	return e.findMatches(expectation, existingMatches)
}

func (e *expectation) findMatches(expectation *expectation, existingMatches []map[string]interface{}) []map[string]interface{} {
	var i = 0

ExpectationMatchLoop:
	for {
		expectationData := expectation.getData()
		if reflect.DeepEqual(e.getData(), expectationData) {
			if len(existingMatches) == 0 {
				existingMatches = append(existingMatches, e.getData())
				break
			}

			var found bool
			for _, v := range existingMatches {
				if reflect.DeepEqual(v, e.getData()) {
					found = true
					if i == e.countMatches(existingMatches) {
						existingMatches = append(existingMatches, e.getData())
						break ExpectationMatchLoop
					}

					i++
				}
			}

			if !found {
				existingMatches = append(existingMatches, e.getData())
			}
		}

		if expectation.isTerminal() {
			break
		}

		expectation = expectation.next()
	}

	return existingMatches
}

func (e *expectation) isTerminal() bool {
	if e.nextExpector == nil {
		return true
	}

	return false
}

func (e *expectation) getData() map[string]interface{} {
	return e.data
}

func (e *expectation) next() *expectation {
	return e.nextExpector
}

func (e *expectation) countMatches(matches []map[string]interface{}) int {
	var skips = 0

	for _, v := range matches {
		if reflect.DeepEqual(v, e.getData()) {
			skips++
		}
	}

	return skips
}
