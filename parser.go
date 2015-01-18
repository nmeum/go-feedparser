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
	"github.com/nmeum/freddie/feed"
	"github.com/nmeum/freddie/rss"
	"io/ioutil"
	"net/http"
	"sort"
)

type FeedFunc func([]byte) (feed.Feed, error)

var parsers = []FeedFunc{
	rss.Parse,
	atom.Parse,
}

func ParseFunc(url string, fn FeedFunc) (f feed.Feed, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	f, err = fn(body)
	if err != nil {
		return
	}

	sort.Sort(byDate(f.Items))
	return
}

func Parse(url string) (f feed.Feed, err error) {
	for _, p := range parsers {
		f, err = ParseFunc(url, p)
		if err == nil {
			break
		}
	}

	return
}
