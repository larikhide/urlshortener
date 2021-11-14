package main

import (
	"log"
	"os"
	"time"

	"github.com/larikhide/urlshortener/api/handler"
	"github.com/larikhide/urlshortener/api/routergin"
	"github.com/larikhide/urlshortener/app/repos/urls"
	"github.com/larikhide/urlshortener/db/mem"
	"github.com/larikhide/urlshortener/db/postgresdb"
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

	var storage urls.URLStore

	//stt := os.Getenv("URLSHORTENER_STORE")
	stt := "mem"

	switch stt {
	case "rds":
		//dsn := os.Getenv("DATABASE_URL")
		dsn := "redis://user:password@localhost:6789/3?dial_timeout=3&db=1&read_timeout=6s&max_retries=2"
		rds, err := redisdb.NewDB(dsn)
		if err != nil {
			log.Fatal("Cannot initialize database: %w", err)
		}
		storage = rds
	case "pgst":
		//dsn := os.Getenv("DATABASE_URL")
		dsn := "postgres://postgres:1110@localhost/test?sslmode=disable"
		pgs, err := postgresdb.NewDB(dsn)
		if err != nil {
			log.Fatal("Cannot initialize database: %w", err)
		}
		defer pgs.Close()
		storage = pgs
	case "mem":
		storage = mem.NewDB()
	default:
		log.Fatal("unknown URLSHORTENER_STORE = ", stt)
	}

	st := urls.NewURLs(storage)
	hs := handler.NewHandlers(st)
	router := routergin.NewRouterGin(hs)

	err := router.Run(":9808")
	if err != nil {
		log.Fatal("Failed to start the web server - Error: %w", err)
	}

}
