// Copyright (C) 2018  Sachin Saini

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"image/png"
	"os"
	"path"
)

// decode and compress png file
func decodeAndCompressPng(q png.CompressionLevel, f string, pwd string) {
	imgFile, err := os.Open(f)
	check(err)
	img, err := png.Decode(imgFile)
	check(err)

	imgC, err := os.Create(path.Join(pwd, "dist", f))
	check(err)
	enc := png.Encoder{CompressionLevel: q}
	err = enc.Encode(imgC, img)
	check(err)

	err = imgC.Close()
	check(err)

	err = imgFile.Close()
	check(err)
	fmt.Println("successfully compressed", f)
}
