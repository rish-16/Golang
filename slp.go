package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var X []float64
	var y []float64

	for i := 0; i < 1000; i++ {
		X = append(X, float64(i))
		y = append(y, float64(2*i))
	}

	m, c := linearRegression(X, y)
	fmt.Println(m, c)
}

func hypothesis(X float64, m float64, c float64) float64 {
	return (X * m) + c
}

func loss(X []float64, y []float64, m float64, c float64) float64 {
	var totalLoss float64

	for i := 0; i < len(X); i++ {
		xi := X[i]
		yi := y[i]

		yHat := hypothesis(xi, m, c)
		loss := 0.5 * (yHat - yi) * (yHat - yi)

		totalLoss += loss
	}

	return totalLoss
}

func getDerivs(m float64, c float64, X []float64, y []float64) (float64, float64) {
	var dm float64
	var dc float64

	for i := 0; i < len(X); i++ {
		xi := X[i]
		yi := y[i]

		yHat := hypothesis(xi, m, c)
		dc += yHat - yi
		dm += (yHat - yi) * xi
	}

	dm = dm / float64(len(X))
	dc = dc / float64(len(X))

	return dm, dc
}

func update(m float64, c float64, X []float64, y []float64, alpha float64) (float64, float64) {
	dm, dc := getDerivs(m, c, X, y)

	m = m - (alpha * dm)
	c = m - (alpha * dc)

	return m, c
}

func linearRegression(X []float64, y []float64) (float64, float64) {
	rand.Seed(time.Now().UnixNano())

	var m float64 = 3
	var c float64 = 4

	fmt.Println(m, c)

	var errors []float64

	epochs := 100
	alpha := 0.01

	for i := 0; i < epochs; i++ {
		error := loss(X, y, m, c)
		m, c = update(m, c, X, y, alpha)
		fmt.Println(m, c)
		errors = append(errors, error)
	}

	fmt.Println(errors)

	return m, c
}
