//Package mymath provides math solutions relevent
package mymath

import "sort"

//CenteredAvg computes the avrage of a list of number
//after removing the largest and smallest values
func CenteredAvg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}
	f := float64(n) / float64(len(xi))
	return f
}
