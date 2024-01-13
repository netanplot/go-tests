package main

import (
	"fmt"
	"math"
)



func options(n, max int) (r []int) {
	a := int(math.Pow(2,float64(n-1)))
	for i:=0; i<max; i++ {
		if i & a > 0 {
			r = append(r, i)
		}
	}
	return
}

func main() {

	cards := 20

	m := math.Pow(2,float64(cards))

	arr:= make( [][]int, cards)
	for i:=1; i<=cards; i++ {
		arr[i-1] = options(i, int(m))	
		fmt.Println("Card:", i)
		fmt.Println(arr[i-1])
	}

 
}

    