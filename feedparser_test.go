package feedparser

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var (
	parsers = []FeedFunc{testParser}
	items   = []Item{
		{Title: "1", Date: time.Unix(0, 0)},
		{Title: "0", Date: time.Now()},
	}
)

func testParser(data []byte) (f Feed, err error) {
	return Feed{Title: string(data), Items: items}, nil
}

func TestParse(t *testing.T) {
	fp := filepath.Join("testdata", "testParse.txt")
	file, err := os.Open(fp)
	if err != nil {
		t.Fatal(err)
	}

	feed, err := Parse(file, parsers)
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello World!\n"
	if feed.Title != expected {
		t.Fatalf("Expected %q - got %q", expected, feed.Title)
	}

	if len(feed.Items) != 2 {
		t.Fatalf("Expected %d - got %d", 2, len(feed.Items))
	}

	for i, item := range feed.Items {
		stringIndex := fmt.Sprintf("%d", i)
		if stringIndex != item.Title {
			t.Fatalf("Expected %q - got %q", item.Title, stringIndex)
		}
	}
}
