// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/zhwei820/g.

package gutil_test

import (
	"testing"

	"github.com/zhwei820/g/frame/g"

	"github.com/zhwei820/g/test/gtest"
	"github.com/zhwei820/g/util/gutil"
)

func Test_StructToSlice(t *testing.T) {
	type A struct {
		K1 int
		K2 string
	}
	gtest.C(t, func(t *gtest.T) {
		a := &A{
			K1: 1,
			K2: "v2",
		}
		s := gutil.StructToSlice(a)
		t.Assert(len(s), 4)
		t.AssertIN(s[0], g.Slice{"K1", "K2", 1, "v2"})
		t.AssertIN(s[1], g.Slice{"K1", "K2", 1, "v2"})
		t.AssertIN(s[2], g.Slice{"K1", "K2", 1, "v2"})
		t.AssertIN(s[3], g.Slice{"K1", "K2", 1, "v2"})
	})
	gtest.C(t, func(t *gtest.T) {
		s := gutil.StructToSlice(1)
		t.Assert(s, nil)
	})
}
