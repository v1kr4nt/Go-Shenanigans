package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	username := fetchUser()
	// ------------ before ------------
	// likes := fetchUserLikes(username)
	// match := fetchUserMatch(username)
	//-----------------------------

	//---------------after go routines------------
	// creating channel to get access of likes and match
	respch := make(chan any, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go fetchUserLikes(username, respch, wg)
	go fetchUserMatch(username, respch, wg)
	wg.Wait() //block until 2 wg.Done()
	close(respch)
	// fmt.Println("likes: ", likes)
	// fmt.Println("match: ", match)
	for resp := range respch {
		fmt.Println("resp: ", resp)
	}
	fmt.Println("took us: ", time.Since(start))
}
func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "BOB"
}
func fetchUserLikes(username string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	respch <- 11
	wg.Done()
}
func fetchUserMatch(username string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- "ANNA"
	wg.Done()
}
