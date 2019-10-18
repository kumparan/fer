package deploy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateTagTime(t *testing.T) {
	tagTime := CreateTagTime()
	assert.Equal(t, len("v20191018.1571371141"), len(tagTime))
}
