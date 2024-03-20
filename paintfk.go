package main

import (
	"fmt"
	"strings"
)

type Paintfuck struct {
	dataGrid         [][]int
	pointer          [2]int
	width            int
	height           int
	code             string
	char             byte
	readPos          int
	iterations       int
	commandsExecuted int
	bracketsIndex    map[int]int
}

func NewPaintfuck(code string, iterations, width, height int) *Paintfuck {
	pf := &Paintfuck{
		dataGrid:         make([][]int, height),
		pointer:          [2]int{0, 0},
		width:            width,
		height:           height,
		code:             code,
		char:             0,
		readPos:          0,
		iterations:       iterations,
		commandsExecuted: 0,
		bracketsIndex:    make(map[int]int),
	}

	for i := 0; i < height; i++ {
		pf.dataGrid[i] = make([]int, width)
	}

	return pf
}

func (pf *Paintfuck) Run() string {
	pf.ReadCode()

	for pf.commandsExecuted < pf.iterations && pf.readPos <= len(pf.code) {
		pf.ExecuteCode()
	}

	return pf.PrintDatagrid()
}

func (pf *Paintfuck) ReadCode() {
	if pf.readPos < len(pf.code) {
		pf.char = pf.code[pf.readPos]
		pf.readPos++
	}
}

func (pf *Paintfuck) CreateBracketPairs() {
	a := []byte(pf.code)
	stack := []int{}

	for i, c := range a {
		if c == '[' {
			stack = append(stack, i)
		}
		if c == ']' {
			openIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			closeIndex := i
			pf.bracketsIndex[openIndex] = closeIndex
			pf.bracketsIndex[closeIndex] = openIndex
		}
	}
}

func (pf *Paintfuck) ExecuteCode() {
	switch pf.char {
	case 'n':
		pf.pointer = pf.GetNorthPointer()
	case 's':
		pf.pointer = pf.GetSouthPointer()
	case 'e':
		pf.pointer = pf.GetEastPointer()
	case 'w':
		pf.pointer = pf.GetWestPointer()
	case '*':
		pf.FlipBit()
	case '[':
		if !pf.CheckBit() {
			pf.JumpToMatchingBracket()
		}
	case ']':
		if pf.CheckBit() {
			pf.JumpToMatchingBracket()
		}
	default:
		pf.ReadCode()
		return
	}

	pf.commandsExecuted++
	pf.ReadCode()
}

func (pf *Paintfuck) JumpToMatchingBracket() {
	pf.readPos = pf.bracketsIndex[pf.readPos-1] + 1
}

func (pf *Paintfuck) FlipBit() {
	if pf.CheckBit() {
		pf.dataGrid[pf.pointer[0]][pf.pointer[1]] = 0
	} else {
		pf.dataGrid[pf.pointer[0]][pf.pointer[1]] = 1
	}
}

func (pf *Paintfuck) CheckBit() bool {
	return pf.dataGrid[pf.pointer[0]][pf.pointer[1]] == 1
}

func (pf *Paintfuck) GetWestPointer() [2]int {
	if pf.pointer[1] == 0 {
		return [2]int{pf.pointer[0], pf.width - 1}
	} else {
		return [2]int{pf.pointer[0], pf.pointer[1] - 1}
	}
}

func (pf *Paintfuck) GetEastPointer() [2]int {
	if pf.pointer[1] == pf.width-1 {
		return [2]int{pf.pointer[0], 0}
	} else {
		return [2]int{pf.pointer[0], pf.pointer[1] + 1}
	}
}

func (pf *Paintfuck) GetNorthPointer() [2]int {
	if pf.pointer[0] == 0 {
		return [2]int{pf.height - 1, pf.pointer[1]}
	} else {
		return [2]int{pf.pointer[0] - 1, pf.pointer[1]}
	}
}

func (pf *Paintfuck) GetSouthPointer() [2]int {
	if pf.pointer[0] == pf.height-1 {
		return [2]int{0, pf.pointer[1]}
	} else {
		return [2]int{pf.pointer[0] + 1, pf.pointer[1]}
	}
}

func (pf *Paintfuck) PrintDatagrid() string {
	var str []string
	for _, row := range pf.dataGrid {
		str = append(str, strings.Replace(fmt.Sprint(row), " ", "", -1))
	}
	return strings.Join(str, "\r\n")
}

func PaintfuckInterpreter(code string, iterations, width, height int) string {
	pf := NewPaintfuck(code, iterations, width, height)
	return pf.Run()
}

func main() {
	fmt.Println(PaintfuckInterpreter("*[s[e]*]", 100, 5, 5))
}
