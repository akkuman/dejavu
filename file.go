// DejaVu - Data snapshot and sync.
// Copyright (c) 2022-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package dejavu

import (
	"bytes"
	"strconv"
)

type File struct {
	Hash    string   `json:"hash"`
	Path    string   `json:"path"`
	Size    int64    `json:"size"`
	Updated int64    `json:"updated"`
	Body    []string `json:"body"` // Chunk IDs
}

func (f *File) ID() string {
	if "" != f.Hash {
		return f.Hash
	}

	buf := bytes.Buffer{}
	buf.WriteString(f.Path)
	buf.WriteString(strconv.FormatInt(f.Size, 10))
	buf.WriteString(strconv.FormatInt(f.Updated, 10))
	for _, c := range f.Body {
		buf.WriteString(c)
	}
	f.Hash = Hash(buf.Bytes())
	return f.Hash
}
