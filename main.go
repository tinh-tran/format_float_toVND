package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	m := FormatAmount(2500000)

	fmt.Printf("%v\n", m )
}

func FormatAmount(m float64) string {
	m = math.Round(m)
	if m == 0 {
		return "-"
	}
	n := int64(m)
	in := strconv.FormatInt(n, 10)
	out := make([]byte, len(in)+(len(in)-2+int(in[0]/'0'))/3)
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = ','
		}
	}
}

