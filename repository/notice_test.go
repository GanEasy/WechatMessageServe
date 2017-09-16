// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import (
	"fmt"
	"os"
	"testing"
)

func Test_SendText(t *testing.T) {
	// SendText("id", "hello")
	logname := os.Getenv("MAX_WORKERS")
	fmt.Printf("MAX_WORKERS is %s\n", logname)
}
