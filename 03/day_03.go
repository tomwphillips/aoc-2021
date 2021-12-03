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
}
