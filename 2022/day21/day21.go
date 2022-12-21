package day21

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	util "github.com/blaine-t-bush/advent-of-code/util"
)

type monkey struct {
	name       string
	num        int
	calculated bool
	parent1    string
	parent2    string
	operation  string
}

func (m monkey) hasNum() bool {
	return m.calculated
}

func parseMonkeys(lines []string) []monkey {
	monkeys := make([]monkey, len(lines))
	for i, line := range lines {
		var m monkey
		// two options: num or operation
		if strings.ContainsAny(line, "+-*/") {
			// is operation
			r, err := regexp.Compile(`(\w+): (\w+) (\+|\-|\*|\/) (\w+)`)
			util.CheckErr(err)
			matches := r.FindStringSubmatch(line)
			if len(matches) != 5 {
				log.Fatal("parseMonkeys: could not parse operation line")
			}
			m.name = matches[1]
			m.parent1 = matches[2]
			m.parent2 = matches[4]
			m.operation = matches[3]
			m.calculated = false
		} else {
			// is num
			r, err := regexp.Compile(`(\w+): (\d+)`)
			util.CheckErr(err)
			matches := r.FindStringSubmatch(line)
			num, err := strconv.Atoi(matches[2])
			util.CheckErr(err)
			m.name = matches[1]
			m.num = num
			m.calculated = true
		}

		monkeys[i] = m
	}

	return monkeys
}

func createMap(monkeys []monkey) map[string]monkey {
	monkeysMap := make(map[string]monkey, len(monkeys))
	for _, m := range monkeys {
		monkeysMap[m.name] = m
	}

	return monkeysMap
}

func calcNum(name string, monkeysMap map[string]monkey) (int, map[string]monkey) {
	var num int
	m := monkeysMap[name]
	if m.hasNum() {
		num = m.num
	} else {
		val1, monkeysMap := calcNum(m.parent1, monkeysMap)
		val2, monkeysMap := calcNum(m.parent2, monkeysMap)
		switch m.operation {
		case "+":
			num = val1 + val2
		case "-":
			num = val1 - val2
		case "*":
			num = val1 * val2
		case "/":
			num = val1 / val2
		default:
			log.Fatal("calcNum: could not parse operation")
		}
	}

	m.num = num
	m.calculated = true
	monkeysMap[name] = m
	return num, monkeysMap
}

func copyMap(monkeysMap map[string]monkey) map[string]monkey {
	copy := make(map[string]monkey, len(monkeysMap))
	for k, v := range monkeysMap {
		copy[k] = v
	}

	return copy
}

func SolvePartOne(inputFile string) {
	input := util.ReadInput(inputFile)
	monkeys := parseMonkeys(input)
	monkeysMap := createMap(monkeys)
	root, _ := calcNum("root", monkeysMap)
	fmt.Println(root)
}

func SolvePartTwo(inputFile string) {
	input := util.ReadInput(inputFile)
	monkeys := parseMonkeys(input)
	monkeysMap := createMap(monkeys)
	name1, name2 := monkeysMap["root"].parent1, monkeysMap["root"].parent2

	// insert new guess
	var mapCopy map[string]monkey
	var updated monkey

	for humn := 3330805250000; humn < 3330805300000; humn += 1 {
		// insert new guess
		mapCopy = copyMap(monkeysMap)
		updated = mapCopy["humn"]
		updated.num = humn
		mapCopy["humn"] = updated

		// update map with values
		_, mapCopy := calcNum("root", mapCopy)

		// check for match
		if mapCopy[name1].num == mapCopy[name2].num {
			fmt.Printf("\nsuccess with guess %d\n", humn)
			break
		}

		fmt.Printf("current guess %d for diff of %d\n", humn, mapCopy[name1].num-mapCopy[name2].num)
	}
}
