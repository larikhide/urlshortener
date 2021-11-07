package redisdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var testRedisDB = &RedisDB{}

func init() {
	_store, err := NewDB()
	if err != nil {
		fmt.Errorf("unexpected error: %w", err)
	}
	testRedisDB = _store
}

func TestNewDB(t *testing.T) {
	assert.True(t, testRedisDB.db != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	uid := uuid.New()
	shortURL := "Jsz4k57oAX"
	ctx := context.Background()

	// Persist data mapping
	testRedisDB.SaveUrlMapping(ctx, shortURL, initialLink, uid)

	// Retrieve initial URL
	actual, _ := testRedisDB.RetrieveInitialUrl(ctx, shortURL)
	assert.Equal(t, initialLink, actual)
}
