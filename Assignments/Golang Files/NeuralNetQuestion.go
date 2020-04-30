// A1_Q2_8179721
package main

import (
	"fmt"
	"math"
)

var alpha01 float64 = 0.1
var alpha11 float64 = 0.3
var alpha12 float64 = 0.4

var alpha02 float64 = 0.5
var alpha21 float64 = 0.8
var alpha22 float64 = 0.3

var alpha03 float64 = 0.7
var alpha31 float64 = 0.6
var alpha32 float64 = 0.6

func CalculateSigmoid(number float64) float64 {
	sigmoid := 1 / (1 + (1 / math.Pow(math.E, number)))
	return sigmoid
}

func CalculateTheta(neurons []float64, chanTheta chan []float64) {
	var thetas []float64
	i := 0
	for i <= len(neurons)-1 {
		theta := CalculateSigmoid((0.5 + (0.3 * neurons[i]) + (0.7 * neurons[i+1]) + (0.1 * neurons[i+2])))
		thetas = append(thetas, theta)
		i += 3
	}
	chanTheta <- thetas
	return
}

func CalculateZeta(x1 []float64, x2 []float64, chanZeta chan []float64) {
	var neuZeta []float64
	for i := 0; i <= len(x1)-1; i++ {
		zeta1 := CalculateSigmoid(alpha01 + (alpha11 * x1[i]) + (alpha12)*x2[i])
		neuZeta = append(neuZeta, zeta1)
		zeta2 := CalculateSigmoid(alpha02 + (alpha21 * x1[i]) + (alpha22)*x2[i])
		neuZeta = append(neuZeta, zeta2)
		zeta3 := CalculateSigmoid(alpha03 + (alpha31 * x1[i]) + (alpha32)*x2[i])
		neuZeta = append(neuZeta, zeta3)
	}
	chanZeta <- neuZeta
	return
}

func CalculateHiddenLayersX1(neurons int, ch chan []float64) {
	var neuSin []float64
	var sinVar float64
	for i := 1; i <= neurons; i++ {
		sinVar = math.Sin(float64((2 * math.Pi) * (float64(i - 1.0)) / float64(neurons)))
		neuSin = append(neuSin, sinVar)
	}
	ch <- neuSin
	return
}

func CalculateHiddenLayersX2(neurons int, ch2 chan []float64) {
	var neuCos []float64
	var cosVar float64
	for i := 1; i <= neurons; i++ {
		cosVar = math.Cos(float64((2 * math.Pi) * (float64(i - 1.0)) / float64(neurons)))
		neuCos = append(neuCos, cosVar)
	}
	ch2 <- neuCos
	return
}

func main() {
	ch := make(chan []float64)
	ch1 := make(chan []float64)
	ch2 := make(chan []float64)
	ch3 := make(chan []float64)

	var userInput int
	fmt.Println("Enter an integer value : ")
	_, err := fmt.Scanf("%d", &userInput)
	if err != nil {
		fmt.Println(err)
	}

	go CalculateHiddenLayersX1(userInput, ch)
	x1 := <-ch
	go CalculateHiddenLayersX2(userInput, ch1)
	x2 := <-ch1
	go CalculateZeta(x1, x2, ch2)
	x3 := <-ch2
	go CalculateTheta(x3, ch3)
	x4 := <-ch3
	fmt.Println(" X1    X2    T1")
	for i := 0; i <= len(x1)-1; i++ {
		fmt.Printf("|%-6.3f|%-6.3f|%-6.3f|", x1[i], x2[i], x4[i])
		fmt.Println()
	}
}
