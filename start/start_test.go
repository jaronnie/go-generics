package start

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNumber1(t *testing.T) {
	sum1 := AddNumber1(2.5, 2.5)
	assert.Equal(t, float64(5), sum1)

	sum2 := AddNumber1(int64(2), int64(2))
	assert.Equal(t, int64(4), sum2)

	// not support MyInt64, MyFloat64
}

func TestAddNumber2(t *testing.T) {
	sum1 := AddNumber2(2.5, 2.5)
	assert.Equal(t, float64(5), sum1)

	sum2 := AddNumber2(int64(2), int64(2))
	assert.Equal(t, int64(4), sum2)

	// support MyInt64, MyFloat64
	sum3 := AddNumber2(MyFloat64(2.5), MyFloat64(2.5))
	assert.Equal(t, MyFloat64(5), sum3)
}
