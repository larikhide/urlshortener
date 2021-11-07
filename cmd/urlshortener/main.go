package main

import (
	"log"
	"os"
	"time"

	"github.com/larikhide/urlshortener/api/handler"
	"github.com/larikhide/urlshortener/api/routergin"
	"github.com/larikhide/urlshortener/app/repos/urls"
	"github.com/larikhide/urlshortener/db/redisdb"
)

func main() {

	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("error loading location '%s': %v\n", tz, err)
		}
	}

	tnow := time.Now()
	tz, _ := tnow.Zone()
	log.Printf("Local time zone %s. Service started at %s", tz,
		tnow.Format("2006-01-02T15:04:05.000 MST"))

	db, err := redisdb.NewDB()
	if err != nil {
		log.Fatal("Cannot initialize database: %w", err)
	}

	storage := urls.NewURLs(db)
	hs := handler.NewHandlers(storage)
	router := routergin.NewRouterGin(hs)

	err = router.Run(":9808")
	if err != nil {
		log.Fatal("Failed to start the web server - Error: %w", err)
	}

}
