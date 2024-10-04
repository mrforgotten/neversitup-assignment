package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Permutation(input string) []string {
	inputLen := len(input)
	ansMap := map[string]bool{}
	ans := []string{}

	//declare function
	var permute func(toPermute []rune, positionToSwap int)
	// implement
	permute = func(toPermute []rune, positionToSwap int) {
		// check for last index, so no swap position left
		if inputLen == positionToSwap {
			ansStr := string(toPermute)
			if !ansMap[ansStr] {
				ansMap[ansStr] = true
				ans = append(ans, ansStr)
				return
			}

		}

		//actual permutation
		for i := positionToSwap; i < inputLen; i++ {
			// swap first for making next string to permute
			toPermute[i], toPermute[positionToSwap] = toPermute[positionToSwap], toPermute[i]
			permute(toPermute, positionToSwap+1)
			// swap back, so no lost in previous permute
			toPermute[i], toPermute[positionToSwap] = toPermute[positionToSwap], toPermute[i]
		}
	}

	permute([]rune(input), 0)

	return ans
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str := readLine(reader)

	fmt.Println("function return:", Permutation(str))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()

	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\n\r")
}
