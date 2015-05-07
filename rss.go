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
//
// This is a slightly modified version of 'encoding/xml/read_test.go'.
// Copyright 2009 The Go Authors. All rights reserved. Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package feedparser

import (
	"encoding/xml"
)

// RssFeed represents an rss web feed.
type RssFeed struct {
	// XMLName.
	XMLName xml.Name `xml:"rss"`

	// Name of the channel (required).
	Title string `xml:"channel>title"`

	// URL to the website (required).
	Link string `xml:"channel>link"`

	// Description for the channel (required).
	Description string `xml:"channel>description"`

	// Items for the feed (required).
	Items []RssItem `xml:"channel>item"`

	// Language the channel is written in (optional).
	Language string `xml:"channel>language"`

	// Copyright notice for the content (optional).
	Copyright string `xml:"channel>copyright"`

	// Email address of the editor (optional).
	Editor string `xml:"channel>managingEditor"`

	// Email address of the web master (optional).
	WebMaster string `xml:"channel>webMaster"`

	// Publication date for the content (optional).
	PubDate string `xml:"channel>pubDate"`

	// Last time the content was updated (optional).
	LastBuildDate string `xml:"channel>lastBuildDate"`

	// Categories the feed belongs to (optional).
	Categories []RssCategory `xml:"channel>category"`

	// Program used to generate the channel (optional).
	Generator string `xml:"channel>generator"`

	// URL that points to documentation for the used format (optional).
	Docs string `xml:"channel>docs"`

	// Cloud for update notifications (optional).
	Cloud RssCloud `xml:"channel>cloud"`

	// How long the channel can be cached (optional).
	TTL int `xml:"channel>ttl"`

	// Image that can be displayed with the channel (optional).
	Image RssImage `xml:"channel>image"`

	// PICS rating for the channel (optional).
	Rating string `xml:"channel>rating"`

	// Text input box related to the channel (optional).
	TextInput RssTextInput `xml:"channel>textInput"`

	// Hint for aggregators telling them which hours can be skipped (optional).
	SkipHours []RssHour `xml:"channel>skipHours"`

	// Hint for aggregators telling them which days can be skipped (optional).
	SkipDays []RssDay `xml:"channel>skipDays"`
}

// RssItem represents an rss item.
type RssItem struct {
	// Title of the item (required if description isn't present).
	Title string `xml:"title"`

	// The item synopsis (required if title isn't present).
	Description string `xml:"description"`

	// The URL of the item (optional).
	Link string `xml:"link"`

	// Email address of the author of the item (optional).
	Author string `xml:"author"`

	// Includes item in one or more categories (optional).
	Categories []RssCategory `xml:"category"`

	// URL to a page for comments (optional).
	Comments string `xml:"comments"`

	// Media object that is attached to the item (optional).
	Enclosure string `xml:"enclosure"`

	// String that uniquely identifies the item (optional).
	GUID string `xml:"guid"`

	// Time the item was published (optional).
	PubDate string `xml:"pubDate"`

	// The RSS channel the item came from (optional).
	Source RssSource `xml:"source"`
}

// RssImage represents an rss image.
type RssImage struct {
	// URL to image that represents the channel (required).
	URL string `xml:"url"`

	// Title which describes the image (required).
	Title string `xml:"title"`

	// URL of the site itself (required).
	Link string `xml:"link"`

	// Width of the image (optional).
	Width int `xml:"width"`

	// Height of the image (optional).
	Height int `xml:"height"`

	// Additional description of the image (optional).
	Description string `xml:"description"`
}

// RssCloud represents the rss cloud tag.
type RssCloud struct {
	// Domain cloud service is running on (required).
	Domain string `xml:"domain,attr"`

	// Port to use for TCP socket connection (required).
	Port int `xml:"port,attr"`

	// Path to use for the request (required).
	Path string `xml:"path,attr"`

	// Register procedure which should be used (required).
	RegisterProcedure string `xml:"registerProcedure,attr"`

	// Protocol used for registration et cetera (required).
	Protocol string `xml:"protocol,attr"`
}

// RssCategory represents the rss category tag.
type RssCategory struct {
	// Human readable category name (required).
	Name string `xml:",chardata"`

	// Domain that identifies categorization taxonomy (optional).
	Domain string `xml:"domain,attr"`
}

// RssTextInput represents the rss textInput tag.
type RssTextInput struct {
	// The label of the Submit button in the text input area (required).
	Title string `xml:"title"`

	// Explains the text input area (required).
	Description string `xml:"description"`

	// The name of the text object in the text input area (required).
	Name string `xml:"name"`

	// The URL of the CGI script that processes text input requests (required).
	Link string `xml:"link"`
}

// RssSource represents the rss source tag.
type RssSource struct {
	// URL which links to the XMLization source (required).
	URL string `xml:"url,attr"`

	// Source name (required).
	Name string `xml:",chardata"`
}

// RssHour represents the hour tag, a subelement of the skipHours tag.
type RssHour struct {
	// Number between 0 and 23 representing time in GMT (required).
	Hour string `xml:"hour"`
}

// RssDay represents the day tag, a subelement of the skipDays tag.
type RssDay struct {
	// Weekday (e.g Monday) (required).
	Day string `xml:"day"`
}

// parseRss parses an rss feed and returns a generic feed.
func parseRss(data []byte) (f Feed, err error) {
	var origFeed RssFeed
	if err = unmarshal(data, &origFeed); err != nil {
		return
	}

	f = Feed{
		Type:        "rss",
		Title:       origFeed.Title,
		Link:        origFeed.Link,
		Description: origFeed.Description,
		Image:       origFeed.Image.URL,
		Generator:   origFeed.Generator,
		Rights:      origFeed.Copyright,
		Author:      origFeed.Editor,
	}

	f.Updated, err = parseTime(origFeed.LastBuildDate)
	if err != nil {
		return
	}

	for _, category := range origFeed.Categories {
		f.Categories = append(f.Categories, category.Name)
	}

	for _, entry := range origFeed.Items {
		item := Item{
			ID:         entry.GUID,
			Title:      entry.Title,
			Link:       entry.Link,
			Content:    entry.Description,
			Attachment: entry.Enclosure,
			Author:     entry.Author,
		}

		for _, category := range entry.Categories {
			item.Categories = append(item.Categories, category.Name)
		}

		item.PubDate, err = parseTime(entry.PubDate)
		if err != nil {
			return
		}
	}

	return
}
