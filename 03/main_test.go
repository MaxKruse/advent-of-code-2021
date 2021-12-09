package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

type Sample struct {
	bits    string
	bitsNum int
}

// implement parseSample
func parseSample(sample string) Sample {
	// split sample by space
	converted, err := strconv.Atoi(sample)
	if err != nil {
		panic(err)
	}
	return Sample{sample, converted}
}

func ReadSamplesFromFile(fileName string) ([]Sample, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var samples []Sample
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		samples = append(samples, parseSample(scanner.Text()))
	}
	return samples, scanner.Err()
}

func bitArrayToDecimal(bitArray []int) int64 {
	mostCommonBitsString := ""
	for _, bit := range bitArray {
		mostCommonBitsString += strconv.Itoa(bit)
	}

	log.Printf("mostCommonBitsString: %s", mostCommonBitsString)

	mostCommonBitsNum, _ := strconv.ParseInt(mostCommonBitsString, 2, 64)

	return mostCommonBitsNum
}

func AdventCheck(samples []Sample) int64 {
	// array of bit counts of length len(samples[0])
	oneCounts := make([]int, len(samples[0].bits))
	zeroCounts := make([]int, len(samples[0].bits))

	mostCommonBits := make([]int, len(samples[0].bits))

	// parse samples
	for _, sample := range samples {

		for i, bitChar := range sample.bits {
			if bitChar == '1' {
				oneCounts[i]++
			} else {
				zeroCounts[i]++
			}
		}
	}

	// fill mostcommonbits depending on what is more common at that index
	for i := range oneCounts {
		if oneCounts[i] > zeroCounts[i] {
			mostCommonBits[i] = 1
		} else {
			mostCommonBits[i] = 0
		}
	}

	// Convert the most common to an actual int
	gamma := bitArrayToDecimal(mostCommonBits)
	log.Printf("gamma: %d", gamma)

	// get the epsilon

	// bitflip every bit in mostCommonBits
	for i := range mostCommonBits {
		if mostCommonBits[i] == 1 {
			mostCommonBits[i] = 0
		} else {
			mostCommonBits[i] = 1
		}
	}

	epsilon := bitArrayToDecimal(mostCommonBits)
	log.Printf("epsilon: %d", epsilon)

	return gamma * epsilon
}

func TestAdventCheck(t *testing.T) {
	testRes, err := ReadSamplesFromFile("test.txt")
	if err != nil {
		t.Error(err)
	}

	var expected int64 = 198
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
