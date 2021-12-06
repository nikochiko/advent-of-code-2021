package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/nikochiko/advent-of-code/solution"
)

var solutions = map[int]map[int]func(){
	1: {
		1: solution.Day1First,
		2: solution.Day1Second,
	},
	2: {
		1: solution.Day2First,
		2: solution.Day2Second,
	},
	3: {
		1: solution.Day3First,
	},
}

func main() {
	if len(os.Args) != 3 {
		err := errors.New("should be run like `./advent-of-code DAY_NUMBER PROBLEM_NUMBER`, followed by any inputs in the standard input")
		panic(err)
	}

	dayNumberInput := os.Args[1]
	problemNumberInput := os.Args[2]

	dayNumber, err := strconv.Atoi(dayNumberInput)
	if err != nil {
		panic(err)
	}

	problemNumber, err := strconv.Atoi(problemNumberInput)
	if err != nil {
		panic(err)
	}

	daysSolutions, ok := solutions[dayNumber]
	if !ok {
		err := fmt.Errorf("no solutions found for day %d", dayNumber)
		panic(err)
	}

	if soln, ok := daysSolutions[problemNumber]; ok {
		soln()
	} else {
		err := fmt.Errorf("no solution found for problem %d of day %d", problemNumber, dayNumber)
		panic(err)
	}
}
