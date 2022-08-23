package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var debug bool

func processing(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)

	if !sc.Scan() {
		return
	}

	// dictSize, _ = strconv.ParseInt(sc.Text(), 10, 64)

	// for i := 0; i < int(dictSize); i++ {
	// 	if !sc.Scan() {
	// 		break
	// 	}
	// 	s := sc.Text()

	// }

	// if !sc.Scan() {
	// 	return
	// }
	// wordCount, _ = strconv.ParseInt(sc.Text(), 10, 64)

}

func main() {
	t := time.Now()
	if len(os.Args) > 1 {
		debug = true
	}
	processing(os.Stdin, os.Stdout)

	if len(os.Args) > 1 {
		fmt.Println()
		fmt.Println(time.Since(t))
	}
}
