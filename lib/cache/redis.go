package cache

import (
	"errors"
	"time"

	"github.com/gomodule/redigo/redis"
)

//Redis redis cache
type Redis struct {
	conn *redis.Pool
}

type RedisOpts struct {
	Host        string
	Password    string
	Database    int
	MaxIdle     int
	MaxActive   int
	IdleTimeout int32
}

func NewRedisCache() Cache {
	return &Redis{}
}

func (r *Redis) SetConn(conn *redis.Pool) {
	r.conn = conn
}

func (r *Redis) Get(key string) interface{} {
	conn := r.conn.Get()
	defer conn.Close()

	if v, err := redis.Bytes(conn.Do("GET", key)); err == nil {
		return v
	}
	return nil
}

func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	conn := r.conn.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, int64(timeout/time.Second), val)
	return err
}

func (r *Redis) Delete(key string) error {
	conn := r.conn.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}

func (r *Redis) IsExist(key string) bool {
	conn := r.conn.Get()
	defer conn.Close()

	a, _ := conn.Do("EXISTS", key)
	i := a.(int64)
	if i > 0 {
		return true
	}
	return false
}

func (r *Redis) Init(cfg interface{}) error {
	var opts *RedisOpts
	if val, ok := cfg.(*RedisOpts); !ok {
		return errors.New("interface not type RedisOpts")
	} else {
		opts = val
	}
	pool := &redis.Pool{
		MaxActive:   opts.MaxActive,
		MaxIdle:     opts.MaxIdle,
		IdleTimeout: time.Second * time.Duration(opts.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", opts.Host,
				redis.DialDatabase(opts.Database),
				redis.DialPassword(opts.Password),
			)
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
	r.conn = pool
	return nil
}

func init() {
	Register("redis", NewRedisCache())
}
