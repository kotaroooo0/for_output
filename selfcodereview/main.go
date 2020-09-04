package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"

	"golang.org/x/sync/errgroup"
)

func randomResponse() (*http.Response, error) {
	if rand.Float32() < 0.1 {
		return nil, fmt.Errorf("error")
	}
	return &http.Response{}, nil
}

func getResponses() ([]*http.Response, error) {
	var responses []*http.Response
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			res, err := randomResponse()
			if err != nil {
				// TODO: エラーを伝播できない
				log.Fatal(err)
			}
			mutex.Lock()
			responses = append(responses, res)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return responses, nil
}

func getResponsesError() ([]*http.Response, error) {
	var responses []*http.Response
	mutex := sync.Mutex{}
	var e error
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			res, err := randomResponse()
			if err != nil {
				e = err
			}
			mutex.Lock()
			responses = append(responses, res)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	if e != nil {
		return nil, e
	}
	return responses, nil
}

func getResponsesErrGroups() ([]*http.Response, error) {
	var responses []*http.Response
	mutex := sync.Mutex{}
	eg := errgroup.Group{}

	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			res, err := randomResponse()
			if err != nil {
				return err
			}
			mutex.Lock()
			responses = append(responses, res)
			mutex.Unlock()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return responses, nil
}

type Profile struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func main() {
	// prof := Profile{
	// 	Name:   "Yamada Tarou",
	// 	Age:    18,
	// 	Height: 178.5,
	// }
	// var p Profile
	// p := Profile{}
	var p []string
	// p := []string{}
	profJson, _ := json.Marshal(p)
	fmt.Println(profJson)
}

// func main() {
// 	rs, err := getResponses()
// 	fmt.Println(len(rs))
// 	fmt.Println(err)
// }
