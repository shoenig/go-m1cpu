//go:build darwin && arm64

package m1cpu

import (
	"testing"

	"github.com/shoenig/test/must"
)

func Test_PCoreHz(t *testing.T) {
	hz := PCoreHz()
	must.Greater(t, 3_000_000_000, hz)
}

func Test_ECoreHz(t *testing.T) {
	hz := ECoreHz()
	must.Greater(t, 1_000_000_000, hz)
}

func Test_PCoreGHz(t *testing.T) {
	ghz := PCoreGHz()
	must.Greater(t, 3, ghz)
}

func Test_ECoreGHz(t *testing.T) {
	ghz := ECoreGHz()
	must.Greater(t, 1, ghz)
}

func Test_Show(t *testing.T) {
	t.Log("pCore Hz", PCoreHz())
	t.Log("eCore Hz", ECoreHz())
	t.Log("pCore GHz", PCoreGHz())
	t.Log("eCore GHz", ECoreGHz())
}