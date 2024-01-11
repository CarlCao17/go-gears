package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	smallNMax  = 100
	mediumNMax = 10_000
	largeNMax  = 1000_000
)

func main() {
	path := "./pkg/sort/test/test_cases.txt"
	GenerateTestCases(path)
}

func GenerateTestCases(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	start := time.Now()
	var costInt, costFloat, costStr, costCopy time.Duration
	defer func() {
		cost := time.Since(start)
		fmt.Printf("GenerateTestCases: generate rand test case\n\ttotal cost %dms\n", cost.Milliseconds())
		fmt.Printf("\tint cost %dms\n", costInt.Milliseconds())
		fmt.Printf("\tfloat cost %dms\n", costFloat.Milliseconds())
		fmt.Printf("\tstring cost %dms\n", costStr.Milliseconds())
		fmt.Printf("\tcopy cost %dms\n", costCopy.Milliseconds())
	}()
	rand.NewSource(time.Now().UnixNano())

	smallN := rand.Intn(smallNMax)
	mediumN := rand.Intn(mediumNMax)
	largeN := rand.Intn(largeNMax)

	var smallInt, mediumInt, largeInt []int
	var smallFloat, mediumFloat, largeFloat []float64
	var smallStr, mediumStr, largeStr []string
	wg.Add(3)
	go func() {
		startInt := time.Now()
		smallInt = NewRandSlice(smallN, rand.Int)
		mediumInt = NewRandSlice(mediumN, rand.Int)
		largeInt = NewRandSlice(largeN, rand.Int)
		costInt = time.Since(startInt)
		wg.Done()
	}()
	go func() {
		startFloat := time.Now()
		smallFloat = NewRandSlice(smallN, rand.Float64)
		mediumFloat = NewRandSlice(mediumN, rand.Float64)
		largeFloat = NewRandSlice(largeN, rand.Float64)
		costFloat = time.Since(startFloat)
		wg.Done()
	}()
	go func() {
		smallStringLen := rand.Intn(1024 + 1)         // <= 1KB
		mediumStringLen := rand.Intn(100*1024 + 1)    // <= 100KB
		largeStringLen := rand.Intn(10*1024*1024 + 1) // <= 10MB

		startStr := time.Now()
		smallStr = NewRandStringSlice(smallN, smallStringLen, NewRandString)
		mediumStr = NewRandStringSlice(mediumN, mediumStringLen, NewRandString)
		largeStr = NewRandStringSlice(largeN, largeStringLen, NewRandString)
		costStr = time.Since(startStr)
		wg.Done()
	}()
	wg.Wait()

	storeTestCasesIntoFile(file, smallInt, mediumInt, largeInt, smallFloat, mediumFloat, largeFloat, smallStr, mediumStr, largeStr)
}

func storeTestCasesIntoFile(file *os.File, smallInt, mediumInt, largeInt []int, smallFloat, mediumFloat, largeFloat []float64, smallStr, mediumStr, largeStr []string) {
	encoder := json.NewEncoder(file)
	err := encoder.Encode(smallInt)
	if err != nil {
		panic(fmt.Sprintf("could not store smallInt into file %s", file.Name()))
	}
	err = encoder.Encode(mediumInt)
	if err != nil {
		panic(fmt.Sprintf("could not store mediumInt into file %s", file.Name()))
	}
	err = encoder.Encode(largeInt)
	if err != nil {
		panic(fmt.Sprintf("could not store largeInt into file %s", file.Name()))
	}
	err = encoder.Encode(smallFloat)
	if err != nil {
		panic(fmt.Sprintf("could not store smallFloat into file %s", file.Name()))
	}
	err = encoder.Encode(mediumFloat)
	if err != nil {
		panic(fmt.Sprintf("could not store mediumFloat into file %s", file.Name()))
	}
	err = encoder.Encode(largeFloat)
	if err != nil {
		panic(fmt.Sprintf("could not store largeFloat into file %s", file.Name()))
	}
	err = encoder.Encode(smallStr)
	if err != nil {
		panic(fmt.Sprintf("could not store smallStr into file %s", file.Name()))
	}
	err = encoder.Encode(mediumStr)
	if err != nil {
		panic(fmt.Sprintf("could not store mediumStr into file %s", file.Name()))
	}
	err = encoder.Encode(largeStr)
	if err != nil {
		panic(fmt.Sprintf("could not store largeStr into file %s", file.Name()))
	}
}

func NewRandSlice[T any](n int, gen func() T) []T {
	s := make([]T, n)
	for i := 0; i < n; i++ {
		s[i] = gen()
	}
	return s
}

func NewRandStringSlice(n int, m int, gen func(int2 int) string) []string {
	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = gen(m)
	}
	return s
}

var (
	someASCIIRunes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz,.;{}\"':*!()[]-+&^%$#@\\/"
)

func NewRandString(m int) string {
	var s strings.Builder
	for i := 0; i < m; i++ {
		s.WriteRune(randASCIIRune(someASCIIRunes))
	}
	return s.String()
}

func randASCIIRune(source string) rune {
	return rune(source[rand.Intn(len(source))])
}
