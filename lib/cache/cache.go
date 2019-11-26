package cache

import (
	"fmt"
	"time"
)

//Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	Delete(key string) error
	IsExist(key string) bool
	Init(cfg interface{}) error
}

var adapters = make(map[string]Cache)

func Register(name string, adapter Cache) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

func NewCache(adapterName string, config interface{}) (adapter Cache, err error) {
	adapter, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	err = adapter.Init(config)
	if err != nil {
		adapter = nil
	}
	return
}
