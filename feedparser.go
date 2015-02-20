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
	"io"
	"io/ioutil"
	"sort"
	"time"
)

// FeedFunc describes a function which implements a feed parser,
// a FeedFunc should take a byte slice as an argument and should return
// a generic Feed struct and and error. If the error is nil it is
// assumed that the feed was parsed successfully.
type FeedFunc func([]byte) (Feed, error)

// Feed represents a generic feed.
type Feed struct {
	// The feed title.
	Title string

	// The feed type, should be the name of the feed standard.
	// For example "atom" or "rss".
	Type string

	// The feed link.
	Link string

	// Feed items.
	Items []Item
}

// Item represents generic feed items.
type Item struct {
	// Title of the item.
	Title string

	// Link for the item.
	Link string

	// Time the item was created.
	Date time.Time

	// Attachment (if any).
	Attachment string
}

// byDate sorts a generic Item slice by the time there were published.
type byDate []Item

// Functions below are required to implement sort.Interface.
func (b byDate) Len() int           { return len(b) }
func (b byDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byDate) Less(i, j int) bool { return b[i].Date.After(b[j].Date) }

// Parse parses the given reader. It invokes each given FeedFunc and if
// a FeedFunc returns no error the yielded feed is returned.
func Parse(r io.Reader, funcs []FeedFunc) (f Feed, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	for _, fn := range funcs {
		f, err = fn(data)
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
