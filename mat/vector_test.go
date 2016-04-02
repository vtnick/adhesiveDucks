package mat

import (
	"fmt"
	"testing"
)

func TestZeroVector(t *testing.T) {
	v := ZeroVector(5, Row)
	fmt.Println(v)
}
