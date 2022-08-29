package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func basketsum(basket map[int64]int64) (total int64) {
	for price := range basket {
		disc := basket[price] / 3
		left := basket[price] % 3

		prdisc := disc * 2 * price
		prleft := left * price

		total += (prdisc + prleft)
	}
	return
}

func processing(r io.Reader, w io.Writer) {

	var numSets int
	fmt.Fscan(r, &numSets)

	for i := 0; i < numSets; i++ {

		var j int64
		fmt.Fscan(r, &j)

		basket := make(map[int64]int64)

		var price int64
		for j > 0 {
			j--
			fmt.Fscan(r, &price)
			basket[price] = basket[price] + 1
		}

		fmt.Fprintln(w, basketsum(basket))

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
