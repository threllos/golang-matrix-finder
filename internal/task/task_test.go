package task

import (
	"testing"

	"github.com/threllos/golang-matrix-finder/pkg/matrix"
)

func BenchmarkTask(b *testing.B) {
	m := matrix.NewMatrix(10, 10)
	t := Task{Matrix: m}

	for i := 0; i < b.N; i++ {
		t.Solve()
	}
}
