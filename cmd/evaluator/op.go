package eval

func Add(e Env, x, y Expr) float64 {
	return x.Eval(e) + y.Eval(e)
}

func Sub(e Env, x, y Expr) float64 {
	return x.Eval(e) - y.Eval(e)
}

func Multi(e Env, x, y Expr) float64 {
	return x.Eval(e) * y.Eval(e)
}

func Div(e Env, x, y Expr) float64 {
	return x.Eval(e) / y.Eval(e)
}
