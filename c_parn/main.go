package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

// sc := bufio.NewScanner(in)

// if sc.Scan() {
// 	n, err := strconv.ParseInt(sc.Text(), 10, 64)
// 	if err == nil {
// 	}
// }

func processing(in io.Reader, out io.Writer) {

	var numSets int
	fmt.Fscan(in, &numSets)

	var allDevs [][]int

	for ; numSets > 0; numSets-- {
		var numDevs, devLevel int
		var devs []int
		fmt.Fscan(in, &numDevs)

		for ; numDevs > 0; numDevs-- {

			fmt.Fscan(in, &devLevel)
			devs = append(devs, devLevel)
		}
		allDevs = append(allDevs, devs)
	}

	for _, devs := range allDevs {
		for {
			if !printPairs(out, devs) {
				break
			}
		}
		fmt.Fprintf(out, "\n")
	}
}

func printPairs(out io.Writer, devs []int) (result bool) {
	dev1 := -1
	dev2 := -1
	lastmin := math.MaxFloat64
	for i, level := range devs {
		if dev1 < 0 {
			if level >= 0 {
				dev1 = i
				continue
			}
		} else {
			if level >= 0 {
				diff := math.Abs(float64(devs[dev1] - level))
				if diff < lastmin {
					dev2 = i
					result = true
					lastmin = diff
					continue
				}
			}
		}
	}
	if result {
		devs[dev1] = -1
		devs[dev2] = -1
		fmt.Fprintf(out, "%d %d\n", dev1+1, dev2+1)
	}
	return
}

func main() {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	processing(os.Stdin, f)
}
