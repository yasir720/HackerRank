//Problem: https://www.hackerrank.com/challenges/ctci-recursive-staircase/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=recursion-backtracking

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'stepPerms' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER n as parameter.
 */

func stepPerms(n int32) int32 {
    // Write your code here
    m := make(map[int32]int32)
    var inner func(b int32) int32

    m[1] = 1
    m[2] = 2
    m[3] = 4
    

    inner = func(b int32) int32 {
        _, doneBefore := m[b]

        if doneBefore {
            return m[b]
        }

        m[b] = inner(b-1) + inner(b-2) + inner(b-3)
        return m[b]

    }

    m[n] = inner(n)
    return m[n]

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    sTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    s := int32(sTemp)

    for sItr := 0; sItr < int(s); sItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        res := stepPerms(n)

        fmt.Fprintf(writer, "%d\n", res)
    }

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
