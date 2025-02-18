package server

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func Start(config *Config) error {

	mux := http.NewServeMux()

	mux.HandleFunc("/whois", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HYYYYPE")
	})

	// используется rs/cors, только так cors разрешает post запрос с фронтенда
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{config.frontendUrl}, //адресса, имеющие доступ к серверу
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true, //для cookie (вроде)
	}).Handler(mux)

	return http.ListenAndServe(config.binaddr, corsHandler)
}
