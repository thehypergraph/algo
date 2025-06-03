package graphs

import (
	"testing"
)

func TestDijkstra(t *testing.T) {
	g := NewGraph(5)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 3)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 2)
	g.AddEdge(2, 1, 4)
	g.AddEdge(2, 3, 8)
	g.AddEdge(2, 4, 2)
	g.AddEdge(3, 4, 7)

	result := g.Dijkstra(0)

	expectedDistances := []int{0, 7, 3, 9, 5}
	for i := range expectedDistances {
		if result.Distances[i] != expectedDistances[i] {
			t.Errorf("Expected distance to vertex %d to be %d, got %d", i, expectedDistances[i], result.Distances[i])
		}
	}
}
