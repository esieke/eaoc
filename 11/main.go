package main

import (
	"fmt"
	"sort"
)

func main() {
	//ms := example()
	ms := puzzle()

	ins := make([]int, len(ms))
	for round := 0; round < 10000; round++ {
		for mi, m := range ms {
			for _, i := range m.items {
				ins[mi] += 1
				w := m.operation(i)
				//w = w % (17 * 13 * 19 * 23) // example
				w = w % (2 * 7 * 3 * 17 * 11 * 19 * 5 * 13) // puzzle
				throwTo := m.test(w)
				ms[throwTo].items = append(ms[throwTo].items, w)
			}
			ms[mi].items = []int{}
		}
	}
	sort.Ints(ins)
	r := ins[len(ins)-1] * ins[len(ins)-2]

	fmt.Printf("%d\n", r)
}
