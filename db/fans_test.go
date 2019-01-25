// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"testing"
	"time"
)

func Test_CheckSign(t *testing.T) {

	var fans Fans

	fans.GetFansByOpenID(`as213sadassaklamsdlk213123`)
	// t.Fatal(fans.AllFans())
	t.Fatal(fans.DoSign())
	t.Fatal(fans.CheckSign())
}

func Test_GetTodaySignFans(t *testing.T) {

	var fans Fans
	t.Fatal(fans.GetTodaySignFans())
}

func Test_GetTodaySignFansW(t *testing.T) {

	var fans Fans
	t.Fatal(fans.GetTodaySignFansW(time.Now()))
}

func Test_XID(t *testing.T) {

	//time.Parse("2006-01-02 15:04:05", x)
	x, _ := time.Parse("2006-01-02 15:04:05", "2018-12-31 12:03:04")
	// x, _ := time.Parse("2006-01-02 15:04:05", "2019-01-01 12:03:04")
	t.Fatal(XID(x))
	t.Fatal(XID(time.Now()))
}
