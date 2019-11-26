package dingtalk

import (
	"sync"

	"github.com/fromiuan/dingtalk/cache"
)

type Client struct {
	AppKey    string
	AppSecret string
	Debug     bool

	Cache *cache.Cache
	Tlock *sync.RWMutex
}

func NewClient(appkey, appsecret string) *Client {
	cli := &Client{
		AppKey:    appkey,
		AppSecret: appsecret,
	}
	defaultCacheCfg := cache.MemoryOpts{Interval: 1 * 60 * 60}
	cli.Cache = cache.NewCache("memory", defaultCacheCfg)
	return cli
}

func (c *Client) SetCache(key string, cfg interface{}) {
	c.Cache = cache.NewCache(key, cfg)
}

func (c *Client) SetDebug(b bool) {
	c.Debug = b
}
