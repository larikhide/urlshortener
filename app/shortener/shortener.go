package shortener

//TODO: интерфейс для работы с другими реализациями , например, https://pkg.go.dev/google.golang.org/api/urlshortener/v1
//попробую реализовать позже
// type URLShortener interface {
// 	Cut(ctx context.Context, u urls.URL) (*urls.URL, error)
// 	Expand(ctx context.Context, u urls.URL) (*urls.URL, error)
// 	GetStat(ctx context.Context, uid uuid.UUID) (*urls.URL, error)
// }

// type Shortener struct {
// 	shortener URLShortener
// }

// func NewShortener(shortener URLShortener) *Shortener {
// 	return &Shortener{
// 		shortener: shortener,
// 	}
// }
