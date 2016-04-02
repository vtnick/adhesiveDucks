package mat

import (
	"fmt"
	"testing"
)

func TestOnes(t *testing.T) {
	m := Ones(3, 3)
	fmt.Println(m)

}

func TestZeros(t *testing.T) {
	m := Zeros(3, 4)
	fmt.Println(m)

}

func TestLU(t *testing.T) {
	m := Matrix{
		{1, 1, -1},
		{2, -1, 3},
		{3, 1, -1},
	}

	l, u := m.LU()

	fmt.Println("m:", m)
	fmt.Println("l:", l)
	fmt.Println("u:", u)
}

func TestDet(t *testing.T) {
	m := Matrix{
		{1, 1, 2},
		{2, 0, -1},
		{-2, 1, 3},
	}

	det, err := m.Det()

	if err != nil {
		fmt.Println("Error Det:", err)
	} else {
		fmt.Println("det:", det)
	}
}
