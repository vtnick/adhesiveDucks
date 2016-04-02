package mat

// Add matrices together
func Add(a, b Matrix) (Matrix, error) {

	// Check matrix dimensions
	sizeA := a.Size()
	sizeB := b.Size()
	if (sizeA[0] != sizeB[0]) || sizeA[1] != sizeB[1] {
		return Matrix{}, ErrDimMisMatch
	}

	c := Zeros(sizeA[0], sizeA[1])

	for i := range a {
		for j := range a[i] {
			c[i][j] = a[i][j] + b[i][j]
		}
	}

	return c, nil
}

// Sub subtracts matrices
func Sub(a, b Matrix) (Matrix, error) {

	// Check matrix dimensions
	sizeA := a.Size()
	sizeB := b.Size()
	if (sizeA[0] != sizeB[0]) || sizeA[1] != sizeB[1] {
		return Matrix{}, ErrDimMisMatch
	}

	c := Zeros(sizeA[0], sizeA[1])

	for i := range a {
		for j := range a[i] {
			c[i][j] = a[i][j] - b[i][j]
		}
	}

	return c, nil
}

// Mldm stands for MATLAB Dot Multiply 'c = a .* b;'
func Mldm(a, b Matrix) (Matrix, error) {

	// Check matrix dimensions
	sizeA := a.Size()
	sizeB := b.Size()
	if (sizeA[0] != sizeB[0]) || sizeA[1] != sizeB[1] {
		return Matrix{}, ErrDimMisMatch
	}

	c := Zeros(sizeA[0], sizeA[1])

	for i := range a {
		for j := range a[i] {
			c[i][j] = a[i][j] * b[i][j]
		}
	}

	return c, nil
}

// Det returns the determinant of the matrix
func (m Matrix) Det() (float64, error) {

	var det float64
	var err error

	// Check if matrix is square
	if !m.IsSquare() {
		return det, ErrMatNotSq
	}

	// Get matrix dimensions
	size := m.Size()
	nRows := size[0]

	// calc determinant
	switch nRows {
	case 1:
		det = m[0][0]
	case 2:
		det = m[0][0]*m[1][1] - m[0][1]*m[1][0]
	default:
		det, err = nxnDet(m)
	}

	return det, err
}

// nxnDet performs the determinant operation
func nxnDet(m Matrix) (float64, error) {

	// init determinants
	var det, lDet, uDet float64

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

// LU Decomposition
func (m Matrix) LU() (Matrix, Matrix) {

	// Initialize L and U
	n := m.Size()[0]
	l := Zeros(n, n)
	u := Zeros(n, n)

	val := 0.0

	// Calculate LU
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j < i {
				l[j][i] = 0.0
				continue
			}
			l[j][i] = m[j][i]
			for k := 0; k < i; k++ {
				val = l[j][i] - l[j][k]*u[k][i]
				l[j][i] = val
			}
			for j := 0; j < n; j++ {
				if j < i {
					u[i][j] = 0.0
				} else if j == i {
					u[i][j] = 1.0
				} else {
					val = m[i][j] / l[i][i]
					u[i][j] = val
					for k := 0; k < i; k++ {
						val = u[i][j] - (l[i][j]*u[k][j])/l[i][i]
						u[i][j] = val
					}
				}
			}
		}
	}

	return l, u
}
