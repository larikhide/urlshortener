package redisdb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testRedisDB = &RedisDB{}

func init() {
	_store, _ := NewDB("redis://user:password@localhost:6789/test?db=0")
	testRedisDB = _store
}

func TestNewDB(t *testing.T) {
	assert.True(t, testRedisDB.db != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL := "Jsz4k57oAX"
	ctx := context.Background()

	// Persist data mapping
	testRedisDB.SaveUrlMapping(ctx, shortURL, initialLink)

	// Retrieve initial URL
	actual, _ := testRedisDB.RetrieveInitialUrl(ctx, shortURL)
	assert.Equal(t, initialLink, actual)
}
