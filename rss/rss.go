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

package rss

import (
	"encoding/xml"
	"github.com/nmeum/freddie"
)

type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Title   string   `xml:"channel>title"`
	Link    string   `xml:"channel>link"`
	Items   []Item   `xml:"channel>item"`
}

type Item struct {
	PubDate   string    `xml:"pubDate"`
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Type string `xml:"type,attr"`
	URL  string `xml:"url,attr"`
}

func Parse(data []byte) (f freddie.Feed, err error) {
	var rss Feed
	if err = xml.Unmarshal(data, &rss); err != nil {
		return
	}

	f = freddie.Feed{
		Title: rss.Title,
		Type:  "rss",
		Link:  rss.Link,
	}

	for _, i := range rss.Items {
		item := freddie.Item{
			Title:      i.Title,
			Link:       i.Link,
			Attachment: i.Enclosure.URL,
		}

		item.Date, err = freddie.ParseTime(i.PubDate)
		if err != nil {
			panic(err)
		}

		f.Items = append(f.Items, item)
	}

	return
}
