package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func sockMerchant(n int, items []int) map[int]int {
	var i int = 0
	m := make(map[int]int)
	for _, x := range items {
		_, isThere := m[x]; if isThere {
			i++
			m[x]--
		} else {
			m[x] ++
		}
		
	}
	fmt.Println("i:", i)
	return m
}

func main () {
	//slice := generateSlice(5)

	slice2 := make([]int, 10)
	slice2[0] = 1
	slice2[1] = 1
	slice2[2] = 1
	slice2[3] = 2
	slice2[4] = 3
	slice2[5] = 3
	slice2[6] = 4
	slice2[7] = 4
	slice2[8] = 5
	slice2[9] = 5

	ans := sockMerchant(10, slice2)
	fmt.Println(slice2)
	fmt.Println(ans)

}