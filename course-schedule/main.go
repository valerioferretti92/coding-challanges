package main

import "fmt"

func main() {
	prerequisites := [][]int{{}}
	fmt.Println(canFinish(3, prerequisites))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)

	for _, pre := range prerequisites {
		if len(pre) < 2 {
			continue
		}

		p := pre[1]
		t := pre[0]

		if adj[t] == nil {
			adj[t] = make([]int, numCourses)
		}

		adj[t][p] = 1
	}

	freeCourses := getFreeCourses(adj)
	for len(freeCourses) != 0 {
		offset := 0
		for _, course := range freeCourses {
			adj = removeRow(adj, course-offset)
			adj = removeColumn(adj, course-offset)
			offset++
		}

		freeCourses = getFreeCourses(adj)
	}

	return len(adj) == 0
}

func removeColumn(adj [][]int, j int) [][]int {
	if len(adj) == 0 {
		return adj
	}

	for i := 0; i < len(adj); i++ {
		if adj[i] == nil {
			continue
		}

		if j == len(adj[i])-1 {
			adj[i] = adj[i][:len(adj[i])-1]
			continue
		}

		if j == 0 {
			adj[i] = adj[i][1:]
			continue
		}

		adj[i] = append(adj[i][:j], adj[i][j+1:]...)
	}

	return adj
}

func removeRow(adj [][]int, i int) [][]int {
	if len(adj) == 0 {
		return adj
	}

	if i == len(adj)-1 {
		return adj[:len(adj)-1]
	}

	if i == 0 {
		return adj[1:]
	}

	return append(adj[:i], adj[i+1:]...)
}

func getFreeCourses(adj [][]int) []int {
	free := make([]int, 0)

	for i, pres := range adj {
		if pres == nil {
			free = append(free, i)
			continue
		}

		sum := 0
		for _, i := range pres {
			sum = sum + i
		}

		if sum == 0 {
			free = append(free, i)
		}
	}

	return free
}
