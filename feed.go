package freddie

import (
	"time"
)

type Feed struct {
	Title string
	Type  string
	Link  string
	Items []Item
}

type Item struct {
	Title      string
	Link       string
	Date       time.Time
	Attachment string
}
