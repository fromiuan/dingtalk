package cache

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Cache struct {
	conn     *memcache.Client
	conninfo []string
}

type MemOpts struct {
	Conn string
}

func NewMemCache() cache.Cache {
	return &Cache{}
}

func (rc *Cache) Get(key string) interface{} {
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

func (rc *Cache) Set(key string, val interface{}, timeout time.Duration) error {
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

func (rc *Cache) Delete(key string) error {
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return err
		}
	}
	return rc.conn.Delete(key)
}

func (rc *Cache) IsExist(key string) bool {
	if rc.conn == nil {
		if err := rc.connectInit(); err != nil {
			return false
		}
	}
	_, err := rc.conn.Get(key)
	return !(err != nil)
}

func (rc *Cache) Init(cfg *MemOpts) error {
	var opts *MemOpts
	if opts, ok := cfg.(*MemOpts); !ok {
		return errors.New("interface not type MemOpts")
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

func (rc *Cache) connectInit() error {
	rc.conn = memcache.New(rc.conninfo...)
	return nil
}

func init() {
	cache.Register("memcache", NewMemCache)
}
