package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type id struct {
	ID uint64
	i  uint64
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	var timeNow uint64
	IDs := make([]id, 0, 1024)

	for s.Scan() {
		l := s.Text()
		if timeNow == 0 {
			t, err := strconv.Atoi(l)
			timeNow = uint64(t)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			continue
		}
		sIDs := strings.Split(l, ",")
		for i, sID := range sIDs {
			if sID != "x" {
				ID, err := strconv.Atoi(sID)
				if err != nil {
					fmt.Println(err)
					os.Exit(2)
				}
				IDs = append(IDs, id{ID: uint64(ID), i: uint64(i)})
			}
		}

	}

	// Mode one: mode == false / time increment f(max(ID1, ID2, ... IDn))
	// Mode two: mode == true  / time increment f(lsm(ID1, ID2, ... IDn))
	// In mode two we search the time after which have the exact same state.
	// The offset should not further increase! n_Rounds1 * ID1 = n_Rounds2 * ID2 = t
	var mode bool

	// max step size for mode one
	bus := struct {
		iD  uint64
		dep uint64
	}{}
	for _, ID := range IDs {
		if ID.ID > bus.iD {
			bus.iD = ID.ID
			bus.dep = ID.i
		}
	}

	timeNow = 0
	timeModeOne := (100000000000000 / bus.iD) * bus.iD
	for {
		var timeModeTwo uint64
		var valid uint64
		valid = 0
		for _, ID := range IDs {
			nextDep := ((timeNow + ID.i + 1) % ID.ID)
			if nextDep == 0 {
				valid++
				if valid == 1 {
					timeModeTwo = ID.ID
				} else {
					// update mode two time step
					timeModeTwo = calcLcm(timeModeTwo, ID.ID)
					mode = true
				}
			}
		}
		if valid == uint64(len(IDs)) {
			break
		}
		if mode && valid > 1 {
			timeNow += timeModeTwo
		} else {
			timeModeOne += bus.iD
			timeNow = timeModeOne - bus.dep - 1
		}

	}
	fmt.Println(timeNow + IDs[0].i + 1)
}

func ggt(m, n uint64) uint64 {
	if n == 0 {
		return m
	}
	return ggt(n, m%n)
}

// least common multiple,
func calcLcm(m, n uint64) uint64 {
	o := ggt(m, n)
	p := (m * n) / o
	return p
}
