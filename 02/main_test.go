package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

type Direction struct {
	direction string
	value     int
}

// implement parseSample
func parseSample(sample string) Direction {
	// split sample by space
	results := strings.Split(sample, " ")
	converted, err := strconv.Atoi(results[1])
	if err != nil {
		panic(err)
	}
	return Direction{results[0], converted}
}

func ReadSamplesFromFile(fileName string) ([]Direction, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var samples []Direction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		samples = append(samples, parseSample(scanner.Text()))
	}
	return samples, scanner.Err()
}

func AdventCheck(directions []Direction) int {
	var horizontal, depth int

	for _, direction := range directions {
		switch direction.direction {
		case "forward":
			horizontal += direction.value
		case "down":
			depth += direction.value
		case "up":
			depth -= direction.value
		}
	}

	return horizontal * depth
}

func TestAdventCheck(t *testing.T) {
	testRes, err := ReadSamplesFromFile("test.txt")
	if err != nil {
		t.Error(err)
	}

	expected := 150
	if AdventCheck(testRes) != expected {
		t.Errorf("Expected %d, got %d", expected, AdventCheck(testRes))
	}
}

func TestAdventActual(t *testing.T) {
	testRes, err := ReadSamplesFromFile("samples.txt")
	if err != nil {
		t.Error(err)
	}

	t.Logf("Result: %d", AdventCheck(testRes))
}
