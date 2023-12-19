package puzzles

import (
	"bytes"
	"container/heap"
	"os"
)

func Solve_part_1() int {
	grid, err := readFile("puzzles/inputs/day17.txt")
	if err != nil {
		return -1
	}

	return shortest(grid, 1, 3)
}

func Solve_part_2() int {
	grid, err := readFile("puzzles/inputs/day17.txt")
	if err != nil {
		return -1
	}

	return shortest(grid, 4, 10)
}

func (g *Graph) getEdges(vertex *Vertex, minSteps int, maxSteps int) []*Vertex {
	path := make([]*Vertex, 0, 6)

	if vertex.direction == 1 || vertex.direction == 2 {
		for heatloss, dy := 0, 1; dy <= maxSteps; dy++ {
			v := g.getVertex(vertex.position.x, vertex.position.y+dy, 0)
			if v != nil {
				heatloss += v.heatLoss
				if dy >= minSteps {
					v.calculatedHeatLoss = heatloss
					path = append(path, v)
				}
			}
		}

		for heatloss, dy := 0, 1; dy <= maxSteps; dy++ {
			v := g.getVertex(vertex.position.x, vertex.position.y-dy, 0)
			if v != nil {
				heatloss += v.heatLoss
				if dy >= minSteps {
					v.calculatedHeatLoss = heatloss
					path = append(path, v)
				}
			}
		}
	}

	if vertex.direction == 0 || vertex.direction == 2 {
		for heatloss, dx := 0, 1; dx <= maxSteps; dx++ {
			v := g.getVertex(vertex.position.x+dx, vertex.position.y, 1)
			if v != nil {
				heatloss += v.heatLoss
				if dx >= minSteps {
					v.calculatedHeatLoss = heatloss
					path = append(path, v)
				}
			}
		}

		for heatloss, dx := 0, 1; dx <= maxSteps; dx++ {
			v := g.getVertex(vertex.position.x-dx, vertex.position.y, 1)
			if v != nil {
				heatloss += v.heatLoss
				if dx >= minSteps {
					v.calculatedHeatLoss = heatloss
					path = append(path, v)
				}
			}
		}
	}

	return path
}

func NewGraph(grid [][]int) Graph {
	graph := Graph{}
	vertices := make([]Vertex, 0, len(grid)*len(grid)*2)
	graph.height = len(grid)

	for y := range grid {
		graph.width = len(grid[y])
		for x := range grid[y] {
			vertices = append(vertices, Vertex{
				position:  Point{x, y},
				direction: 0,
				total:     1 << 30,
				heatLoss:  grid[y][x],
			})

			vertices = append(vertices, Vertex{
				position:  Point{x, y},
				direction: 1,
				total:     1 << 30,
				heatLoss:  grid[y][x],
			})
		}
	}

	graph.vertices = vertices
	return graph
}

func (g *Graph) getVertex(x int, y int, plane int) *Vertex {
	if x < 0 || y < 0 || y >= g.height || x >= g.width {
		return nil
	}

	return &g.vertices[y*2*g.width+x*2+plane]
}

func shortest(grid [][]int, minSteps int, maxSteps int) int {
	graph := NewGraph(grid)
	vertices := graph.vertices

	vertices[0].total = 0
	vertices[0].direction = 2

	queue := make(PriorityQueue, len(vertices))
	for i := 0; i < len(vertices); i++ {
		vertices[i].index = i
		queue[i] = &vertices[i]
	}

	heap.Init(&queue)

	var vertex *Vertex
	var next = &vertices[len(vertices)-1]
	for {
		vertex = heap.Pop(&queue).(*Vertex)

		if vertex.position.x == next.position.x && vertex.position.y == next.position.y {
			break
		}

		vertex.visited = true

		for _, e := range graph.getEdges(vertex, minSteps, maxSteps) {
			if vertex.total+e.calculatedHeatLoss < e.total {
				e.total = vertex.total + e.calculatedHeatLoss
				queue.update(e, e.total)
			}
		}
	}

	return vertex.total
}

func readFile(filepath string) ([][]int, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(bytes.TrimSpace(file), []byte("\n"))
	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))

		for j, character := range line {
			n := int(character) - '0'
			grid[i][j] = n
		}
	}

	return grid, nil
}
