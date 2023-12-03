package eval

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return u.x.Eval(env)
	case '-':
		return -(u.x.Eval(env))
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

var (
	binaryOpFunc = map[rune]func(e Env, x, y Expr) float64{
		'+': Add,
		'-': Sub,
		'*': Multi,
		'/': Div,
	}
)

func (b binary) Eval(env Env) float64 {
	f, ok := binaryOpFunc[b.op]
	if !ok {
		panicF("unsupported binary operation: %q", b.op)
	}
	return f(env, b.x, b.y)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unsupported binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

var (
	callFunc = map[string]func(e Env, args []Expr) float64{
		"sqrt": func(e Env, args []Expr) float64 {
			x := args[0].Eval(e)
			return math.Sqrt(x)
		},
		"pow": func(e Env, args []Expr) float64 {
			x, y := args[0].Eval(e), args[1].Eval(e)
			return math.Pow(x, y)
		},
		"sin": func(e Env, args []Expr) float64 {
			x := args[0].Eval(e)
			return math.Sin(x)
		},
	}
)

func (c call) Eval(env Env) float64 {
	f, ok := callFunc[c.fn]
	if !ok {
		panicF("unsupported function call: %q", c.fn)
	}
	return f(env, c.args)
}

var (
	numParams = map[string]int{
		"sqrt": 1,
		"pow":  2,
		"sin":  1,
	}
)

func (c call) Check(vars map[Var]bool) error {
	num, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if num != len(c.args) {
		return fmt.Errorf("call to %s has %d args, want %d", c.fn, len(c.args), num)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}
