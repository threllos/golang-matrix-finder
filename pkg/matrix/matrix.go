package matrix

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	Rows    int
	Columns int
	Data    [][]int
}

func NewMatrix(n, m int) Matrix {
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, m)
	}

	return Matrix{
		Rows:    n,
		Columns: m,
		Data:    d,
	}
}

func (m Matrix) FillRandom(n int) Matrix {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			m.Data[i][j] = rand.Intn(n) + 1
		}
	}

	return m
}

func (m Matrix) Print() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Columns; j++ {
			if m.Data[i][j] > 0 {
				fmt.Printf("%d ", m.Data[i][j])
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
}
