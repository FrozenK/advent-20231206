package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time     int
	distance int
}

func main() {
	f, err := os.Open("input4.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	races := make([]race, 4)
	fmt.Println(races)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		splits := strings.Split(line, ":")
		switch splits[0] {
		case "Time":
			for i, v := range strings.Fields(splits[1]) {
				vn, _ := strconv.Atoi(v)
				races[i].time = vn
			}
		case "Distance":
			for i, v := range strings.Fields(splits[1]) {
				vn, _ := strconv.Atoi(v)
				races[i].distance = vn
			}
		}
	}

	sum := 0
	for _, r := range races {
		if r.time == 0 {
			continue
		}
		num := 0
		for dh := 1; dh < r.time; dh++ {
			timeToMove := r.time - dh
			if r.distance < dh*timeToMove {
				num++
			}
		}
		if num > 0 {
			if sum == 0 {
				sum = num
			} else {
				sum *= num
			}
		}
	}
	fmt.Println(races)
	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
