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
			time.Sleep(time.Second)
			rand := rand.Float64()
			fmt.Println(rand)
			if rand < 0.2 {
				q <- struct{}{}
			}

			fmt.Println("Waiting...")
			time.Sleep(1 * time.Second)
		}
	}()
	<-q

	fmt.Println("Finish")
}
