package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("numbers")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	numberStrings := strings.Split(string(data), "\n")
	var numbers []int
	for _, i := range numberStrings {
		if i == "" {
			break
		}
		num, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		numbers = append(numbers, num)
	}
	fmt.Println(part1(numbers))
	fmt.Println(part2(numbers))
}

func part1(numbers []int) int {
	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == 2020 {
				return i * j
			}
		}
	}
	return 0
}

func part2(numbers []int) int {
	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == 2020 {
					return i * j * k
				}
			}
		}
	}
	return 0
}
