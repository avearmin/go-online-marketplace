package main

import (
	"log"
	"net/http"
	"time"

	"github.com/avearmin/gorage-sale/internal/api"
	_ "github.com/lib/pq"
)

func main() {
	apiCfg, err := api.CreateConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/status", api.HandleStatus)

	mux.HandleFunc("/v1/auth/google/login", apiCfg.HandleOAuthGoogleLogin)
	mux.HandleFunc("/v1/auth/google/callback", apiCfg.HandleOAuthGoogleCallback)

	mux.HandleFunc("/v1/users", apiCfg.HandleUsers)
	mux.HandleFunc("/v1/items", apiCfg.HandleItems)

	corsMux := api.MiddlewareCors(mux)

	srv := http.Server{
		Addr:         ":" + apiCfg.Port,
		Handler:      corsMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Serving on port: " + apiCfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
