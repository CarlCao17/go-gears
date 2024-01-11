package eval

import "fmt"

func panicF(format string, a ...any) {
	panic(fmt.Sprintf(format, a))
}
