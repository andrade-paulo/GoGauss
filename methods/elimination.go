package methods

import ("errors")

func lineOperation(line []float64, pivotLine []float64, position int) []float64 {
    newLine := make([]float64, len(line))

    for i := 0; i < len(line); i++ {
        // Fill the lower triangle with 0s
        newLine[i] = ( line[i] * (pivotLine[position] / line[position]) ) - pivotLine[i]
        // fmt.Println("Equation: ", line[i], "*", pivotLine[position], "/", line[position], "-", pivotLine[i])
    }

    return newLine
}


func gaussElimination(mat [][]float64) ([][]float64, error) {
    size := len(mat)

    // Previous ordenation
    for i := 0; i < size; i++ {
        
        // If 0 on the pivot position
        if mat[i][i] == 0 && (i+1) != size {         

            // Search a proper vector for swap
            for j := i+1; mat[i][i] == 0; j = (j + 1) % size {

                // If I reach vector "i" again, then there's a column of 0s
                if j == i {
                    return mat, errors.New("The determinant is 0")
                }

                // If not 0 on the pivot position of both vectors
                if mat[j][i] != 0 && mat[i][j] != 0 {  
                    mat[i], mat[j] = mat[j], mat[i]  // Swap 
                }
            }
        
        }

    }
    
    for i := 0; i < size; i++ {
        // Search only the relevant elements
        for j := 0; j < i; j++ {
            if mat[i][j] != 0 {
                // fmt.Println("i:", i, "| j:", j)

                mat[i] = lineOperation(mat[i], mat[j], j)
            }
        // fmt.Println()
        }
    }

    return mat, nil
}
