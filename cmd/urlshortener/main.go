package main

import (
	"log"
	"os"
	"time"

	"github.com/larikhide/urlshortener/api/routergin"
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

	router := routergin.NewRouterGin()

	err := router.Run(":9808")
	if err != nil {
		log.Fatal("Failed to start the web server - Error: %w", err)
	}

}
