package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// sc := bufio.NewScanner(in)

// if sc.Scan() {
// 	n, err := strconv.ParseInt(sc.Text(), 10, 64)
// 	if err == nil {
// 	}
// }

// fmt.Fscan(in, &numDevs)

// for ; numDevs > 0; numDevs-- {

type MyTime struct {
	T1 time.Time
	T2 time.Time
}

func (tt MyTime) IsCross(other MyTime) bool {
	if tt.T1.Equal(other.T1) || tt.T1.Equal(other.T2) || tt.T2.Equal(other.T1) || tt.T2.Equal(other.T2) {
		return true
	}
	return other.T1.After(tt.T1) && other.T1.Before(tt.T2) || (other.T2.After(tt.T1) && other.T2.Before(tt.T2)) ||
		tt.T1.After(other.T1) && tt.T1.Before(other.T2) || (tt.T2.After(other.T1) && tt.T2.Before(other.T2))
}

func (tt MyTime) Valid() bool {
	// левая граница отрезка находится не позже его правой границы (но границы могут быть равны);
	if tt.T1.Equal(tt.T2) {
		return true
	}
	return tt.T2.After(tt.T1)
}

func processing(in io.Reader, w io.Writer) {
	r := bufio.NewReader(in)
	var numSets int
	fmt.Fscan(r, &numSets)

	for ; numSets > 0; numSets-- {
		var numTimes int
		fmt.Fscan(r, &numTimes)

		var knownTimes []MyTime

		valid := true

		for ; numTimes > 0; numTimes-- {

			var val string
			fmt.Fscan(r, &val)

			if !valid {
				//fmt.Println("!!!skip")
				continue
			}
			vals := strings.Split(val, "-")

			var t MyTime
			t1, err1 := time.Parse("15:04:05", vals[0])
			t2, err2 := time.Parse("15:04:05", vals[1])
			t.T1 = t1
			t.T2 = t2
			if nil != err1 || nil != err2 {
				// fmt.Println("!!!format")
				valid = false
			} else {
				knownTimes = append(knownTimes, t)

				if !t.Valid() {
					// fmt.Println("!!!wrong", t)
					valid = false
				}
			}
		}

		//tn := time.Now()
		if valid {
			for _, t := range knownTimes {
				cross := false
				var tcr MyTime
				for _, tcr = range knownTimes {
					if t.T1 != tcr.T1 && t.T2 != tcr.T2 {
						if t.IsCross(tcr) {
							cross = true
							break
						}
					}
				}
				if cross {
					//fmt.Println("!!!cross\n", t, "\n", tcr)
					valid = false
					break
				}
			}
		}

		//fmt.Fprintln(w, "calculated", time.Since(tn))

		if valid {
			fmt.Fprintln(w, "YES")
		} else {
			fmt.Fprintln(w, "NO")
		}

	}

}

func main() {

	processing(os.Stdin, os.Stdout)

}
