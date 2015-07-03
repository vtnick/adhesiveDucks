// defines the vector type and properties
package vector

import (
	"fmt"
	"math"
)

// Base Vector Type
type Vector struct {
	Data []float64
	Len  int
}

// Vector of all zeros
func Zeros(n int) (v Vector) {

	// Make data slice to represent the vector
	v.Data = make([]float64, n)
	v.Len = n

	return v
}

// Vector of all ones
func Ones(n int) (v Vector) {

	// Build zero vecotr first
	v = Zeros(n)

	// Loop over values and make 1
	for i := range v.Data {
		v.Data[i] = 1.0
	}

	return v
}

// Magnitude of the vector
func (v *Vector) Mag() (mag float64) {
	sqv := 0.0
	for _, val := range v.Data {
		sqv += val * val
	}
	mag = math.Sqrt(sqv)

	return mag
}

// Product of vector elements
func (v *Vector) Prod() (p float64) {
	p = 1.0
	for _, val := range v.Data {
		p *= val
	}

	return p
}

// Unit vector
func (v *Vector) UnitVec() (uv Vector) {
	mag := v.Mag()
	uv = Ones(v.Len)

	for i, val := range v.Data {
		uv.Data[i] = val / mag
	}

	return uv
}

// String print format for vector
func (v Vector) String() string {

	// init string
	vectorString := "["

	// loop through values
	for _, val := range v.Data {
		vectorString += fmt.Sprintf("%f, ", val)
	}

	// append closing brace
	vectorString += "]\n"

	return vectorString

}
