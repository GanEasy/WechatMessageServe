// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package orm

import (
	"fmt"
	"testing"
)

func Test_GetDB(t *testing.T) {
	fmt.Println(DB())
}
