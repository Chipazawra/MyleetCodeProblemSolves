package main

import (
	"fmt"

	"github.com/Chipazawra/MyleetCodeProblemSolves/problems/containerWithMostWater"
	mergeintervals "github.com/Chipazawra/MyleetCodeProblemSolves/problems/mergeIntervals"
	"github.com/Chipazawra/MyleetCodeProblemSolves/problems/numberOfIslands"
	topKFrequentWords "github.com/Chipazawra/MyleetCodeProblemSolves/problems/topKFrequentWords"
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

	fmt.Println(topKFrequentWords.TopKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))

	fmt.Println(containerWithMostWater.MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	fmt.Println(containerWithMostWater.MaxArea([]int{1, 0, 0, 0, 0, 0, 0, 2, 2}))

}
