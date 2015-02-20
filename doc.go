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
Package freddie implements a simple library for parsing feeds.

The only function of direct interest is probably the Parse function
besides an arbitrary byte slice it also requires a FeedFunc as an
argument.

A FeedFunc can of cause be implemented by the caller itself but
FeedFuncs for the most common feed Formats (RSS and ATOM) are already
shipped by this package, however, the those FeedFuncs have to be
imported manually. If you want to parse RSS and ATOM feeds consider the
following example:

	import (
		// ...
		"github.com/nmeum/go-feedparser"
		"github.com/nmeum/go-feedparser/atom"
		"github.com/nmeum/go-feedparser/rss"
		"os"
		// ...
	)

	parsers := []feedparser.FeedFunc{rss.Parse, atom.Parse}
	file, err := os.Open("feed.xml")
	if err != nil {
		panic(err)
	}

	feed, err := feedparser.Parse(file, parsers)
	if err != nil {
		panic(err)
	}

The Parse function will automatically chose the appropriate FeedFunc for
the given reader. However, if you already know which FeedFunc will be
required to parse the feed you can only pass the required one to the
Parse function.

If you need specific feed informations which are not part of the generic
feedparser.Feed struct that you can also call the appropriate FeedFunc
directly. Consider the following example for an ATOM feed:

	import (
		// ...
		"github.com/nmeum/go-feedparser/atom"
		"io/ioutil"
		"os"
		// ...
	)

	data, err := ioutil.ReadFile("feed.atom")
	if err != nil {
		panic(err)
	}

	atomFeed, err := atom.Parse(data)
	if err != nil {
		panic(err)
	}

The ATOM and RSS FeedFuncs can also be used to create ATOM / RSS feeds.
However, this does not apply to all FeedFuncs and is therefore
documented in the ATOM / RSS packages.

If you want to implement your own FeedFunc take a look at the util
subpackage. An example FeedFunc can be found in the ATOM and RSS
subpackages.
*/
package feedparser
