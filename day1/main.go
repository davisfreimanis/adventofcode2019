package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input2.txt")
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
		sum += (i/3 - 2)
	}
	fmt.Println(sum)
}
