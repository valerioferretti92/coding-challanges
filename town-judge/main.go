package main

import "fmt"

func main() {
	trust := [][]int{{2, 3}, {3, 1}, {2, 1}, {4, 1}}
	fmt.Println(findJudge(4, trust))
}

func findJudge(n int, trust [][]int) int {
	adj := make([][]int, n)

	for _, t := range trust {
		a := t[0] - 1
		b := t[1] - 1

		if adj[a] == nil {
			adj[a] = make([]int, n)
		}

		adj[a][b] = 1
	}

	for i := 0; i < n; i++ {
		sumr := sumRow(adj, i)
		sumc := sumColumn(adj, i)

		if sumr == 0 && sumc == n-1 {
			return i + 1
		}
	}
	return -1
}

func sumRow(adj [][]int, i int) int {
	if adj[i] == nil {
		return 0
	}

	sum := 0
	for _, t := range adj[i] {
		sum = sum + t
	}

	return sum
}

func sumColumn(adj [][]int, j int) int {
	sum := 0

	for i := 0; i < len(adj); i++ {
		if adj[i] == nil {
			continue
		}

		sum = sum + adj[i][j]
	}

	return sum
}
