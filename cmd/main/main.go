package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gastighost/book-management-system/pkg/routes"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Group(routes.RegisterBookStoreRoutes)

	fmt.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
