package dingtalk

import (
	"sync"

	"github.com/fromiuan/dingtalk/lib/cache"
)

type Client struct {
	AppKey    string // 企业内部应用appKey
	AppSecret string // 企业内部应用appSecret
	Debug     bool

	Cache cache.Cache
	Tlock *sync.RWMutex
}

func NewClient(appkey, appsecret string) *Client {
	cli := &Client{
		AppKey:    appkey,
		AppSecret: appsecret,
		Tlock:     new(sync.RWMutex),
	}
	defaultCacheCfg := &cache.MemoryOpts{Interval: 1 * 60 * 60}
	cacheAdapter, err := cache.NewCache("memory", defaultCacheCfg)
	if err != nil {
		panic("not find memory cache")
	}
	cli.Cache = cacheAdapter
	return cli
}

func (c *Client) SetCache(key string, cfg interface{}) {
	c.Cache, _ = cache.NewCache(key, cfg)
}

func (c *Client) SetDebug(b bool) {
	c.Debug = b
}
