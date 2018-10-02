package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isPalindrome(s string) bool {
	length := len(s)
	for k, _ := range s {
		if s[k] != s[length-k-1] {
			return false
		}
	}
	return true

}

// Complete the palindromeIndex function below.
func palindromeIndex(s string) int32 {
	if isPalindrome(s) {
		return -1
	}

	for k, _ := range s {
		newString := s[:k] + s[k+1:]
		if isPalindrome(newString) {
			return int32(k)
		}
	}

	return -1
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := palindromeIndex(s)

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
