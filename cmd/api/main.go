package main

import (
	"fmt"
	"net/http"

	"github.com/Omar-Gebal/go-simple-api/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	const PORT string = "localhost:8000"
	fmt.Println("Starting GO API service...")
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Printf("API is running at: http://%s\n", PORT)
	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Error(err)
	}
}
