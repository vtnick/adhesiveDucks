// Defines errors for the matrix package
package matrix

import "errors"

// Error for matrices are not square
var ErrMatNotSq = errors.New("Matrix is not square")

// Error for row/col out of bounds
var ErrInpOutBnds = errors.New("Row/Column input out of bounds")

var ErrRowColSlc = errors.New("Row/Column input does not match slice length")
