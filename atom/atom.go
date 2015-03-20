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

package atom

import (
	"encoding/xml"
	"github.com/nmeum/go-feedparser"
	"github.com/nmeum/go-feedparser/util"
)

// Feed represents an ATOM feed.
type Feed struct {
	// XMLName.
	XMLName xml.Name `xml:"feed"`

	// Feed title.
	Title string `xml:"title"`

	// Feed links.
	Links []Link `xml:"link"`

	// Feed entries.
	Entries []Entry `xml:"entry"`
}

// Entry represents an ATOM feed entry.
type Entry struct {
	// Time the entry was published.
	Published string `xml:"published"`

	// Title of the entry.
	Title string `xml:"title"`

	// Links for the entry.
	Links []Link `xml:"link"`
}

// Link represents an ATOM feed link.
type Link struct {
	// Link type.
	Type string `xml:"type,attr"`

	// Link URL.
	Href string `xml:"href,attr"`

	// Link rel.
	Rel string `xml:"rel,attr"`
}

// Parse parses an ATOM feed. It implements feedparser.FeedFunc.
func Parse(data []byte) (f feedparser.Feed, err error) {
	var atom Feed
	if err = util.Unmarshal(data, &atom); err != nil {
		return
	}

	f = feedparser.Feed{
		Title: atom.Title,
		Type:  "atom",
		Link:  findLink(atom.Links).Href,
	}

	for _, e := range atom.Entries {
		item := feedparser.Item{
			Title:      e.Title,
			Link:       findLink(e.Links).Href,
			Attachment: findAttachment(e.Links).Href,
		}

		item.Date, err = util.ParseTime(e.Published)
		if err != nil {
			return
		}

		f.Items = append(f.Items, item)
	}

	return
}

// findLink attempts to find the most relevant link. This is necessary
// because the generic feedparser.Feed struct doesn't support more than
// one link.
func findLink(links []Link) Link {
	var score int
	var match Link

	for _, link := range links {
		switch {
		case link.Rel == "alternate" && link.Type == "text/html":
			return link
		case score < 3 && link.Type == "text/html":
			score = 3
			match = link
		case score < 2 && link.Rel == "self":
			score = 2
			match = link
		case score < 1 && link.Rel == "":
			score = 1
			match = link
		case match == nil:
			match = link
		}
	}

	return match
}

// findLink attempts to find a link which represents an attachment.
func findAttachment(links []Link) Link {
	for _, link := range links {
		if link.Rel == "enclosure" {
			return link
		}
	}

	return Link{}
}
