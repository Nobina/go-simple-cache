package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"context"
)

type memoryCache struct {
	items map[string][]byte
}

func (c *memoryCache) Delete(ctx context.Context, k string) error {
	delete(c.items, k)

	return nil
}

func (c *memoryCache) Flush(ctx context.Context) error {
	c.items = map[string][]byte{}

	return nil
}

func (c *memoryCache) Get(ctx context.Context, k string, v interface{}) error {
	buf := c.items[k]
	if buf == nil {
		return ErrCacheMiss
	}

	if err := json.Unmarshal(buf, v); err != nil {
		return err
	}

	return nil
}

func (c *memoryCache) Set(ctx context.Context, k string, v interface{}, expire time.Duration) error {
	buf, err := json.Marshal(v)
	if err != nil {
		return err
	}

	c.items[k] = buf

	return nil
}

func (c *memoryCache) Nearby(ctx context.Context, k string, lon, lat, radius float64) ([]Location, error) {
	return nil, fmt.Errorf("not supported")
}

func (c *memoryCache) GeoAdd(ctx context.Context, k string, locations ...Location) error {
	return fmt.Errorf("not supported")
}

func NewMemoryCache() Client {
	return &memoryCache{
		items: map[string][]byte{},
	}
}
