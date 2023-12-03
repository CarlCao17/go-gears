package main

import (
	"flag"
	"fmt"
	"strings"
)

func echo(args []string, n bool) {
	fmt.Print(strings.Join(args, " "))
	if !n {
		fmt.Println()
	}
}

func main() {
	var n bool

	flag.BoolVar(&n, "n", false, "Do not print the trailing newline character.  This may also be achieved by appending ‘\\c’ to the end of the string, as is done by iBCS2 compatible systems.  Note that this option as well as the effect of ‘\\c’ are\n           implementation-defined in IEEE Std 1003.1-2001 (“POSIX.1”) as amended by Cor. 1-2002.  Applications aiming for maximum portability are strongly encouraged to use printf(1) to suppress the newline character.")
	flag.Parse()

	args := flag.Args()
	if len(args) > 1 {
		end := args[len(args)-1]
		if end[len(end)-2:] == "\\c" {
			n = true
		}
	}
	echo(args, n)
}
