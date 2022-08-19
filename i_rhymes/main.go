package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

var dictSize, wordCount int64
var threads, uniqwords int
var result map[word]string

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

func (w *word) maxrhyme() (result int) {
	result = (10 - bytes.LastIndex(w[:], []byte{0x00})) - 2
	return
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
	rmax := s.maxrhyme()

	for word = range dict {
		val, eq := rhymeValue(s, dict[word])
		if !eq {
			found2 = word
			if val > max {
				f++
				max = val
				found = word
				if max == rmax {
					break
				}
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
	wordCount, _ = strconv.ParseInt(sc.Text(), 10, 64)

	words := make([]word, wordCount)

	uniqw := make(map[string]struct{})

	for i := 0; i < int(wordCount); i++ {
		if !sc.Scan() {
			break
		}
		s := sc.Text()
		words[i].set(s)
		uniqw[s] = struct{}{}
	}

	uniqwords = len(uniqw)

	result = make(map[word]string)
	res := parallel(words, result)

	for i := range res {
		fmt.Fprintln(w, res[i])
	}

}

func parallel(words []word, result map[word]string) []string {
	var mut sync.RWMutex
	var mut2 sync.Mutex

	r := make([]string, len(words))

	proc := func(wg *sync.WaitGroup, part []word, res []string) {

		for i := range part {
			mut.RLock()
			s, ok := result[part[i]]
			mut.RUnlock()
			if !ok {
				s = maxRhymeWord(part[i])
				mut.Lock()
				result[part[i]] = s
				mut.Unlock()
			}
			mut2.Lock()
			res[i] = s
			mut2.Unlock()
		}
		wg.Done()
	}
	var wg sync.WaitGroup

	packsize := len(words) / 16

	for endOffset := len(words); endOffset > 0; endOffset -= packsize {
		beginOffset := endOffset - packsize
		if beginOffset < 0 {
			beginOffset = 0
		}
		wg.Add(1)
		threads++
		go proc(&wg, words[beginOffset:endOffset], r[beginOffset:endOffset])
	}

	wg.Wait()

	return r
}

func main() {
	t := time.Now()
	processing(os.Stdin, os.Stdout)
	if len(os.Args) > 1 {
		fmt.Println()
		fmt.Println(time.Since(t))
		fmt.Printf("threads: %v\n", threads)
		fmt.Printf("dictSize: %v\n", dictSize)
		fmt.Printf("wordCount: %v\n", wordCount)
		fmt.Printf("uniqwords: %v\n", uniqwords)
		fmt.Printf("result size: %v\n", len(result))
	}
}
