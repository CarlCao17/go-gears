package sort

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"
	"time"
)

var (
	smallInt        []int
	mediumInt       []int
	largeInt        []int
	expectSmallInt  []int
	expectMediumInt []int
	expectLargeInt  []int

	smallFloat        []float64
	mediumFloat       []float64
	largeFloat        []float64
	expectSmallFloat  []float64
	expectMediumFloat []float64
	expectLargeFloat  []float64

	smallStr        []string
	mediumStr       []string
	largeStr        []string
	expectSmallStr  []string
	expectMediumStr []string
	expectLargeStr  []string
)

func BenchmarkQSortIntSlice(b *testing.B) {
	b.Run("smallInt", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(smallInt, func(i, j int) bool {
				return smallInt[i] <= smallInt[j]
			})
		}
	})
	b.Run("smallInt-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sort.Slice(expectSmallInt, func(i, j int) bool {
				return expectSmallInt[i] <= expectSmallInt[j]
			})
		}
	})
	b.Run("mediumInt", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(mediumInt, func(i, j int) bool {
				return mediumInt[i] <= mediumInt[j]
			})
		}
	})
	b.Run("mediumInt-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			sort.Slice(expectMediumInt, func(i, j int) bool {
				return expectMediumInt[i] <= expectMediumInt[j]
			})
		}
	})
	b.Run("largeInt", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(largeInt, func(i, j int) bool {
				return largeInt[i] <= largeInt[j]
			})
		}
	})
	b.Run("largeInt-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(expectLargeInt, func(i, j int) bool {
				return expectLargeInt[i] <= expectLargeInt[j]
			})
		}
	})
}

func BenchmarkQSortFloatSlice(b *testing.B) {
	b.Run("smallFloat", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(smallFloat, func(i, j int) bool {
				return smallFloat[i] <= smallFloat[j]
			})
		}
	})
	b.Run("smallFloat-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(expectSmallFloat, func(i, j int) bool {
				return expectSmallFloat[i] <= expectSmallFloat[j]
			})
		}
	})
	b.Run("mediumFloat", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(mediumFloat, func(i, j int) bool {
				return mediumFloat[i] <= mediumFloat[j]
			})
		}
	})
	b.Run("mediumFloat-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(expectMediumFloat, func(i, j int) bool {
				return expectMediumFloat[i] <= expectMediumFloat[j]
			})
		}
	})
	b.Run("largeFloat", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(largeFloat, func(i, j int) bool {
				return largeFloat[i] <= largeFloat[j]
			})
		}
	})
	b.Run("largeFloat-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(expectLargeFloat, func(i, j int) bool {
				return expectLargeFloat[i] <= expectLargeFloat[j]
			})
		}
	})
}

func BenchmarkQSortStrSlice(b *testing.B) {
	b.Run("smallStr", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(smallStr, func(i, j int) bool {
				return smallStr[i] <= smallStr[j]
			})
		}
	})
	b.Run("smallStr-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(expectSmallStr, func(i, j int) bool {
				return expectSmallStr[i] <= expectSmallStr[j]
			})
		}
	})
	b.Run("mediumStr", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(mediumStr, func(i, j int) bool {
				return mediumStr[i] <= mediumStr[j]
			})
		}
	})
	b.Run("mediumStr-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(expectMediumStr, func(i, j int) bool {
				return expectMediumStr[i] <= expectMediumStr[j]
			})
		}
	})
	b.Run("largeStr", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(largeStr, func(i, j int) bool {
				return largeStr[i] <= largeStr[j]
			})
		}
	})
	b.Run("largeStr-std", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(expectLargeStr, func(i, j int) bool {
				return expectLargeStr[i] <= expectLargeStr[j]
			})
		}
	})
}

func TestQSortIntSlice(t *testing.T) {
	sort.Slice(expectSmallInt, func(i, j int) bool {
		return expectSmallInt[i] <= expectSmallInt[j]
	})
	sort.Slice(expectMediumInt, func(i, j int) bool {
		return expectMediumInt[i] <= expectMediumInt[j]
	})
	sort.Slice(expectLargeInt, func(i, j int) bool {
		return expectLargeInt[i] <= expectLargeInt[j]
	})

	type args struct {
		name   string
		slice  []int
		expect []int
	}

	tests := []args{
		{
			name:   "ascend order",
			slice:  []int{1, 2, 3, 4, 5},
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "descend order",
			slice:  []int{5, 4, 3, 2, 1},
			expect: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "smallInt",
			slice:  smallInt,
			expect: expectSmallInt,
		},
		{
			name:   "MediumInt",
			slice:  mediumInt,
			expect: expectMediumInt,
		},
		{
			name:   "largeInt",
			slice:  largeInt,
			expect: expectLargeInt,
		},
	}

	for _, tt := range tests {
		Slice(tt.slice, func(i, j int) bool {
			return tt.slice[i] <= tt.slice[j]
		})
		if !reflect.DeepEqual(tt.slice, tt.expect) {
			t.Errorf("should be equal, got=%v, expect=%v", tt.slice, tt.expect)
		}
	}
}

func TestQSortFloatSlice(t *testing.T) {
	sort.Slice(expectSmallFloat, func(i, j int) bool {
		return expectSmallFloat[i] <= expectSmallFloat[j]
	})
	sort.Slice(expectMediumFloat, func(i, j int) bool {
		return expectMediumFloat[i] <= expectMediumFloat[j]
	})
	sort.Slice(expectLargeFloat, func(i, j int) bool {
		return expectLargeFloat[i] <= expectLargeFloat[j]
	})
	type args struct {
		name   string
		slice  []float64
		expect []float64
	}

	tests := []args{
		{
			name:   "smallFloat",
			slice:  smallFloat,
			expect: expectSmallFloat,
		},
		{
			name:   "MediumFloat",
			slice:  mediumFloat,
			expect: expectMediumFloat,
		},
		{
			name:   "largeFloat",
			slice:  largeFloat,
			expect: expectLargeFloat,
		},
	}

	for _, tt := range tests {
		Slice(tt.slice, func(i, j int) bool {
			return tt.slice[i] <= tt.slice[j]
		})
		if !reflect.DeepEqual(tt.slice, tt.expect) {
			t.Errorf("should be equal, got=%v, expect=%v", tt.slice, tt.expect)
		}
	}
}

func TestQSortStrSlice(t *testing.T) {
	sort.Slice(expectSmallStr, func(i, j int) bool {
		return expectSmallStr[i] <= expectSmallStr[j]
	})
	sort.Slice(expectMediumStr, func(i, j int) bool {
		return expectMediumStr[i] <= expectMediumStr[j]
	})
	sort.Slice(expectLargeStr, func(i, j int) bool {
		return expectLargeStr[i] <= expectLargeStr[j]
	})
	type args struct {
		name   string
		slice  []string
		expect []string
	}

	tests := []args{
		{
			name:   "smallStr",
			slice:  smallStr,
			expect: expectSmallStr,
		},
		{
			name:   "MediumStr",
			slice:  mediumStr,
			expect: expectMediumStr,
		},
		{
			name:   "largeStr",
			slice:  largeStr,
			expect: expectLargeStr,
		},
	}

	for _, tt := range tests {
		Slice(tt.slice, func(i, j int) bool {
			return tt.slice[i] <= tt.slice[j]
		})
		if !reflect.DeepEqual(tt.slice, tt.expect) {
			t.Errorf("should be equal, got=%v, expect=%v", tt.slice, tt.expect)
		}
	}
}

func init() {
	start := time.Now()
	// 由于 GenerateTestCases 随机生成非常耗时，因此默认从序列化中的数据读取，而非重新生成
	pathName := "./pkg/sort/test/test_cases.txt"
	file, err := os.Open(pathName)
	if err != nil {
		panic(fmt.Sprintf("could not load from %q, will generate random case(Very Cost time!!!)", pathName))
	}
	loadTestCasesFromFile(file)
	fmt.Printf("init: cost %dms", time.Since(start).Milliseconds())
}

func loadTestCasesFromFile(file *os.File) {
	start := time.Now()
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&smallInt)
	if err != nil {
		panic(fmt.Sprintf("could not load smallInt from file %s", file.Name()))
	}
	err = decoder.Decode(&mediumInt)
	if err != nil {
		panic(fmt.Sprintf("could not load mediumInt from file %s", file.Name()))
	}
	err = decoder.Decode(&largeInt)
	if err != nil {
		panic(fmt.Sprintf("could not load largeInt from file %s", file.Name()))
	}
	err = decoder.Decode(&smallFloat)
	if err != nil {
		panic(fmt.Sprintf("could not load smallFloat from file %s", file.Name()))
	}
	err = decoder.Decode(&mediumFloat)
	if err != nil {
		panic(fmt.Sprintf("could not load mediumFloat from file %s", file.Name()))
	}
	err = decoder.Decode(&largeFloat)
	if err != nil {
		panic(fmt.Sprintf("could not load largeFloat from file %s", file.Name()))
	}
	err = decoder.Decode(&smallStr)
	if err != nil {
		panic(fmt.Sprintf("could not load smallStr from file %s", file.Name()))
	}
	err = decoder.Decode(&mediumStr)
	if err != nil {
		panic(fmt.Sprintf("could not load mediumStr from file %s", file.Name()))
	}
	err = decoder.Decode(&largeStr)
	if err != nil {
		panic(fmt.Sprintf("could not load largeStr from file %s", file.Name()))
	}

	// 并不要求测试时完全相同，只是生成时使用了相同的长度
	assertEqual(len(smallInt), len(smallFloat), len(smallStr))
	assertEqual(len(mediumInt), len(mediumFloat), len(mediumStr))
	assertEqual(len(largeInt), len(largeInt), len(largeInt))

	copyToExpect(len(smallInt), len(mediumInt), len(largeInt))

	fmt.Printf("load test case from file: cost %dms\n", time.Since(start).Milliseconds())
}

func copyToExpect(smallN, mediumN, largeN int) {
	expectSmallInt = make([]int, smallN)
	copy(expectSmallInt, smallInt)
	expectMediumInt = make([]int, mediumN)
	copy(expectMediumInt, mediumInt)
	expectLargeInt = make([]int, largeN)
	copy(expectLargeInt, largeInt)

	expectSmallFloat = make([]float64, smallN)
	copy(expectSmallFloat, smallFloat)
	expectMediumFloat = make([]float64, mediumN)
	copy(expectMediumFloat, mediumFloat)
	expectLargeFloat = make([]float64, largeN)
	copy(expectLargeFloat, largeFloat)

	expectSmallStr = make([]string, smallN)
	copy(expectSmallStr, smallStr)
	expectMediumStr = make([]string, mediumN)
	copy(expectMediumStr, mediumStr)
	expectLargeStr = make([]string, largeN)
	copy(expectLargeStr, largeStr)
}

func assertEqual(nums ...int) {
	if len(nums) == 0 {
		return
	}
	origin := nums[0]
	for i, num := range nums[1:] {
		if origin != num {
			panic(fmt.Sprintf("assertEqual: index=%d(value=%d) is not equal to the previous(value=%d)", i, num, origin))
		}
	}
}
