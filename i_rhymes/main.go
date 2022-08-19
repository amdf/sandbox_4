package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
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

var dict map[string]word

func rhymeValue(s1, s2 word) (result int, equal bool) {
	equal = true
	for i := 9; i >= 0; i-- {
		if 0 == s1[i] || 0 == s2[i] || s1[i] != s2[i] {
			if !(0 == s1[i] && 0 == s2[i]) {
				equal = false
			}

			break
		}

		result++

	}

	return
}

func maxRhymeWord(s word) (result string) {
	var max, f int
	var word, found, found2 string

	for word = range dict {
		val, eq := rhymeValue(s, dict[word])
		if !eq {
			found2 = word
			if val > max {
				f++
				max = val
				found = word
			}
		}
	}

	result = found
	if 0 == len(result) {
		result = found2
	}

	if 0 == len(found) && 0 == len(found2) {
		panic("!")
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

	dict = make(map[string]word, dictSize)
	for i := 0; i < int(dictSize); i++ {
		if !sc.Scan() {
			break
		}
		t := sc.Text()
		var w word
		w.set(t)
		dict[t] = w
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

	for _, s := range words {
		fmt.Fprintf(w, maxRhymeWord(s)+"\n")
	}

}

func main() {
	t := time.Now()
	processing(os.Stdin, os.Stdout)
	if len(os.Args) > 1 {
		fmt.Println(time.Since(t))
	}
}
