package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func computeGammaRate(report []string) string {
	bits := len(report[0])
	totals := make([]int, bits)
	for _, entry := range report {
		for i := 0; i < bits; i++ {
			totals[i] += int(entry[i] - '0')
		}
	}

	gamma_rate := make([]int, bits)
	for i := 0; i < bits; i++ {
		switch {
		case totals[i] > len(report)/2:
			gamma_rate[i] = 1
		default:
			gamma_rate[i] = 0
		}
	}

	return strings.Replace(strings.Trim(fmt.Sprint(gamma_rate), "[]"), " ", "", -1)
}

func computeEpsilonRate(gamma_rate string) string {
	var epsilon_rate string
	for _, bit := range gamma_rate {
		if bit == '1' {
			epsilon_rate += "0"
		} else {
			epsilon_rate += "1"
		}
	}
	return epsilon_rate
}

func binaryToInt(s string) int {
	i, err := strconv.ParseInt(s, 2, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}

func filterValues(report []string, position int, mostCommon bool) []string {
	var sum int
	for i := range report {
		x, err := strconv.Atoi(string(report[i][position]))
		if err != nil {
			log.Fatal(err)
		}
		sum += x
	}

	var condition bool
	if mostCommon {
		condition = sum >= len(report)-sum
	} else {
		condition = sum < len(report)-sum
	}

	var filter_bit rune
	if condition {
		filter_bit = '1'
	} else {
		filter_bit = '0'
	}

	filtered_report := make([]string, 0, len(report))
	for i := range report {
		if report[i][position] == byte(filter_bit) {
			filtered_report = append(filtered_report, report[i])
		}
	}
	return filtered_report
}

func computeOxygenGeneratorRating(report []string) int {
	position := 0
	for {
		if len(report) == 1 {
			return binaryToInt(report[0])
		}
		report = filterValues(report, position, true)
		position++
	}
}

func computeCO2ScrubberRating(report []string) int {
	position := 0
	for {
		if len(report) == 1 {
			return binaryToInt(report[0])
		}
		report = filterValues(report, position, false)
		position++
	}
}

func main() {
	var report []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		report = append(report, scanner.Text())
	}
	gamma_rate := computeGammaRate(report)
	fmt.Printf("Gamma rate: %v (%d)\n", gamma_rate, binaryToInt(gamma_rate))

	epsilon_rate := computeEpsilonRate(gamma_rate)
	fmt.Printf("Epsilon rate: %v (%d)\n", epsilon_rate, binaryToInt(epsilon_rate))

	fmt.Printf("Power consumption: %d\n", binaryToInt(gamma_rate)*binaryToInt(epsilon_rate))

	oxygen_rating := computeOxygenGeneratorRating(report)
	fmt.Printf("Oxygen rating: %d\n", oxygen_rating)

	co2_scrubber_rating := computeCO2ScrubberRating(report)
	fmt.Printf("CO2 scrubber rating: %d\n", co2_scrubber_rating)

	fmt.Printf("Life support rating: %d\n", co2_scrubber_rating*oxygen_rating)
}
