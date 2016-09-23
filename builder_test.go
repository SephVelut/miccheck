package miccheck

import "github.com/stretchr/testify/assert"
import "testing"

func TestBuildsExpectation(t *testing.T) {
	b := &builder{}
	data := map[string]interface{}{"key": "value"}
	data2 := map[string]interface{}{"key": "value2"}

	b.assemble(data)
	b.assemble(data2)
	exp := b.build()

	assert.True(t, exp.getData()["key"] == "value2")
	assert.True(t, exp.next().getData()["key"] == "value")
}
