package main

import (
	"fmt"
	"math/rand"
)

func makeVectorZeros(n int) []float64 {
	array := make([]float64, n)
	return array
}

func makeRandomVector(n int) []float64 {
	randomArray := make([]float64, n)
	for i := 0; i < n; i++ {
		randomArray[i] = rand.Float64()
	}
	return randomArray
}

func makeZeroMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
	}
	return matrix
}

func makeRandomMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = rand.Float64()
		}
	}
	return matrix
}

func multiplyVectorXMatrix(matrixSize int, vectorSize int) []float64 {
	matrix := makeRandomMatrix(matrixSize)
	vector := makeRandomVector(vectorSize)
	resultVector := makeVectorZeros(vectorSize)
	for i := range matrix {
		var value float64
		for j := range vector {
			value += matrix[i][j] * vector[j]
		}
		resultVector[i] = value
	}
	return resultVector
}

func ecuationSystemResolver() {
	// Missing
}

func main() {
	zeroVector := makeVectorZeros(10)
	fmt.Println("ZERO ARRAY:", zeroVector)
	randomVector := makeRandomVector(10)
	fmt.Println("RANDOM ARRAY:", randomVector)
	zeroMatrix := makeZeroMatrix(3)
	fmt.Println("ZERO MATRIX:", zeroMatrix)
	randomMatrix := makeRandomMatrix(3)
	fmt.Println("RANDOM MATRIX:", randomMatrix)
	resultVector := multiplyVectorXMatrix(3, 3)
	fmt.Println("MULTIPLY VECTOR:", resultVector)

}
