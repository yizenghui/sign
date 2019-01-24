// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"testing"
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
