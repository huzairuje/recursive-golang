package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumDigit(a int, n []int) interface{} {
	if len(n) > 0 {
		s := len(n) - 1
		a = a + n[s]
		n = n[:s]
		return sumDigit(a, n)
	} else {
		x := strconv.Itoa(a)
		y := strings.Split(x, "")
		if len(y) > 1 {
			z := make([]int, len(y))
			for i := range z {
				z[i], _ = strconv.Atoi(y[i])
			}
			return sumDigit(0, z)
		} else {
			return a
		}
	}
}

func main() {
	arrayDigit := []int{18278172812, 91823908912, 8172837123, 918298391283, 9898298323}
	a := sumDigit(0, arrayDigit)
	fmt.Printf("%d\n", a)
}