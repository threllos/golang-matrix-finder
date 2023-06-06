package main

import "github.com/threllos/golang-matrix-finder/internal/task"

func main() {
	t := task.NewTask()
	t = t.Solve()
	t.Result()
}
