package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

func main() {
	cost_file, err := os.Create("cost_file.dat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in creating file: %v\n", err)
	}
	paramsFile, err := os.Create("params_file.dat")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in creating file: %v\n", err)
	}

	var N int = 1000
	x1, x2, y := generateDataToMultiplyTwoNumbersWithLog(N)
	var alpha float64 = 0.01
	w1 := rand.Float64() * 100
	w2 := rand.Float64() * 100
	b := rand.Float64() * 100

	var number_of_epochs int = 10000

	for i := 0; i < number_of_epochs; i++ {
		var cost float64 = 0.0
		for j := 0; j < N; j++ {
			z := w1*x1[j] + w2*x2[j] + b
			y_hat := linear(z)
			loss := (y[j] - y_hat)
			cost += math.Pow(y[j]-y_hat, 2)

			dw1 := ((-2.0 / float64(N)) * loss) * x1[j]
			dw2 := ((-2.0 / float64(N)) * loss) * x2[j]
			db := ((-2.0 / float64(N)) * loss)

			w1 = w1 - alpha*dw1
			w2 = w2 - alpha*dw2
			b = b - alpha*db

		}
		cost = cost * (1 / float64(N))
		fmt.Fprintf(cost_file, "%d\t%f\n", i, cost)
	}
	fmt.Fprintf(paramsFile, "Weight 1: %f\nWeight 2: %f\nBias: %f\n", w1, w2, b)

	var test_x float64 = 20.0
	var test_y float64 = 5.0
	pred := prediction(test_x, test_y, w1, w2, b)
	actual_value := test_x * test_y
	fmt.Fprintf(os.Stdout, "First Number: %f\nSecond Number: %f\nProduct: %f\nPredicted Product: %f\n", test_x, test_y, actual_value, pred)
}

func linear(x float64) float64 {
	return x
}

func prediction(test_x float64, test_y float64, w1 float64, w2 float64, b float64) float64 {
	prediction_to_exponentiate := w1*math.Log(test_x) + w2*math.Log(test_y) + b
	pred := math.Exp(prediction_to_exponentiate)
	return pred
}
