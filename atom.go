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

package feedparser

import (
	"encoding/xml"
)

// AtomFeed represents an atom web feed.
type AtomFeed struct {
	// XMLName.
	XMLName xml.Name `xml:"feed"`

	// Universally unique feed ID (required).
	ID string `xml:"id"`

	// Human readable title for the feed (required).
	Title AtomText `xml:"title"`

	// Last time the feed was significantly modified (required).
	Updated string `xml:"updated"`

	// Entries for the feed (required).
	Entries []AtomEntry `xml:"entry"`

	// Authors of the feed (recommended).
	Authors []AtomPerson `xml:"author"`

	// Links which identify related web pages (recommended).
	Links []AtomLink `xml:"link"`

	// Categories the feed belongs to (optional).
	Categories []AtomCategory `xml:"category"`

	// Contributors to the feed (optional).
	Contributors []AtomPerson `xml:"contributor"`

	// Software used to generate the feed (optional).
	Generator AtomGenerator `xml:"generator"`

	// Small icon used for visual identification (optional).
	Icon string `xml:"icon"`

	// Larger logo for visual identification (optional).
	Logo string `xml:"logo"`

	// Information about rights, for example copyrights (optional).
	Rights AtomText `xml:"rights"`

	// Human readable description or subtitle (optional).
	Subtitle AtomText `xml:"subtitle"`
}

// AtomEntry represents an atom entry.
type AtomEntry struct {
	// Universally unique feed ID (required).
	ID string `xml:"id"`

	// Human readable title for the entry (required).
	Title AtomText `xml:"title"`

	// Last time the feed was significantly modified (required).
	Updated string `xml:"updated"`

	// Authors of the entry (recommended).
	Authors []AtomPerson `xml:"author"`

	// Content of the entry (recommended).
	Content AtomText `xml:"content"`

	// Links which identify related web pages (recommended).
	Links []AtomLink `xml:"link"`

	// Short summary, abstract or excerpt of the entry (recommended).
	Summary AtomText `xml:"summary"`

	// Categories the entry belongs too (optional).
	Categories []AtomCategory `xml:"category"`

	// Contributors to the entry (optional).
	Contributors []AtomPerson `xml:"contributor"`

	// Time of the initial creation of the entry (optional).
	Published string `xml:"published"`

	// FIXME
	// Feed's metadata, only used when entry was copied from another feed (optional).
	// Source AtomFeed `xml:"source"`

	// Information about rights, for example copyrights (optional).
	Rights AtomText `xml:"rights"`
}

// AtomLink represents the atom link tag.
type AtomLink struct {
	// Hypertext reference (required).
	Href string `xml:"href,attr"`

	// Single Link relation type (optional).
	Rel string `xml:"rel,attr"`

	// Media type of the resource (optional).
	Type string `xml:"type,attr"`

	// Language of referenced resource (optional).
	HrefLang string `xml:"hreflang,attr"`

	// Human readable information about the link (optional).
	Title string `xml:"title,attr"`

	// Length of the resource in bytes (optional).
	Length string `xml:"length,attr"`
}

// AtomPerson represents a person, corporation, et cetera.
type AtomPerson struct {
	// Human readable name for the person (required).
	Name string `xml:"name"`

	// Home page for the person (optional).
	URI string `xml:"uri"`

	// Email address for the person (optional).
	Email string `xml:"email"`
}

// AtomCategory identifies the category.
type AtomCategory struct {
	// Identifier for this category (required).
	Term string `xml:"term,attr"`

	// Categorization scheme via a URI (optional).
	Scheme string `xml:"scheme,attr"`

	// Human readable label for display (optional).
	Label string `xml:"label,attr"`
}

// AtomGenerator identifies the generator.
type AtomGenerator struct {
	// Generator name (required).
	Name string `xml:",chardata"`

	// URI for this generator (optional).
	URI string `xml:"uri,attr"`

	// Version for this generator (optional).
	Version string `xml:"version,attr"`
}

// AtomText identifies human readable text.
type AtomText struct {
	// Text body (required).
	Body string `xml:",chardata"`

	// InnerXML data (optional).
	InnerXML string `xml:",innerxml"`

	// Text type (optional).
	Type string `xml:"type,attr"`

	// URI where the content can be found (optional for <content>).
	URI string `xml:"uri,att"`
}

// parseAtom parses an atom feed and returns a generic feed.
func parseAtom(data []byte) (f Feed, err error) {
	var origFeed AtomFeed
	if err = unmarshal(data, &origFeed); err != nil {
		return
	}

	f = Feed{
		Type:        "atom",
		Title:       origFeed.Title.Body,
		Link:        findLink(origFeed.Links).Href,
		Description: origFeed.Subtitle.Body,
		Image:       origFeed.Logo,
		Generator:   origFeed.Generator.Name,
		Rights:      origFeed.Rights.Body,
	}

	if len(origFeed.Authors) > 0 {
		f.Author = origFeed.Authors[0].Email
	}

	f.Updated, err = parseTime(origFeed.Updated)
	if err != nil {
		return
	}

	for _, category := range origFeed.Categories {
		f.Categories = append(f.Categories, category.Term)
	}

	for _, entry := range origFeed.Entries {
		item := Item{
			ID:         entry.ID,
			Title:      entry.Title.Body,
			Link:       findLink(entry.Links).Href,
			Content:    entry.Content.Body,
			Attachment: findAttachment(entry.Links).Href,
		}

		if len(entry.Authors) > 0 {
			item.Author = entry.Authors[0].Email
		}

		for _, category := range entry.Categories {
			item.Categories = append(item.Categories, category.Term)
		}

		timeStr := entry.Updated
		if len(entry.Published) > 0 {
			timeStr = entry.Published
		}

		item.PubDate, err = parseTime(timeStr)
		if err != nil {
			return
		}

		f.Items = append(f.Items, item)
	}

	return
}

// findLink attempts to find the most relevant link.
func findLink(links []AtomLink) AtomLink {
	var score int
	var match AtomLink

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
		case &match == nil:
			match = link
		}
	}

	return match
}

// findAttachment attempts to find a link which represents an attachment.
func findAttachment(links []AtomLink) AtomLink {
	for _, link := range links {
		if link.Rel == "enclosure" {
			return link
		}
	}

	return AtomLink{}
}
