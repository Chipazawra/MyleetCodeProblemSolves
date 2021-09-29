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

	fmt.Println(mergeintervals.Merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	}))
	fmt.Println(mergeintervals.Merge([][]int{
		[]int{1, 4},
		[]int{4, 5},
	}))
	fmt.Println(mergeintervals.Merge([][]int{
		[]int{1, 4},
		[]int{0, 4},
	}))
	fmt.Println(mergeintervals.Merge([][]int{
		[]int{0, 0},
		[]int{1, 4},
	}))

}
