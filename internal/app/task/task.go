package task

import (
	"fmt"
	"log"

	"github.com/threllos/golang-matrix-finder/pkg/matrix"
)

type Task struct {
	Blocks       matrix.Matrix
	Groups       []Group
	MaxGroups    []Group
	MaxGroupSize int
}

func NewTask() Task {
	var r, c, n int

	fmt.Print("Input sizes row and column: ")
	_, err := fmt.Scanf("%d %d", &r, &c)
	if err != nil {
		log.Fatalf("Incorrect input size: %v\n", err.Error())
	}
	if r < 1 || c < 1 {
		log.Fatalln("Size must over 0")
	}
	data := matrix.NewMatrix(r, c)

	fmt.Print("Input colors size: ")
	_, err = fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatalf("Incorrect input size: %v\n", err.Error())
	}
	if n < 1 {
		log.Fatalln("Colors size must over 0")
	}
	data = data.FillRandom(n)

	return Task{
		Blocks:       data,
		Groups:       make([]Group, 0),
		MaxGroups:    make([]Group, 0),
		MaxGroupSize: 0,
	}
}

func (t Task) Solve() Task {
	directions := [4]Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	maxSize := 0

	queue := NewSetB()
	visited := NewSetP()
	queue = queue.Add(Block{0, 0, t.Blocks.Data[0][0]})

	for queue.Size() > 0 {
		cur, _ := queue.Get()
		group := NewGroup()
		group.Color = cur.Color
		stack := NewStackB()
		stack = stack.Push(cur)

		for stack.Size() > 0 {
			tmp, block, _ := stack.Pop()
			point := Point{block.X, block.Y}
			ok := visited.Exist(point)
			if ok {
				stack = tmp
				continue
			}
			visited = visited.Add(point)
			group.Append(point)

			for _, dir := range directions {
				next := point.Add(dir)
				ok := visited.Exist(next)
				if !ok && next.X >= 0 && next.X < t.Blocks.Columns && next.Y >= 0 && next.Y < t.Blocks.Rows {
					if t.Blocks.Data[next.Y][next.X] == t.Blocks.Data[point.Y][point.X] {
						tmp = tmp.Push(Block{next.X, next.Y, t.Blocks.Data[next.Y][next.X]})
					} else {
						queue = queue.Add(Block{next.X, next.Y, t.Blocks.Data[next.Y][next.X]})
					}
				}
			}
			stack = tmp
		}

		queue = queue.Delete(cur)
		t.Groups = append(t.Groups, group)
		if group.Size > maxSize {
			maxSize = group.Size
		}
	}

	t.MaxGroupSize = maxSize
	for _, group := range t.Groups {
		if group.Size == maxSize {
			t.MaxGroups = append(t.MaxGroups, group)
		}
	}

	return t
}

func (t Task) Result() {
	fmt.Println("Generated matrix:")
	t.Blocks.Print()

	fmt.Printf("Max group size: %d\n", t.MaxGroupSize)
	fmt.Printf("Groups len: %d\n", len(t.MaxGroups))

	for i, group := range t.MaxGroups {
		m := matrix.NewMatrix(t.Blocks.Rows, t.Blocks.Columns)
		for _, p := range group.Group {
			m.Data[p.Y][p.X] = group.Color
		}

		fmt.Printf("Group %d:\n", i+1)
		m.Print()
	}
}
