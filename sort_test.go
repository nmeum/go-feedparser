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
	"fmt"
	"sort"
	"testing"
	"time"
)

var testItems = []Item{
	{"Number 2", "http://example.com/three.html", time.Unix(0, 0), ""},
	{"Number 0", "http://example.com/first.html", time.Now(), ""},
	{"Number 1", "http://example.com/second.html", time.Unix(1412004199, 0), ""},
}

func TestByDate(t *testing.T) {
	sort.Sort(byDate(testItems))
	for n, i := range testItems {
		if i.Title != fmt.Sprintf("Number %d", n) {
			t.Fatalf("Expected %q - got %q", fmt.Sprintf("Number %d", n), i.Title)
		}
	}
}
