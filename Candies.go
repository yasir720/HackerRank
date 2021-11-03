//Problem: https://www.hackerrank.com/challenges/candies/problem?h_l=interview&isFullScreen=false&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dynamic-programming

package main

import (
	"fmt"
)

func main() {

	slice := make([]int32, 10)
	slice[0] = 2
	slice[1] = 4
	slice[2] = 2
	slice[3] = 6
	slice[4] = 1
	slice[5] = 7
	slice[6] = 8
	slice[7] = 9
	slice[8] = 2
	slice[9] = 1

	candies(10, slice)
}


func candies(n int32, arr []int32) int64 {
	dp := make([]int32, n){}














    // Write your code 
	// candyDistribution := make([]int32, n)

	// candyDistribution[0] = 1
	// var prev, current, next int32 = 0, 0, 0
	// var totalCandies int32 = 1

	// for i := 1; int32(i) < n-1; i=i+2 {
	// 	prev = arr[i-1]
	// 	current = arr[i]
	// 	next = arr[i+1]

	// 	if current > prev {
	// 		candyDistribution[i] = candyDistribution[i-1] + 1
	// 	} else if current < prev {
	// 		candyDistribution[i] = candyDistribution[i-1] - 1
	// 	}

	// 	if prev == next











		// if (current > prev && current < next) {
		// 	candyDistribution[i] = candyDistribution[i-1] + 1
		// 	candyDistribution[i+1] = candyDistribution[i-1] + 2
		// } else if (prev == next && current >prev) {
		// 	candyDistribution[i] = candyDistribution[i-1] + 1
		// 	candyDistribution[i+1] = candyDistribution[i-1]
		// }		
	//}
	// fmt.Println(candyDistribution)
	// return totalCandies
}
