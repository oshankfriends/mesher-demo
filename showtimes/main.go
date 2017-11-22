package main

import (
	"log"
	"github.com/oshankfriends/mesher-demo/showtimes/routers"
	"net/http"
	"github.com/oshankfriends/mesher-demo/showtimes/common"
)

func main() {
	router := routers.InitRoutes()

	server := &http.Server{
		Handler: router,
		Addr:    common.AppConfig.Server,
	}
	log.Printf("Listening on %s",common.AppConfig.Server)
	server.ListenAndServe()
}

