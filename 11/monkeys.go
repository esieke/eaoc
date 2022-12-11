package main

type operation func(int) int
type test func(int) int

type monkey struct {
	items     []int
	operation operation
	test      test
}

func example() []monkey {
	return []monkey{
		{
			items:     []int{79, 98},
			operation: func(i int) int { return i * 19 },
			test: func(i int) int {
				if i%23 == 0 {
					return 2
				}
				return 3
			},
		},
		{
			items:     []int{54, 65, 75, 74},
			operation: func(i int) int { return i + 6 },
			test: func(i int) int {
				if i%19 == 0 {
					return 2
				}
				return 0
			},
		},
		{
			items:     []int{79, 60, 97},
			operation: func(i int) int { return i * i },
			test: func(i int) int {
				if i%13 == 0 {
					return 1
				}
				return 3
			},
		},
		{
			items:     []int{74},
			operation: func(i int) int { return i + 3 },
			test: func(i int) int {
				if i%17 == 0 {
					return 0
				}
				return 1
			},
		},
	}
}

func puzzle() []monkey {
	return []monkey{
		{
			items:     []int{80},
			operation: func(i int) int { return i * 5 },
			test: func(i int) int {
				if i%2 == 0 {
					return 4
				}
				return 3
			},
		},
		{
			items:     []int{75, 83, 74},
			operation: func(i int) int { return i + 7 },
			test: func(i int) int {
				if i%7 == 0 {
					return 5
				}
				return 6
			},
		},
		{
			items:     []int{86, 67, 61, 96, 52, 63, 73},
			operation: func(i int) int { return i + 5 },
			test: func(i int) int {
				if i%3 == 0 {
					return 7
				}
				return 0
			},
		},
		{
			items:     []int{85, 83, 55, 85, 57, 70, 85, 52},
			operation: func(i int) int { return i + 8 },
			test: func(i int) int {
				if i%17 == 0 {
					return 1
				}
				return 5
			},
		},
		{
			items:     []int{67, 75, 91, 72, 89},
			operation: func(i int) int { return i + 4 },
			test: func(i int) int {
				if i%11 == 0 {
					return 3
				}
				return 1
			},
		},
		{
			items:     []int{66, 64, 68, 92, 68, 77},
			operation: func(i int) int { return i * 2 },
			test: func(i int) int {
				if i%19 == 0 {
					return 6
				}
				return 2
			},
		},
		{
			items:     []int{97, 94, 79, 88},
			operation: func(i int) int { return i * i },
			test: func(i int) int {
				if i%5 == 0 {
					return 2
				}
				return 7
			},
		},
		{
			items:     []int{77, 85},
			operation: func(i int) int { return i + 6 },
			test: func(i int) int {
				if i%13 == 0 {
					return 4
				}
				return 0
			},
		},
	}
}
