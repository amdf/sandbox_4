package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

var dict map[string]struct{}

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

	words := make([]string, qSize)

	for i := 0; i < int(qSize); i++ {
		if !sc.Scan() {
			break
		}
		words[i] = sc.Text()
	}

	for _, val := range words {
		fmt.Fprintf(w, "%s\n", val)
	}

}

func main() {
	processing(os.Stdin, os.Stdout)
}
