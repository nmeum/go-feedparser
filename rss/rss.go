// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package rss

import (
	"encoding/xml"
	"github.com/nmeum/go-feedparser"
	"github.com/nmeum/go-feedparser/util"
)

// Feed represents an RSS feed.
type Feed struct {
	// XMLName.
	XMLName xml.Name `xml:"rss"`

	// Feed Title.
	Title string `xml:"channel>title"`

	// Link for the feed.
	Link string `xml:"channel>link"`

	// Feed items.
	Items []Item `xml:"channel>item"`
}

// Item represents an RSS item.
type Item struct {
	// Time the Item was published.
	PubDate string `xml:"pubDate"`

	// Item title.
	Title string `xml:"title"`

	// Link for the Item.
	Link string `xml:"link"`

	// Enclosure for the Item.
	Enclosure Enclosure `xml:"enclosure"`
}

// Enclosure represents an RSS enclosure.
type Enclosure struct {
	// Enclosure type.
	Type string `xml:"type,attr"`

	// Enclosure URL.
	URL string `xml:"url,attr"`
}

// Parse parses an RSS feed. It implements feedparser.FeedFunc.
func Parse(data []byte) (f feedparser.Feed, err error) {
	var rss Feed
	if err = util.Unmarshal(data, &rss); err != nil {
		return
	}

	f = feedparser.Feed{
		Title: rss.Title,
		Type:  "rss",
		Link:  rss.Link,
	}

	for _, i := range rss.Items {
		item := feedparser.Item{
			Title:      i.Title,
			Link:       i.Link,
			Attachment: i.Enclosure.URL,
		}

		item.Date, err = util.ParseTime(i.PubDate)
		if err != nil {
			return
		}

		f.Items = append(f.Items, item)
	}

	return
}
