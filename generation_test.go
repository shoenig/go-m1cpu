// Copyright (c) The M1CPU Authors
// SPDX-License-Identifier: MPL-2.0

package m1cpu

import (
	"testing"

	"github.com/shoenig/test/must"
)

func Test_generation(t *testing.T) {
	cases := []struct {
		model string
		exp   int
	}{
		{model: "Apple M1", exp: 1},
		{model: "Apple M1 Pro", exp: 1},
		{model: "Apple M3 Max", exp: 3},
		{model: "Apple M4", exp: 4},
		{model: "Apple M4 Pro", exp: 4},
		{model: "Apple M5", exp: 5},
		{model: "Apple M5 Pro", exp: 5},
		{model: "Apple M5 Max", exp: 5},
		{model: "Apple M10", exp: 10},
		{model: "VirtualApple @ 2.50GHz", exp: 0},
		{model: "", exp: 0},
	}
	for _, tc := range cases {
		t.Run(tc.model, func(t *testing.T) {
			must.Eq(t, tc.exp, generation(tc.model))
		})
	}
}
