//Problem: https://www.hackerrank.com/challenges/ctci-ransom-note/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the checkMagazine function below.
func checkMagazine(magazine []string, note []string) {
	words := make(map[string]int, len(magazine))
	for _, w := range magazine {
		words[w] += 1
	}
	for _, w := range note {
		if _, ok := words[w]; !ok || words[w] == 0 {
			fmt.Println("No")
			return
		}
		words[w] -= 1
	}
	fmt.Println("Yes")

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mn := strings.Split(readLine(reader), " ")

	mTemp, err := strconv.ParseInt(mn[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(mn[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	magazineTemp := strings.Split(readLine(reader), " ")

	var magazine []string

	for i := 0; i < int(m); i++ {
		magazineItem := magazineTemp[i]
		magazine = append(magazine, magazineItem)
	}

	noteTemp := strings.Split(readLine(reader), " ")

	var note []string

	for i := 0; i < int(n); i++ {
		noteItem := noteTemp[i]
		note = append(note, noteItem)
	}

	checkMagazine(magazine, note)
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
