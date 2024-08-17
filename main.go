package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ikechukwu-peter/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

const port int = 8080

func main() {

	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handlers(r)

	fmt.Println("Starting Go API service...")
	fmt.Printf("Server started on port: %d", port)

	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", port), r)

	if err != nil {
		log.Error(err)
		return
	}

}
