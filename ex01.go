package main

import (
	"fmt"
	"time"
)

func SimpleFunc() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 20)
		fmt.Printf("%v	Func 1\n", time.Now().UnixMilli())
	}
	wg.Done()
}

func ValFunc(i int) {
	for j := 1; j <= 5; j++ {
		time.Sleep(time.Millisecond * 28)
		fmt.Printf("%v	%v	Func 2\n", time.Now().UnixMilli(), i*j)
	}
	wg.Done()
}
