package miccheck

import "reflect"

type expectation struct {
	data         map[string]interface{}
	nextExpector *expectation
}

func (e *expectation) validate(expectation *expectation) []map[string]interface{} {
	var matches = []map[string]interface{}{}
	var ourData = e.getData()

	if !e.isTerminal() {
		previousMatches := e.nextExpector.validate(expectation)
		matches = append(matches, previousMatches...)
	}

	skips := 0
	for _, v := range matches {
		if reflect.DeepEqual(v, ourData) {
			skips++
		}
	}

	var i = 0
	var expectationToMatch = expectation
	var expectationData map[string]interface{}

ExpectationMatchLoop:
	for {
		expectationData = expectationToMatch.getData()
		if reflect.DeepEqual(ourData, expectationData) {
			if len(matches) == 0 {
				matches = append(matches, ourData)
				break ExpectationMatchLoop
			}

			var found bool
			for _, v := range matches {
				if reflect.DeepEqual(v, ourData) {
					found = true
					if i == skips {
						matches = append(matches, ourData)
						break ExpectationMatchLoop
					}

					i++
				}
			}

			if !found {
				matches = append(matches, ourData)
			}
		}

		if expectationToMatch.isTerminal() {
			break ExpectationMatchLoop
		}

		expectationToMatch = expectationToMatch.next()
	}

	return matches
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
