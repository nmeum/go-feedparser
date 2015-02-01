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
	"encoding/xml"
	"io/ioutil"
	"os"
	"testing"
)

type testpair struct {
	XMLName xml.Name `xml:"g"`
	Group   struct {
		Char string `xml:"a,attr"`
	} `xml:"f"`
}

func TestUnmarshal(t *testing.T) {
	file, err := os.Open("testdata/iso-8859-1.xml")
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	var doc testpair
	if err := Unmarshal(data, &doc); err != nil {
		t.Fatal(err)
	}

	if doc.Group.Char != "é" {
		t.Fatalf("Expected %q - got %q", "é", doc.Group.Char)
	}
}
