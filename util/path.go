package util

import (
	"log"
	"math"
)

type Coord struct {
	X int
	Y int
}

type DescendantMap map[Coord][]Coord

type PredecessorMap map[Coord][]Coord

type Queue []Coord

func (d DescendantMap) ToPredecessorMap() PredecessorMap {
	p := PredecessorMap{}
	for predecessor, adjacents := range d {
		for _, adjacent := range adjacents {
			if len(p[adjacent]) > 0 {
				p[adjacent] = append(p[adjacent], predecessor)
			} else {
				p[adjacent] = []Coord{predecessor}
			}
		}
	}

	return p
}

func (d DescendantMap) ToQueue() Queue {
	q := Queue{}
	for c := range d {
		q = append(q, c)
	}

	return q
}

func (q Queue) Has(c Coord) bool {
	for _, existing := range q {
		if existing == c {
			return true
		}
	}

	return false
}

func (q Queue) Remove(c Coord) Queue {
	new := Queue{}
	for _, existing := range q {
		if existing != c {
			new = append(new, existing)
		}
	}

	return new
}

func (p PredecessorMap) ToDescendantMap() DescendantMap {
	d := DescendantMap{}
	for adjacent, predecessors := range p {
		for _, predecessor := range predecessors {
			if len(p[predecessor]) > 0 {
				d[predecessor] = append(p[predecessor], adjacent)
			} else {
				d[predecessor] = []Coord{adjacent}
			}
		}
	}

	return d
}

func (d DescendantMap) HasDescendant(start Coord, end Coord) bool {
	if connectedCoords, exists := d[start]; exists {
		for _, c := range connectedCoords {
			if c == end {
				return true
			}
		}
	} else {
		log.Fatalf("DescendantMap.HasDescendant(): coord %v not found in descendant map\n", start)
	}

	return false
}

func (p PredecessorMap) HasPredecessor(start Coord, end Coord) bool {
	if predecessorCoords, exists := p[start]; exists {
		for _, c := range predecessorCoords {
			if c == end {
				return true
			}
		}
	} else {
		log.Fatalf("PredecessorMap.HasPredecessor(): coord %v not found in predecessor map\n", start)
	}

	return false
}

func (d DescendantMap) ShortestPath(start Coord, end Coord) float64 {
	// check that start has descendants
	if _, exists := d[start]; !exists {
		return -1
	}

	// Dijkstra's algorithm
	// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Pseudocode
	distances := map[Coord]float64{}
	previous := map[Coord]Coord{}
	q := d.ToQueue()
	for c := range d {
		distances[c] = math.Inf(1)
	}
	distances[start] = 0

	for {
		if len(q) == 0 {
			break
		}

		minDistance := math.Inf(1)
		minDistanceCoord := Coord{}
		found := false
		for _, c := range q {
			if distances[c] < minDistance {
				found = true
				minDistance = distances[c]
				minDistanceCoord = c
			}
		}

		if !found {
			break
		}

		if minDistanceCoord == end {
			break
		}

		q = q.Remove(minDistanceCoord)

		for _, neighbor := range q {
			if InSlice[Coord](neighbor, d[minDistanceCoord]) {
				alt := distances[minDistanceCoord] + 1
				if alt < distances[neighbor] {
					distances[neighbor] = alt
					previous[neighbor] = minDistanceCoord
				}
			}
		}
	}

	if distances[end] == math.Inf(1) {
		return -1
	}

	return distances[end]
}
