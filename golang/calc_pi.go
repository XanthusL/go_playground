package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("5000\n", pi(5000))

	fmt.Println("500000\n", calcPi(5000, 100))

	fmt.Println("5000000\n", calcPi(5000, 1000))

	fmt.Println("10000000\n", calcPi(5000, 2000))

	fmt.Println("20000000\n", calcPi(10000, 2000))
}

func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

// ------------------- MY VERSION ------------------------
// c concurrency
// g granularity
func calcPi(c, g int) float64 {

	ch := make(chan float64)
	for k := 0; k <= c; k++ {
		go term2(ch, float64(k*g), g)
	}
	f := 0.0
	for k := 0; k <= c; k++ {
		f += <-ch
	}
	return f

}
func term2(ch chan float64, k float64, g int) {
	f := float64(0)
	for i := 0; i < g; i++ {
		j := k + float64(i)
		f += 4 * math.Pow(-1, j) / (2*j + 1)

	}
	ch <- f

}
