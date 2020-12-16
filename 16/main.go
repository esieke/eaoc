package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type lim struct {
	min int
	max int
}

type rule struct {
	name string
	lims []lim
}

// nearby tickets:

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	rules := make([]rule, 0, 1024)
	myTicket := make([]int, 0, 1024)
	nearbyTickets := make([][]int, 0, 1024)

	state := 0
	for s.Scan() {
		l := s.Text()
		if l == "your ticket:" {
			state = 1
			continue
		}
		if l == "nearby tickets:" {
			state = 2
			continue
		}

		if state == 0 {
			// class: 1-3 or 5-7
			ls := strings.Split(l, ":")
			if len(ls) != 2 {
				fmt.Println("parse rules failed")
				continue
			}
			r := rule{
				name: ls[0],
			}
			ls = strings.Split(ls[1], " ")
			for _, raw := range ls {
				if raw != "or" && raw != "" {
					rawMinMax := strings.Split(raw, "-")
					if len(rawMinMax) != 2 {
						fmt.Println("parse limit raw failed")
						continue
					}
					min, err := strconv.Atoi(rawMinMax[0])
					if err != nil {
						fmt.Println("parse min rule failed")
						continue
					}
					max, err := strconv.Atoi(rawMinMax[1])
					if err != nil {
						fmt.Println("parse max rule failed")
						continue
					}
					r.lims = append(r.lims, lim{
						min: min,
						max: max,
					})
				}
			}
			rules = append(rules, r)
			continue
		}
		if state == 1 {
			if l == "" {
				continue
			}
			ls := strings.Split(l, ",")
			if len(ls) > 0 {
				for _, num := range ls {
					n, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println("parse ticket failed")
						continue
					}
					myTicket = append(myTicket, n)
				}
			}
			continue
		}
		if state == 2 {
			if l == "" {
				continue
			}
			ls := strings.Split(l, ",")
			if len(ls) > 0 {
				t := make([]int, 0, 1024)
				for _, num := range ls {
					n, err := strconv.Atoi(num)
					if err != nil {
						fmt.Println("parse ticket failed")
						continue
					}
					t = append(t, n)
				}
				nearbyTickets = append(nearbyTickets, t)
			}
			continue
		}
	}

	result := 0
	for _, nt := range nearbyTickets {
		for _, val := range nt {
			result += validate(val, rules)
		}
	}
	fmt.Println(result)
}

func validate(val int, rules []rule) int {
	for _, rule := range rules {
		for _, lim := range rule.lims {
			if val >= lim.min && val <= lim.max {
				return 0
			}

		}
	}
	return val
}
