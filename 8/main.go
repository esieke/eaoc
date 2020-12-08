package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instr struct {
	addr int
	oper string
	arg  int
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	stack := []instr{}
	addr := 0
	for s.Scan() {
		l := s.Text()
		s := strings.Split(l, " ")
		if len(s) != 2 {
			fmt.Printf("unexpected format skip line\n")
			continue
		}
		arg, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Printf("argument must be of type int skip line with error %v\n", err)
		}
		stack = append(stack, instr{
			addr: addr,
			oper: s[0],
			arg:  arg,
		})
		addr++
	}

	// puzzle one
	accu := 0
	trace := make(map[int]bool)
	addr = 0
	for {
		// second time
		if trace[addr+1] {
			break
		}

		if stack[addr].oper == "jmp" {
			addr += stack[addr].arg
			continue
		}
		if stack[addr].oper == "acc" {
			accu += stack[addr].arg
		}
		// inc addr for nop and acc instruction
		addr++
		trace[addr] = true
	}
	fmt.Printf("result of puzzle one: %d\n", accu)
}
