package ogxbig_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/niconical/ogx/extra/ogxbig"
)

func TestInt(t *testing.T) {
	a := big.NewInt(100)
	b := big.NewInt(200)

	t.Run("multiply", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		y := ogxbig.FromMathBig(b)
		// 100 * 200 = 20000
		assert.Equal(t, ogxbig.FromMathBig(big.NewInt(20000)), x.Mul(y))
	})
	t.Run("add", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		y := ogxbig.FromMathBig(b)
		// 100 + 200 = 300
		assert.Equal(t, ogxbig.FromMathBig(big.NewInt(300)), x.Add(y))
	})

	t.Run("sub", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		y := ogxbig.FromMathBig(b)
		// 100 -200 = -100
		assert.Equal(t, ogxbig.FromMathBig(big.NewInt(-100)), x.Sub(y))
	})

	t.Run("div", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		y := ogxbig.FromMathBig(b)
		// 200 / 100 = 2
		assert.Equal(t, ogxbig.FromMathBig(big.NewInt(2)), y.Div(x))
	})

	t.Run("negation", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		assert.Equal(t, ogxbig.FromMathBig(big.NewInt(-100)), x.Neg())
	})

	t.Run("int64", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		assert.Equal(t, int64(-100), x.Neg().ToInt64())
	})

	t.Run("uint64", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		assert.Equal(t, uint64(100), x.ToUInt64())
	})
	t.Run("toString", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)
		assert.Equal(t, "100", x.String())
	})
	t.Run("fromString", func(t *testing.T) {
		x, err := ogxbig.NewInt().FromString("100")
		assert.Nil(t, err)
		assert.Equal(t, "100", x.String())
	})
	t.Run("fromInt64", func(t *testing.T) {
		x := ogxbig.FromInt64(100000000)
		assert.Equal(t, int64(100000000), x.ToInt64())
	})

	t.Run("Abs", func(t *testing.T) {
		x := ogxbig.FromMathBig(a)

		assert.Equal(t, x.Neg().Abs(), x)
	})
	t.Run("compare: ", func(t *testing.T) {
		x := ogxbig.FromMathBig(a) // 100
		y := ogxbig.FromMathBig(b) // 200

		cmp := x.Cmp(y)

		t.Run("eq ?", func(t *testing.T) {
			assert.Equal(t, cmp.Eq(), false)
		})
		t.Run("lt ?", func(t *testing.T) {
			assert.Equal(t, cmp.Lt(), true)
		})
		t.Run("gt ?", func(t *testing.T) {
			assert.Equal(t, cmp.Gt(), false)
		})
		t.Run("leq ?", func(t *testing.T) {
			assert.Equal(t, cmp.Leq(), true)
		})
		t.Run("geq ?", func(t *testing.T) {
			assert.Equal(t, cmp.Geq(), false)
		})
	})

	t.Run("empty string ", func(t *testing.T) {

		x, err := ogxbig.NewInt().FromString("")
		assert.Nil(t, err)
		assert.Equal(t, x.ToInt64(), int64(0))
	})

}

func TestFloat(t *testing.T) {

	cases := []struct {
		f1   float64
		f2   float64
		diff float64
		mul  float64
		sum  float64
		div  float64
		eq   bool
		geq  bool
		leq  bool
		lt   bool
		gt   bool
	}{
		{
			f1:   1.01,
			f2:   1.02,
			diff: 0.01,
			mul:  1.0302,
			sum:  2.03,
			div:  1,
			eq:   false,
			geq:  true,
			lt:   false,
			gt:   true,
			leq:  false,
		},
		{
			f1:   10.001,
			f2:   10.01,
			diff: 0.009,
			sum:  20.011,
			mul:  100.11001,
			div:  1,
			eq:   false,
			geq:  true,
			lt:   false,
			gt:   true,
			leq:  false,
		},
		{
			f1:   1,
			f2:   1,
			diff: 0,
			sum:  2,
			mul:  1,
			div:  1,
			eq:   true,
			geq:  true,
			leq:  true,
			lt:   false,
			gt:   false,
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%f , %f ", c.f1, c.f2), func(t *testing.T) {
			f1 := big.NewFloat(c.f1)
			f2 := big.NewFloat(c.f2)
			diff := big.NewFloat(c.diff)
			mul := big.NewFloat(c.mul)
			sum := big.NewFloat(c.sum)
			div := big.NewFloat(c.div)

			ogxF1 := ogxbig.NewFloat().FromMathFloat(f1)
			ogxF2 := ogxbig.NewFloat().FromMathFloat(f2)
			ogxDiff := ogxbig.NewFloat().FromMathFloat(diff)
			ogxMul := ogxbig.NewFloat().FromMathFloat(mul)
			ogxSum := ogxbig.NewFloat().FromMathFloat(sum)
			ogxDiv := ogxbig.NewFloat().FromMathFloat(div)

			cmp := ogxF2.Cmp(ogxF1)

			assert.Equal(t, ogxDiff.String(), ogxF2.Sub(ogxF1).String())
			assert.Equal(t, ogxMul.String(), ogxF2.Mul(ogxF1).String())
			assert.Equal(t, ogxSum.String(), ogxF2.Add(ogxF1).String())
			assert.Equal(t, ogxDiv.String(), ogxF2.Div(ogxF2).String())

			assert.Equal(t, cmp.Eq(), c.eq)
			assert.Equal(t, cmp.Geq(), c.geq)
			assert.Equal(t, cmp.Gt(), c.gt)
			assert.Equal(t, cmp.Leq(), c.leq)
			assert.Equal(t, cmp.Lt(), c.lt)
		})
	}

	f, err := ogxbig.NewFloat().FromString("-100")
	assert.NoError(t, err)

	assert.Equal(t, f.Abs().String(), "100")

	f2 := ogxbig.NewFloat().FromMathFloat(big.NewFloat(100))

	assert.Equal(t, f.String(), f2.Neg().String())
}

func TestFixture(t *testing.T) {
	data := `
- id: 1
  name: ethereum
  base: wei
  equals: 1000000000000000000	
- id: 2
  name: bitcoin
  base: satoshi
  equals: 1000000000
`
	type CryptoNetwork struct {
		ID     int
		Name   string
		Base   string
		Equals *ogxbig.Int
	}

	cryptoNet := []CryptoNetwork{}

	err := yaml.Unmarshal([]byte(data), &cryptoNet)

	assert.NoError(t, err)

	// @Todo
	// we expect that the decoded values become convertible to ogxbig.Int
}
