//Problem: https://www.hackerrank.com/challenges/max-array-sum/problem?h_l=interview&isFullScreen=false&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dynamic-programming

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func theBiggerNumber (i, j int32) int32 {
    if i >= j {
        return i
    } else {
        return j
    }
}
// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
    s := make([]int32, 2)
    
    s[0] = theBiggerNumber(0, arr[0])
    s[1] = theBiggerNumber(s[0], arr[1])

    max := theBiggerNumber(s[0], s[1])

    for i := 2; i < len(arr); i++ {
        k := theBiggerNumber((s[i-1]), (arr[i] + s[i-2]))
        s = append(s, k)

        if k > max {
            max = k
        }
    }
    fmt.Println(s)
    return max


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

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := maxSubsetSum(arr)

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
