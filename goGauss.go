package main

import (
	"fmt"
	"math/rand"

	"github.com/andrade-paulo/GoGauss/methods"
)

func populate(size int) [][]float64 {
	mat := make([][]float64, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size+1; j++ {
			mat[i] = append(mat[i], 0+rand.Float64()*10)
		}
	}

	return mat
}

func solveMatrix(mat [][]float64) []float64 {
	ans := make([]float64, len(mat))

	// Starts from the last element
	for i := len(mat) - 1; i >= 0; i-- {
		ans[i] = mat[i][len(mat)] // b_i

		// Subtraction of other known elements
		for j := 0; j < len(mat); j++ {
			if j != i {
				// a_{i,j} * x_j
				ans[i] -= mat[i][j] * ans[j]
			}
		}

		// Coefficient of line i
		ans[i] /= mat[i][i]
	}

	return ans
}

func main() {
	A := [][]float64{
		{10, 2, 1},
		{1, 5, 1},
		{2, 3, 10},
	}

	b := []float64{7, -8, 6}

	x0 := []float64{0.7, -1.6, 0.6}

	fmt.Println("Matrix:", A)
	fmt.Println()

	answer_jacobi, err_jacobi := methods.GaussJacobi(A, b, x0, 0.0001)
	answer_seidel, err_seidel := methods.GaussSeidel(A, b, x0, 0.0001)

	fmt.Println("-=- Iterative Methods -=-")

	fmt.Print("Gauss-Jacobi: ")
	if err_jacobi != nil {
		fmt.Println(err_jacobi)
	} else {
		fmt.Println(answer_jacobi)
	}

	fmt.Print("Gauss-Seidel: ")
	if err_seidel != nil {
		fmt.Println(err_seidel)
	} else {
		fmt.Println(answer_seidel)
	}
}
