// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

import "testing"

func Test_SendText(t *testing.T) {
	SendText("o7UTkjr7if4AQgcPmveQ5wJ5alsA", "hello")
}

func Test_SendArticle(t *testing.T) {
	SendArticle("o7UTkjr7if4AQgcPmveQ5wJ5alsA", "title", "description", "", "")
}
