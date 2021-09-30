package knapsackproblem

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//KnapsackProblemSolve([]int{4, 3, 2, 1}, []int{28, 20, 13, 6}, 10)
	//KnapsackProblemSolve([]int{6, 5, 4, 3, 2, 1}, []int{42, 35, 28, 20, 13, 6}, 10)

	println("Test case #1")
	println(x5RiteilKnapsackProblemSolve(1000, []int{100, 200, 300, 400}, []int{50, 100, 500, 100}, []int{4, 5, 7, 3}, []int{0, 100, 5000, 1000}))

	println("Test case #2")
	println(x5RiteilKnapsackProblemSolve(1000, []int{100, 200, 300, 400}, []int{50, 100, 500, 100}, []int{4, 5, 7, 3}, []int{1000, 100, 5000, 1000}))

	println("Test case #3")
	println(x5RiteilKnapsackProblemSolve(80000, []int{100, 200, 300, 400}, []int{50, 100, 500, 100}, []int{17, 36, 70, 45}, []int{1000, 500, 5000, 1000}))

	println("Test case #4")
	println(x5RiteilKnapsackProblemSolve(28000, []int{100, 200, 3000, 400}, []int{50, 100, 500, 100}, []int{17, 36, 70, 45}, []int{1000, 500, 5000, 1000}))

	println("Test case #5")
	println(x5RiteilKnapsackProblemSolve(28000, []int{500, 200, 3000, 400}, []int{50, 100, 500, 100}, []int{17, 36, 70, 45}, []int{1000, 500, 5000, 0}))

	println("Test case #6")
	println(x5RiteilKnapsackProblemSolve(8000, []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{50, 100, 500, 100, 50, 100, 500, 100, 50}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{10, 20, 30, 40, 50, 60, 70, 80, 90}))
}

func KnapsackProblemSolve(W []int, P []int, L int) {

	tableFl, tableFx := make([][]int, len(W)+1), make([][]int, len(W)+1)
	for i, _ := range tableFl {
		tableFl[i] = make([]int, L+1)
		tableFx[i] = make([]int, L+1)
	}

	var currentW, currentP, currentP_1, MaxP, MaxX int

	for n := 1; n <= len(W); n++ {
		for l := 0; l <= L; l++ {
			for x := 0; x <= L/W[n-1]; x++ {
				currentW := x * W[n-1]
				currentP := P[n-1] * x
				currentP_1, _, _ := max(tableFl[n-1], tableFx[n-1], l-currentW)
				fmt.Printf("currentW = %v currentP = %v currentP_1 = %v\n", currentW, currentP, currentP_1)

				if currentW <= l && currentP+currentP_1 > MaxP {
					MaxP = currentP + currentP_1
					MaxX = x
				}

			}
			tableFl[n][l] = MaxP
			tableFx[n][l] = MaxX
			MaxP, MaxX = 0, 0
			fmt.Printf("L = %v n = %v max = %v\n", l, n, 1)
			printTwoTable(tableFl, tableFx)
			println()
		}
	}
	fmt.Printf("currentW = %v currentP = %v currentP_1 = %v\n", currentW, currentP, currentP_1)

	currentL, currentW, currentX, currentIdx := L, 0, 0, 0

	for i := len(tableFl) - 1; i > 0; i-- {
		currentW, currentX, currentIdx = max(tableFl[i], tableFx[i], currentL)
		fmt.Printf("currentW = %v currentX(%v) = %v currentIdx = %v\n", currentW, i, currentX, currentIdx)
		currentL -= W[i-1] * currentX
	}
}

func x5RiteilKnapsackProblemSolve(n int, price []int, amount []int, count []int, delivery []int) string {

	//init
	r := len(count) + 1
	step, _ := minOfSlice(amount)
	//step = 500
	c := (n / step) + 1
	tableCount := initTable(r, c)
	tablePrice := initTable(r, c)
	tableAmount := initTable(r, c)

	SuppliersUnitPrice := make(map[int]float64)
	// Расчитываем удельную стоимость единицы для определения порядка закупки
	for idx := 0; idx < len(count); idx++ {
		SuppliersUnitPrice[idx] = float64(delivery[idx]+price[idx]*count[idx]) / float64(count[idx]*amount[idx])
	}
	keys := make([]int, 0, len(SuppliersUnitPrice))
	for k := range SuppliersUnitPrice {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return SuppliersUnitPrice[keys[i]] < SuppliersUnitPrice[keys[j]]
	})

	// Т.к жадный алгоритм даёт неудовлетворительный результат, от него возьмём лишь сортировку
	// Основной расчет будем проводить с помощью метода динамического программирования
	for i := 1; i < r; i++ {

		currentSupplier := keys[i-1]

		for currentN := step; currentN <= n; currentN = currentN + step {

			j := (currentN / step)
			scopeAmount := *new([]int)
			scopeCount := *new([]int)
			scopePrice := *new([]int)

			for currentCount := 0; currentCount <= min(count[currentSupplier], currentN/amount[currentSupplier]); currentCount++ {

				minPrice_1 := tablePrice[i-1][j-currentCount*amount[currentSupplier]/step]
				minAmount_1 := tableAmount[i-1][j-currentCount*amount[currentSupplier]/step]

				currentPrice := currentCount*price[currentSupplier] + minPrice_1

				if currentCount > 0 {
					currentPrice += delivery[currentSupplier]
				}

				currentAmount := currentCount*amount[currentSupplier] + minAmount_1

				if currentAmount == currentN {
					scopeAmount = append(scopeAmount, currentAmount)
					scopeCount = append(scopeCount, currentCount)
					scopePrice = append(scopePrice, currentPrice)
				}
			}

			if len(scopeAmount) > 0 {

				_, idx := minOfSlice(scopePrice)

				tableAmount[i][j] = scopeAmount[idx]
				tableCount[i][j] = scopeCount[idx]
				tablePrice[i][j] = scopePrice[idx]

			}
		}
	}

	// for i, v := range keys {
	// 	fmt.Printf("index - %v supplier - %v \n", i, v)
	// }

	// printTable(tableAmount)
	// printTable(tableCount)
	// printTable(tablePrice)
	//Восстанавливаем решение из таблицы решений обратным обходом

	result := make([][]byte, 0)
	currentN := len(tableAmount[0]) - 1

	for row := len(tableAmount) - 1; row > 0; row-- {
		var col int
		for _col := currentN; _col > 0; _col-- {
			if tableAmount[row][_col] != 0 {
				col = _col
				break
			}
		}

		currenCount := tableCount[row][col]

		if currenCount == 0 {
			continue
		} else {
			currentSupplier := keys[row-1]
			result = append(result, []byte(strconv.Itoa(currentSupplier)))
			currentN -= currenCount * amount[currentSupplier] / step
		}

	}
	// Реультат в слайсе задом наперёд, разворачиваем
	// Препенд не работал для [][]byte
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(bytes.Join(result, []byte(",")))
}

func printTwoTable(table1, table2 [][]int) {
	for i, _ := range table1 {
		fmt.Printf("%v - %3v\t %3v\n", i, table1[i], table2[i])
	}
}

func printTable(table1 [][]int) {
	for i, _ := range table1 {
		fmt.Printf("%v - %6v\n", i, table1[i])
	}
	fmt.Printf("\n")
}

func minOfSlice(slice []int) (int, int) {

	min_v, min_i := slice[len(slice)-1], len(slice)-1
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] < min_v {
			min_v = slice[i]
			min_i = i
		}
	}

	return min_v, min_i
}

func initTable(r, c int) [][]int {

	table := make([][]int, r)

	for i, _ := range table {
		table[i] = make([]int, c)
	}

	return table
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(Fl, Fx []int, L int) (int, int, int) {
	maxW, maxX, idx := 0, 0, 0
	for i := 0; i <= L; i++ {
		if maxW < Fl[i] {
			maxW, maxX, idx = Fl[i], Fx[i], i
		}
	}
	return maxW, maxX, idx
}
