package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func getSimbolEq(k []float64, c float64) (string, error) {
	var buf []string
	alphabet := []string{"x", "y", "z"}
	for i, j := range k {
		if j > 0 && i > 0 {
			buf = append(buf, "+")
		}

		if j < 0 && i == 0 {
			if j == -1 {
				buf = append(buf, "-"+alphabet[i])
			} else {
				buf = append(buf, strconv.FormatFloat(j, 'f', -1, 64)+alphabet[i])
			}

			continue
		}

		if j < 0 && i > 0 {
			buf = append(buf, "-")
		}

		if j == 0 {
			continue
		}

		if j == 1 || j == -1 {
			buf = append(buf, alphabet[i])

			continue
		}

		buf = append(buf, strconv.FormatFloat(math.Abs(j), 'f', -1, 64)+alphabet[i])
	}
	return strings.Join(buf, " ") + " = " + strconv.FormatFloat(c, 'f', -1, 64), nil
}

func getMyRand(irange, frange int64) float64 {
	x := rand.Float64() * float64(irange)
	if frange == 0 {
		return math.Trunc(x)
	}
	d := float64(frange)
	return math.Trunc(x) + (math.Trunc(x*d-math.Trunc(math.Trunc(x)*d)))/d
}

func getSign(i, m float64) float64 {
	if i >= m {
		return 1
	}
	return -1
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var x, y, z float64
	var a1, b1, bb1 float64
	var a2, b2, bb2 float64
	var a3, b3, bb3 float64
	var c1 float64
	var c2 float64
	var c3 float64
	success := 0
	for success < 10 {
		x = getMyRand(100, 100)
		y = getMyRand(100, 100)
		z = getMyRand(100, 100)
		a1 = getMyRand(10, 10)
		a2 = getMyRand(10, 10)
		a3 = getMyRand(10, 10)
		b1 = getMyRand(10, 10)
		b2 = getMyRand(10, 10)
		b3 = getMyRand(10, 10)
		b1 = b1 * getSign(b1, 5)
		b2 = b2 * getSign(b2, 5)
		b3 = b3 * getSign(b3, 5)
		bb1 = getMyRand(10, 10)
		bb2 = getMyRand(10, 10)
		bb3 = getMyRand(10, 10)
		bb1 = bb1 * getSign(bb1, 5)
		bb2 = bb2 * getSign(bb2, 5)
		bb3 = bb3 * getSign(bb3, 5)

		c1 = a1*x + b1*y + bb1*z
		c2 = a2*x + b2*y + bb2*z
		c3 = a3*x + b3*y + bb3*z

		if c1 < 0 || c2 < 0 || c3 < 0 {
			continue
		}

		cc1 := strconv.FormatFloat(c1, 'f', -1, 64)
		cc2 := strconv.FormatFloat(c2, 'f', -1, 64)
		cc3 := strconv.FormatFloat(c3, 'f', -1, 64)

		if len(cc1) > 9 || len(cc2) > 9 || len(cc3) > 9 {
			continue
		}

		success++
		eq1, _ := getSimbolEq([]float64{a1, b1, bb1}, c1)
		fmt.Println(eq1)
		eq2, _ := getSimbolEq([]float64{a2, b2, bb2}, c2)
		fmt.Println(eq2)
		eq3, _ := getSimbolEq([]float64{a3, b3, bb3}, c3)
		fmt.Println(eq3)
		fmt.Println("Answer: {", x, ",", y, ",", z, "}")
		fmt.Println("===")
	}
}
