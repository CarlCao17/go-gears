package main

import (
	"fmt"
	"sort"
)

func main() {
	var prereqs = map[string][]string{}
	for i, course := range TopoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// TopoSort using DFS
func TopoSort(m map[string][]string) []string {
	var order []string
	var seen map[string]bool
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
