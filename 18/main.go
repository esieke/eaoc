package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type tokenEnum int

const (
	TokenInt tokenEnum = iota
	TokenMul tokenEnum = iota
	TokenAdd tokenEnum = iota
	TokenBB  tokenEnum = iota
	TokenBE  tokenEnum = iota
	TokenEnd tokenEnum = iota
)

type stateEnum int

const (
	StateInit      stateEnum = iota
	StateRun       stateEnum = iota
	StateEnd       stateEnum = iota
	StateMul       stateEnum = iota
	StateAdd       stateEnum = iota
	StateWaitForOp stateEnum = iota
)

type token struct {
	tokenType tokenEnum
	val       interface{}
}

type expression struct {
	ptr    int
	tokens []token
}

func (e *expression) getNextToken() *token {
	e.ptr++
	return &e.tokens[e.ptr-1]
}

func (e *expression) calc() int {
	num := 0
	var state stateEnum
	state = StateInit
	for {
		t := e.getNextToken()
		switch t.tokenType {
		case TokenInt:
			switch state {
			case StateInit:
				num = t.val.(int)
			case StateMul:
				num *= t.val.(int)
			case StateAdd:
				num += t.val.(int)
			default:
				panic("syntax error")
			}
			state = StateWaitForOp
		case TokenAdd:
			if state != StateWaitForOp {
				panic("syntax error operator add")
			}
			state = StateAdd
		case TokenMul:
			if state != StateWaitForOp {
				panic("syntax error operator mul")
			}
			state = StateMul
		case TokenBB:
			switch state {
			case StateInit:
				num = e.calc()
			case StateMul:
				num *= e.calc()
			case StateAdd:
				num += e.calc()
			default:
				panic("syntax error")
			}
			state = StateWaitForOp
		case TokenBE:
			return num
		case TokenEnd:
			return num
		}
	}
}

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	s := bufio.NewScanner(input)

	e := &expression{}
	var scanNum []byte
	result := 0
	for s.Scan() {
		l := []byte(s.Text())
		stateScanNum := 0
		for _, b := range l {
			// 0-9
			isDig := b >= 48 && b <= 57
			if isDig || stateScanNum == 1 {
				if isDig {
					if stateScanNum == 0 {
						scanNum = make([]byte, 0, 1024)
						stateScanNum = 1
					}
					scanNum = append(scanNum, b)
				} else {
					val, err := strconv.Atoi(string(scanNum))
					if err != nil {
						panic("parsing integer value")
					}
					t := token{
						tokenType: TokenInt,
						val:       val,
					}
					e.tokens = append(e.tokens, t)
					stateScanNum = 0
				}
			}
			if b == '*' {
				t := token{
					tokenType: TokenMul,
				}
				e.tokens = append(e.tokens, t)
			}
			if b == '+' {
				t := token{
					tokenType: TokenAdd,
				}
				e.tokens = append(e.tokens, t)
			}
			if b == '(' {
				t := token{
					tokenType: TokenBB,
				}
				e.tokens = append(e.tokens, t)
			}
			if b == ')' {
				t := token{
					tokenType: TokenBE,
				}
				e.tokens = append(e.tokens, t)
			}
		}
		if stateScanNum == 1 {
			val, err := strconv.Atoi(string(scanNum))
			if err != nil {
				panic("parsing integer value")
			}
			t := token{
				tokenType: TokenInt,
				val:       val,
			}
			e.tokens = append(e.tokens, t)
		}
		t := token{
			tokenType: TokenEnd,
		}
		e.tokens = append(e.tokens, t)

		result += e.calc()

	}

	fmt.Println(result)
}
