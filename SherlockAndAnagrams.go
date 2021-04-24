//Problem: https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=dictionaries-hashmaps

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'sherlockAndAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	map_s := make(map[string]int32)
	for length_sub := 1; length_sub <= len(s); length_sub++ {
		for i := 0; i <= len(s)-length_sub; i++ {
			j := i + length_sub - 1
			var sub_string, parent_key string = "", ""
			var map_tmp = make(map[string]int)
			keys := make([]string, 0)
			for k := i; k <= j; k++ {
				sub_string += string(s[k])
				map_tmp[string(s[k])]++
			}
			for k, _ := range map_tmp {
				keys = append(keys, k)
			}
			sort.Strings(keys) // we sort to get same keys-values for 1 Anagram
			for _, k := range keys {
				v := strconv.Itoa(map_tmp[k])
				parent_key += k + "-" + v
			}
			map_s[parent_key]++ //count number frequency of anagram
		}
	}
	var count int32 = 0
	for _, v := range map_s {
		count += v * (v - 1) / 2
	}
	return count

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
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
