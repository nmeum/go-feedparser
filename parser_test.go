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
		{"http://www.heise.de/developer/rss/news-atom.xml", "atom"},
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
