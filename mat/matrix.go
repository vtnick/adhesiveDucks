package mat

import "math"

// Matrix is typed as a 2d slice -- rows x columns
type Matrix [][]float64

// Size returns the size of the matrix
func (m Matrix) Size() []int {
	s := make([]int, 2)

	s[0] = len(m)
	s[1] = len(m[0])

	return s
}

// IsSquare checks for square matrix
func (m Matrix) IsSquare() bool {
	// Get the size of the matrix
	s := m.Size()

	if s[0] == s[1] {
		return true
	}

	return false
}

// Min will return the minimum value in the matrix as well as its location
func (m Matrix) Min() (float64, []int) {
	// init location slice
	loc := make([]int, 2)

	// init min value
	min := math.MaxFloat64

	for i := range m {
		for j := range m[i] {
			if m[i][j] < min {
				min = m[i][j]
				loc[0] = i
				loc[1] = j
			}
		}
	}
	return min, loc
}

// Max will return the maximum value in the matrix as well as its location
func (m Matrix) Max() (float64, []int) {
	// init location slice
	loc := make([]int, 2)

	// init min value
	max := -math.MaxFloat64

	for i := range m {
		for j := range m[i] {
			if m[i][j] > max {
				max = m[i][j]
				loc[0] = i
				loc[1] = j
			}
		}
	}
	return max, loc
}

// Zeros will create a matrix of zeros of size mxn
func Zeros(m, n int) Matrix {
	z := make([][]float64, m)

	for i := range z {
		z[i] = make([]float64, n)
	}

	return z
}

// Ones will create a matrix of ones of size mxn
func Ones(m, n int) Matrix {
	z := make([][]float64, m)

	for i := range z {
		z[i] = make([]float64, n)
		for j := range z[i] {
			z[i][j] = 1.0
		}
	}

	return z
}

// Transpose will perform a tranoose of the matrix
func (m Matrix) Transpose() Matrix {
	// Get dimensions
	size := m.Size()

	t := Zeros(size[1], size[0])

	for i := range m {
		for j, val := range m[i] {
			t[j][i] = val
		}
	}

	return t
}

// Diag will return the diagonal of a square matrix.
func (m Matrix) Diag() (Matrix, error) {
	// Check if matrix is square
	if !m.IsSquare() {
		return Matrix{}, ErrMatNotSq
	}

	// Get matrix dimensions
	size := m.Size()
	nRows := size[0]

	d := Zeros(nRows, 1)

	for i := range m {
		d[i][0] = m[i][i]
	}

	return d, nil
}
