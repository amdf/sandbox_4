package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type word [10]byte

func (w *word) set(s string) {
	if len(s) <= 10 {
		i := 10 - len(s)
		copy(w[:i], "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"[:i])
		copy(w[i:], s)
	} else {
		panic("length")
	}
}

var dict map[string]struct{}

func rhymeValue(s1, s2 word) (result int, equal bool) {
	equal = true
	for i := 9; i >= 0; i-- {
		if 0 == s1[i] || 0 == s2[i] || s1[i] != s2[i] {
			equal = false
			break
		}

		result++

	}

	return
}

func processing(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)

	var dictSize int64

	if !sc.Scan() {
		return
	}

	dictSize, _ = strconv.ParseInt(sc.Text(), 10, 64)

	dict = make(map[string]struct{}, dictSize)
	for i := 0; i < int(dictSize); i++ {
		if !sc.Scan() {
			break
		}
		dict[sc.Text()] = struct{}{}
	}

	if !sc.Scan() {
		return
	}
	qSize, _ := strconv.ParseInt(sc.Text(), 10, 64)

	words := make([]word, qSize)

	for i := 0; i < int(qSize); i++ {
		if !sc.Scan() {
			break
		}
		words[i].set(sc.Text())
	}

	for _, val := range words {
		fmt.Fprintf(w, "%s\n", val)
	}

}

func main() {
	processing(os.Stdin, os.Stdout)
}
