package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	lines := map[string]bool{}

	if len(args) == 0 {
		dupLines, err := getDupLines(os.Stdin, lines)
		if err != nil {
			log.Fatalf("duplines: %v\n", err)
		}
		for no, str := range dupLines {
			fmt.Printf("line %d: %s\n", no, str)
		}
		return
	}
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			fmt.Printf("duplines: open file=%s failed, err=%v\n", arg, err)
			continue
		}
		dupLines, err := getDupLines(file, lines)
		if err != nil {
			fmt.Printf("duplines: file=%s, err=%v\n", arg, err)
			continue
		}
		for no, str := range dupLines {
			fmt.Printf("%s line %d: %s\n", arg, no, str)
		}
	}
}

func getDupLines(file *os.File, lines map[string]bool) (dupLines map[int]string, err error) {
	scanner := bufio.NewScanner(file)
	dupLines = map[int]string{}

	idx := 0
	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := lines[line]; ok {
			dupLines[idx] = line
		} else {
			lines[line] = true
		}
		idx++
	}
	err = scanner.Err()
	return
}
