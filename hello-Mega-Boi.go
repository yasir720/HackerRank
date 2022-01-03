package main

import "fmt"

func alternatingCharacters(s string) int32 {
    // Write your code here
	numOfDeletions := 0

	for i := 0; i < len(s); i++ {
		if s[i] == s[i + 1] {
			numOfDeletions++
		}
	}

}

func main() {
	// for i := 0; i < 3; i++ {
	// }

	s := make([]int, 5)

	s[0] = 0
	s[1] = 1
	s[2] = 2
	s[3] = 3
    s[4] = 4

	one := s[len(s) - 1]
	fmt.Println(one)
    s[4] = 10
    fmt.Println(one)



}
