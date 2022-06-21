package main

import "fmt"

func main() {
	var (
		arr    = []int{4, 1, 4, -4, 6, 3, 8, 8}
		result []int
		exist  = make(map[int]struct{})
	)

	for i := range arr {
		if _, ok := exist[arr[i]]; !ok {
			exist[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}

	fmt.Println(result)
}
