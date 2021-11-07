package redisdb

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/larikhide/urlshortener/app/repos/urls"
)

var _ urls.URLStore = &RedisDB{}

type RedisDB struct {
	db *redis.Client
}

// Note that in a real world usage, the cache duration shouldn't have
// an expiration time, an LRU policy config should be set where the
// values that are retrieved less often are purged automatically from
// the cache and stored back in RDBMS whenever the cache is full
const CacheDuration = 6 * time.Hour

func NewDB() (*RedisDB, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     "localhost",
		Password: "",
		DB:       0,
	})

	// TODO: нормально ли здесь создавать контекст для пинга? Он вроде пустой и ни к чему не обязывает
	ctx := context.Background()
	_, err := db.Ping(ctx).Result()
	if err != nil {
		db.Close()
		return nil, err
	}

	us := &RedisDB{
		db: db,
	}

	return us, nil
}

func (rds *RedisDB) SaveUrlMapping(ctx context.Context, shortUrl string, longUrl string) error {
	err := rds.db.Set(ctx, shortUrl, longUrl, CacheDuration).Err()
	if err != nil {
		//log.Printf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)
		return fmt.Errorf("failed saving key url | error: %v - shorturl: %s - originalurl: %s", err, shortUrl, longUrl)
	}
	return nil
}

func (rds *RedisDB) RetrieveInitialUrl(ctx context.Context, shortUrl string) (*urls.URL, error) {
	result, err := rds.db.Get(ctx, shortUrl).Result()
	if err != nil {
		//panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
		return &urls.URL{}, fmt.Errorf("failed retrieve url | error: %v - shorturl: %s", err, shortUrl)
	}
	return &urls.URL{
		LongURL:  result,
		ShortURL: shortUrl,
	}, nil
}
