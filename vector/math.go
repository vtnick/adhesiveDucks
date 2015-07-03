// Package contains basic vector math operations
package vector

// Dot product of two vectors
func DotProd(a, b Vector) (c float64, err error) {

	// Check if a and b are the same length
	if a.Len != b.Len {
		return c, ErrLenNotEq
	}

	// Set tmp variables to speed up access
	dataA := a.Data
	dataB := b.Data

	// loop over elements and add their product to the sum
	for i := range dataA {
		c += dataA[i] * dataB[i]
	}

	return c, nil
}

// Cross Product of two vectors
func CrossProd(a, b Vector) (c Vector, err error) {
	// Check Vector dimensions
	if a.Len != b.Len {
		return c, ErrLenNotEq
	} else if a.Len < 2 || a.Len > 3 {
		return c, ErrBadDimXProd
	}

	// Set tmp variables to speed up access
	dataA := a.Data
	dataB := b.Data

	// If vector is 2d append 0 for z dimension
	if a.Len == 2 {
		dataA = append(dataA, 0.0)
		dataB = append(dataB, 0.0)
	}

	// fill out result vector
	c = Zeros(a.Len)

	// compute cross product
	c.Data[0] = dataA[1]*dataB[2] - dataA[2]*dataB[1]
	c.Data[1] = dataA[0]*dataB[2] - dataA[2]*dataB[0]
	c.Data[2] = dataA[0]*dataB[1] - dataA[1]*dataB[0]

	return c, nil
}

// Mldm stands for MATLAB Dot Multiply 'c = a .* b;'
func Mldm(a, b Vector) (c Vector, err error) {

	// Check Vector dimensions
	if a.Len != b.Len {
		return c, ErrLenNotEq
	}

	// fill out result vector
	c = Zeros(a.Len)

	// loop over elements and multiply them together
	for i := range a.Data {
		c.Data[i] = a.Data[i] * b.Data[i]
	}

	return c, nil
}
