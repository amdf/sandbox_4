package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Tyle struct {
	c       byte
	tl, tr  *Tyle
	bl, br  *Tyle
	l, r    *Tyle
	visited bool
}

type Row []Tyle

type HexMap struct {
	colors  map[byte]int
	found   map[byte]bool
	rows    []Row
	alt     bool
	visited int
	total   int
}

func NewHexMap(alt bool) *HexMap {
	m := HexMap{}
	m.colors = make(map[byte]int)
	m.found = make(map[byte]bool)
	m.alt = alt
	return &m
}

func (m *HexMap) getLine(line string) (result []byte) {

	for _, c := range []byte(line) {
		if c >= 'A' && c <= 'Z' {
			result = append(result, c)
		}
	}
	if debug {
		fmt.Printf("result: %v\n", result)
	}

	return
}

func (m *HexMap) newRow(line string) (row Row) {
	chars := m.getLine(line)

	for _, c := range chars {
		color := c
		m.colors[color] = m.colors[color] + 1 //count colors
		m.total++                             //count total
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
	m.rows = append(m.rows, row)
}

func (m *HexMap) addOdd(line string) {
	top := m.rows[len(m.rows)-1]
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

	m.rows = append(m.rows, row)
}

func (m *HexMap) addEven(line string) {
	top := m.rows[len(m.rows)-1]
	row := m.newRow(line)

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

	m.rows = append(m.rows, row)
}

func (m *HexMap) Add(line string) {
	odd := (0 == len(m.rows)%2)
	if odd {
		if 0 == len(m.rows) {
			if debug {
				fmt.Println("addFirst")
			}
			m.addFirst(line)
		} else {
			if debug {
				fmt.Println("addOdd")
			}
			m.addOdd(line)
		}
	} else {
		if debug {
			fmt.Println("addEven")
		}
		m.addEven(line)
	}
}

func (m *HexMap) FirstNonVisited() (result *Tyle) {
	for i := range m.rows {
		for j := range m.rows[i] {
			if !m.rows[i][j].visited {
				result = &m.rows[i][j]
				return
			}
		}
	}
	return
}

func (m *HexMap) FillFrom(color byte, t *Tyle) (size int) {

	if nil == t {
		// fmt.Println("nil")
		return
	}

	if t.visited {
		// fmt.Println("visited")
		return
	}

	if t.c == color {
		t.visited = true
		size = 1

		// fmt.Println("right:")
		size += m.FillFrom(color, t.r)
		// fmt.Println("bottom right:")
		size += m.FillFrom(color, t.br)
		// fmt.Println("bottom left:")
		size += m.FillFrom(color, t.bl)
		// fmt.Println("left:")
		size += m.FillFrom(color, t.l)
		// fmt.Println("top left:")
		size += m.FillFrom(color, t.tl)
		// fmt.Println("top right:")
		size += m.FillFrom(color, t.tr)
	}
	return
}

func (m *HexMap) IsOK() string {
	for m.visited < m.total {
		// 		ищем непосещённую
		tyle := m.FirstNonVisited()
		if nil == tyle {
			break
		}
		// получаем цвет

		// если цвет в списке найденных ответ НЕТ
		if m.found[tyle.c] {
			if debug {
				fmt.Println("color found!!!", tyle.c)
			}
			return "NO"
		}
		// иначе
		// заполняем область - fill() - получаем количество в области
		if debug {
			fmt.Println("====TYLE=====")
		}
		size := m.FillFrom(tyle.c, tyle)
		m.visited += size
		// если количество в области менее количества в цвете - ответ НЕТ
		if size < m.colors[tyle.c] {
			if debug {
				fmt.Println("color", string(tyle.c), "found", size, "must be", m.colors[tyle.c])
			}
			return "NO"
		}
		// иначе
		// отмечаем цвет как найденный
		m.found[tyle.c] = true
		if debug {
			fmt.Println("color", string(tyle.c), ":", size)
		}
	}
	return "YES"
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
		if debug {
			fmt.Printf("lines: %v\n", lines)
			fmt.Printf("strlen: %v\n", strlen)
		}

		m := NewHexMap((0 == strlen%2))
		maps = append(maps, m)

		for j := 0; j < lines; j++ {
			line, _ := r.ReadString('\n')
			m.Add(line)
		}
	}

	for i := 0; i < numMaps; i++ {
		if debug {
			fmt.Println("=========================")
		}
		fmt.Fprintln(w, maps[i].IsOK())
	}

}

func main() {
	t := time.Now()
	if len(os.Args) > 1 {
		debug = true
		//debug2 = true
	}
	if debug2 {
		f, _ := os.Open("tests/01")

		processing(f, os.Stdout)
	} else {
		processing(os.Stdin, os.Stdout)
	}

	if debug {
		fmt.Println()
		fmt.Println(time.Since(t))

	}
}
