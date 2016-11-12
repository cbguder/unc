package models

import "time"

type Note struct {
	Title    string
	Body     string
	Tags     []string
	Created  time.Time
	Modified time.Time
}
