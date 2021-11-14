package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/larikhide/urlshortener/app/repos/urls"
	"github.com/larikhide/urlshortener/app/shortener"
)

type URL struct {
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}

type Handlers struct {
	db *urls.URLs
}

func NewHandlers(db *urls.URLs) *Handlers {
	r := &Handlers{
		db: db,
	}
	return r
}

func (hs *Handlers) GenShortLink(ctx context.Context, u URL) (URL, error) {

	bu := urls.URL{
		ShortURL: u.ShortURL,
		LongURL:  u.LongURL,
	}

	shortUrl, err := shortener.GenShortLink(bu.LongURL)
	if err != nil {
		return URL{}, fmt.Errorf("error when generating: %w", err)
	}

	bu.ShortURL = shortUrl

	nbu, err := hs.db.Create(ctx, bu)
	if err != nil {
		return URL{}, fmt.Errorf("error when creating: %w", err)
	}

	return URL{
		ShortURL: nbu.ShortURL,
		LongURL:  nbu.LongURL,
	}, nil
}

func (hs *Handlers) HandleShortUrlRedirect(ctx context.Context, shortUrl string) (URL, error) {
	// if reflect.TypeOf(shortUrl) == reflect.TypeOf("") {
	// 	return URL{}, fmt.Errorf("bad request: input isnt string")
	// }

	nbu, err := hs.db.Read(ctx, shortUrl)
	if err != nil {
		return URL{}, errors.New("url not found")
	}
	return URL{
		LongURL:  nbu.LongURL,
		ShortURL: nbu.ShortURL,
	}, nil
}
