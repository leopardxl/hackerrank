package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func beautifulSet(num int) {
	fmt.Printf("%d\n", num)
	arr := make([]int, num)

	//construct the array
	for x := 0; x < num; x++ {
		arr[x] = x
	}

	//rotate the array
	for x := 0; x < num; x++ {
		_begin := arr[num-x:]
		for _, v := range _begin {
			fmt.Printf("%d ", v)
		}
		_end := arr[:num-x]
		for _, v := range _end {
			fmt.Printf("%d ", v)
		}
		fmt.Println()

	}
}

func main() {
	//cardinality, _ := strconv.Atoi(os.Args[1])

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	cardinality, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Atoi Error: %v", err)
		os.Exit(2)
	}
	beautifulSet(cardinality)
}
