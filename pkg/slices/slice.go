package slices

import (
	"sort"
	"strconv"

	"go_src/pkg/constraints"
)

func NewSliceWith[T any](val T, n int) []T {
	slice := make([]T, n)
	for i := 0; i < n; i++ {
		slice[i] = val
	}
	return slice
}

func Sum[T constraints.RealNumber](s []T) T {
	return Reduce(s, RFSum[T])
}

func Multi[T constraints.RealNumber](s []T) T {
	return ReduceWithAcc(s, 1, RFMulti[T])
}

func Max[T constraints.RealNumber](s []T) T {
	return Reduce(s, RFMax[T])
}

func Min[T constraints.RealNumber](s []T) T {
	if len(s) == 0 {
		panic("Min: should have length at least one")
	}
	if len(s) == 1 {
		return s[0]
	}
	return ReduceWithAcc(s, s[0], RFMin[T])
}

func Reduce[T any](s []T, f func(acc T, item T) T) T {
	var acc T
	return ReduceWithAcc(s, acc, f)
}

func ReduceWithAcc[T any](s []T, acc T, f func(acc T, item T) T) T {
	for _, item := range s {
		acc = f(acc, item)
	}
	return acc
}

func MapReduce[T any, V any](s []T, mapFunc func(T) V, initAcc V, reduceFunc func(acc V, item V) V) V {
	acc := initAcc
	for _, item := range s {
		acc = reduceFunc(acc, mapFunc(item))
	}
	return acc
}

type ReduceFunc[T any] func(acc T, item T) T

type MapFunc[T any, V any] func(item T) V

func RFSum[T constraints.RealNumber](acc T, item T) T {
	return acc + item
}

func RFMulti[T constraints.RealNumber](acc T, item T) T {
	return acc * item
}

func RFMax[T constraints.RealNumber](acc T, item T) T {
	if item > acc {
		return item
	}
	return acc
}

func RFMin[T constraints.RealNumber](acc T, item T) T {
	if item < acc {
		return item
	}
	return acc
}

func Map[T any, V any](s []T, mapFunc func(T) V) []V {
	res := make([]V, len(s))
	for i, ss := range s {
		res[i] = mapFunc(ss)
	}
	return res
}

func StrToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func StrToInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, strconv.IntSize)
	return int(i)
}

func GroupBy[T any, V comparable](s []T, f func(T) V) map[V][]T {
	m := make(map[V][]T)
	for _, t := range s {
		key := f(t)
		m[key] = append(m[key], t)
	}
	return m
}

func GroupByAndSort[T any, V comparable](s []T, f func(T) V, less func(T, T) bool) map[V][]T {
	groups := GroupBy(s, f)
	for k, v := range groups {
		b := by[T]{s: v, less: less}
		sort.Stable(b)
		groups[k] = b.s
	}
	return groups
}

type by[T any] struct {
	s    []T
	less func(T, T) bool
}

func (b by[T]) Len() int { return len(b.s) }
func (b by[T]) Swap(i, j int) {
	b.s[i], b.s[j] = b.s[j], b.s[i]
}
func (b by[T]) Less(i, j int) bool {
	return b.less(b.s[i], b.s[j])
}
