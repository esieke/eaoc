package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type list struct {
	val  int
	next *list
}

func (l *list) pop(n int) *list {
	buf := l.next
	for i := 0; i < n-1; i++ {
		buf = buf.next
	}
	ret := l.next
	l.next = buf.next
	buf.next = nil
	return ret
}

func (l *list) push(new *list) {
	next := l.next
	l.next = new
	for {
		if l.next == nil {
			l.next = next
			break
		}
		l = l.next
	}
}

func (l *list) find(val int) (*list, bool) {
	this := l
	for {
		if l.val == val {
			return l, true
		}
		if l.next == nil || l.next == this {
			break
		}
		l = l.next
	}
	return nil, false
}

func (l *list) offset(n int) *list {
	for i := 0; i < n; i++ {
		l = l.next
	}
	return l
}

func (l *list) print() {
	this := l
	for {
		fmt.Printf("%d", int(l.val))
		if l.next == nil || l.next == this {
			break
		}
		l = l.next
	}
	fmt.Printf("\n")
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	cups := &list{}
	cupsSlice := make([]*list, 1000000)
	nextCup := cups
	init := true
	line := 0
	for s.Scan() {
		l := []byte(s.Text())
		for _, b := range l {
			if init {
				cupsSlice[line] = cups
				nextCup.val = int(b) - 48
				nextCup.next = nextCup
				init = false
			} else {
				cupsSlice[line] = &list{
					val: int(b) - 48,
				}
				nextCup.push(cupsSlice[line])
			}
			nextCup = nextCup.next
			line++
		}
	}
	for i := 10; i <= 1000000; i++ {
		cupsSlice[i-1] = &list{
			val: i,
		}
		nextCup.push(cupsSlice[i-1])
		nextCup = nextCup.next
	}

	sort.SliceStable(cupsSlice, func(i, j int) bool {
		return cupsSlice[i].val < cupsSlice[j].val
	})
	nMoves := 10000000
	for i := 0; i < nMoves; i++ {
		pickUps := cups.pop(3)
		destVal := cups.val - 1

		for {
			if destVal < 1 {
				destVal = 1000000
			}
			_, found := pickUps.find(destVal)
			if !found {
				break
			}
			destVal--
		}

		i := sort.Search(len(cupsSlice)-1, func(i int) bool { return cupsSlice[i].val >= destVal })
		if !(i < len(cupsSlice) && cupsSlice[i].val == destVal) {
			panic("value must be in list")
		}
		cupsSlice[i].push(pickUps)
		cups = cups.next
		//cups.offset(9 - (i+1)%9).print()
	}

	i := sort.Search(len(cupsSlice)-1, func(i int) bool { return cupsSlice[i].val >= 1 })
	if !(i < len(cupsSlice) && cupsSlice[i].val == 1) {
		panic("value must be in list")
	}

	r1 := cupsSlice[i].pop(1).val
	r2 := cupsSlice[i].pop(1).val
	fmt.Println(r1 * r2)
}
