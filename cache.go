package cache

import (
	"errors"
	"time"

	"context"
)

var (
	ErrCacheMiss = errors.New("cache miss")
)

type Client interface {
	Delete(ctx context.Context, k string) error
	Flush(ctx context.Context) error
	Get(ctx context.Context, k string, v interface{}) error
	Set(ctx context.Context, k string, v interface{}, expire time.Duration) error
	Nearby(ctx context.Context, k string, lon, lat, radius float64) ([]Location, error)
	GeoAdd(ctx context.Context, k string, locations ...Location) error
}

type Location struct {
	Name      string
	Longitude float64
	Latitude  float64
	Distance  float64
}
