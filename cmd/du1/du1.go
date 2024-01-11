package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

func main() {
	verbose := flag.Bool("v", false, "show verbose progress messages")
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	filesize := make(chan int64)
	wg := sync.WaitGroup{}
	for _, root := range roots {
		wg.Add(1)
		go walkdir(root, filesize, &wg)
	}
	go func() {
		wg.Wait()
		close(filesize)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var fileNum, totalSize int64
loop:
	for {
		select {
		case x, ok := <-filesize:
			if !ok {
				break loop
			}
			fileNum++
			totalSize += x
		case <-tick:
			printDiskUsage(fileNum, totalSize)
		}
	}
	printDiskUsage(fileNum, totalSize)
}

func printDiskUsage(fileNum, totalSize int64) {
	fmt.Printf("du1: file num=%d\ttotal size=%.2fGb\n", fileNum, float64(totalSize)/1e9)
}

func walkdir(dir string, filesize chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Printf("read dir %s failed: %v", dir, entries)
		return
	}
	for _, e := range entries {
		subdir := path.Join(dir, e.Name())
		if e.IsDir() {
			wg.Add(1)
			go walkdir(subdir, filesize, wg)
		} else {
			info, err := e.Info()
			if err != nil {
				//log.Printf("read sub entry %s: %v", subdir, err)
				continue
			}
			filesize <- info.Size()
		}
	}
}
