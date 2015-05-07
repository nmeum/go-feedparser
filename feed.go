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
	"io"
	"io/ioutil"
	"sort"
	"time"
)

// parseFunc describes a function which implements a feed parser.
type parseFunc func([]byte) (Feed, error)

// parsers lists all known feed parsers.
var parsers = []parseFunc{parseAtom, parseRss}

// Feed represents a generic feed.
type Feed struct {
	// Title for the feed.
	Title string

	// Feed type (either atom or rss).
	Type string

	// URL to the website.
	Link string

	// Description or subtitle for the feed.
	Description string

	// Categories the feed belongs to.
	Categories []string

	// Email address of the feed author.
	Author string

	// Last time the feed was updated.
	Updated time.Time

	// URL to image for the feed.
	Image string

	// Software used to generate the feed.
	Generator string

	// Information about rights, for example copyrights.
	Rights string

	// Feed Items
	Items []Item
}

// Item represents a generic feed item.
type Item struct {
	// Universally unique item ID.
	ID string

	// Title of the item.
	Title string

	// URL for the item.
	Link string

	// Content of the item.
	Content string

	// Email address of the item author.
	Author string

	// Categories the item belongs to.
	Categories []string

	// Time the item was published.
	PubDate time.Time

	// URL to media attachment.
	Attachment string
}

// Parse tries to parse the content of the given reader. It also sorts all items
// by there publication date. Meaning that the first item is guaranteed to be
// the most recent one.
func Parse(r io.Reader) (f Feed, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	for _, p := range parsers {
		f, err = p(data)
		if err == nil {
			break
		}
	}

	if err != nil {
		return
	}

	sort.Sort(byDate(f.Items))
	return
}
