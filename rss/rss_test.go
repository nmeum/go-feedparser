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
	"io/ioutil"
	"testing"
)

var rssFeed Feed

func TestFeed(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/feed.rss")
	if err != nil {
		t.Fatal(err)
	}

	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		t.Fatal(err)
	}

	if rssFeed.Title != "Some Title" {
		t.Fatalf("Expected %q - got %q", "Some Title", rssFeed.Title)
	}

	if rssFeed.Link != "http://example.org" {
		t.Fatalf("Expected %q - got %q", "http://example.org", rssFeed.Link)
	}
}

func TestItem(t *testing.T) {
	item := rssFeed.Items[0]

	if item.PubDate != "Tue, 20 May 2003 08:56:02 GMT" {
		t.Fatalf("Expected %q - got %q", "Tue, 20 May 2003 08:56:02 GMT", item.PubDate)
	}

	if item.Title != "Test Post" {
		t.Fatalf("Expected %q - got %q", "Test Post", item.Title)
	}

	if item.Link != "http://example.org/posts/test.html" {
		t.Fatalf("Expected %q - got %q", "http://example.org/posts/test.html", item.Link)
	}
}

func TestEnclosure(t *testing.T) {
	enclosure := rssFeed.Items[0].Enclosure

	if enclosure.Type != "audio/ogg" {
		t.Fatalf("Expected %q - got %q", "audio/ogg", enclosure.Type)
	}

	if enclosure.URL != "http://example.org/posts/test.ogg" {
		t.Fatalf("Expected %q - got %q", "http://example.org/posts/test.ogg", enclosure.URL)
	}
}
