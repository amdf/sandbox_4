package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func processing(r io.Reader, w io.Writer) {

	var numSets int
	fmt.Fscan(r, &numSets)

	for i := 0; i < numSets; i++ {

		var a, b int64
		fmt.Fscan(r, &a, &b)
		fmt.Fprintln(w, a+b)
	}

}

func main() {
	// t := time.Now()
	r := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	processing(r, f)
	// fmt.Fprintln(f, "time", time.Since(t))
}
