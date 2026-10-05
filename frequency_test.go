// Copyright (c) The M1CPU Authors
// SPDX-License-Identifier: MPL-2.0

package m1cpu

import (
	"testing"

	"github.com/shoenig/test/must"
)

func Test_eCoreVoltageState(t *testing.T) {
	cases := []struct {
		model string
		exp   int
	}{
		{model: "Apple M1", exp: 1},
		{model: "Apple M3 Max", exp: 1},
		{model: "Apple M4 Pro", exp: 1},
		{model: "Apple M5", exp: 1},
		{model: "Apple M5 Pro", exp: 22},
		{model: "Apple M5 Max", exp: 22},
		{model: "", exp: 1},
	}
	for _, tc := range cases {
		t.Run(tc.model, func(t *testing.T) {
			must.Eq(t, tc.exp, eCoreVoltageState(tc.model))
		})
	}
}

func Test_frequencyConversion(t *testing.T) {
	cases := []struct {
		name  string
		model string
		raw   uint64
		hz    uint64
		ghz   float64
	}{
		{name: "M1 performance", model: "Apple M1", raw: 3_204_000_000, hz: 3_204_000_000, ghz: 3.204},
		{name: "M3 efficiency", model: "Apple M3 Max", raw: 2_748_000_000, hz: 2_748_000_000, ghz: 2.748},
		{name: "M4 performance", model: "Apple M4 Pro", raw: 4_512_000, hz: 4_512_000_000, ghz: 4.512},
		{name: "M5 performance", model: "Apple M5", raw: 4_608_000, hz: 4_608_000_000, ghz: 4.608},
		{name: "M5 Pro performance", model: "Apple M5 Pro", raw: 4_608_000, hz: 4_608_000_000, ghz: 4.608},
		{name: "M5 Max performance", model: "Apple M5 Max", raw: 4_608_000, hz: 4_608_000_000, ghz: 4.608},
		{name: "M5 Max efficiency", model: "Apple M5 Max", raw: 4_380_000, hz: 4_380_000_000, ghz: 4.38},
		{name: "missing frequency", model: "Apple M5 Max", raw: 0, hz: 0, ghz: 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			must.Eq(t, tc.hz, toHz(tc.raw, tc.model))
			must.Eq(t, tc.ghz, toGhz(tc.raw, tc.model))
		})
	}
}
