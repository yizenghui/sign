// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package db

import (
	"testing"
)

func Test_PosterJsonMap(t *testing.T) {

	// Poster 海报json数据
	type Poster struct {
		Width  int64                    `json:"width"`
		Height int64                    `json:"height"`
		Clear  bool                     `json:"clear"`
		Views  []map[string]interface{} `json:"views"`
	}

	m1 := map[string]interface{}{"name": "John", "age": 10}
	m2 := map[string]interface{}{"name": "Alex", "age": 12}

	mbg := map[string]interface{}{"type": "image", "url": "https://signapi.readfollow.com/static/images/bg.jpg", "top": 0, "left": 0, "width": 375, "height": 360}

	js2 := Poster{
		375,
		360,
		true,
		[]map[string]interface{}{
			mbg,
			m1,
			m2,
		},
	}
	t.Fatal(js2)
}

func Test_PosterJson(t *testing.T) {
	type Item struct {
		Type      string
		URL       string
		Content   string
		Top       int64
		Left      int64
		Width     int64
		Height    int64
		FontSize  int64
		Color     string
		TextAlign string
		Bolder    bool
	}
	// Fans 粉丝数据信息
	type Poster struct {
		Width  int64
		Height int64
		Clear  bool
		Views  []Item
	}

	js := Poster{
		375,
		360,
		true,
		[]Item{
			Item{
				`image`,
				`https://signapi.readfollow.com/static/images/bg.jpg`,
				``,
				0,
				0,
				375,
				360,
				14,
				``,
				``,
				false,
			},
		},
	}

	// Fans 粉丝数据信息
	type Poster2 struct {
		Width  int64
		Height int64
		Clear  bool
		Views  []map[string]interface{}
	}

	m1 := map[string]interface{}{"name": "John", "age": 10}
	m2 := map[string]interface{}{"name": "Alex", "age": 12}

	js2 := Poster2{
		375,
		360,
		true,
		[]map[string]interface{}{
			m1,
			m2,
		},
	}
	t.Fatal(js2)
	t.Fatal(js)
}
