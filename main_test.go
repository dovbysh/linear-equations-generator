package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataGetSimbol = []struct {
	eq string
	k  []float64
	c  float64
}{
	{"x + 5y = 4", []float64{1, 5}, 4},
	{"x - 5y = 4", []float64{1, -5}, 4},
	{"x - 5y = -4", []float64{1, -5}, -4},
	{"-2x - 5y = -4", []float64{-2, -5}, -4},
	{"-x - 5y = -4", []float64{-1, -5}, -4},
	{"x - 5z = -4", []float64{1, 0, -5}, -4},
	{"x - 5.5z = -4", []float64{1, 0, -5.5}, -4},
}

func TestGetSimbolEq(t *testing.T) {
	for _, d := range dataGetSimbol {
		s, err := getSimbolEq(d.k, d.c)

		assert.Nil(t, err)
		assert.Equal(t, s, d.eq)
	}
}

var dataGetSign = []struct {
	exp, a, b float64
}{
	{-1, 4, 5},
	{1, 6, 5},
}

func TestGetSign(t *testing.T) {
	for _, d := range dataGetSign {
		actual := getSign(d.a, d.b)

		assert.Equal(t, actual, d.exp)
	}
}
