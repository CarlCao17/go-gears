package sort

import (
	"context"

	"go_src/pkg/async_worker"
)

type Sortable interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

type sortReq struct {
	aw     async_worker.AsyncWorker[sortReq]
	data   Sortable
	lo, hi int
}

type bySlice[T any] struct {
	s    []T
	less func(i, j int) bool
}

func (b bySlice[T]) Len() int {
	return len(b.s)
}

func (b bySlice[T]) Swap(i, j int) {
	b.s[i], b.s[j] = b.s[j], b.s[i]
}

func (b bySlice[T]) Less(i, j int) bool {
	return b.less(i, j)
}

func Slice[T any](data []T, less func(i, j int) bool) {
	QSort(bySlice[T]{data, less})
}

func QSort(data Sortable) {
	lo, hi := 0, data.Len()
	aw := async_worker.NewAsyncWorker(context.Background(), hi+1, func(ctx context.Context, req sortReq) {
		qsort(req.aw, req.data, req.lo, req.hi)
	})
	aw.Submit(sortReq{aw, data, lo, hi})
	aw.Wait()
}

func qsort(aw async_worker.AsyncWorker[sortReq], data Sortable, lo, hi int) {
	pivot := partition(data, lo, hi)
	aw.Submit(sortReq{aw, data, lo, pivot - 1})
	aw.Submit(sortReq{aw, data, pivot + 1, hi})
}

// [lo, hi)
// [lo, i] <= data[pivot]
// [i+1, j) > data[pivot]
// [j, hi-1) to be scanned, pivot=hi-1 also need to be scanned and swapped
func partition(data Sortable, lo, hi int) (pivot int) {
	i, j := lo-1, lo
	pivot = hi - 1
	for j < hi-1 {
		if data.Less(j, pivot) {
			i++
			data.Swap(i, j)
		}
		j++
	}
	data.Swap(i+1, pivot)
	i++ // In this time, [lo, pivot) <= data[pivot], pivot, [pivot+1， j] > data[pivot]
	return pivot
}
