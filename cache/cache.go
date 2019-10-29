package cache

import (
	"bytes"
	"github.com/kumparan/fer/config"
	"gopkg.in/djherbis/fscache.v0"
	"io"
)

type (
	// FerSetter :nodoc:
	FerSetter func() (interface{}, error)

	// FerKeeper :nodoc:
	FerKeeper interface {
		GetOrSetFunc(key string, fn FerSetter) (interface{}, error)
	}

	keeper struct {
		cacher fscache.Cache
	}
)

// GetOrSetFunc :nodoc:
func (k *keeper) GetOrSetFunc(key string, fn FerSetter) (val interface{}, err error) {
	r, w, err := k.cacher.Get(key)
	if err != nil {
		return
	}
	defer func() {
		if r != nil {
			_ = r.Close()
		}
		if w != nil {
			_ = w.Close()
		}
	}()
	// No cache
	if w != nil {
		val, err = fn()
		if err != nil {
			return
		}
		_, err = w.Write(val.([]byte))
		return
	}
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, r)
	if err != nil {
		return
	}
	return buf.Bytes(), nil
}

// New :nodoc:
func New(dir string) (FerKeeper, error) {
	c, err := fscache.New(dir, config.CacheDirPerm, config.CacheTTLFerVersion)
	if err != nil {
		return nil, err
	}
	return &keeper{cacher: c}, nil
}
