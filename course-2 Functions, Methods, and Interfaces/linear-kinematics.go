package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5 * a * (t * t)
	}
}

// formula: s =½ a t2 + vot + so

func main() {
	var acceleration float64
	var velocity float64
	var displacement float64
	var err error

	fmt.Println("Please, enter an acceleration:")
	_, err = fmt.Scan(&acceleration)
	fmt.Println("Please, enter initial velocity:")
	_, err = fmt.Scan(&velocity)
	fmt.Println("Please, enter initial displacement:")
	_, err = fmt.Scan(&displacement)
	if err != nil {
		fmt.Print("Something went wrong!")
	}
}
