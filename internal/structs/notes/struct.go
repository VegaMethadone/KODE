package notes

import "time"

type Note struct {
	Id    int       `json:"id"`
	Owner int       `json:"owner"`
	Date  time.Time `json:"date"`
	Title string    `json:"title"`
	Body  string    `json:"body"`
}
