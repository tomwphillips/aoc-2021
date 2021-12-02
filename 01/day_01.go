package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var last_measurement int = 0
	var depth_increases int = -1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		measurement, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if measurement > last_measurement {
			depth_increases+=1
		}
		last_measurement = measurement
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(depth_increases)
}
