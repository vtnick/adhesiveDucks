package mat

// defines the vector type and properties. vector is defined as a 1-d matrix so
// all functions will work on matrices
import (
	"fmt"
	"math"
)

// Row and Col describe the type of vector
const (
	Row = 0
	Col = 1
)

// ZeroVector returns a zero vector of length m in either Row or Col format
func ZeroVector(m, vType int) Matrix {
	// If this is a row vector
	if vType == Row {
		return Zeros(0, m)
	}
	return Zeros(m, 0)
}

// OnesVector Returns a ones vector of length m in either Row or Col format
func OnesVector(m, vType int) Matrix {
	// If this is a row vector
	if vType == Row {
		return Ones(0, m)
	}
	return Ones(m, 0)
}

// Mag - Magnitude of the vector
func (v Matrix) Mag() float64 {
	sqv := 0.0
	// loop through vector
	for _, rowdata := range v {
		for _, val := range rowdata {
			sqv += val * val
		}
	}

	return math.Sqrt(sqv)
}

// Prod - Product of vector elements
func (v Matrix) Prod() float64 {
	p := 1.0
	// loop through vector
	for _, rowdata := range v {
		for _, val := range rowdata {
			p *= val
		}
	}

	return p
}

// Len returns the length of the vector
func (v Matrix) Len() int {

	// get size
	s := v.Size()

	// return appropriate dimension
	if s[Row] == 1 && s[Col] > 0 {
		return s[Col]
	}

	return s[Row]

}

// UnitVec - Unit vector
func (v Matrix) UnitVec() Matrix {
	mag := v.Mag()
	uv := v

	// loop through vector
	for row, rowdata := range v {
		for col, val := range rowdata {
			uv[row][col] = val / mag
		}
	}

	return uv
}

// VType - is the vector type
func (v Matrix) VType() int {
	// get size
	s := v.Size()

	// if row dim is one then column vector
	if s[Row] == 1 && s[Col] > 0 {
		return Col
	}

	return Row
}

// String print format for vector
func (v Matrix) String() string {

	// init string
	vectorString := "["

	// loop through values
	for _, rowdata := range v {
		for _, val := range rowdata {
			vectorString += fmt.Sprintf("%f, ", val)
		}
	}

	// append closing brace
	vectorString += "]\n"

	return vectorString

}
