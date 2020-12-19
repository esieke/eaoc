package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	id int
	//  or nodes
	refs  [][]int
	rules []string
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	rules := make(map[int]*rule)
	msgs := []string{}
	state := 0
	for s.Scan() {
		l := s.Text()

		if l == "" {
			state = 2
			continue
		}

		// rules
		or := 0
		r := rule{
			refs:  [][]int{},
			rules: make([]string, 0, 1024),
		}
		if state == 0 {
			l := strings.Split(l, ":")
			if len(l) != 2 {
				fmt.Println("parse line failed")
				continue
			}

			id, err := strconv.Atoi(l[0])
			if err != nil {
				fmt.Printf("convert to integer failed with error %v\n", err)
				continue
			}
			r.id = id
			l = strings.Split(l[1], " ")
			for _, v := range l {
				if v == "" {
					continue
				}
				if []byte(v)[0] == '"' {
					// check for rule
					s := strings.ReplaceAll(v, "\"", "")
					r.rules = append(r.rules, s)
					continue
				}

				if v == "|" {
					or = 1
					r.refs = append(r.refs, []int{})
					continue
				}

				// rule reference
				ref, err := strconv.Atoi(v)
				if err != nil {
					fmt.Printf("convert to integer failed with error %v\n", err)
					break

				}
				if len(r.refs) == 0 {
					r.refs = append(r.refs, []int{})
				}
				r.refs[or] = append(r.refs[or], ref)
			}
			rules[r.id] = &r
		}

		// messages
		if state == 2 {
			msgs = append(msgs, l)
		}
	}

	// do the work

	// search relevant rules -> ;) all rules are required ;)
	// rulesFilt := make(map[int]rule)
	// filter(0, rules, rulesFilt)

	// expand rules
	expandRules(0, rules)

	result := 0
	for _, m := range msgs {
		for _, r := range rules[0].rules {
			if strings.Compare(m, r) == 0 {
				result++
				break
			}
		}
	}

	fmt.Println(result)
}

func filter(id int, src, dst map[int]rule) {
	for _, ids := range src[id].refs {
		for _, i := range ids {
			dst[i] = src[i]
			filter(i, src, dst)
		}
	}
}

func expandRules(id int, r map[int]*rule) {
	for _, ids := range r[id].refs {
		for _, i := range ids {
			if len(r[i].rules) == 0 {
				expandRules(i, r)
			}
		}
	}

	for i := range r[id].refs {
		if i > 1 {
			panic("more then one or in rule")
		}
		// three rules
		if len(r[id].refs[0]) == 3 {
			for _, v0 := range r[r[id].refs[i][0]].rules {
				for _, v1 := range r[r[id].refs[i][1]].rules {
					for _, v2 := range r[r[id].refs[i][2]].rules {
						s := strings.Join([]string{v0, v1}, "")
						s = strings.Join([]string{s, v2}, "")
						rl := r[id]
						rl.rules = append(rl.rules, s)
					}
				}
			}
		}
		// two rules
		if len(r[id].refs[0]) == 2 {
			for _, v0 := range r[r[id].refs[i][0]].rules {
				for _, v1 := range r[r[id].refs[i][1]].rules {
					s := strings.Join([]string{v0, v1}, "")
					rl := r[id]
					rl.rules = append(rl.rules, s)
				}
			}
		}
		// one rule
		if len(r[id].refs[0]) == 1 {
			for _, v := range r[r[id].refs[i][0]].rules {
				rl := r[id]
				rl.rules = append(rl.rules, v)
			}
		}
	}
}
