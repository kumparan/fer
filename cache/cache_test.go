package cache

import (
	"github.com/kumparan/fer/config"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestKeeper_GetOrSetFunc(t *testing.T) {
	tempDir, err := ioutil.TempDir("", config.TempDirPrefix)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tempDir)
	}()

	c, err := New(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	key := "somekey"
	cachedValue := []byte("Yep this is the shit")

	v1, err := c.GetOrSetFunc(key, func() (interface{}, error) {
		return cachedValue, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if v1 == nil {
		t.Fatal("v should not be nil")
	}
	assert.Equal(t, v1.([]byte), cachedValue)

	v2, err := c.GetOrSetFunc(key, func() (interface{}, error) {
		t.Fatal("This should not be called")
		return nil, nil
	})
	assert.NotNil(t, v2)
	assert.Equal(t, v2.([]byte), cachedValue)
}
