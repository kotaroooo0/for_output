package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	q := make(chan struct{})

	go func() {
		for {
			time.Sleep(time.Second * 1)
			r := rand.Int() / 100000000
			fmt.Println(r)
			if r < 10372007942 {
				q <- struct{}{}
			}
		}
	}()

	// q に何か入るまで待つ
	<-q
	fmt.Println("fuga")
}
