package day19

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	util "github.com/blaine-t-bush/advent-of-code/util"
	"github.com/dominikbraun/graph"
)

const (
	inputFile  = "./2022/day19/example_input.txt"
	maxMinutes = 24
)

type oreBot struct {
	oreCost int
}

type clayBot struct {
	oreCost int
}

type obsidianBot struct {
	oreCost  int
	clayCost int
}

type geodeBot struct {
	oreCost      int
	obsidianCost int
}

type blueprint struct {
	id          int
	oreBot      oreBot
	clayBot     clayBot
	obsidianBot obsidianBot
	geodeBot    geodeBot
}

type resources struct {
	ore      int
	clay     int
	obsidian int
	geodes   int
}

type bots struct {
	oreBots      int
	clayBots     int
	obsidianBots int
	geodeBots    int
}

type state struct {
	minute    int
	blueprint blueprint
	bots      bots
	resources resources
}

func parseBlueprint(line string) blueprint {
	r, err := regexp.Compile(`Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 8 {
		log.Fatal("parseBlueprint: incorrect number of regex matches")
	}

	id, err := strconv.Atoi(m[1])
	util.CheckErr(err)
	oreBotOreCost, err := strconv.Atoi(m[2])
	util.CheckErr(err)
	clayBotOreCost, err := strconv.Atoi(m[3])
	util.CheckErr(err)
	obsidianBotOreCost, err := strconv.Atoi(m[4])
	util.CheckErr(err)
	obsidianBotClayCost, err := strconv.Atoi(m[5])
	util.CheckErr(err)
	geodeBotOreCost, err := strconv.Atoi(m[6])
	util.CheckErr(err)
	geodeBotObsidianCost, err := strconv.Atoi(m[7])
	util.CheckErr(err)

	return blueprint{
		id: id,
		oreBot: oreBot{
			oreCost: oreBotOreCost,
		},
		clayBot: clayBot{
			oreCost: clayBotOreCost,
		},
		obsidianBot: obsidianBot{
			oreCost:  obsidianBotOreCost,
			clayCost: obsidianBotClayCost,
		},
		geodeBot: geodeBot{
			oreCost:      geodeBotOreCost,
			obsidianCost: geodeBotObsidianCost,
		},
	}
}

func parseBlueprints(lines []string) []blueprint {
	blueprints := []blueprint{}
	for _, line := range lines {
		blueprints = append(blueprints, parseBlueprint(line))
	}
	return blueprints
}

func gatherResources(s state) resources {
	for i := 0; i < s.bots.oreBots; i++ {
		s.resources.ore++
	}
	for i := 0; i < s.bots.clayBots; i++ {
		s.resources.clay++
	}
	for i := 0; i < s.bots.obsidianBots; i++ {
		s.resources.obsidian++
	}
	for i := 0; i < s.bots.geodeBots; i++ {
		s.resources.geodes++
	}

	return s.resources
}

func buyBot(botType string, s state) (resources, bots) {
	switch botType {
	case "ore":
		if s.resources.ore >= s.blueprint.oreBot.oreCost {
			s.resources.ore -= s.blueprint.oreBot.oreCost
			s.bots.oreBots++
		}
	case "clay":
		if s.resources.ore >= s.blueprint.clayBot.oreCost {
			s.resources.ore -= s.blueprint.clayBot.oreCost
			s.bots.clayBots++
		}
	case "obsidian":
		if s.resources.ore >= s.blueprint.obsidianBot.oreCost && s.resources.clay >= s.blueprint.obsidianBot.clayCost {
			s.resources.ore -= s.blueprint.obsidianBot.oreCost
			s.resources.clay -= s.blueprint.obsidianBot.clayCost
			s.bots.obsidianBots++
		}
	case "geode":
		if s.resources.ore >= s.blueprint.geodeBot.oreCost && s.resources.obsidian >= s.blueprint.geodeBot.obsidianCost {
			s.resources.ore -= s.blueprint.geodeBot.oreCost
			s.resources.obsidian -= s.blueprint.geodeBot.obsidianCost
			s.bots.geodeBots++
		}
	default:
		log.Fatal("buyBot: invalid bot type")
	}

	return s.resources, s.bots
}

func (s state) qualityLevel() int {
	return s.blueprint.id * s.resources.geodes
}

func stateHash(s state) state {
	return s
}

func instantiateState(b blueprint) state {
	// define starting conditions
	return state{
		minute:    0,
		blueprint: b,
		bots: bots{
			oreBots:      1,
			clayBots:     0,
			obsidianBots: 0,
			geodeBots:    0,
		},
		resources: resources{
			ore:      0,
			clay:     0,
			obsidian: 0,
			geodes:   0,
		},
	}
}

func instantiateGraph(s state) graph.Graph[state, state] {
	// instantiate graph
	g := graph.New(stateHash, graph.Directed(), graph.Acyclic())
	_ = g.AddVertex(s)

	return g
}

func expandGraph(current state, g graph.Graph[state, state], finalScores []int) (graph.Graph[state, state], []int) {
	// determine new possible states
	futureStates, finished := getFutureStates(current)

	// create connections to new states
	if !finished {
		for _, futureState := range futureStates {
			// some optimization. need to cut some corners to shrink graph.
			// 1. if there are no geodes collected by minute 20,
			//    this path is likely fruitless
			if futureState.resources.geodes == 0 && futureState.minute >= 20 {
				continue
			}

			// add vertices for all future states and edges to them from current state
			_ = g.AddVertex(futureState)
			_ = g.AddEdge(stateHash(current), stateHash(futureState))

			// continue the search if there is time remaining.
			// otherwise, append the final score for that state.
			if futureState.minute < maxMinutes {
				g, finalScores = expandGraph(futureState, g, finalScores)
			} else {
				finalScores = append(finalScores, futureState.blueprint.id*futureState.resources.geodes)
			}
		}
	}

	return g, finalScores
}

func getFutureStates(s state) ([]state, bool) {
	if s.minute >= maxMinutes {
		return []state{}, true
	}

	// determine possibly buying options
	buyOptions := []string{"none"}
	// optimizations:
	//   never buy anything in last minute.
	if s.minute < maxMinutes {
		if s.resources.ore >= s.blueprint.oreBot.oreCost {
			buyOptions = append(buyOptions, "ore")
		}
		if s.resources.ore >= s.blueprint.clayBot.oreCost {
			buyOptions = append(buyOptions, "clay")
		}
		if s.resources.ore >= s.blueprint.obsidianBot.oreCost && s.resources.clay >= s.blueprint.obsidianBot.clayCost {
			buyOptions = append(buyOptions, "obsidian")
		}
		if s.resources.ore >= s.blueprint.geodeBot.oreCost && s.resources.obsidian >= s.blueprint.geodeBot.obsidianCost {
			buyOptions = append(buyOptions, "geode")
		}
	}

	futureStates := []state{}
	for _, buyOption := range buyOptions {
		newState := s
		var tempBots bots
		switch buyOption {
		case "none":
			// gather resources
			newState.resources = gatherResources(newState)
		default:
			// start building
			newState.resources, tempBots = buyBot(buyOption, newState)
			// gather resources
			newState.resources = gatherResources(newState)
			// finish building
			newState.bots = tempBots
		}

		// advance minute
		newState.minute++
		futureStates = append(futureStates, newState)
	}

	return futureStates, false
}

func SolvePartOne() {
	input := util.ReadInput(inputFile)
	blueprints := parseBlueprints(input)
	qualityLevels := []int{}
	for _, b := range blueprints {
		fmt.Printf("analyzing blueprint %d...\n", b.id)
		s := instantiateState(b)
		g := instantiateGraph(s)
		g, finalScores := expandGraph(s, g, []int{})
		qualityLevel := util.MaxIntsSlice(finalScores)
		qualityLevels = append(qualityLevels, qualityLevel)

		fmt.Printf("  graph size: %d\n", g.Size())
		fmt.Printf("  best score: %d\n", qualityLevel)
	}

	fmt.Println(qualityLevels)
}

func SolvePartTwo() {
	input := util.ReadInput(inputFile)
	fmt.Println(len(input))
}
