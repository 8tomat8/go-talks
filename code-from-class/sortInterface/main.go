package main

import (
	"fmt"
	"sort"
)

func main() {
	x := Data{[]int{5, 2, 67, 8, 3, 2, 5, 7, 8}}
	sort.Sort(&x)
	fmt.Println(x)
}

type Data struct {
	x []int
}

func (d *Data) Len() int {
	return len(d.x)
}

func (d *Data) Less(i, j int) bool {
	return d.x[i] > d.x[j]
}

func (d *Data) Swap(i, j int) {
	d.x[i], d.x[j] = d.x[j], d.x[i]
}
