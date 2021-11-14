package urls

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type URL struct {
	ID       uuid.UUID
	Counter  uint64
	LongURL  string
	ShortURL string
}

type URLStore interface {
	// Create(ctx context.Context, u URL) (*uuid.UUID, error)
	// Read(ctx context.Context, uid uuid.UUID) (*URL, error)
	// Delete(ctx context.Context, uid uuid.UUID) error
	// SearchUsers(ctx context.Context, s string) (chan URL, error)

	SaveUrlMapping(ctx context.Context, shortUrl string, longUrl string) error
	RetrieveInitialUrl(ctx context.Context, shortUrl string) (*URL, error)
}

type URLs struct {
	urlstore URLStore
}

func NewURLs(urlstore URLStore) *URLs {
	return &URLs{
		urlstore: urlstore,
	}
}

func (us *URLs) Create(ctx context.Context, u URL) (*URL, error) {
	if err := us.urlstore.SaveUrlMapping(ctx, u.ShortURL, u.LongURL); err != nil {
		return nil, fmt.Errorf("save url error: %w", err)
	}
	return &u, nil
}

func (us *URLs) Read(ctx context.Context, shortUrl string) (*URL, error) {
	u, err := us.urlstore.RetrieveInitialUrl(ctx, shortUrl)
	if err != nil {
		return nil, fmt.Errorf("read url error: %w", err)
	}
	return u, nil
}
