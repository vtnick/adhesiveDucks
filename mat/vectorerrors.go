package mat

import "errors"

// ErrLenNotEq - Error for vectors that do not have the same length
var ErrLenNotEq = errors.New("Vectors are not the same length")

// ErrBadDimXProd - Error for vectors whose dimensions do not meet cross product rules
var ErrBadDimXProd = errors.New("Cannot compute cross product with these vector dimensions")
