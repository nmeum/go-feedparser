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
	"time"
)

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
