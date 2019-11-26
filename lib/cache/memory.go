package cache

import (
	"errors"
	"sync"
	"time"
)

var (
	DefaultEvery = 60 // 1 minute
)

type MemoryOpts struct {
	Interval int
}

type MemoryItem struct {
	val         interface{}
	createdTime time.Time
	lifespan    time.Duration
}

type MemoryCache struct {
	sync.RWMutex
	dur   time.Duration
	items map[string]*MemoryItem
	Every int
}

func NewMemoryCache() Cache {
	cache := MemoryCache{items: make(map[string]*MemoryItem)}
	return &cache
}

func (bc *MemoryCache) Get(name string) interface{} {
	bc.RLock()
	defer bc.RUnlock()
	if itm, ok := bc.items[name]; ok {
		if itm.isExpire() {
			return nil
		}
		return itm.val
	}
	return nil
}

func (bc *MemoryCache) Set(name string, value interface{}, lifespan time.Duration) error {
	bc.Lock()
	defer bc.Unlock()
	bc.items[name] = &MemoryItem{
		val:         value,
		createdTime: time.Now(),
		lifespan:    lifespan,
	}
	return nil
}

func (bc *MemoryCache) Delete(name string) error {
	bc.Lock()
	defer bc.Unlock()
	if _, ok := bc.items[name]; !ok {
		return errors.New("key not exist")
	}
	delete(bc.items, name)
	if _, ok := bc.items[name]; ok {
		return errors.New("delete key error")
	}
	return nil
}

func (bc *MemoryCache) IsExist(name string) bool {
	bc.RLock()
	defer bc.RUnlock()
	if v, ok := bc.items[name]; ok {
		return !v.isExpire()
	}
	return false
}

func (bc *MemoryCache) Init(cfg interface{}) error {
	var opts *MemoryOpts
	if val, ok := cfg.(*MemoryOpts); !ok {
		return errors.New("interface not type MemoryOpts")
	} else {
		opts = val
	}

	if opts.Interval == 0 {
		opts.Interval = DefaultEvery
	}

	dur := time.Duration(opts.Interval) * time.Second
	bc.Every = opts.Interval
	bc.dur = dur
	go bc.vacuum()
	return nil
}

func (bc *MemoryCache) vacuum() {
	bc.RLock()
	every := bc.Every
	bc.RUnlock()

	if every < 1 {
		return
	}
	for {
		<-time.After(bc.dur)
		if bc.items == nil {
			return
		}
		if keys := bc.expiredKeys(); len(keys) != 0 {
			bc.clearItems(keys)
		}
	}
}

func (bc *MemoryCache) expiredKeys() (keys []string) {
	bc.RLock()
	defer bc.RUnlock()
	for key, itm := range bc.items {
		if itm.isExpire() {
			keys = append(keys, key)
		}
	}
	return
}

func (bc *MemoryCache) clearItems(keys []string) {
	bc.Lock()
	defer bc.Unlock()
	for _, key := range keys {
		delete(bc.items, key)
	}
}

func (mi *MemoryItem) isExpire() bool {
	if mi.lifespan == 0 {
		return false
	}
	return time.Now().Sub(mi.createdTime) > mi.lifespan
}

func init() {
	Register("memory", NewMemoryCache())
}
