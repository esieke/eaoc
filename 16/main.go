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
	id   int
}

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
	ruleID := 0
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
			if l == "" {
				continue
			}
			ls := strings.Split(l, ":")
			if len(ls) != 2 {
				fmt.Println("parse rules failed")
				continue
			}
			r := rule{
				name: ls[0],
				id:   -1, //ruleID,
			}
			ruleID++
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

	invalid := make([]int, 0, 1024)
	for tid, nt := range nearbyTickets {
		for _, val := range nt {
			if r := validate(val, rules); r > 0 {
				invalid = append(invalid, tid)
			}
		}
	}
	validTickets := popAll(invalid, nearbyTickets)
	findAndSetIds(validTickets, rules)

	result := 1
	for _, r := range rules {
		if strings.Contains(r.name, "departure") {
			result *= myTicket[r.id]
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

func popAll(all []int, s [][]int) [][]int {
	l := len(all) - 1
	for i := range all {
		s = append(s[:all[l-i]], s[all[l-i]+1:]...)
	}
	return s
}

func findAndSetIds(tickets [][]int, rules []rule) {
	fountAllCtr := 0
	for {
		for rid := range rules {
			founds := 0
			foundTid := 0
			if rules[rid].id != -1 {
				continue
			}
			for tid := range tickets[0] {
				valid := true
				for _, ticket := range tickets {
					if v := validate(ticket[tid], rules[rid:rid+1]); v > 0 {
						valid = false
						break
					}
				}
				if valid {
					setCtr := true
					for _, r := range rules {
						if r.id == tid {
							setCtr = false
							break
						}
					}
					if setCtr {
						founds++
						foundTid = tid
					}
				}
			}
			if founds == 1 {
				rules[rid].id = foundTid
				fountAllCtr++
			}
		}
		if fountAllCtr >= len(rules) {
			break
		}
	}
}
