package main

import ("fmt"; "math/rand"; "errors")


func populate(size int) [][]float64 {
    mat := make([][]float64, size)
    
    for i := 0; i < size; i++ {
        for j := 0; j < size+1; j++ {
            mat[i] = append(mat[i], 10 + rand.Float64() * 1)
        }
    }

    return mat
}


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


func solveMatrix(mat [][]float64) []float64 {
    ans := make([]float64, len(mat))

    // Starts from the last element
    for i := len(mat) - 1; i >= 0; i-- {
        ans[i] = mat[i][len(mat)]  // b_i
        
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
    var size int

    fmt.Print("Size of random square matrix: ")
    fmt.Scanln(&size)

    mat := populate(size)

    //mat := [][]float64 {
    //    {1, 2, 3, 9},
    //    {4, 5, 6, 3},
    //    {7, 8, 9, -2},
    //}

    fmt.Println("Matrix:", mat)
    fmt.Println()
    
    gaussMatrix, err := gaussElimination(mat)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Gauss Matrix:", gaussMatrix)
    }
    fmt.Println()

    ans := solveMatrix(gaussMatrix)
    
    fmt.Println("Solution:", ans)
}
