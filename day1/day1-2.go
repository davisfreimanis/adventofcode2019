package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		tmp := scanner.Text()
		i, err := strconv.Atoi(tmp)
		if err != nil {
			log.Fatal(err)
		}
		sum += calculateFuel(i)
	}
	fmt.Println(sum)
}

func calculateFuel(fuel int) int {
	totalFuel := 0
	for fuel >= 3 {
		fuel = fuel/3 - 2
		if fuel < 0 {
			fuel = 0
		}
		totalFuel += fuel
	}
	return totalFuel
}
