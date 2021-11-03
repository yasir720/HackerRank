package main

import "fmt"

func main() {
	// for i := 0; i < 3; i++ {
	//     fmt.Println("Enter text: ")
	//     var input string
	//     fmt.Scanln(&input)
	//     fmt.Println(string(input[1]))
	//     fmt.Println(input)
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
