package day19

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"

	util "github.com/blaine-t-bush/advent-of-code/util"
	"github.com/dominikbraun/graph"
)

const (
	inputFile = "./2022/day19/input.txt"
)

var maxMinutes int

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

func (s state) geodeCount() int {
	return s.resources.geodes
}

func stateHash(s state) state {
	return s
}

func (s1 state) hasMoreOrEqualBots(s2 state) bool {
	return s1.bots.oreBots >= s2.bots.oreBots && s1.bots.clayBots >= s2.bots.clayBots && s1.bots.obsidianBots >= s2.bots.obsidianBots && s1.bots.geodeBots >= s2.bots.geodeBots
}

func (s1 state) hasMoreOrEqualResources(s2 state) bool {
	return s1.resources.ore >= s2.resources.ore && s1.resources.clay >= s2.resources.clay && s1.resources.obsidian >= s2.resources.obsidian && s1.resources.geodes >= s2.resources.geodes
}

func (s1 state) strictlyWorseThanExistingState(existingStates []state) bool {
	for _, s2 := range existingStates {
		if s1.minute >= s2.minute && s2.hasMoreOrEqualBots(s1) && s2.hasMoreOrEqualResources(s1) {
			return true
		}
	}

	return false
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

func expandGraph(current state, g graph.Graph[state, state], possibleGeodeCounts []int) (graph.Graph[state, state], []int) {
	// determine new possible states
	futureStates, finished := getFutureStates(current)

	// create connections to new states
	if !finished {
		for _, futureState := range futureStates {
			// don't bother with states that have already been searched
			if _, err := g.Vertex(futureState); err == nil {
				continue
			}

			// don't bother with state if there is a strictly better version of it out there

			// add vertices for all future states and edges to them from current state
			_ = g.AddVertex(futureState)
			_ = g.AddEdge(stateHash(current), stateHash(futureState))

			// some optimization. need to cut some corners to shrink graph.
			// 1. if there are no geodes collected by minute 20,
			//    this path is likely fruitless
			// 2. if late enough, no point in buying more bots.
			if futureState.bots.geodeBots == 0 && futureState.minute >= maxMinutes-4 {
				possibleGeodeCounts = append(possibleGeodeCounts, futureState.geodeCount())
				continue
			}
			if futureState.bots.obsidianBots == 0 && futureState.minute >= maxMinutes-6 {
				possibleGeodeCounts = append(possibleGeodeCounts, futureState.geodeCount())
				continue
			}
			if futureState.bots.clayBots == 0 && futureState.minute >= maxMinutes-8 {
				possibleGeodeCounts = append(possibleGeodeCounts, futureState.geodeCount())
				continue
			}

			// continue the search if there is time remaining.
			// otherwise, append the final score for that state.
			if futureState.minute < maxMinutes {
				g, possibleGeodeCounts = expandGraph(futureState, g, possibleGeodeCounts)
			} else {
				possibleGeodeCounts = append(possibleGeodeCounts, futureState.geodeCount())
			}
		}
	}

	return g, possibleGeodeCounts
}

func (b blueprint) maxOreCost() int {
	return util.MaxIntsSlice([]int{b.oreBot.oreCost, b.clayBot.oreCost, b.obsidianBot.oreCost, b.geodeBot.oreCost})
}

func (b blueprint) maxClayCost() int {
	return b.obsidianBot.clayCost
}

func (b blueprint) maxObsidianCost() int {
	return b.geodeBot.obsidianCost
}

func getFutureStates(s state) ([]state, bool) {
	if s.minute >= maxMinutes {
		return []state{}, true
	}

	// determine possibly buying options
	buyOptions := []string{"none"}
	// optimizations:
	//   1. never buy anything in last minute.
	//   2. always buy geode bot if possible.
	//   3. if have ore bots equal to highest ore cost, no need to buy more ore bots.
	//      same for clay and obsidian.
	//   4. rush clay bots by buying ore bots as early as possible until can build a clay bot every minute.
	if s.minute < maxMinutes {
		if s.resources.ore >= s.blueprint.oreBot.oreCost && s.bots.oreBots < s.blueprint.maxOreCost() {
			buyOptions = append(buyOptions, "ore")
		}
		if s.resources.ore >= s.blueprint.clayBot.oreCost && s.bots.clayBots < s.blueprint.maxClayCost() {
			buyOptions = append(buyOptions, "clay")
		}
		if s.resources.ore >= s.blueprint.obsidianBot.oreCost && s.resources.clay >= s.blueprint.obsidianBot.clayCost && s.bots.obsidianBots < s.blueprint.maxObsidianCost() {
			buyOptions = append(buyOptions, "obsidian")
		}
		if s.resources.ore >= s.blueprint.geodeBot.oreCost && s.resources.obsidian >= s.blueprint.geodeBot.obsidianCost {
			buyOptions = []string{"geode"}
		}
		// if s.bots.oreBots < s.blueprint.clayBot.oreCost && s.resources.ore >= s.blueprint.oreBot.oreCost {
		// 	buyOptions = []string{"ore"}
		// }
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
	fmt.Println("solving part one...")
	maxMinutes = 24
	input := util.ReadInput(inputFile)
	blueprints := parseBlueprints(input)
	qualityLevels := []int{}

	var wg sync.WaitGroup
	wg.Add(len(blueprints))

	for _, b := range blueprints {
		go func(b blueprint) {
			defer wg.Done()
			fmt.Printf("  blueprint %d starting analysis\n", b.id)
			s := instantiateState(b)
			g := instantiateGraph(s)
			g, possibleGeodeCounts := expandGraph(s, g, []int{})
			qualityLevel := util.MaxIntsSlice(possibleGeodeCounts) * b.id
			qualityLevels = append(qualityLevels, qualityLevel)

			fmt.Printf("  blueprint %d graph size: %d\n", b.id, g.Size())
			fmt.Printf("  blueprint %d best score: %d\n", b.id, qualityLevel)
		}(b)
	}

	wg.Wait()

	fmt.Printf("  best quality levels: %v\n", qualityLevels)

	sum := 0
	for _, q := range qualityLevels {
		sum += q
	}

	fmt.Printf("  sum of best quality levels: %d\n", sum)
}

func SolvePartTwo() {
	fmt.Println()
	fmt.Println("solving part two...")
	maxMinutes = 32
	input := util.ReadInput(inputFile)
	blueprints := parseBlueprints(input)
	geodeCounts := []int{}

	var wg sync.WaitGroup
	wg.Add(len(blueprints))

	for _, b := range blueprints {
		go func(b blueprint) {
			defer wg.Done()
			fmt.Printf("  blueprint %d starting analysis\n", b.id)
			s := instantiateState(b)
			g := instantiateGraph(s)
			g, possibleGeodeCounts := expandGraph(s, g, []int{})
			geodeCount := util.MaxIntsSlice(possibleGeodeCounts)
			geodeCounts = append(geodeCounts, geodeCount)

			fmt.Printf("  blueprint %d graph size: %d\n", b.id, g.Size())
			fmt.Printf("  blueprint %d best score: %d\n", b.id, geodeCount)
		}(b)
	}

	wg.Wait()

	fmt.Printf("  best geode counts: %v\n", geodeCounts)

	product := 1
	for _, q := range geodeCounts {
		product *= q
	}

	fmt.Printf("  product of best geode counts: %d\n", product)
}
