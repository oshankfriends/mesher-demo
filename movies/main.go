package main

import (
	"log"
	"github.com/oshankfriends/mesher-demo/movies/routers"
	"net/http"
	"github.com/oshankfriends/mesher-demo/movies/common"
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

