package main

import "os"

func main() {
	breadthFirst(crawl, os.Args[1:])
}
