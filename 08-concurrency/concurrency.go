package main


import (
	"fmt"
	"time"
	"math/rand"
	"sync"  // Provides shared-memory synchronization utilities.
)


// Go is born for concurrency! Easy to use & safety ensured.
func main() {

	//
	// Goroutines.
	//

	// A *goroutine* starts a function's execution in a new thread (in the
	// same address space).
	var wg sync.WaitGroup
	go delayedMsg("hello", &wg)
	go delayedMsg("world", &wg)
	go delayedMsg("salut", &wg)

	// Main thread won't wait for goroutines to finish by default. To wait
	// for a goroutine to finish:
	//   - Waiting for a single goroutine: use a channel `chan bool`.
	//   - Waiting for multiple goroutines: use `sync.WaitGroup`.
	wg.Add(3)
	wg.Wait()


	//
	// Channels.
	//

	// Channels provide primitive support for interthread communication.
	//
	// Channels will block by default until both sides are ready, and do
	// not need explicit synchronization control.
	data := []int{3, 4, 9, 1, 6, 7}
	ch := make(chan int)
	go calcSum(data[:len(data)/2], ch)
	go calcSum(data[len(data)/2:], ch)
	s1, s2 := <-ch, <-ch 	// Receive from channel.
	fmt.Println("Sum:", s1+s2)

	// Buffered channels can queue a certain amount of messages and will
	// block only if sending to a full queue / receiving from an empty
	// queue.
	chb := make(chan bool, 2) 	// Max queue size = 2.
	chb <- true
	chb <- false
	// chb <- false 	// This won't work.
	fmt.Println(<-chb, <-chb)

	// Sender can explicitly "close" the channel to indicate the receiver.
	che := make(chan string)
	finished := make(chan bool)
	go echoAll(che, finished)
	che <- "Start!"
	che <- "I love Rust."
	che <- "You say Go."
	che <- "He swims in the C."
	che <- "We drink Java."
	close(che)
	<-finished

	// The `select` statement selects from a channel which is ready.
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: 	// The default branch is selected if no one is ready.
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}


func delayedMsg(msg string, wg *sync.WaitGroup) {
	defer wg.Done() 	// Good habit: defer a `Done()` at the beginning.

	time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
	fmt.Println(msg)
}


func calcSum(data []int, ch chan int) {
	sum := 0
	for _, v := range data {
		sum += v
	}
	ch <- sum 	// Send to channel.
}


func echoAll(ch chan string, finished chan bool) {
	msg, ok := <- ch 	// One-time test on whether channel is closed.
	fmt.Println(msg, ok)

	for msg := range ch { 	// Range on channel loops until closed.
		fmt.Println(msg)
	}

	finished <- true
}


// NOTE: Check the `sync` package for a lot more useful utilities!
