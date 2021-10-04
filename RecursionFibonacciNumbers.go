//Problem: https://www.hackerrank.com/challenges/ctci-fibonacci-numbers/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=recursion-backtracking

package main
import "fmt"

func fibonacci(n int) int {

    // Write your code here.
    if n <= 1 {
        return n
    } else {return fibonacci(n-1) + fibonacci(n-2)}
}

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    fmt.Println(fibonacci(n))
}