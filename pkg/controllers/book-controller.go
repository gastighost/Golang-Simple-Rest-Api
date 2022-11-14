package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gastighost/book-management-system/pkg/models"
	"github.com/gastighost/book-management-system/pkg/utils"

	"github.com/go-chi/chi/v5"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		utils.CreateJson(w, http.StatusBadRequest, "There was an input error", []string{})
		return
	}

	book, err := newBook.Create()
	if err != nil {
		utils.CreateJson(w, http.StatusBadRequest, err.Error(), []string{})
		return
	}
	utils.CreateJson(w, http.StatusCreated, "Book successfully created!", book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	utils.CreateJson(w, http.StatusOK, "All books successfully retrieved!", books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	n, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		utils.CreateJson(w, http.StatusBadRequest, "There is an error with your Id", []string{})
		return
	}

	book, _ := models.GetBookById(n)
	utils.CreateJson(w, http.StatusOK, "Book successfully retrieved!", book)
}

func EditBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	n, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		utils.CreateJson(w, http.StatusBadRequest, "There is an error with your Id", []string{})
		return
	}

	var bookUpdating models.Book

	err = json.NewDecoder(r.Body).Decode(&bookUpdating)
	if err != nil {
		utils.CreateJson(w, http.StatusBadRequest, "There was an input error", []string{})
		return
	}

	bookDetails, db := models.GetBookById(n)

	if bookUpdating.Name != "" {
		bookDetails.Name = bookUpdating.Name
	}

	if bookUpdating.Author != "" {
		bookDetails.Author = bookUpdating.Author
	}

	if bookUpdating.Publication != "" {
		bookDetails.Publication = bookUpdating.Publication
	}

	db.Save(&bookDetails)

	utils.CreateJson(w, http.StatusOK, "Book successfully updated!", bookDetails)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "bookId")
	n, err := strconv.ParseInt(bookId, 10, 64)

	if err != nil {
		utils.CreateJson(w, http.StatusBadRequest, "There is an error with your Id", []string{})
		return
	}

	book := models.DeleteBook(n)

	utils.CreateJson(w, http.StatusOK, "Book was successfully deleted!", book)
}
