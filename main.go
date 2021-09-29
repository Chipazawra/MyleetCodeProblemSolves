package main

import (
	"fmt"

	mergeintervals "github.com/Chipazawra/MyleetCodeProblemSolves/problems/mergeIntervals"
	"github.com/Chipazawra/MyleetCodeProblemSolves/problems/numberOfIslands"
	"github.com/Chipazawra/MyleetCodeProblemSolves/problems/validParentheses"
)

/**
 * Implement function getResult
 */

func main() {

	fmt.Println(validParentheses.IsValid("()[]{}"))

	fmt.Println(numberOfIslands.NumberOfIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

	fmt.Println(numberOfIslands.NumberOfIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))

	fmt.Println(mergeintervals.Merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))
	fmt.Println(mergeintervals.Merge([][]int{
		{1, 4},
		{4, 5},
	}))
	fmt.Println(mergeintervals.Merge([][]int{
		{1, 4},
		{0, 4},
	}))
	fmt.Println(mergeintervals.Merge([][]int{
		{0, 0},
		{1, 4},
	}))

}
