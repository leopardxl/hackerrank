package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int32) int32 {
	sightings := make(map[int32]int)

	for _, v := range arr {
		sightings[v]++
	}

	var i int32 = 6
	var max int

	for k, v := range sightings {
		if v > max {
			max = v
			i = k
		} else if v == max && i > k {
			max = v
			i = k
		}

		fmt.Printf("%d\t%d\n", k, v)
	}

	return i
}

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

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
