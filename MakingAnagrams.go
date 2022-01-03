//Problem: https://www.hackerrank.com/challenges/ctci-making-anagrams/problem?h_l=interview&isFullScreen=false&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
    // Write your code here
    var deletions int32
    hashMap := map[rune]int32{}
    
    for _, char := range a {
        hashMap[char]++
    }
    
    for _, char := range b {
        if hashMap[char] == 0 {
            deletions++
        } else {
            hashMap[char]--
        }
    }
    
    for _, frequency := range hashMap {
            deletions += frequency
    }
    
    return deletions
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    a := readLine(reader)

    b := readLine(reader)

    res := makeAnagram(a, b)

    fmt.Fprintf(writer, "%d\n", res)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
