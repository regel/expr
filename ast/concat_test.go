package ast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatFloat64(t *testing.T) {
	a := []float64{1.0}
	b := []float64{-1.0}
	expected := Args{a, b}
	actual := concat(a, b)
	assert.Equal(t, expected, actual)
}
