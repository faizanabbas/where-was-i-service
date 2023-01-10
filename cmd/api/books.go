package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/faizanabbas/where-was-i-service/internal/data"
)

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		GBooksVolID    string `json:"gBooksVolId"`
		LastPageNumber int64  `json:"lastPageNumber"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) readBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		app.serverErrorResponse(w, r, err)
	}
}
