package main

import (
	"fmt"
	"sort"
)

func main() {
	//ms := example()
	ms := puzzle()

	ins := make([]int, len(ms))
	for round := 0; round < 20; round++ {
		for mi, m := range ms {
			for _, i := range m.items {
				ins[mi] += 1
				w := m.operation(i)
				w = w / 3
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
