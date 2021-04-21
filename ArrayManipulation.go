//Problem: https://www.hackerrank.com/challenges/crush/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=arrays

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
	"math"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	var max, accumulated int64 = 0, 0
	res := make([]int64, n+2)

	// Let's do O(M)
	for _, q := range queries {

		a, b, k := q[0], q[1], int64(q[2])

		res[a] += k
		res[b+1] -= k
	}

	// So, for this case:
	// 2 6 8
	// 3 5 7
	// 1 8 1
	// 5 9 15
	// res will look like:
	// [0 1 8 7 0 15 -7 -8 0 -1 -15 0]

	// Now we do O(N) and find max accumulated value
	for _, v := range res {
		accumulated += v
		max = int64(math.Max(float64(max), float64(accumulated)))
	}

	return max


}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    var queries [][]int32
    for i := 0; i < int(m); i++ {
        queriesRowTemp := strings.Split(readLine(reader), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != int(3) {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    result := arrayManipulation(n, queries)

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
