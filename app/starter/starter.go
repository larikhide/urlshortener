package starter

import (
	"context"
	"sync"

	"github.com/larikhide/urlshortener/app/repos/urls"
	"github.com/larikhide/urlshortener/app/shortener"
)

type App struct {
	us *urls.URLs
	sh *shortener.Shortener
}

func NewApp(ust urls.URLStore, sht shortener.URLShortener) *App {
	a := &App{
		us: urls.NewURLs(ust),
		sh: shortener.NewShortener(sht),
	}
	return a
}

type HTTPServer interface {
	Start(us *urls.URLs, sh *shortener.Shortener)
	Stop()
}

func (a *App) Serve(ctx context.Context, wg *sync.WaitGroup, hs HTTPServer) {
	defer wg.Done()
	hs.Start(a.us, a.sh)
	<-ctx.Done()
	hs.Stop()
}
