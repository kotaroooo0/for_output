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
// 			fmt.Println("Waiting...")
// 			if requestApi() {
// 				q <- struct{}{}
// 			}
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()

// 	for {
// 		if len(q) > 0 {
// 			break
// 		}

// 		// q に溜まるまで他の事をしたい
// 		time.Sleep(time.Second)
// 		fmt.Println("Do something")
// 	}

// 	fmt.Println("Finish")
// }

func requestApi() bool {
	time.Sleep(time.Second)
	if rand.Float64() < 0.2 {
		return true
	}
	return false
}

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
// 	defer cancel()

// 	fmt.Println("goroutine start")
// 	go func() {
// 		for {
// 			if requestApi() {
// 				fmt.Println("goroutine done")
// 				cancel()
// 			}
// 		}
// 	}()

// 	<-ctx.Done()
// 	fmt.Println("main done")
// }

func main() {
	q := make(chan struct{})

	go func() {
		for {
			fmt.Println("Waiting...")
			rand := rand.Float64()
			fmt.Println(rand)
			if rand < 0.5 {
				q <- struct{}{} // qに入れる
			}
			time.Sleep(1 * time.Second)
		}
	}()

	<-q // q に何か入るまで待つ
	fmt.Println("Finish")
}
