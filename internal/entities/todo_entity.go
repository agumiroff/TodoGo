package entities

import "time"

type Todo struct {
	ID    int64     `db:"id" json:"id"`
	Title string    `db:"head" json:"title"`
	Todo  string    `db:"todo" json:"todo"`
	Date  time.Time `db:"created_at" json:"date"`
}
