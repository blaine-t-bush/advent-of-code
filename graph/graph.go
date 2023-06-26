package graph

import (
	"errors"
	"log"
	"math"
)

type Vertex[T comparable] struct {
	Value T
}

type Edges[T comparable] map[*Vertex[T]]float64

type Queue[T comparable] []*Vertex[T]

type DirectedGraph[T comparable] map[*Vertex[T]]Edges[T]

type UndirectedGraph[T comparable] map[*Vertex[T]]Edges[T]

func NewGraph[T comparable]() DirectedGraph[T] {
	return DirectedGraph[T]{}
}

func (g *DirectedGraph[T]) HasVertex(v *Vertex[T]) bool {
	_, exists := (*g)[v]
	return exists
}

func (g *DirectedGraph[T]) AddVertex(v *Vertex[T]) {
	if g.HasVertex(v) {
		return
	}

	(*g)[v] = Edges[T]{}
}

func (g *DirectedGraph[T]) HasDirectConnection(from *Vertex[T], to *Vertex[T]) bool {
	if g.HasVertex(from) {
		if from == to {
			return true
		}

		for descendant := range (*g)[from] {
			if descendant == to {
				return true
			}
		}
	}

	return false
}

func (g *DirectedGraph[T]) AddEdge(from *Vertex[T], to *Vertex[T], distance float64) {
	if !g.HasVertex(from) {
		log.Fatalf("DirectedGraph.AddEdge: graph does not have vertex %v", *from)
		return
	}

	(*g)[from][to] = distance
}

func (g *DirectedGraph[T]) Queue() Queue[T] {
	queue := []*Vertex[T]{}
	for v := range *g {
		queue = append(queue, v)
	}
	return queue
}

func (q *Queue[T]) Remove(v *Vertex[T]) {
	// TODO add check that v exists in q
	deref := *q
	var index int
	for i, existing := range deref {
		if existing == v {
			index = i
		}
	}

	deref = append(deref[:index], deref[index+1:]...)
	*q = deref
}

func (g *DirectedGraph[T]) GetDistances(start *Vertex[T]) (map[*Vertex[T]]float64, error) {
	// check that start has descendants
	if _, exists := (*g)[start]; !exists {
		return nil, errors.New("start vertex does not exist in graph")
	}

	distances := map[*Vertex[T]]float64{}
	previous := map[*Vertex[T]]*Vertex[T]{}
	queue := g.Queue()
	for v := range *g {
		distances[v] = math.Inf(1)
	}
	distances[start] = 0

	for {
		if len(queue) == 0 {
			break
		}

		minDistance := math.Inf(1)
		minDistanceVertex := &Vertex[T]{}
		found := false
		for _, v := range queue {
			if distances[v] < minDistance {
				found = true
				minDistance = distances[v]
				minDistanceVertex = v
			}
		}

		if !found {
			break
		}

		queue.Remove(minDistanceVertex)

		for _, descendant := range queue {
			if g.HasDirectConnection(descendant, minDistanceVertex) {
				possibleDistance := distances[minDistanceVertex] + (*g)[minDistanceVertex][descendant]
				if possibleDistance < distances[descendant] {
					distances[descendant] = possibleDistance
					previous[descendant] = minDistanceVertex
				}
			}
		}
	}

	return distances, nil
}

func (g *DirectedGraph[T]) GetShortestDistance(start *Vertex[T], end *Vertex[T]) (float64, error) {
	distances, err := g.GetDistances(start)
	if err != nil {
		return 0, err
	}

	if !g.HasVertex(end) {
		return 0, errors.New("end vertex does not exist in graph")
	}

	return distances[end], nil
}
