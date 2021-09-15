//Problem: https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
    // Write your code here
    m := make(map[rune]int, 0)
    var i, j int = -1, -1
    n := make(map[int]int, 0)    
    
    for _, letter := range s {
        m[letter] ++
                
    }    

    for _, frequency := range m {
        n[frequency] ++
        fmt.Println(frequency)
        if i != frequency {
            j = i
            i = frequency            
        }
    }

    if len(n) == 1 {
        return "YES"
    } else if len(n) > 2 {
        return "NO"
    } else if len(n) == 2 {
        switch {
            case (i == 1 && n[i] == 1):
                return "YES"
            case (j == 1 && n[j] == 1):
                return "YES"
            case (i - j == 1 && n[i] == 1):
                return    "YES"
            case (i - j == -1 && n[j] == 1):
                return    "YES"
        }        
    }
    return "NO"

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := isValid(s)

    fmt.Fprintf(writer, "%s\n", result)

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
