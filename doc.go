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

/*
Package go-feedparser implements a simple RSS and ATOM feed parser.

Tho primary function of interest is the Parse function. You can pass an
arbitrary Reader to this function and it will return the corresponding
feed. The following demonstrates and example use case (reading a feed
from a file):

	file, err := os.Open("feed.xml");
	if err != nil {
		panic(err)
	}
	defer file.Close()

	feed, err := feedparser.Parse(file);
	if err != nil {
		panic(err)
	}

	switch (feed.Type) {
	case "rss":
		fmt.Println("RSS feed!")
	case "atom":
		fmt.Println("ATOM feed!")
	default:
		fmt.Println("Unknown feed format")
	}
*/
package feedparser
