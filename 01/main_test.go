package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func AdventCheck(samples []int) int {
	// set previous to the highest possible int
	previous := int(^uint(0) >> 1)

	// counter
	counter := 0

	// loop through the samples
	for _, sample := range samples {
		// if the sample is higher than the previous, count up
		if sample > previous {
			counter++
		}

		// set previous to the current sample
		previous = sample
	}

	return counter
}

func TestAdventCheck(t *testing.T) {
	samples := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	expected := 7

	if actual := AdventCheck(samples); actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	} else {
		t.Logf("AdventCheck() = %d", actual)
	}
}

func ReadSamplesFromFile(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var samples []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		samples = append(samples, parseSample(scanner.Text()))
	}
	return samples, scanner.Err()
}

// implement parseSample
func parseSample(sample string) int {
	val, err := strconv.Atoi(sample)
	if err != nil {
		panic(err)
	}
	return val
}

func TestAdventActual(t *testing.T) {

	// read samples from file
	samples, err := ReadSamplesFromFile("samples.txt")
	if err != nil {
		t.Errorf("Error reading samples from file: %s", err)
	}

	actual := AdventCheck(samples)

	t.Logf("AdventCheck() = %d", actual)

}
