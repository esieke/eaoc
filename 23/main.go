package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	nextCup := cups
	init := true
	for s.Scan() {
		l := []byte(s.Text())
		for _, b := range l {
			if init {
				nextCup.val = int(b) - 48
				nextCup.next = nextCup
				init = false
			} else {
				nextCup.push(&list{
					val: int(b) - 48,
				})
			}
			nextCup = nextCup.next
		}
	}
	for i := 10; i <= 1000000; i++ {
		nextCup.push(&list{
			val: i,
		})
		nextCup = nextCup.next
	}

	nMoves := 10000000
	for i := 0; i < nMoves; i++ {
		pickUps := cups.pop(3)
		destVal := cups.val - 1

		for {
			if destVal < 1 {
				destVal = 9
			}
			_, found := pickUps.find(destVal)
			if !found {
				break
			}
			destVal--
		}

		destList, found := cups.find(destVal)
		if !found {
			panic("value must be in list")
		}
		destList.push(pickUps)
		cups = cups.next

		if i%1000 == 0 {
			fmt.Println(i)
		}

		//cups.offset(9 - (i+1)%9).print()
	}
	result, _ := cups.find(1)
	result.print()
	r1 := result.pop(1).val
	r2 := result.pop(1).val
	fmt.Println(r1 * r2)
}
