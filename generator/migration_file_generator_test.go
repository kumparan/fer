package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUniqueTime(t *testing.T) {
	m := migration{}
	res := m.createUniqueTime()
	t.Log(res)
	assert.NotEmpty(t, res)
	assert.Equal(t, len("20190201101112"), len(res))
}
