package main

import (
	"fmt"
	"sync"
)

func main() {

	thing := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	// A channel of strings/chars, could also be [1]string if you wanted to be
	// pedantic
	letters := make(chan string)

	// Anonymous function. This will block if not using a goroutine.
	go func() {
		for _, v := range thing {
			letters <- v
		}
		// You can drain a closed channel. This
		close(letters)
	}()

	var wg sync.WaitGroup

	// 3 is arbitrary, but is used to show that 3 "workers" will print 11 values
	for i := 0; i < 3; i++ {
		// Let the waitgroup know that we are starting another job.
		wg.Add(1)
		go letterWriter(letters, &wg)
	}

	// We wait here until all of the goroutines have returned
	wg.Wait()
	fmt.Println("Waitgroup is done")

}

func letterWriter(in chan string, wg *sync.WaitGroup) {
	// defer is called after return, which, in this case, is when the channel is
	// empty. wg.Done() removes 1 from the wg struct, getting us closer to 0,
	// which is the break condition for Wait()
	defer wg.Done()

	// Iterate over the channel, for as long as there is data. If the channel
	// doesn't close, this will block while waiting for input.
	for thing := range in {
		fmt.Println("letter:", thing)
	}
}
