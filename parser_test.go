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

package freddie

import (
	"github.com/nmeum/freddie/atom"
	"github.com/nmeum/freddie/rss"
	"testing"
)

type testpair struct {
	URL  string
	Type string
}

func TestParseFunc1(t *testing.T) {
	feed, err := ParseFunc("http://taz.de/rss.xml", rss.Parse)
	if err != nil {
		t.Fatal(err)
	}

	if feed.Title != "taz.de - taz.de" {
		t.Fatalf("Expected %q - got %q", "taz.de - taz.de", feed.Title)
	}
}

func TestParseFunc2(t *testing.T) {
	feed, err := ParseFunc("http://blog.golang.org/feed.atom", atom.Parse)
	if err != nil {
		t.Fatal(err)
	}

	if feed.Title != "The Go Programming Language Blog" {
		t.Fatalf("Expected %q - got %q", "The Go Programming Language Blog", feed.Title)
	}
}

func TestParse(t *testing.T) {
	tests := []testpair{
		{"http://cyber.law.harvard.edu/rss/examples/rss2sample.xml", "rss"},
		{"http://blog.case.edu/news/feed.atom", "atom"},
	}

	for _, test := range tests {
		feed, err := Parse(test.URL)
		if err != nil {
			t.Fatal(err)
		}

		if feed.Type != test.Type {
			t.Fatalf("Expected %q - got %q", test.Type, feed.Type)
		}
	}
}
