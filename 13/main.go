package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var timeNow int
	IDs := make([]int, 0, 1024)
	for s.Scan() {
		l := s.Text()
		if timeNow == 0 {
			t, err := strconv.Atoi(l)
			timeNow = t
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			continue
		}
		sIDs := strings.Split(l, ",")
		for _, sID := range sIDs {
			if sID != "x" {
				ID, err := strconv.Atoi(sID)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				IDs = append(IDs, ID)
			}
		}

	}

	bus := struct {
		iD  int
		dep int
	}{
		dep: 1000,
	}

	for _, ID := range IDs {
		nextDep := ID - (timeNow % ID)
		if nextDep < bus.dep {
			bus.iD = ID
			bus.dep = nextDep
		}
	}
	fmt.Println(bus.iD * bus.dep)
}
