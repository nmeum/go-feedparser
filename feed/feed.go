// This file is part of Freddie.
//
// Freddie is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Freddie is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Freddie. If not, see <http://www.gnu.org/licenses/>.

package feed

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

var dateFormats = []string{
	time.RFC1123Z, time.RFC1123, time.RFC822Z,
	time.RFC822, time.ANSIC, time.RFC3339,
	time.RFC850, time.RubyDate, time.UnixDate,
	"2 January 2006 15:04:05 -0700", "2 January 2006 15:04:05 MST",
	"2 Jan 2006 15:04:05 -0700", "2 Jan 2006 15:04:05 MST",
	"Mon, 2 Jan 2006 15:04:05 -0700", "Mon, 2 Jan 2006 15:04:05 MST",
	"2006-01-02T15:04:05", "2006-01-02T15:04:05Z",
}

func ParseDate(data string) (date time.Time, err error) {
	for _, format := range dateFormats {
		date, err = time.Parse(format, data)
		if err == nil {
			return
		}
	}

	return
}
