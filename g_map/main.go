package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Tyle struct {
	c      byte
	tl, tr *Tyle
	bl, br *Tyle
	l, r   *Tyle
}

type Row []Tyle

type HexMap struct {
	n   map[byte]int
	r   []Row
	alt bool
}

func NewHexMap(alt bool) *HexMap {
	m := HexMap{}
	m.n = make(map[byte]int)
	m.alt = alt
	return &m
}

func (m *HexMap) getLine(line string) (result []byte) {

	for _, c := range []byte(line) {
		if c >= 'A' && c <= 'Z' {
			result = append(result, c)
		}
	}
	fmt.Printf("result: %v\n", result)

	return
}

func (m *HexMap) newRow(line string) (row Row) {
	chars := m.getLine(line)

	for _, c := range chars {
		color := c
		m.n[color] = m.n[color] + 1 //count colors
		t := Tyle{c: color}
		row = append(row, t)
	}

	for i := 0; i < len(row)-1; i++ {
		row[i].r = &row[i+1]
		row[i+1].l = &row[i]
	}
	return
}

func (m *HexMap) addFirst(line string) {
	row := m.newRow(line)
	m.r = append(m.r, row)
}

func (m *HexMap) addOdd(line string) {
	top := m.r[len(m.r)-1]
	row := m.newRow(line)

	if !m.alt {
		for i := 0; i < len(top); i++ { //TOP bottom lefts
			top[i].bl = &row[i]
		}
		for i := 0; i < len(top); i++ { //ROW top rights
			row[i].tr = &top[i]
		}
		for i := 0; i < len(top); i++ { //TOP bottom rights
			top[i].br = &row[i+1]
		}
		for i := 0; i < len(top); i++ { //ROW  top lefts
			row[i+1].tl = &top[i]
		}

	} else {
		for i := 0; i < len(top); i++ { //TOP bottom lefts
			top[i].bl = &row[i]
		}
		for i := 0; i < len(row); i++ { //ROW top rights
			row[i].tr = &top[i]
		}
		for i := 0; i < len(top)-1; i++ { //TOP bottom rights
			top[i].br = &row[i+1]
		}
		for i := 1; i < len(row); i++ { //ROW top lefts
			row[i].tl = &top[i-1]
		}
	}

	m.r = append(m.r, row)
}

func (m *HexMap) addEven(line string) {
	top := m.r[len(m.r)-1]
	row := m.newRow(line)
	fmt.Printf("len(top): %v\n", len(top))
	fmt.Printf("len(row): %v\n", len(row))

	if !m.alt {
		for i := 0; i < len(row); i++ {
			row[i].tl = &top[i]
			top[i].br = &row[i]
			row[i].tr = &top[i+1]
			top[i+1].bl = &row[i]
		}
	} else {
		for i := 1; i < len(top); i++ { //TOP bottom lefts
			top[i].bl = &row[i-1]
		}
		for i := 0; i < len(row)-1; i++ { //ROW top rights
			row[i].tr = &top[i+1]
		}
		for i := 0; i < len(top); i++ { //TOP bottom rights
			top[i].br = &row[i]
		}
		for i := 1; i < len(row); i++ { //ROW top lefts
			row[i].tl = &top[i]
		}
	}

	m.r = append(m.r, row)
}

func (m *HexMap) Add(line string) {
	odd := (0 == len(m.r)%2)
	if odd {
		if 0 == len(m.r) {
			fmt.Println("addFirst")
			m.addFirst(line)
		} else {
			fmt.Println("addOdd")
			m.addOdd(line)
		}
	} else {
		fmt.Println("addEven")
		m.addEven(line)
	}
}

///////////////////////////////////////////////////////////////////////////////////////

var debug bool
var debug2 bool
var numMaps int
var lines, strlen int
var maps []*HexMap

func processing(inp io.Reader, w io.Writer) {
	r := bufio.NewReader(inp)

	fmt.Fscanln(r, &numMaps)

	for i := 0; i < numMaps; i++ {
		fmt.Fscanf(r, "%d %d\n", &lines, &strlen)
		fmt.Printf("lines: %v\n", lines)
		fmt.Printf("strlen: %v\n", strlen)

		m := NewHexMap((0 == strlen%2))
		maps = append(maps, m)

		for j := 0; j < lines; j++ {
			line, _ := r.ReadString('\n')
			//fmt.Fscanln(r, &line)
			m.Add(line)
		}
	}

	for _, m := range maps {
		fmt.Printf("colors: %v\n", m.n)
	}

}

func main() {
	t := time.Now()
	if len(os.Args) > 1 {
		debug = true
		//debug2 = true
	}
	processing(os.Stdin, os.Stdout)

	if debug {
		fmt.Println()
		fmt.Println(time.Since(t))

	}
}
