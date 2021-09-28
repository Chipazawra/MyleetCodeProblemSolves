package main

import (
	"fmt"

	"github.com/Chipazawra/MyleetCodeProblemSolves/problems/numberOfIslands"
	"github.com/Chipazawra/MyleetCodeProblemSolves/problems/validParentheses"
)

/**
 * Implement function getResult
 */

func main() {

	fmt.Println(validParentheses.IsValid("()[]{}"))

	fmt.Println(numberOfIslands.NumberOfIslands([][]byte{
		[]byte{'1', '1', '1', '1', '0'},
		[]byte{'1', '1', '0', '1', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '0', '0', '0'},
	}))

	fmt.Println(numberOfIslands.NumberOfIslands([][]byte{
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'1', '1', '0', '0', '0'},
		[]byte{'0', '0', '1', '0', '0'},
		[]byte{'0', '0', '0', '1', '1'},
	}))

}
