package cache

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

// for some very simple function
// goos: darwin
// goarch: amd64
// pkg: caozhengsheng.carl/go_src/cache
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_Memo_Get_Add/Memo_Get_Add-12              2605            415371 ns/op
// Benchmark_Memo_Get_Add/Add_nocache-12            7286642               163.9 ns/op
// Benchmark_Memo_Get_Multi/Memo_Get_Add-12             966           1227491 ns/op
// Benchmark_Memo_Get_Multi/Add_nocache-12          7265521               163.9 ns/op
// PASS
// ok      caozhengsheng.carl/go_src/cache 6.451s

// goos: darwin
// goarch: amd64
// pkg: caozhengsheng.carl/go_src/cache
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// Benchmark_Memo_Get_Add/Memo_Get_Add-12                 2         500155025 ns/op
// Benchmark_Memo_Get_Add/Add_nocache-12                  2         500201040 ns/op
// Benchmark_Memo_Get_Multi/Memo_Get_Add-12               1        1000421153 ns/op
// Benchmark_Memo_Get_Multi/Add_nocache-12                1        1000171165 ns/op
// PASS
// ok      caozhengsheng.carl/go_src/cache 156.693s

func Test(t *testing.T) {
	addMemo := New(func(args ...interface{}) (string, error) {
		a := args[0].(int32)
		b := args[1].(int64)
		return fmt.Sprintf("%d + %d = ", a, b), nil
	}, func(args ...interface{}) (interface{}, error) {
		a := args[0].(int32)
		b := args[1].(int64)
		return Add(a, b), nil
	})

	start := time.Now()
	val, err := addMemo.Get(int32(1), int64(2))
	dur := time.Since(start)
	fmt.Println(val, err, dur)
	start = time.Now()
	val, err = addMemo.Get(int32(1), int64(2))
	dur = time.Since(start)
	fmt.Println(val, err, dur)
}

func Benchmark_Memo_Get_Handle_Add(b *testing.B) {
	cases, _ := prepareAddTestCases(100)

	b.Run("Memo_Get_Add", func(b *testing.B) {
		addMemo := New(func(args ...interface{}) (string, error) {
			a := args[0].(int32)
			b := args[1].(int64)
			return fmt.Sprintf("%d + %d = ", a, b), nil
		}, func(args ...interface{}) (interface{}, error) {
			a := args[0].(int32)
			b := args[1].(int64)
			return Add(a, b), nil
		})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cs := cases[0]
			_, _ = addMemo.Get(cs.a, cs.b)
		}
		addMemo.Close()
	})
	b.Run("Memo_Get_Add_handle", func(b *testing.B) {
		addMemo := New2(func(args ...interface{}) (string, error) {
			a := args[0].(int32)
			b := args[1].(int64)
			return fmt.Sprintf("%d + %d = ", a, b), nil
		}, func(args ...interface{}) (interface{}, error) {
			a := args[0].(int32)
			b := args[1].(int64)
			return Add(a, b), nil
		})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			//idx := i % len(cases)
			cs := cases[0]
			_, _ = addMemo.Get(cs.a, cs.b)
		}
		addMemo.Close()
	})
}

// TODO: 如何控制重复度，来确定结构
func Benchmark_Memo_Get_Add(b *testing.B) {
	cases, _ := prepareAddTestCases(100)

	b.Run("Memo_Get_Add", func(b *testing.B) {
		addMemo := New(func(args ...interface{}) (string, error) {
			a := args[0].(int32)
			b := args[1].(int64)
			return fmt.Sprintf("%d + %d = ", a, b), nil
		}, func(args ...interface{}) (interface{}, error) {
			a := args[0].(int32)
			b := args[1].(int64)
			return Add(a, b), nil
		})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cs := cases[1]
			_, _ = addMemo.Get(cs.a, cs.b)
		}
		addMemo.Close()
	})
	b.Run("Add_nocache", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			cs := cases[1]
			_ = Add(cs.a, cs.b)
		}
	})
}

func Benchmark_Memo_Get_Multi(b *testing.B) {
	cases, _ := prepareMultiTestCases(100)

	b.Run("Memo_Get_Add", func(b *testing.B) {
		multiMemo := New(func(args ...interface{}) (string, error) {
			a := args[0].(float64)
			b := args[1].(int32)
			return fmt.Sprintf("%.2f * %d = ", a, b), nil
		}, func(args ...interface{}) (interface{}, error) {
			a := args[0].(float64)
			b := args[1].(int32)
			return Multi(a, b), nil
		})
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			idx := i % len(cases)
			cs := cases[idx]
			_, _ = multiMemo.Get(cs.a, cs.b)
		}
		multiMemo.Close()
	})
	b.Run("Add_nocache", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			idx := i % len(cases)
			cs := cases[idx]
			_ = Multi(cs.a, cs.b)
		}
	})
}

func TestMemo_demo_Add(t *testing.T) {
	addMemo := New(func(args ...interface{}) (string, error) {
		a := args[0].(int32)
		b := args[1].(int64)
		return fmt.Sprintf("%d + %d = ", a, b), nil
	}, func(args ...interface{}) (interface{}, error) {
		a := args[0].(int32)
		b := args[1].(int64)
		return Add(a, b), nil
	})

	ch := make(chan stat, 100)
	addCases, stats := prepareAddTestCases(10)

	wg := sync.WaitGroup{}
	for i, cs := range addCases {
		wg.Add(1)
		go func(i int, cs *addCase) {
			start := time.Now()
			val, err := addMemo.Get(cs.a, cs.b)
			dur := time.Since(start)
			assert(err == nil)
			assertEqual(val.(int64), cs.expect)
			ch <- stat{
				str:  genAddStatement(cs),
				idx:  i,
				cost: dur,
			}
			wg.Done()
		}(i, cs)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for st := range ch {
		stats[st.str][st.idx] = st.cost
	}
	for str, m := range stats {
		fmt.Printf("| %s | idx |   cost  |\n", str)
		for idx, cost := range m {
			fmt.Printf("|%*s|", len(str)+2, " ")
			fmt.Printf(" %d |%v|\n", idx, cost)
		}
	}
}

func TestMemo_demo_Multi(t *testing.T) {
	multiMemo := New(func(args ...interface{}) (string, error) {
		a := args[0].(float64)
		b := args[1].(int32)
		return fmt.Sprintf("%.2f * %d = ", a, b), nil
	}, func(args ...interface{}) (interface{}, error) {
		a := args[0].(float64)
		b := args[1].(int32)
		return Multi(a, b), nil
	})
	wg := sync.WaitGroup{}
	ch := make(chan stat, 10)
	multiCases, stats := prepareMultiTestCases(17)
	for i, cs := range multiCases {
		wg.Add(1)
		go func(i int, cs *multiCase) {
			start := time.Now()
			val, err := multiMemo.Get(cs.a, cs.b)
			dur := time.Since(start)
			assert(err == nil)
			assertEqual(val.(float64), cs.expect)
			ch <- stat{
				str:  genMultiStatement(cs),
				idx:  i,
				cost: dur,
			}
			wg.Done()
		}(i, cs)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for st := range ch {
		stats[st.str][st.idx] = st.cost
	}
	for str, m := range stats {
		fmt.Printf("| %s | idx |   cost  |\n", str)
		for idx, cost := range m {
			fmt.Printf("|%*s|", len(str)+2, " ")
			fmt.Printf(" %d |%v|\n", idx, cost)
		}
	}
}

type stat struct {
	str  string
	idx  int
	cost time.Duration
}

type addCase struct {
	a       int32
	b       int64
	expect  int64
	isFirst bool
}

func Add(a int32, b int64) int64 {
	time.Sleep(100 * time.Millisecond)
	return int64(a) + b
}

type multiCase struct {
	a       float64
	b       int32
	expect  float64
	isFirst bool
}

func Multi(a float64, b int32) float64 {
	time.Sleep(2000 * time.Millisecond)
	return a * float64(b)
}

func prepareAddTestCases(n int) ([]*addCase, map[string]map[int]time.Duration) {
	res := make([]*addCase, 2*n)
	temp := make([]addCase, n)
	for i := 0; i < n; i++ {
		temp[i].a, temp[i].b = rand.Int31(), rand.Int63()
		temp[i].expect = Add(temp[i].a, temp[i].b)
	}

	for i := 0; i < n; i++ {
		res[i] = &temp[i]
	}
	for i := n; i < 2*n; i++ {
		res[i] = &temp[rand.Intn(n)]
	}
	rand.Shuffle(2*n, func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	stats := make(map[string]map[int]time.Duration, n)
	for i, c := range res {
		str := genAddStatement(c)
		if _, exist := stats[str]; !exist {
			stats[str] = make(map[int]time.Duration)
		}
		stats[str][i] = 0
	}
	return res, stats
}

func genAddStatement(c *addCase) string {
	return fmt.Sprintf("%d + %d = %d", c.a, c.b, c.expect)
}

func genMultiStatement(c *multiCase) string {
	return fmt.Sprintf("%.2f * %d = %.2f", c.a, c.b, c.expect)
}

func prepareMultiTestCases(n int) ([]*multiCase, map[string]map[int]time.Duration) {
	res := make([]*multiCase, 2*n)
	temp := make([]multiCase, n)
	for i := 0; i < n; i++ {
		temp[i].a, temp[i].b = rand.Float64(), rand.Int31()
		temp[i].expect = Multi(temp[i].a, temp[i].b)
	}
	for i := 0; i < n; i++ {
		res[i] = &temp[i]
	}
	for i := n; i < 2*n; i++ {
		res[i] = &temp[rand.Intn(n)]
	}
	rand.Shuffle(2*n, func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})
	stats := make(map[string]map[int]time.Duration, n)
	for i, c := range res {
		str := genMultiStatement(c)
		if _, exist := stats[str]; !exist {
			stats[str] = make(map[int]time.Duration)
		}
		stats[str][i] = 0
	}
	return res, stats
}

func assert(b bool) {
	if !b {
		panic("assert failed")
	}
}

func assertEqual(left interface{}, right interface{}) {
	lT, rT := reflect.TypeOf(left), reflect.TypeOf(right)
	if lT != rT {
		panic(fmt.Sprintf("type is not equal, left type is %s, right type is %s", lT.String(), rT.String()))
	}
	if !reflect.DeepEqual(left, right) {
		panic(fmt.Sprintf("value is not equal, left value is %v, right value is %v", left, right))
	}
}
