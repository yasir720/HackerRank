//Problem: https://www.hackerrank.com/challenges/count-triplets-1/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
    // sort.Slice(arr, func(i, j int) bool { return dirRange[i] < dirRange[j] })
    // result := 0
    
    // for i := 0; i < len(arr); i++ {
    //     if i + 2 >= len(arr) {
    //         break
    //     }
    //     if (arr[i] * r == arr[i + 1]) && (arr[i + 1] * r == arr[i + 2]) {
    //         result += 1
    //     } else {
    //         continue
    //     }
    // }
    // return int64(result)
    
    x := int64(0)
    p := make(map[int64]int64)
    s := make(map[int64]int64)

    for i := len(arr) - 1; i >= 0; i-- {
        item := arr[i]
        next := item * r
        if v, ok := p[next]; ok {
            x += v
        }
        if v, ok := s[next]; ok {
            p[item] += v
        }

        s[item]++
    }
    return x


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nr := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(nr[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    r, err := strconv.ParseInt(nr[1], 10, 64)
    checkError(err)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int64

    for i := 0; i < int(n); i++ {
        arrItem, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arr = append(arr, arrItem)
    }

    ans := countTriplets(arr, r)

    fmt.Fprintf(writer, "%d\n", ans)

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

}
