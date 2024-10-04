package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func OddInt(input []int) int {
	countMap := map[int]int{}

	for _, v := range input {
		_, exist := countMap[v]
		if !exist {
			countMap[v] = 0
		}
		countMap[v]++
	}

	oddCountNumber := 0

	for k, v := range countMap {

		if v%2 != 0 {
			oddCountNumber = k
			break
		}
	}

	return oddCountNumber
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str := readLine(reader)
	// input always array -> have []
	str = str[1 : len(str)-1]
	numStrArr := strings.Split(str, ",")
	arr := []int{}
	for _, v := range numStrArr {
		// assume no invalid input
		newNum, _ := strconv.Atoi(v)
		arr = append(arr, newNum)
	}

	fmt.Println(OddInt(arr))
}

func readLine(reader *bufio.Reader) string {
	s, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(s), "\r\n")
}
