package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        fmt.Println("Enter text: ")
        var input string
        fmt.Scanln(&input)
        fmt.Println(string(input[1]))
        fmt.Println(input)
    }
    
}
