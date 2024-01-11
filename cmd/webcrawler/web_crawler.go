package main

import (
	"fmt"
	"log"

	"github.com/CarlCao17/go-gears/pkg/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.ExtractLinks(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
