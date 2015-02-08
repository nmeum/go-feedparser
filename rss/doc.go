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

package rss

/*
Package rss implements a FeedFunc for RSS feeds.

This package can also be used to create RSS feeds. Consider the
following example to do so:

	items := []rss.Item{
		{"Monday, January 2, 2006 15:04:05 MST", "Bar", "http://example.org/bar", ""},
		{"Monday, August 9, 2012 11:02:23 MST", "Foo", "http://example.org/foo", ""},
	}

	feed := rss.Feed{
		Title: "Foobar",
		Link: "http://example.org",
		Items: items,
	}

	output, err := xml.MarshalIndent(feed, "  ", "    ")
	if err != nil {
		panic(err)
	}

	os.Stdout.Write(output)
*/
