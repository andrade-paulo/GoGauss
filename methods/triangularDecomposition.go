package methods

import ("fmt"; "math"; "errors")


func transpose(mat [][]float64) [][]float64 {
    size := len(mat)
    T := make([][]float64, size)

    for i := range size {
        T[i] = make([]float64, size)
    }

    for i := range size {
        for j := range size {
            T[i][j] = mat[j][i]
        }    
    }

    return T
}


func DooLittle(mat [][]float64) {
    size := len(mat)

    L := make([][]float64, size)
    U := make([][]float64, size)
    
    for i := range size {
        L[i] = make([]float64, size)
        U[i] = make([]float64, size)
    }
    
    var sum float64

    // For eache line
    for i := range size {
        // Only valid positions of the lower triangular matrix
        for j := 0; j <= i; j++ {
            if i == j {
                L[i][i] = 1  // Diagonal of L is 1
            } else {
                sum = 0

                // Generalization of element L[i][j]
                for k := 0; k < i; k++ {
                    sum += L[i][k] * U[k][j]
                }

                L[i][j] = (mat[i][j] - sum) / U[j][j]
            }
        }
        
        // Only valid positions of the upper triangular matrix
        for j := i; j < size; j++ {
            sum = 0

            // Generalization of element U[i][j]
            for k := 0; k < i; k++ {
                sum += L[i][k] * U[k][j]
            }
            
            U[i][j] = mat[i][j] - sum
        }        
    }

   fmt.Println(mat, " = ", L, " x ", U) 
}



func Crout(mat [][]float64) {
    size := len(mat)

    L := make([][]float64, size)
    U := make([][]float64, size)
    
    for i := range size {
        L[i] = make([]float64, size)
        U[i] = make([]float64, size)
    }
    
    var sum float64

    // For eache line
    for i := range size { 
        // Only valid positions of the lower triangular matrix
        for j := 0; j <= i; j++ {
            sum = 0

            // Generalization of element L[i][j]
            for k := 0; k < i; k++ {
                sum += L[i][k] * U[k][j]
            }
            
            L[i][j] = mat[i][j] - sum
        }

        // Only valid positions of the upper triangular matrix
        for j := i; j < size; j++ {
            if i == j {
                U[i][i] = 1  // Diagonal of U is 1
            } else {
                sum = 0

                // Generalization of element U[i][j]
                for k := 0; k < i; k++ {
                    sum += L[i][k] * U[k][j]
                }

                U[i][j] = (mat[i][j] - sum) / L[i][i]
            }
        }
    }

   fmt.Println(mat, " = ", L, " x ", U) 
}


func Cholesky(mat [][]float64) error {
    size := len(mat)

    G := make([][]float64, size)
    
    for i := range size {G[i] = make([]float64, size)}
    
    var sum float64

    // For eache line
    for i := range size { 
        sum = 0
        
        // Sum for G[i][i]
        for j := 0; j < i; j++ {
            sum += math.Pow(G[i][j], 2)            
        }

        // Check if it is defined positive
        r := mat[i][i] - sum

        if r <= 0 {
            return errors.New("Matrix is not defined positive")
        }

        // Determining G[i][i] that will be used further
        G[i][i] = math.Sqrt(r)

        for j := i+1; j < size; j++ {
            sum = 0

            for k := range i {
                sum += G[j][k] * G[i][k]
            }

            G[j][i] = (mat[j][i] - sum) / G[i][i]
        }
    }

    G_t := transpose(G)
    fmt.Println(mat, " = ", G, " x ", G_t) 

    return nil
}
