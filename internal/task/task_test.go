package task

import (
	"testing"

	"github.com/threllos/golang-matrix-finder/pkg/matrix"
)

func BenchmarkSolve10x10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := matrix.NewMatrix(10, 10)
		m.FillRandom(2)
		t := Task{Matrix: m}
		b.StartTimer()

		t.Solve()
	}
}

func BenchmarkSolve100x100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := matrix.NewMatrix(100, 100)
		m.FillRandom(10)
		t := Task{Matrix: m}
		b.StartTimer()

		t.Solve()
	}
}

func BenchmarkSolve1000x1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := matrix.NewMatrix(1000, 1000)
		m.FillRandom(10)
		t := Task{Matrix: m}
		b.StartTimer()

		t.Solve()
	}
}
