package day16

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	graph "github.com/blaine-t-bush/advent-of-code/graph"
	util "github.com/blaine-t-bush/advent-of-code/util"
)

// TODO
// add weighting to graph library
// (rooms with 0 flow rate should increase the distance to subsequent rooms but
// not be included in the graph)

const maxMinutes = 30

type valve struct {
	name    string
	flow    int
	tunnels map[string]float64 // key is connected room name, value is distance to it
}

type state struct {
	room       string
	openValves map[string]int // key is room name, value is minute at which opened
	minute     int
}

func parseLine(line string) valve {
	r := regexp.MustCompile(`Valve (\w\w) has flow rate=(\d+); tunnels? leads? to valves? ((?:\w+,? ?)*)`)
	m := r.FindStringSubmatch(line)
	if len(m) != 4 {
		log.Fatalf("parseInput: could not parse line %s\n", line)
	}

	// convert flow rate to int
	flowRate, err := strconv.Atoi(m[2])
	util.CheckErr(err)

	// create map of connections
	tunnels := map[string]float64{}
	for _, name := range strings.Split(m[3], ", ") {
		tunnels[name] = 1
	}

	return valve{
		name:    m[1],
		flow:    flowRate,
		tunnels: tunnels,
	}
}

func parseInput(lines []string) (map[string]*graph.Vertex[string], map[string]valve) {
	// create map of valves
	valves := map[string]valve{}
	for _, line := range lines {
		v := parseLine(line)
		valves[v.name] = v
	}

	// create map of vertices (only care about ones remaining after collapse)
	vertices := map[string]*graph.Vertex[string]{}
	for _, v := range valves {
		vertex := &graph.Vertex[string]{Value: v.name}
		vertices[v.name] = vertex
	}

	return vertices, valves
}

func createGraph(vertices map[string]*graph.Vertex[string], valves map[string]valve) graph.DirectedGraph[string] {
	g := graph.NewGraph[string]()

	// add vertices to graph
	for _, vertex := range vertices {
		g.AddVertex(vertex)
	}

	// add edges to graph
	for name, v := range valves {
		for tunnel, distance := range v.tunnels {
			g.AddEdge(vertices[name], vertices[tunnel], distance)
		}
	}

	return g
}

func (s state) isOpen(name string) bool {
	_, exists := s.openValves[name]
	return exists
}

func (s state) releasedPressure(currentMinute int, valves map[string]valve) int {
	pressure := 0
	for name, minute := range s.openValves {
		pressure += valves[name].flow * (currentMinute - minute)
	}
	return pressure
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	vertices, valves := parseInput(input)
	g := createGraph(vertices, valves)
	for name, vertex := range vertices {
		dist, err := g.GetShortestDistance(vertices["AA"], vertex)
		util.CheckErr(err)
		fmt.Printf("distance from AA to %s: %.0f\n", name, dist)
	}

	initialState := state{
		room:       "AA",
		openValves: map[string]int{},
		minute:     0,
	}
	states := []state{initialState}

	// start in room AA
	// options are move to an adjacent room or open valve if not already open
	// append each possible option and update state.
	// continue until have reached 30 steps
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
