package puzzles

import "container/heap"

type Graph struct {
	vertices []Vertex
	width    int
	height   int
}

type Point struct {
	x int
	y int
}

type Vertex struct {
	position  Point
	direction int

	visited  bool
	heatLoss int

	calculatedHeatLoss int
	total              int

	index int
}

type PriorityQueue []*Vertex

func (queue PriorityQueue) Len() int { return len(queue) }

func (queue PriorityQueue) Less(i, j int) bool {
	return queue[i].total < queue[j].total
}

func (queue *PriorityQueue) Push(x any) {
	length := len(*queue)
	item := x.(*Vertex)
	item.index = length
	*queue = append(*queue, item)
}

func (queue *PriorityQueue) Pop() any {
	old := *queue
	length := len(old)
	item := old[length-1]
	old[length-1] = nil
	item.index = -1
	*queue = old[0 : length-1]

	return item
}

func (queue PriorityQueue) Swap(i, j int) {
	queue[i], queue[j] = queue[j], queue[i]
	queue[i].index = i
	queue[j].index = j
}

func (queue *PriorityQueue) update(item *Vertex, priority int) {
	heap.Fix(queue, item.index)
}
