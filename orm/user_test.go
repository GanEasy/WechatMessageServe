// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package orm

import (
	"testing"

	"fmt"
)

func init() {

	DB().AutoMigrate(&User{})
}

func Test_CreateUser(t *testing.T) {

	var user User
	user.GetUserByEmail("245561237@qq.com")

	user.Invited = true
	user.Registered = true
	user.Subscribed = true
	user.OpenID = "o7UTkjr7if4AQgcPmveQ5wJ5alsA"
	user.Save()
	// user.GetUserFollowBooks()
	fmt.Println(user)
}

func Test_GetUser(t *testing.T) {
	var user User
	DB().First(&user, 1)
	// user.GetUserFollowBooks()
	fmt.Println(user)
}

func Test_GetUserByOpenID(t *testing.T) {
	var user User
	user.GetUserByOpenID("o7UTkjr7if4AQgcPmveQ5wJ5alsA")
	// user.GetUserFollowBooks()
	DB().Save(&user)
	fmt.Println(user)

}
