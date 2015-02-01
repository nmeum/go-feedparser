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

package util

import (
	"testing"
)

func TestParseDate(t *testing.T) {
	testFormat := "Thu, 27 Feb 2014 18:46:18 +0100"
	var timestamp int64 = 1393523178

	date, err := ParseTime(testFormat)
	if err != nil {
		t.Fatal(err)
	}

	if date.Unix() != timestamp {
		t.Fatalf("Expected %d - got %d", timestamp, date.Unix())
	}
}
