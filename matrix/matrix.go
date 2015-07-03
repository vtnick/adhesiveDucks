package matrix

import (
	"adhesiveDucks/vector"
	"fmt"
)

// Base Matrix Type
type Matrix struct {
	Rows  int       // rows
	Cols  int       // columns
	Total int       // total number of elements
	Data  []float64 // matrix
}

// Matrix of all ones
func Ones(row, col int) (m Matrix) {

	// Make zeros matrix first
	m = Zeros(row, col)

	// Set all values to 1
	for i := range m.Data {
		m.Data[i] = 1.0
	}

	return m
}

// Matrix of all zeros
func Zeros(row, col int) (m Matrix) {

	// Store number of rows and columns for matrix
	m.Rows = row
	m.Cols = col
	m.Total = row * col

	// Make data slice to represent the matrix
	m.Data = make([]float64, row*col)

	return m
}

// Matrix from slice
func FromSlice(row, col int, s []float64) (m Matrix, err error) {

	// Make sure dimensions work
	if row*col != len(s) {
		return m, ErrRowColSlc
	}

	// start with zeros
	m = Zeros(row, col)

	// Copy slice to data
	m.Data = s

	return m, nil
}

// Diagonal matrix
func Diag(row, col int, x float64) (m Matrix) {

	// Make zeros matrix first
	m = Zeros(row, col)

	// Loop down diagonal
	colId := 0
	for row := 0; row < m.Rows; row++ {
		m.Set(row, colId, x)
		colId++
	}

	return m
}

// Set value in matrix
func (m *Matrix) Set(row, col int, val float64) {

	// Get the index
	i := row*m.Rows + col

	// Set the index
	m.Data[i] = val

}

// Get value from matrix
func (m Matrix) Get(row, col int) (val float64) {

	// Get the index
	i := row*m.Rows + col

	// Extract the index
	val = m.Data[i]

	return val
}

// get diag of square matrix
func (m Matrix) Diag() (v vector.Vector, err error) {
	v = vector.Zeros(m.Rows)

	if !m.IsSq() {
		return v, ErrMatNotSq
	}

	for i := 0; i < m.Rows; i++ {
		v.Data[i] = m.Get(i, i)
	}

	return v, nil
}

// whether the matrix is square
func (m Matrix) IsSq() bool {
	return m.Rows == m.Cols
}

func (m Matrix) IdToRC(id int) (rc [2]int, err error) {
	if id > len(m.Data) {
		return rc, ErrInpOutBnds
	}

	// get number of rows
	rc[0] = id / m.Rows

	// get number of columns
	rc[1] = id % m.Rows

	return rc, nil
}

// String representation of the matrix
func (m Matrix) String() string {

	matString := "[\n"

	// Print each row
	for row := 0; row < m.Rows; row++ {
		matString += "["
		for col := 0; col < m.Cols; col++ {
			matString += fmt.Sprintf("%f", m.Get(row, col))
			matString += ", "
		}
		matString += "]\n"
	}
	matString += "]\n"
	return matString
}
