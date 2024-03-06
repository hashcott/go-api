package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hashcott/go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go API service...")
	fmt.Println(`
	______     ______        ______     ______   __    
	/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
	\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
	 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
		\/_____/   \/_____/      \/_/\/_/   \/_/     \/_/
	`)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}

}
