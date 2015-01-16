package freddie

import (
	"github.com/nmeum/freddie/feed"
)

type byDate []feed.Item

func (b byDate) Len() int {
	return len(b)
}

func (b byDate) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byDate) Less(i, j int) bool {
	return b[i].Date.After(b[j].Date)
}
