package matrix

import "adhesiveDucks/vector"

// Get row slice from matrix
func (m *Matrix) GetRow(row int) (vals vector.Vector) {

	// Make vals slice
	vals = vector.Zeros(m.Cols)

	for col := 0; col < m.Cols; col++ {
		vals.Data[col] = m.Get(row, col)
	}

	return vals
}

// Get row slice from matrix, only certain columns
func (m *Matrix) GetRowByCol(row int, cols []int) (vals vector.Vector) {

	// Make vals slice
	vals = vector.Zeros(len(cols))

	// Loop over columns
	for i, col := range cols {
		vals.Data[i] = m.Get(row, col)
	}

	return vals
}
