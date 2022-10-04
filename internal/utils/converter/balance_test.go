package converter

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestConvertI2F(t *testing.T) {
	res := Int64ToFloat64(1000003)
	assert.Equal(t, res, 1000003.0/math.Pow10(8))
}
func TestConvertF2I(t *testing.T) {
	res := Float64ToInt64(3.00003423)
	assert.Equal(t, res, int64(300003423))
}
func TestConvertLI2F(t *testing.T) {
	res := Int64ToFloat64(math.MaxInt64)
	assert.Equal(t, res, 92233720368.54775807)
}
func TestConvertLF2I(t *testing.T) {
	res := Float64ToInt64(math.MaxFloat64)
	assert.Equal(t, res, int64(300003423))
}
