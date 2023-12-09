package main

import (
	"log"
    "net/http"

    // rename main to package name
    "main/routes"
    
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)


func main() {
    var PORT string = ":5000"
    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    routes.ConfigureRoutes(r)

    log.Println("Starting server on port ", PORT)
    http.ListenAndServe(PORT, r)
}
