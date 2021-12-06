package solution

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Day1First() {
	reader := bufio.NewReader(os.Stdin)

	previous := int(^uint(0) >> 1)

	fmt.Printf("Largest integer: %d\n", previous)

	result := 0
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			panic(err)
		}

		if previous < num {
			result++
		}

		previous = num
	}

	fmt.Println(result)
}

func Day1Second() {
	reader := bufio.NewReader(os.Stdin)

	largestInt := int(^uint(0) >> 1)
	last3 := make([]int, 3)

	addToLast3 := func(num int) {
		last3[2] = last3[1]
		last3[1] = last3[0]
		last3[0] = num
	}

	for i := 0; i < 3; i++ {
		addToLast3(largestInt)
	}

	result := 0
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		num, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			panic(err)
		}

		if last3[2] < num {
			result++
		}

		addToLast3(num)
	}

	fmt.Println(result)
}

func Day2First() {
	type pos struct {
		Horizontal int
		Depth      int
	}
	position := pos{0, 0}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		trimmedInput := strings.TrimSpace(input)
		parts := strings.SplitN(trimmedInput, " ", 2)
		command, inputMagnitude := parts[0], parts[1]

		magnitude, err := strconv.Atoi(inputMagnitude)
		if err != nil {
			panic(err)
		}

		switch command {
		case "forward":
			position.Horizontal += magnitude
		case "up":
			position.Depth -= magnitude
		case "down":
			position.Depth += magnitude
		default:
			err := fmt.Errorf("unknown command: %s", command)
			panic(err)
		}
	}

	fmt.Printf("Position: {%d, %d}\n", position.Horizontal, position.Depth)

	answer := position.Horizontal * position.Depth
	fmt.Printf("Answer: %d\n", answer)
}

func Day2Second() {
	type pos struct {
		Horizontal int
		Depth      int
		Aim        int
	}
	position := pos{0, 0, 0}

	reader := bufio.NewReader(os.Stdin)

	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		trimmedInput := strings.TrimSpace(input)
		parts := strings.SplitN(trimmedInput, " ", 2)
		command, inputMagnitude := parts[0], parts[1]

		magnitude, err := strconv.Atoi(inputMagnitude)
		if err != nil {
			panic(err)
		}

		switch command {
		case "forward":
			position.Horizontal += magnitude
			position.Depth += position.Aim * magnitude
		case "up":
			position.Aim -= magnitude
		case "down":
			position.Aim += magnitude
		default:
			err := fmt.Errorf("unknown command: %s", command)
			panic(err)
		}
	}

	fmt.Printf("Position and Aim: {%d, %d, %d}\n", position.Horizontal, position.Depth, position.Aim)

	answer := position.Horizontal * position.Depth
	fmt.Printf("Answer: %d\n", answer)
}

func Day3First() {
	var onBitOccurences []int

	logCount := 0

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		logCount++

		binaryLog := strings.TrimSpace(input)
		if onBitOccurences == nil {
			onBitOccurences = make([]int, len(binaryLog))
		}

		for i, ch := range binaryLog {
			if ch == '1' {
				onBitOccurences[i]++
			} else if ch == '0' {
				// do nothing
			} else {
				err := fmt.Errorf("unknown element in binary log: %c", ch)
				panic(err)
			}
		}
	}

	gammaRateBinary := ""

	for _, onCount := range onBitOccurences {
		if onCount >= logCount/2.0 {
			gammaRateBinary += "1"
		} else {
			gammaRateBinary += "0"
		}
	}

	gammaRate := 0
	_, err := fmt.Sscanf(gammaRateBinary, "%b", &gammaRate)
	if err != nil {
		panic(err)
	}

	allOnesBinary := strings.Repeat("1", len(onBitOccurences))
	allOnes := 0
	_, err = fmt.Sscanf(allOnesBinary, "%b", &allOnes)
	if err != nil {
		panic(err)
	}

	epsilonRate := allOnes ^ gammaRate
	fmt.Printf("Gamma Rate: %d, Epsilon Rate: %d\n", gammaRate, epsilonRate)

	answer := gammaRate * epsilonRate
	fmt.Printf("Answer: %d\n", answer)
}
