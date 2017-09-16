// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"testing"
)

func Test_HashID(t *testing.T) {

	e := Encode([]int{45, 434, 1313, 99, 2017091416})
	fmt.Println(len(e), e)
	d := Decode(e)
	fmt.Println(d)
}
