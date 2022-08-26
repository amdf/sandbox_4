package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

///////////////////////////////////////////////////////////////////////////////////////

var debug bool
var debug2 bool

func processing(inp io.Reader, w io.Writer) {
	r := bufio.NewReader(inp)

	// fmt.Fscanln(r, &ProcCount, &TaskCount)

	// for i := 0; i < ProcCount; i++ {
	// 	// fmt.Fscan(r, &energy)

	// }

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
