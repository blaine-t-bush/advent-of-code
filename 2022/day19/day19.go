package day19

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"
	"time"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

const (
	geodeWeight    = 1.0
	obsidianWeight = 0.1
	clayWeight     = 0.01
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

func (s state) heuristic() float64 {
	return geodeWeight * float64(s.resources.geodes+s.bots.geodeBots*(maxMinutes-s.minute)) // + obsidianWeight*float64(s.resources.obsidian+s.bots.obsidianBots*maxMinutes-s.minute) + clayWeight*float64(s.resources.clay+s.bots.clayBots*maxMinutes-s.minute)
}

func averageHeuristic(states []state) float64 {
	sum := 0.0
	for _, s := range states {
		sum += s.heuristic()
	}
	return sum / float64(len(states))
}

func filterStates(states []state) []state {
	// filter out states that are strictly not better than all other states
	cutoff := 0.5 * averageHeuristic(states)
	filtered := []state{}
	for _, s := range states {
		if s.heuristic() >= cutoff {
			filtered = append(filtered, s)
		}
	}

	return filtered
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

func getMaxGeodeCount(start state, minOffset int) int {
	// for each terminus state in graph (i.e. has no edges from),
	// get next possible states and add them to graph.
	// repeat until we have added states up through maxMinutes.
	terminusStates := []state{start}
	maxGeodeCount := 0
	maxExpectedGeodeCount := 0
	for {
		// if no more terminus states, we've reached maxMinutes.
		if len(terminusStates) == 0 {
			break
		}

		fmt.Printf("  blueprint %d minute %d\n", start.blueprint.id, terminusStates[0].minute)

		// for each terminus state, get possible next-minute states.
		// for each possible next-minute state, add it to list if it's not
		// strictly worse than a previously-explore state.
		futureStates := []state{}
		for _, terminusState := range terminusStates {
			futureStatesTemp, _ := getFutureStates(terminusState)
			for _, futureState := range futureStatesTemp {
				// optimizations:
				//   1. if at maxMinutes-4 with no geode bots, assume fruitless.
				//      same for obsidian and clay at -5 and -6 minutes.

				if futureState.minute >= maxMinutes-minOffset && futureState.bots.geodeBots == 0 {
					continue
				}
				if futureState.minute >= maxMinutes-minOffset-2 && futureState.bots.obsidianBots == 0 {
					continue
				}
				if futureState.minute >= maxMinutes-minOffset-4 && futureState.bots.clayBots == 0 {
					continue
				}

				futureStates = append(futureStates, futureState)
				if futureState.resources.geodes > maxGeodeCount {
					maxGeodeCount = futureState.resources.geodes
					fmt.Printf("  blueprint %d current maxGeodeCount %d with state %v\n", start.blueprint.id, maxGeodeCount, futureState)
				}
				if futureState.bots.obsidianBots >= futureState.blueprint.maxObsidianCost() {
					expectedGeodeCount := futureState.resources.geodes
					for m := futureState.minute + 1; m <= maxMinutes; m++ {
						expectedGeodeCount += futureState.bots.geodeBots + 1
					}
					if expectedGeodeCount > maxExpectedGeodeCount {
						maxExpectedGeodeCount = expectedGeodeCount
					}
					fmt.Printf("  blueprint %d sustaining obsidian engine reached with state %v. expected final geode count %d\n", start.blueprint.id, futureState, expectedGeodeCount)
				}
			}
		}

		if len(futureStates) > 0 && futureStates[0].minute >= 20 {
			terminusStates = filterStates(futureStates)
		} else {
			terminusStates = futureStates
		}
	}

	return maxGeodeCount
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

func SolvePartOne(inputFile string) {
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
			maxGeodeCount := getMaxGeodeCount(s, 4)
			qualityLevel := maxGeodeCount * b.id
			qualityLevels = append(qualityLevels, qualityLevel)

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

func SolvePartTwo(inputFile string) {
	startTimeUnixMilli := time.Now().UnixMilli()
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
			maxGeodeCount := getMaxGeodeCount(s, 4)
			geodeCounts = append(geodeCounts, maxGeodeCount)

			fmt.Printf("  blueprint %d best score: %d\n", b.id, maxGeodeCount)
		}(b)
	}

	wg.Wait()

	fmt.Printf("  best geode counts: %v\n", geodeCounts)

	product := 1
	for _, q := range geodeCounts {
		product *= q
	}

	fmt.Printf("  product of best geode counts: %d\n", product)
	fmt.Printf("runtime: %d milliseconds", time.Now().UnixMilli()-startTimeUnixMilli)
}
