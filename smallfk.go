package main

import (
	"strings"
)

type Smallfuck struct {
	tape           string
	pointer        int
	code           string
	char           rune
	readPos        int
	bracketMapping map[int]int
}

func (s *Smallfuck) run() string {
	s.createBracketMapping(s.code)
	hasMoreCode := s.readCode()

	for s.readPos <= len(s.code) && s.pointer >= 0 && s.pointer < len(s.tape) && hasMoreCode {
		s.executeCode()
	}
	return s.tape
}

func (s *Smallfuck) readCode() bool {
	if s.readPos != len(s.code) {
		s.char = rune(s.code[s.readPos])
		s.readPos++
		return true
	} else {
		s.readPos++
		return false

	}
}

func (s *Smallfuck) createBracketMapping(code string) map[int]int {
	chars := strings.Split(code, "")
	var stack []int
	var pairs = make(map[int]int)

	for i := 0; i < len(chars); i++ {
		if chars[i] == "[" {
			stack = append(stack, i)
		}

		if chars[i] == "]" {
			openIndex := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			closeIndex := i
			pairs[openIndex] = closeIndex
			pairs[closeIndex] = openIndex
		}
	}
	s.bracketMapping = pairs
	return pairs
}

func (s *Smallfuck) executeCode() bool {
	switch s.char {
	case '>':
		s.pointer++
	case '<':
		s.pointer--
	case '*':
		s.flipBit()
	case '[':
		if !s.checkBit() {
			s.jumpToMatchingBracket()
		}
	case ']':
		if s.checkBit() {
			s.jumpToMatchingBracket()
		}
	}

	hasMoreCode := s.readCode()
	return hasMoreCode
}

func (s *Smallfuck) jumpToMatchingBracket() {
	s.char = rune(s.code[s.bracketMapping[s.readPos-1]])
	s.readPos = s.bracketMapping[s.readPos-1] + 1
}

func (s *Smallfuck) flipBit() {
	out := []rune(s.tape)

	if s.tape[s.pointer] == '0' {
		out[s.pointer] = '1'
		s.tape = string(out)
	} else {
		out[s.pointer] = '0'
		s.tape = string(out)
	}
}

func (s *Smallfuck) checkBit() bool {
	return string(s.tape[s.pointer]) == "1"
}

func newInterpreter(tape string, code string) Smallfuck {
	interpreter := Smallfuck{tape: tape, pointer: 0, code: code, readPos: 0, bracketMapping: make(map[int]int)}
	return interpreter
}

func Interpreter(code string, tape string) string {
	smallfuck := newInterpreter(tape, code)
	return smallfuck.run()
}
