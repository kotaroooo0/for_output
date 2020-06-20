package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	q := make(chan struct{}, 2)

// 	go func() {
// 		for {
// 			time.Sleep(time.Second)
// 			rand := rand.Float64()
// 			fmt.Println(rand)
// 			if rand < 0.2 {
// 				q <- struct{}{}
// 			}

// 			fmt.Println("Waiting...")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	for {
// 		if len(q) > 0 {
// 			break
// 		}

// 		// q に溜まるまで他の事をしたい
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println("並列実行")
// 	}

// 	fmt.Println("Finish")
// }

func main() {
	q := make(chan struct{})

	go func() {
		for {
			rand := rand.Float64()
			fmt.Println(rand)
			if rand < 0.2 {
				q <- struct{}{}
			}

			fmt.Println("Waiting...")
			time.Sleep(1 * time.Second)
		}
	}()

	// q に何か入るまで待つ
	<-q
	fmt.Println("Finish")
}
