package mat

import "errors"

// ErrMatNotSq is the error for matrices are not square
var ErrMatNotSq = errors.New("Matrix is not square")

// ErrDimMisMatch is used when matrix dimensions don't match for operations
var ErrDimMisMatch = errors.New("Matrix dimension mismatch")
