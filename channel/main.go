package main

import (
	"fmt"
	"sync"
	//"time"
	"strconv"
)

var wg sync.WaitGroup // 1

func main() {
	//wg := &sync.WaitGroup{}
	ch := make(chan string, 1)   //bidirectional channel

	//bidirectional channel 
	//func myFunction(ch chan string) {}

	//send only channel
	//func (ch chan <- string)  {	}

	//Receive only channel
	//func (ch <-chan string)  {	}


	wg.Add(2)
	go func(ch <-chan string) {
		for msg := range ch{
			//fmt.Println(<-ch)
			fmt.Println(msg)
		//	go routine(i) // *
		}
    	//wg.Wait() // 4

		wg.Done()
	}(ch)
	go func(ch chan<- string) {
		//close(ch);
		for i :=1; i<10; i++ {
			ch <- "Hello Jack " +  strconv.Itoa(i)
		}
		close(ch)
		//time.Sleep(1 * time.Millisecond)
		//fmt.Println(<-ch)
		wg.Done()
	}(ch)

	wg.Wait()
}

func routine(i int) {
    defer wg.Done() // 3
    fmt.Println("routine finished", strconv.Itoa(i))
}