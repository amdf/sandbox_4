package main

import (
	"bytes"
	"testing"
)

func TestNorm2(t *testing.T) {
	var buf, outbuf bytes.Buffer
	buf.WriteString(`5
6
2 3 2 4 1 3
6
5 4 1 1 4 1
6
3 4 5 4 2 5
8
1 4 2 2 5 5 4 4
2
4 2
`,
	)

	processing(&buf, &outbuf)

	str := outbuf.String()
	needstr :=
		`1 3
2 6
4 5

1 2
3 4
5 6

1 2
3 6
4 5

1 3
2 7
4 8
5 6

1 2

`

	if needstr != str {
		t.Error("something wrong")
	}

}
