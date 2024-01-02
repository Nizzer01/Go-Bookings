package main

import (
	"github.com/Nizzer01/Go-Bookings/pkg/config"
	"github.com/Nizzer01/Go-Bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	//Create file server to get static files (images, js etc
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
