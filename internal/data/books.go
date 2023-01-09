package data

import "time"

type Book struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"-"`
	GBooksVolID    string    `json:"gBooksVolId"`
	LastPageNumber int64     `json:"lastPageNumber"`
	StartDate      Date      `json:"startDate,omitempty"`
	FinishDate     Date      `json:"finishDate,omitempty"`
}
