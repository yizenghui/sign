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

func Test_Date(t *testing.T) {

	t.Fatal(time.Now().AddDate(0, 1, 30).Format(`15:04`))
	t.Fatal(time.Now().AddDate(0, 1, 30).Format(`02`))
	t.Fatal(time.Now().AddDate(0, 1, 30).Format(`01`))

}

func Test_Mate(t *testing.T) {

	var MateArr = map[string]string{
		//
		"01": "一月",
		"02": "二月",
		"03": "三月",
		"04": "四月",
		"05": "五月",
		"06": "六月",
		"07": "七月",
		"08": "八月",
		"09": "九月",
		"10": "十月",
		"11": "十一月",
		"12": "十二月",
	}
	d := time.Now().AddDate(0, 1, 30).Format(`01`)
	t.Fatal(MateArr[d])
}
