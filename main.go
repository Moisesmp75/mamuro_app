package main

import (
	"embed"
	"fmt"
	"io/fs"
	"mamuro_app/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

//go:embed view/dist
var embedFrontend embed.FS

func static(r *chi.Mux) {
	front, err := fs.Sub(embedFrontend, "view/dist")
	if err != nil {
		fmt.Println(err.Error())
	}

	staticServer := http.FileServer(http.FS(front))

	r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		staticServer.ServeHTTP(w, r)
	}))
}

func addCors(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	}))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json", "text/xml"))

	addCors(r)
	static(r)

	apiv1 := chi.NewRouter()

	controllers.AddControllers(apiv1)
	r.Mount("/api/v1", apiv1)

	fmt.Println("Running on port: 3000")
	http.ListenAndServe(fmt.Sprintf(":%v", 3000), r)
}
