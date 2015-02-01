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

type FeedFunc func([]byte) (Feed, error)

type Feed struct {
	Title string
	Type  string
	Link  string
	Items []Item
}

type Item struct {
	Title      string
	Link       string
	Date       time.Time
	Attachment string
}

type byDate []Item

func (b byDate) Len() int           { return len(b) }
func (b byDate) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byDate) Less(i, j int) bool { return b[i].Date.After(b[j].Date) }

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
