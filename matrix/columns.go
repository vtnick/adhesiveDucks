package matrix

import "adhesiveDucks/vector"

// Get column slice from matrix
func (m Matrix) GetCol(col int) (vals vector.Vector) {

	// Make vals slice
	vals = vector.Zeros(m.Rows)

	for row := 0; row < m.Rows; row++ {
		vals.Data[row] = m.Get(row, col)
	}

	return vals
}

// Get column slice from matrix, only certain rows
func (m Matrix) GetColByRow(col int, rows []int) (vals vector.Vector) {

	// Make vals slice
	vals = vector.Zeros(len(rows))

	// Loop over cols
	for i, row := range rows {
		vals.Data[i] = m.Get(row, col)
	}

	return vals
}
