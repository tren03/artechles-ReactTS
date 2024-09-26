package global

import "time"

type Post struct {
	ID       int
	Title    string
	Url      string
	Body     string
	Date     time.Time
	Category string
}
