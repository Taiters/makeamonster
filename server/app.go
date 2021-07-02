package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Taiters/monstermaker/api"
	"github.com/gorilla/mux"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 5000, "Port to listen on")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()

	api.BootstrapServer(router)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	log.Printf("Server running on port %d", port)
	log.Fatal(srv.ListenAndServe())
}
