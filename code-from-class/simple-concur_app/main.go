package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	mul = flag.Int64("mul", 2, "Description")
	max = flag.Int64("max", 100, "max value")
)

func main() {
	flag.Parse()

	results := make([]int64, 0, *max)
	mu := sync.Mutex{}

	wg := &sync.WaitGroup{}
	var i int64 = 0
	// var s string = ""
	// slice, map, chan, interface, pointer = nil
	for ; i < *max; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			result, err := Worker(i, "Woohoo!")
			if err != nil {
				log.Fatal(err)
			}
			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(results)
}

func Worker(num int64, _ string) (int64, error) {
	defer func() {
		fmt.Println(num, "Exit =)")
	}()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return num * *mul, nil
}

