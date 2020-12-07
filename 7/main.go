package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Content is the content of one bag
type Content struct {
	color string
	n     int
}

// Bag is the main object
type Bag struct {
	color    string
	contents []Content
}

// Record to log the state of the bag tree
type Record struct {
	state map[string]bool
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	bags := make([]Bag, 0, 1000)

	for s.Scan() {
		l := s.Text()
		re, err := regexp.Compile("^([a-z]+ [a-z]+) bags contain (.+)$")
		if err != nil {
			fmt.Printf("error regexp compile\n")
			os.Exit(1)
		}
		r := re.FindStringSubmatch(l)
		if len(r) != 3 {
			fmt.Printf("error regexp find string submatch\n")
			os.Exit(1)
		}
		// bag name
		bag := Bag{color: r[1]}
		tail := r[2]
		for {
			re, err := regexp.Compile("([0-9]+) ([a-z]+ [a-z]+)(.*)")
			if err != nil {
				fmt.Printf("error regexp compile\n")
				os.Exit(1)
			}
			r := re.FindStringSubmatch(tail)
			if len(r) != 4 {
				// done
				break
			}
			if bag.contents == nil {
				bag.contents = make([]Content, 0, 100)
			}
			n, err := strconv.Atoi(r[1])
			if err != nil {
				fmt.Printf("error string to int\n")
				os.Exit(1)
			}
			bag.contents = append(bag.contents, Content{color: r[2], n: n})
			tail = r[3]
		}
		bags = append(bags, bag)
	}

	r := NewRecord()
	r.findParents("shiny gold", bags)
	fmt.Printf("result puzzle one: %d\n", r.getSum())
}

// NewRecord init Record struct
func NewRecord() *Record {
	return &Record{
		state: make(map[string]bool),
	}
}

func (r *Record) findParents(color string, bags []Bag) {
	for _, bag := range bags {
		for _, content := range bag.contents {
			if content.color == color {
				r.state[bag.color] = true
				r.findParents(bag.color, bags)
			}
		}
	}
}

func (r *Record) getSum() (res int) {
	for _, s := range r.state {
		if s {
			res++
		}
	}
	return res
}
