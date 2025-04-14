package cache

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/appengine/memcache"
)

type appengineCache struct {
	prefix string
	codec  memcache.Codec
}

func (c *appengineCache) key(k string) string { return c.prefix + "/" + k }

func (c *appengineCache) Delete(ctx context.Context, k string) error {
	return memcache.Delete(ctx, c.key(k))
}

func (c *appengineCache) Flush(ctx context.Context) error {
	return memcache.Flush(ctx)
}

func (c *appengineCache) Get(ctx context.Context, k string, v interface{}) error {
	_, err := c.codec.Get(ctx, c.key(k), v)
	if err == memcache.ErrCacheMiss {
		return ErrCacheMiss
	}

	return err
}

func (c *appengineCache) Set(ctx context.Context, k string, v interface{}, expire time.Duration) error {
	return c.codec.Set(ctx, &memcache.Item{
		Key:        c.key(k),
		Object:     v,
		Expiration: expire,
	})
}

func (c *appengineCache) Nearby(ctx context.Context, k string, lon, lat, radius float64) ([]Location, error) {
	return nil, fmt.Errorf("not supported")
}

func (c *appengineCache) GeoAdd(ctx context.Context, k string, locations ...Location) error {
	return fmt.Errorf("not supported")
}

func NewAppEngineCache(request *http.Request) Client {
	return &appengineCache{
		prefix: os.Getenv("GAE_SERVICE") + "/" + os.Getenv("GAE_VERSION"),
		codec: memcache.Codec{
			Marshal:   json.Marshal,
			Unmarshal: json.Unmarshal,
		},
	}
}
