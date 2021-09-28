package numberOfIslands

import (
	"fmt"
)

func NumberOfIslands(grid [][]byte) {
	numberOfIslands(grid)
}

func numberOfIslands(grid [][]byte) {

	adjacencyList := make(map[int][]int)
	//visitedNodes := make(map[int]bool)

	for i, _ := range grid {
		for j, _ := range grid[i] {

			if grid[i][j] == '0' {
				continue
			}

			fmt.Printf("i - %v j - %v n - %v \n", i, j, IJToIndex(i, j, len(grid[i])))

			_i, _j := IndexToIJ(IJToIndex(i, j, len(grid[i])), len(grid[i]))

			neighbors := make([]int, 0)

			//top
			if i > 0 && grid[i-1][j] == '1' {
				neighbors = append(neighbors, IJToIndex(i-1, j, len(grid[i])))
			}
			//bot
			if i < len(grid)-1 && grid[i+1][j] == '1' {
				neighbors = append(neighbors, IJToIndex(i+1, j, len(grid[i])))
			}
			//left
			if j > 0 && grid[i][j-1] == '1' {
				neighbors = append(neighbors, IJToIndex(i, j-1, len(grid[i])))
			}
			//right
			if j < len(grid[i])-1 && grid[i][j+1] == '1' {
				neighbors = append(neighbors, IJToIndex(i, j+1, len(grid[i])))
			}

			adjacencyList[IJToIndex(i, j, len(grid[i]))] = neighbors

			fmt.Printf("i - %v j - %v\n", _i, _j)
		}
	}

	for node, neighbors := range adjacencyList {
		fmt.Printf("node - %v neighbors - %v\n", node, neighbors)
		for _, neighbor := range neighbors {
			fmt.Printf("\t neighbor - %v\n", neighbor)
		}
	}

}

func IJToIndex(i, j, lenght int) int {
	return i*lenght + j
}

func IndexToIJ(index, lenght int) (int, int) {
	return index / lenght, index % lenght
}
