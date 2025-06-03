package graphs

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight int
}

// Graph represents a weighted directed graph
type Graph struct {
	Vertices int
	AdjList  [][]Edge
}

// NewGraph creates a new graph with the specified number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		AdjList:  make([][]Edge, vertices),
	}
}

// AddEdge adds a weighted edge from vertex 'from' to vertex 'to'
func (g *Graph) AddEdge(from, to, weight int) {
	g.AdjList[from] = append(g.AdjList[from], Edge{To: to, Weight: weight})
}

// AddUndirectedEdge adds an undirected weighted edge between two vertices
func (g *Graph) AddUndirectedEdge(u, v, weight int) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// PriorityQueue implements a min-heap for Dijkstra's algorithm
type PriorityQueue []*Item

type Item struct {
	Vertex   int
	Distance int
	Index    int
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// DijkstraResult holds the result of Dijkstra's algorithm
type DijkstraResult struct {
	Distances []int
	Previous  []int
	Source    int
}

// Dijkstra implements Dijkstra's shortest path algorithm
func (g *Graph) Dijkstra(source int) *DijkstraResult {
	// Initialize distances and previous vertices
	distances := make([]int, g.Vertices)
	previous := make([]int, g.Vertices)
	visited := make([]bool, g.Vertices)

	// Set all distances to infinity and previous to -1
	for i := 0; i < g.Vertices; i++ {
		distances[i] = math.MaxInt32
		previous[i] = -1
	}

	// Distance from source to itself is 0
	distances[source] = 0

	// Create priority queue and add source vertex
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{Vertex: source, Distance: 0})

	for pq.Len() > 0 {
		// Extract vertex with minimum distance
		current := heap.Pop(&pq).(*Item)
		u := current.Vertex

		// Skip if already visited
		if visited[u] {
			continue
		}

		visited[u] = true

		// Check all neighbors of current vertex
		for _, edge := range g.AdjList[u] {
			v := edge.To
			weight := edge.Weight

			// Calculate new distance through current vertex
			newDistance := distances[u] + weight

			// If new distance is shorter, update it
			if newDistance < distances[v] {
				distances[v] = newDistance
				previous[v] = u
				heap.Push(&pq, &Item{Vertex: v, Distance: newDistance})
			}
		}
	}

	return &DijkstraResult{
		Distances: distances,
		Previous:  previous,
		Source:    source,
	}
}

// GetShortestPath reconstructs the shortest path from source to target
func (result *DijkstraResult) GetShortestPath(target int) []int {
	if result.Distances[target] == math.MaxInt32 {
		return nil // No path exists
	}

	path := make([]int, 0)
	current := target

	// Backtrack from target to source
	for current != -1 {
		path = append([]int{current}, path...)
		current = result.Previous[current]
	}

	return path
}

// PrintResult prints the results of Dijkstra's algorithm
func (result *DijkstraResult) PrintResult() {
	fmt.Printf("Shortest distances from vertex %d:\n", result.Source)
	for i, dist := range result.Distances {
		if dist == math.MaxInt32 {
			fmt.Printf("Vertex %d: ∞ (unreachable)\n", i)
		} else {
			fmt.Printf("Vertex %d: %d\n", i, dist)
		}
	}
	fmt.Println()
}

// PrintPath prints the shortest path to a specific vertex
func (result *DijkstraResult) PrintPath(target int) {
	path := result.GetShortestPath(target)
	if path == nil {
		fmt.Printf("No path from vertex %d to vertex %d\n", result.Source, target)
		return
	}

	fmt.Printf("Shortest path from vertex %d to vertex %d: ", result.Source, target)
	for i, vertex := range path {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(vertex)
	}
	fmt.Printf(" (distance: %d)\n", result.Distances[target])
}

func main() {
	// Create a sample graph
	graph := NewGraph(6)

	// Add edges (vertex1, vertex2, weight)
	graph.AddUndirectedEdge(0, 1, 4)
	graph.AddUndirectedEdge(0, 2, 3)
	graph.AddUndirectedEdge(1, 2, 1)
	graph.AddUndirectedEdge(1, 3, 2)
	graph.AddUndirectedEdge(2, 3, 4)
	graph.AddUndirectedEdge(3, 4, 2)
	graph.AddUndirectedEdge(4, 5, 6)
	graph.AddEdge(1, 4, 5) // Directed edge from 1 to 4

	fmt.Println("Graph created with 6 vertices and weighted edges")
	fmt.Println("Running Dijkstra's algorithm from vertex 0...")

	// Run Dijkstra's algorithm from vertex 0
	result := graph.Dijkstra(0)

	// Print all shortest distances
	result.PrintResult()

	// Print shortest paths to specific vertices
	for i := 1; i < graph.Vertices; i++ {
		result.PrintPath(i)
	}
}
