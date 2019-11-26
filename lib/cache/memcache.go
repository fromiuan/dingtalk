package cache

import (
	"errors"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type MemCache struct {
	conn     *memcache.Client
	conninfo []string
}

type MemOpts struct {
	Conn string
}

func NewMemCache() Cache {
	return &MemCache{}
}

func (rc *MemCache) Get(key string) interface{} {
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return err
		}
	}
	if item, err := rc.conn.Get(key); err == nil {
		return item.Value
	}
	return nil
}

func (rc *MemCache) Set(key string, val interface{}, timeout time.Duration) error {
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return err
		}
	}
	item := memcache.Item{Key: key, Expiration: int32(timeout / time.Second)}
	if v, ok := val.([]byte); ok {
		item.Value = v
	} else if str, ok := val.(string); ok {
		item.Value = []byte(str)
	} else {
		return errors.New("val only support string and []byte")
	}
	return rc.conn.Set(&item)
}

func (rc *MemCache) Delete(key string) error {
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return err
		}
	}
	return rc.conn.Delete(key)
}

func (rc *MemCache) IsExist(key string) bool {
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return false
		}
	}
	_, err := rc.conn.Get(key)
	return !(err != nil)
}

func (rc *MemCache) Init(cfg interface{}) error {
	var opts *MemOpts
	if val, ok := cfg.(*MemOpts); !ok {
		return errors.New("interface not type MemOpts")
	} else {
		opts = val
	}
	if opts.Conn == "" {
		return errors.New("config has no conn key")
	}
	rc.conninfo = strings.Split(opts.Conn, ";")
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return err
		}
	}
	return nil
}

func (rc *MemCache) connectInit() error {
	rc.conn = memcache.New(rc.conninfo...)
	return nil
}

func init() {
	Register("memcache", NewMemCache())
}
