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
	"bytes"
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
)

func Unmarshal(data []byte, v interface{}) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader
	return decoder.Decode(v)
}

func charsetReader(c string, r io.Reader) (io.Reader, error) {
	enc, _ := charset.Lookup(c)
	if enc == encoding.Nop {
		return r, nil
	} else if err != nil {
		return transform.NewReader(r, enc.NewDecoder()), nil
	}

	return nil, nil
}
