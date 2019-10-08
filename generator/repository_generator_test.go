package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepositoryGenerator_ToCamelCase(t *testing.T) {
	r := repository{}
	res := r.toCamelCase("ProMoted_LiNk")
	assert.Equal(t, "PromotedLink", res)
}
