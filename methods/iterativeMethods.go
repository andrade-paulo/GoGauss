package methods

import (
	"errors"
	"math"
)

func GaussJacobi(A [][]float64, b []float64, x0 []float64, err float64) ([]float64, error) {
	size := len(A)
	ans := x0
	stopCondition := false

	for !stopCondition {
		new := make([]float64, size)

		for i := range size {
			diagonalSum := 0.0

			// Diagonal = 0
			if A[i][i] == 0 {
				return nil, errors.New("the matrix has a 0 on the diagonal")
			}

			sub := b[i] // b value

			// Sub of elements of the row, except the diagonal
			for j := range size {

				if i != j {
					sub -= A[i][j] * ans[j]
					diagonalSum += math.Abs(A[i][j])
				}
			}

			// Diagonal dominance
			if math.Abs(A[i][i]) <= diagonalSum {
				return nil, errors.New("the matrix is not diagonally dominant")
			}

			// New value for the answer
			new[i] = (sub) / A[i][i]
		}

		// Stop condition
		for i := range size {
			if math.Abs(new[i]-ans[i]) < err {
				stopCondition = true
			}
		}

		// Update the answer
		ans = new

		// Stop iterations
		if stopCondition {
			break
		}
	}

	return ans, nil
}

func GaussSeidel(A [][]float64, b []float64, x0 []float64, err float64) ([]float64, error) {
	size := len(A)
	ans := x0
	stopCondition := false

	for !stopCondition {
		new := make([]float64, len(ans))
		copy(new, ans) // Copy the values of the previous iteration

		for i := range size {
			diagonalSum := 0.0

			// Diagonal = 0
			if A[i][i] == 0 {
				return nil, errors.New("the matrix has a 0 on the diagonal")
			}

			sub := b[i] // b value

			// Sub of elements of the row, except the diagonal
			for j := range size {

				if i != j {
					// Use the new value for the elements that are already calculated
					sub -= A[i][j] * new[j]
					diagonalSum += math.Abs(A[i][j])
				}
			}

			// Diagonal dominance
			if math.Abs(A[i][i]) <= diagonalSum {
				return nil, errors.New("the matrix is not diagonally dominant")
			}

			// New value for the answer
			new[i] = (sub) / A[i][i]
		}

		// Stop condition
		for i := range size {
			if math.Abs(new[i]-ans[i]) < err {
				stopCondition = true
			}
		}

		// Update the answer
		ans = new

		// Stop iterations
		if stopCondition {
			break
		}
	}

	return ans, nil
}
