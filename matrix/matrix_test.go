package matrix

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
	m := Zeros(3, 3)
	m.Data = []float64{
		1, 1, -1,
		2, -1, 3,
		3, 1, -1,
	}

	l, u := m.LU()

	fmt.Println("m:", m)
	fmt.Println("l:", l)
	fmt.Println("u:", u)
}

func TestDet(t *testing.T) {
	m := Zeros(3, 3)
	m.Data = []float64{
		1, 1, 2,
		2, 0, -1,
		-2, 1, 3,
	}

	det, err := m.Det()

	if err != nil {
		fmt.Println("Error Det:", err)
	} else {
		fmt.Println("det:", det)
	}
}

func TestFromSlice(t *testing.T) {
	s := []float64{
		1, 1, 2,
		2, 0, -1,
		-2, 1, 3,
	}

	m, _ := FromSlice(3, 3, s)

	fmt.Println(m)
}

func TestIdToRC(t *testing.T) {
	m := Zeros(3, 3)
	m.Data = []float64{
		1, 1, 2,
		2, 0, -1,
		-2, 1, 3,
	}

	for i := 0; i < len(m.Data); i++ {
		rc, err := m.IdToRC(i)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Index: %d: %f\n", i, m.Get(rc[0], rc[1]))
	}
}
