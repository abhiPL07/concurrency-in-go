package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go SimpleFunc()
	go func() {
		for j := 1; j <= 5; j++ {
			time.Sleep(time.Millisecond * 12)
			fmt.Printf("%v	Func 3\n", time.Now().UnixMilli())
		}
		wg.Done()
	}()
	wg.Add(1)
	go ValFunc(4)
	wg.Wait()
}
