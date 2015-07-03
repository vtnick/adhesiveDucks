// Package contains basic matrix math operations
package matrix

import (
	"errors"
	"fmt"
	"math"
)

type MaxMin struct {
	Val float64
	Loc [2]int
}

func (m *Matrix) Max() (max MaxMin) {
	max.Val = -math.MaxFloat64
	val := 0.0

	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			val = m.Get(i, j)
			if val > max.Val {
				max.Val = val
				max.Loc[0] = i
				max.Loc[1] = j
			}
		}
	}
	return max
}

func (m *Matrix) Min() (min MaxMin) {
	min.Val = math.MaxFloat64

	for i, val := range m.Data {
		if val < min.Val {
			min.Val = val
			rc, err := m.IdToRC(i)
			if err != nil {
				fmt.Println(err)
				return
			}
			min.Loc[0] = rc[0]
			min.Loc[1] = rc[1]
		}
	}

	return min
}

// Add matrices together
func Add(a, b Matrix) (c Matrix, err error) {

	// Check matrix dimensions
	if (a.Rows != b.Rows) || a.Cols != b.Cols {
		return c, errors.New("Matrix dimensions are not equal")
	}

	c = Zeros(a.Rows, a.Cols)

	for i := range a.Data {
		c.Data[i] = a.Data[i] + b.Data[i]
	}

	return c, nil
}

// Subtract matrices together
func Sub(a, b Matrix) (c Matrix, err error) {

	// Check matrix dimensions
	if (a.Rows != b.Rows) || a.Cols != b.Cols {
		return c, errors.New("Matrix dimensions are not equal")
	}

	c = Zeros(a.Rows, a.Cols)

	for i := range a.Data {
		c.Data[i] = a.Data[i] - b.Data[i]
	}

	return c, nil
}

// Mldm stands for MATLAB Dot Multiply 'c = a .* b;'
func Mldm(a, b Matrix) (c Matrix, err error) {

	// Check matrix dimensions
	if (a.Rows != b.Rows) || a.Cols != b.Cols {
		return c, errors.New("Matrix dimensions are not equal")
	}

	c = Zeros(a.Rows, a.Cols)

	for i := range a.Data {
		c.Data[i] = a.Data[i] * b.Data[i]
	}

	return c, nil
}

//------------------------------------------------------------------------------
// Determinant of the matrix
//------------------------------------------------------------------------------
func (m *Matrix) Det() (det float64, err error) {

	det = 0.0
	err = nil

	// Check if matrix is square
	if !m.IsSq() {
		return det, ErrMatNotSq
	}

	// calc determinant
	switch m.Rows {
	case 1:
		det = m.Data[0]
	case 2:
		det = m.Get(0, 0)*m.Get(1, 1) - m.Get(0, 1)*m.Get(1, 0)
	default:
		det, err = nxnDet(m)
	}

	return det, err
}

func nxnDet(m *Matrix) (det float64, err error) {

	// init determinants
	det = 0.0
	lDet := 0.0
	uDet := 0.0

	// calc LU
	l, u := m.LU()

	// det(L) = L.Diag().Prod()
	lDiag, err := l.Diag()
	if err != nil {
		return det, err
	}
	lDet = lDiag.Prod()

	// det(U) = U.Diag().Prod()
	uDiag, err := u.Diag()
	if err != nil {
		return det, err
	}
	uDet = uDiag.Prod()

	// det(a) = det(LU) = det(L)*det(U)
	det = lDet * uDet

	return det, nil
}

//------------------------------------------------------------------------------
// LU Decomposition
//------------------------------------------------------------------------------
func (m *Matrix) LU() (l, u Matrix) {

	// Initialize L and U
	n := m.Rows
	l = Zeros(n, n)
	u = Zeros(n, n)

	val := 0.0

	// Calculate LU
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < i {
				l.Set(j, i, 0.0)
				continue
			}
			l.Set(j, i, m.Get(j, i))
			for k := 0; k < i; k++ {
				val = l.Get(j, i) - l.Get(j, k)*u.Get(k, i)
				l.Set(j, i, val)
			}
			for j := 0; j < n; j++ {
				if j < i {
					u.Set(i, j, 0.0)
				} else if j == i {
					u.Set(i, j, 1.0)
				} else {
					val = m.Get(i, j) / l.Get(i, i)
					u.Set(i, j, val)
					for k := 0; k < i; k++ {
						val = u.Get(i, j) - ((l.Get(i, k) * u.Get(k, j)) / l.Get(i, i))
						u.Set(i, j, val)
					}
				}
			}
		}
	}

	return l, u
}
