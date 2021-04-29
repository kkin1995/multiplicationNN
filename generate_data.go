package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func generateDataToMultiplyTwoNumbers(N int) ([]float64, []float64, []float64) {
	data, err := os.Create("data.dat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Creating File: %v\n", err)
	}
	x1 := make([]float64, N)
	x2 := make([]float64, N)
	product := make([]float64, N)
	for i := 0; i < N; i++ {
		x1[i] = rand.Float64() * 100
		x2[i] = rand.Float64() * 100
		product[i] = x1[i] * x2[i]
		fmt.Fprintf(data, "%f\t%f\t%f\n", x1[i], x2[i], product[i])
	}
	return x1, x2, product
}

func generateDataToMultiplyTwoNumbersWithLog(N int) ([]float64, []float64, []float64) {
	data, err := os.Create("data.dat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error Creating File: %v\n", err)
	}
	x1 := make([]float64, N)
	x2 := make([]float64, N)
	product := make([]float64, N)
	for i := 0; i < N; i++ {
		x1_to_log := rand.Float64() * 100
		x1[i] = math.Log(x1_to_log)
		x2_to_log := rand.Float64() * 100
		x2[i] = math.Log(x2_to_log)
		product_to_log := x1_to_log * x2_to_log
		product[i] = math.Log(product_to_log)
		fmt.Fprintf(data, "%f\t%f\t%f\n", x1[i], x2[i], product[i])
	}
	return x1, x2, product
}
