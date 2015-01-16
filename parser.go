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
			return
		}
	}

	return
}
