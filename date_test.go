package freddie

import (
	"testing"
	"time"
)

func TestParseDatePositive(t *testing.T) {
	testDate := "Thu, 27 Feb 2014 18:46:18 +0100"
	var timestamp int64 = 1393523178

	date, err := parseDate(testDate)
	if err != nil {
		t.Fatal(err)
	}

	if date.Unix() != timestamp {
		t.Fatalf("Expected %d - got %d", timestamp, date.Unix())
	}
}

func TestParseDateNegative(t *testing.T) {
	_, err := parseDate("Die, 42 Feb 1842")
	if _, ok := err.(*time.ParseError); !ok {
		t.Fatal(err)
	}
}
