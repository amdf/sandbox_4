package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

// sc := bufio.NewScanner(in)

// if sc.Scan() {
// 	n, err := strconv.ParseInt(sc.Text(), 10, 64)
// 	if err == nil {
// 	}
// }

// fmt.Fscan(in, &numDevs)

// for ; numDevs > 0; numDevs-- {

type Ans struct {
	Number int
	Result string
}

func processing(in io.Reader, out io.Writer) {

	var numSets int
	fmt.Fscan(in, &numSets)
	results := make([]string, numSets)

	chdone := make(chan int)

	ch := make(chan Ans)
	go func() {
		for x := range ch {
			results[x.Number] = x.Result
		}
		chdone <- 1
	}()

	var wg sync.WaitGroup

	for n := 0; n < numSets; n++ {

		var numTasks int
		fmt.Fscan(in, &numTasks)
		var tasks []int
		for ; numTasks > 0; numTasks-- {
			var d int
			fmt.Fscan(in, &d)
			tasks = append(tasks, d)
		}

		wg.Add(1)
		go func(w *sync.WaitGroup, n int, r []int) {
			var ans Ans
			ans.Number = n
			ans.Result = isReportOk(r)
			ch <- ans
			w.Done()
		}(&wg, n, tasks)
	}

	wg.Wait()
	close(ch)
	<-chdone
	for _, answer := range results {
		fmt.Fprintln(out, answer)
	}

}

func isReportOk(report []int) string {
	taskVariants := make(map[int]struct{})
	prev := -1
	for _, rep := range report {
		_, exists := taskVariants[rep]
		if !exists {
			prev = rep
			taskVariants[rep] = struct{}{}
		} else {
			if rep == prev {
				continue
			}
			return "NO"
		}
	}
	return "YES"
}

func main() {
	//t := time.Now()
	r := bufio.NewReader(os.Stdin)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	processing(r, f)
	//fmt.Fprintln(f, "time", time.Since(t))
}
