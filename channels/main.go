package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//channels = it works like a pipe.. you put stuff into and get stuff out of..

//buffer and unbufferd channel

//buffer -> make(chan string, size)
//size = int
// testch := make(chan string, 2)

//------------------------example1-----------------------------
// func main() {
// 	userch := make(chan string, 1)

// 	userch <- "bob"
// 	userch <- "bobby"

// 	user := <-userch
// 	fmt.Println(user)
// }

//--------------------------

// ---------------example2------------

// type Server struct {
// 	users map[string]string
// }

// func NewServer() *Server {
// 	return &Server{
// 		users: make(map[string]string),
// 	}
// }

// func main() {

// }

// func sendMessage(msgch chan<- string) {
// 	msgch <- "hello"
// }

// func readMessage(msgch <-chan string) {
// 	msg := <-msgch
// 	fmt.Println(msg)
// }

//-----------------------------------

// ---------------example 3 -----------

//channel vs waitgroup

//-----------waitgroup eg--------

// func doWork(d time.Duration, wg *sync.WaitGroup) {
// 	fmt.Println("doing work...")
// 	time.Sleep(d)
// 	fmt.Println("work is done")
// 	wg.Done()
// }

// func main() {
// 	start := time.Now()
// 	wg := sync.WaitGroup{}
// 	wg.Add(2)
// 	go doWork(time.Second * 2, &wg)
// 	go doWork(time.Second * 4, &wg)

// 	wg.Wait()

// 	fmt.Printf("work took %v seconds\n", time.Since(start))
// }

//-----------------------------

//----------channel--------
//channel comes to the rescue when we have to return value from our function

func doWork(d time.Duration, resch chan string) {
	fmt.Println("doing work...")
	time.Sleep(d)
	fmt.Println("work is done")

	resch <- fmt.Sprintf("work %d", rand.Intn(100))
	wg.Done()
}

var wg *sync.WaitGroup

func main() {

	start := time.Now()
	resultch := make(chan string)
	wg = &sync.WaitGroup{}
	wg.Add(2)
	go doWork(time.Second*2, resultch)
	go doWork(time.Second*4, resultch)

	go func() {
		for res := range resultch {
			fmt.Println(res)
		}
		fmt.Printf("work took %v seconds\n", time.Since(start))
	}()
	wg.Wait()
close(resultch)
}
