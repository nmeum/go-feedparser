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

// byDate sorts a generic Item slice by the items date attribute thus
// sorting the items by the date they were published. It implements the
// sort.Interface interface.
type byDate []*Item

func (b byDate) Len() int {
	return len(b)
}

func (b byDate) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byDate) Less(i, j int) bool {
	return b[i].PubDate.After(b[j].PubDate)
}
