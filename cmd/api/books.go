package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/faizanabbas/where-was-i-service/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new book")
}

func (app *application) readBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	book := data.Book{
		ID:             id,
		CreatedAt:      time.Now(),
		GBooksVolID:    "UvK1Slvkz3MC",
		LastPageNumber: 81,
		StartDate:      data.Date{Time: time.Now()},
		FinishDate:     data.Date{},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
