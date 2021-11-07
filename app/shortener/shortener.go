package shortener

import (
	"context"

	"github.com/google/uuid"
	"github.com/larikhide/urlshortener/app/repos/urls"
)

type URLShortener interface {
	Cut(ctx context.Context, u urls.URL) (*urls.URL, error)
	Expand(ctx context.Context, u urls.URL) (*urls.URL, error)
	GetStat(ctx context.Context, uid uuid.UUID) (*urls.URL, error)
	//TODO: Counter method?
}

type Shortener struct {
	shortener URLShortener
}

func NewShortener(shortener URLShortener) *Shortener {
	return &Shortener{
		shortener: shortener,
	}
}
