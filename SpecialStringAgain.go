//Problem: https://www.hackerrank.com/challenges/special-palindrome-again/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=strings

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
    var count int64 = 0
    var spl_count int64 = 0

    for i := 0; i < len(s); i++ {
        count++
        spl_count = 0
        pivot := false
        for j := i + 1; j < len(s); j++ {
            if s[i] == s[j] && pivot == false {
                count++
                spl_count++
            } else if s[i] != s[j] && pivot == true {
                break
            } else if s[i] != s[j] && pivot == false {
                spl_count++
                pivot = true
            } else if s[i] == s[j] && pivot == true {
                spl_count--
                if spl_count == 0 {
                    count++
                }
            }
        }
    }
    return count
}



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    s := readLine(reader)

    result := substrCount(n, s)

    fmt.Fprintf(writer, "%d\n", result)

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
