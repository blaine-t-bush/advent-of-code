package day11

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

var (
	itemsLength = 10
)

type throw struct {
	value  int
	target int
}

type operation struct {
	operator string
	value    string
}

type monkey struct {
	index       int
	items       []int
	operation   operation
	test        int
	pos         int // index of monkey to throw to if test is true
	neg         int // index of monkey to throw to if test is false
	inspections int
}

func (o operation) val() int {
	val, err := strconv.Atoi(o.value)
	util.CheckErr(err)
	return val
}

func (o operation) perform(old int) int {
	// get the value
	var val int
	if o.value == "old" {
		val = old
	} else {
		val = o.val()
	}

	// perform the operation
	var new int
	switch o.operator {
	case "*":
		new = old * val
	case "+":
		new = old + val
	default:
		log.Fatalf("could not perform operation %s on %d", o.operator, val)
	}

	return new
}

func parseMonkeys(lines []string) []monkey {
	monkeyCount := (len(lines) + 1) / 7
	monkeys := make([]monkey, monkeyCount)
	for i := range monkeys {
		startLine := i * 7
		endLine := startLine + 6 // non-inclusive, so first line after this monkey's section
		monkeys[i] = parseMonkey(lines[startLine:endLine])
	}

	return monkeys
}

func parseMonkey(lines []string) monkey {
	if len(lines) != 6 {
		log.Fatal("could not parse monkey")
	}

	return monkey{
		index:       parseMonkeyIndex(lines[0]),
		items:       parseMonkeyItems(lines[1]),
		operation:   parseMonkeyOperation(lines[2]),
		test:        parseMonkeyTest(lines[3]),
		pos:         parseMonkeyTrue(lines[4]),
		neg:         parseMonkeyFalse(lines[5]),
		inspections: 0,
	}
}

func parseMonkeyIndex(line string) int {
	r, err := regexp.Compile(`Monkey (\d+):`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 2 {
		log.Fatal("could not parse monkey index")
	}

	num, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	return num
}

func parseMonkeyItems(line string) []int {
	r, err := regexp.Compile(`  Starting items: ((?:\d+,? ?)+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 2 {
		log.Fatal("could not parse monkey starting items")
	}

	split := strings.Split(m[1], ", ")
	vals := make([]int, 0, itemsLength)
	for _, str := range split {
		val, err := strconv.Atoi(str)
		util.CheckErr(err)
		vals = append(vals, val)
	}

	return vals
}

func parseMonkeyOperation(line string) operation {
	r, err := regexp.Compile(`  Operation: new = old (\*|\+) (.+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 3 {
		log.Fatal("could not parse monkey operation")
	}

	return operation{
		operator: m[1],
		value:    m[2],
	}
}

func parseMonkeyTest(line string) int {
	r, err := regexp.Compile(`  Test: divisible by (\d+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 2 {
		log.Fatal("could not parse monkey test")
	}

	num, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	return num
}

func parseMonkeyTrue(line string) int {
	r, err := regexp.Compile(`    If true: throw to monkey (\d+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 2 {
		log.Fatal("could not parse monkey true condition")
	}

	num, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	return num
}

func parseMonkeyFalse(line string) int {
	r, err := regexp.Compile(`    If false: throw to monkey (\d+)`)
	util.CheckErr(err)

	m := r.FindStringSubmatch(line)
	if len(m) != 2 {
		log.Fatal("could not parse monkey false condition")
	}

	num, err := strconv.Atoi(m[1])
	util.CheckErr(err)

	return num
}

func round(worryDivisor int, monkeys []monkey, multiple int) []monkey {
	for i, monkey := range monkeys {
		throws := []throw{} // key: item, value: monkey to throw to

		for _, item := range monkey.items {
			// calculate worry level
			worryLevel := monkey.operation.perform(item) / worryDivisor

			// scale worry level down if needed
			remainder := worryLevel % multiple

			// check condition
			if worryLevel%monkey.test == 0 {
				throws = append(throws, throw{value: remainder, target: monkey.pos})
			} else {
				throws = append(throws, throw{value: remainder, target: monkey.neg})
			}

			// panic if worryLevel overflows to negative
			if worryLevel < 0 {
				log.Fatal("round: worry level overflowed")
			}
		}

		// update inspection count before removing items
		monkeys[i].inspections = monkey.inspections + len(monkey.items)

		// remove items. they're all being thrown
		monkeys[i].items = make([]int, 0, itemsLength)

		// perform throws
		for _, throw := range throws {
			monkeys[throw.target].items = append(monkeys[throw.target].items, throw.value)
		}
	}

	return monkeys
}

func rounds(worryDivisor int, count int, monkeys []monkey) []monkey {
	divisors := make([]int, len(monkeys))
	for i, m := range monkeys {
		divisors[i] = m.test
	}
	multiple := util.LeastMultiple(divisors)
	fmt.Printf("Product of test divisors: %d\n", multiple)

	for i := 0; i < count; i++ {
		monkeys = round(worryDivisor, monkeys, multiple)
	}

	return monkeys
}

func getMonkeyBusiness(monkeys []monkey) int {
	inspections := make([]int, len(monkeys))
	for i, monkey := range monkeys {
		inspections[i] = monkey.inspections
	}

	val1 := util.MaxIntsSlice(inspections)
	newInspections := util.RemoveIntFromSlice(val1, inspections)
	val2 := util.MaxIntsSlice(newInspections)

	return val1 * val2
}

func SolvePartOne(inputFile string) {
	fmt.Println("Part 1")
	input := util.ReadInput(inputFile)
	monkeys := rounds(3, 20, parseMonkeys(input))
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times\n", monkey.index, monkey.inspections)
	}

	fmt.Printf("Monkey business score: %d\n", getMonkeyBusiness(monkeys))
	fmt.Println()
}

func SolvePartTwo(inputFile string) {
	fmt.Println("Part 2")
	input := util.ReadInput(inputFile)
	monkeys := rounds(1, 10000, parseMonkeys(input))
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times\n", monkey.index, monkey.inspections)
	}

	fmt.Printf("Monkey business score: %d\n", getMonkeyBusiness(monkeys))
	fmt.Println()
}
