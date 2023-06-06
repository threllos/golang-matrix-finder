package task

import (
	"fmt"
	"log"

	"github.com/threllos/golang-matrix-finder/pkg/matrix"
	"github.com/threllos/golang-matrix-finder/pkg/set"
	"github.com/threllos/golang-matrix-finder/pkg/stack"
)

type Task struct {
	Matrix       matrix.Matrix
	Groups       Groups
	MaxGroups    Groups
	MaxGroupSize int
}

func NewTask() Task {
	var r, c, n int

	fmt.Print("Enter the number of rows and columns: ")
	_, err := fmt.Scanf("%d %d", &r, &c)
	if err != nil {
		log.Fatalf("Incorrect entered number: %v\n", err.Error())
	}
	if r < 1 || c < 1 {
		log.Fatalln("Numbers must be greater than 0")
	}
	matrix := matrix.NewMatrix(r, c)

	fmt.Print("Input colors size: ")
	_, err = fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatalf("Incorrect entered number: %v\n", err.Error())
	}
	if n < 1 {
		log.Fatalln("Number must be greater than 0")
	}
	matrix = matrix.FillRandom(n)

	return Task{
		Matrix:       matrix,
		Groups:       NewGroups(),
		MaxGroups:    NewGroups(),
		MaxGroupSize: 0,
	}
}

func (t Task) Solve() Task {
	directions := [4]Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	maxSize := 0
	queue := set.NewSet[Block]()
	visited := set.NewSet[Point]()
	firstBlock := Block{
		Point: Point{0, 0},
		Color: t.Matrix.Data[0][0],
	}
	queue.Add(firstBlock)

	for queue.Size() > 0 {
		cur, _ := queue.Pop()
		group := NewGroup()
		group.Color = cur.Color
		stack := stack.NewStack[Block]()
		stack.Push(cur)

		for stack.Size() > 0 {
			cur, _ := stack.Pop()
			point := cur.Point

			isVisited := visited.Exist(point)
			if isVisited {
				continue
			}

			visited.Add(point)
			group.Append(point)

			for _, dir := range directions {
				next := point.Add(dir)
				isVisited := visited.Exist(next)
				if isVisited {
					continue
				}

				if next.X >= 0 && next.X < t.Matrix.Columns && next.Y >= 0 && next.Y < t.Matrix.Rows {
					nextBlock := Block{
						Point: Point{next.X, next.Y},
						Color: t.Matrix.Data[next.Y][next.X],
					}

					if t.Matrix.Data[next.Y][next.X] == t.Matrix.Data[point.Y][point.X] {
						stack.Push(nextBlock)
					} else {
						queue.Add(nextBlock)
					}
				}
			}
		}
		
		t.Groups.Append(group)
		if group.Size > maxSize {
			maxSize = group.Size
		}
	}

	t.MaxGroupSize = maxSize
	for _, group := range t.Groups {
		if group.Size == maxSize {
			t.MaxGroups.Append(group)
		}
	}

	return t
}

func (t Task) Result() {
	fmt.Println("Generated matrix:")
	t.Matrix.Print()

	fmt.Printf("Max group size: %d\n", t.MaxGroupSize)
	fmt.Printf("Number of groups: %d\n", len(t.MaxGroups))

	for i, group := range t.MaxGroups {
		m := matrix.NewMatrix(t.Matrix.Rows, t.Matrix.Columns)
		for _, p := range group.Group {
			m.Data[p.Y][p.X] = group.Color
		}

		fmt.Printf("Group %d:\n", i+1)
		m.Print()
	}
}
