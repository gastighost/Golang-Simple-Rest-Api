package routes

import (
	"github.com/gastighost/book-management-system/pkg/controllers"
	"github.com/go-chi/chi/v5"
)

var RegisterBookStoreRoutes = func(r chi.Router) {
	r.Post("/books", controllers.CreateBook)
	r.Get("/books", controllers.GetBooks)
	r.Get("/books/{bookId}", controllers.GetBook)
	r.Patch("/books/{bookId}", controllers.EditBook)
	r.Delete("/books/{bookId}", controllers.DeleteBook)
}
