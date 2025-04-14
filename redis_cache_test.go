package cache

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var cacheClient Client

func init() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}
	redisOptions, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err.Error())
	}

	cacheClient = NewRedisCache(redis.NewClient(redisOptions))
}

func TestRedisCacheListPlain(t *testing.T) {
	cacheKey := "list_plain"
	if err := cacheClient.Set(context.Background(), cacheKey, []string{"test"}, time.Minute); err != nil {
		t.Error(err)
	}

	listTest := []string{}
	if err := cacheClient.Get(context.Background(), cacheKey, &listTest); err != nil {
		t.Error(err)
	}

	if len(listTest) == 0 {
		t.Errorf("empty list")
	}
}

func TestRedisCacheListObject(t *testing.T) {
	cacheKey := "list_object"
	if err := cacheClient.Set(context.Background(), cacheKey, []map[string]string{{
		"foo": "bar",
	}}, time.Minute); err != nil {
		t.Error(err)
	}

	listTest := []map[string]string{}
	if err := cacheClient.Get(context.Background(), cacheKey, &listTest); err != nil {
		t.Error(err)
	}

	if len(listTest) == 0 {
		t.Errorf("empty list")
	}
}

func TestRedisCacheGeo(t *testing.T) {
	cacheKey := "geo_list"
	if err := cacheClient.GeoAdd(context.Background(), cacheKey, Location{
		Name:      "test",
		Longitude: 17.0,
		Latitude:  59.0,
	}); err != nil {
		t.Error(err)
	}

	if locations, err := cacheClient.Nearby(context.Background(), cacheKey, 17.0, 59.0, 100); err != nil {
		t.Error(err)
	} else if len(locations) == 0 {
		t.Errorf("empty list")
	}
}
