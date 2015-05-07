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
	"net/http"
	"testing"
)

type testpair struct {
	URL  string
	Type string
}

func TestParse(t *testing.T) {
	tests := []testpair{
		{"http://cyber.law.harvard.edu/rss/examples/rss2sample.xml", "rss"},
		{"http://www.heise.de/developer/rss/news-atom.xml", "atom"},
		{"http://blog.case.edu/news/feed.atom", "atom"},
	}

	for _, test := range tests {
		resp, err := http.Get(test.URL)
		if err != nil {
			t.Fatal(err)
		}

		reader := resp.Body
		defer reader.Close()

		feed, err := Parse(reader)
		if err != nil {
			t.Fatal(err)
		}

		if feed.Type != test.Type {
			t.Fatalf("Expected %q - got %q", test.Type, feed.Type)
		}
	}
}
