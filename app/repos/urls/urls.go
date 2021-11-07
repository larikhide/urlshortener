package urls

import (
	"context"

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

	SaveUrlMapping(ctx context.Context, shortUrl string, longUrl string, uid uuid.UUID) error
	RetrieveInitialUrl(ctx context.Context, shortUrl string) (string, error)
}

type URLs struct {
	urlstore URLStore
}

func NewURLs(urlstore URLStore) *URLs {
	return &URLs{
		urlstore: urlstore,
	}
}
