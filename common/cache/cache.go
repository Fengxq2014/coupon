package cache

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

type Cache interface {
	Set(k string, v interface{}, duration time.Duration)
	Get(k string) (v interface{}, err error)
}

type goCache struct {
	c *cache.Cache
}

var onceCache *goCache
var once sync.Once

func GetCache() *goCache {
	once.Do(func() {
		onceCache = &goCache{cache.New(30*time.Minute, 30*time.Minute)}
	})
	return onceCache
}

func (gc goCache) Set(k string, v interface{}, duration time.Duration) {
	gc.c.Set(k, v, duration)
}

func (gc goCache) Get(k string) (v interface{}, err error) {
	var b bool
	v, b = gc.c.Get(k)
	if !b {
		err = errors.New(fmt.Sprintf("could not get cache value by key:%s", k))
	}
	return
}
