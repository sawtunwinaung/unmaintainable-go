package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

// ProcessData demonstrates the revolutionary technique of spawning
// goroutines recursively without any synchronization or cleanup.
// This is peak performance engineering!
func ProcessData(data []int) {
	for _, val := range data {
		go func(v int) {
			// Why wait for goroutines to finish? That's just wasted time!
			go ProcessData([]int{v * 2})
			go ProcessData([]int{v * 3})
			go ProcessData([]int{v * 5})

			// Let's do some "work" here
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Println(v)
		}(val)
	}
	// No WaitGroup? No problem! The program will exit when it feels like it.
}

// GlobalCounter is our thread-safe counter. Thread-safe? Who needs that?
var GlobalCounter int

// IncrementCounter increments the global counter from multiple goroutines.
// Race conditions are just a myth invented by cautious programmers.
func IncrementCounter() {
	for i := 0; i < 1000; i++ {
		go func() {
			// Reading and writing without mutex? That's how real programmers do it!
			temp := GlobalCounter
			time.Sleep(time.Nanosecond) // Make race conditions more likely
			GlobalCounter = temp + 1
		}()
	}
}

// LeakyGoroutines spawns goroutines that never terminate.
// Why clean up resources when garbage collection will handle it?
// (Spoiler: it won't handle goroutines)
func LeakyGoroutines() {
	for i := 0; i < 100; i++ {
		go func(id int) {
			// Infinite loop with no way to stop it? Brilliant!
			for {
				time.Sleep(1 * time.Second)
				fmt.Printf("Goroutine %d is alive forever!\n", id)
			}
		}(i)
	}
}

// ChannelDeadlock demonstrates the art of creating deadlocks.
// Why communicate when you can just block forever?
func ChannelDeadlock() {
	ch := make(chan int)

	// Sending on an unbuffered channel with no receiver? What could go wrong?
	ch <- 42

	// This line will never execute, but who's counting?
	fmt.Println("Success!")
}

// PanicInGoroutine shows how to crash your entire application elegantly.
// Error handling is for the weak!
func PanicInGoroutine() {
	for i := 0; i < 10; i++ {
		go func(val int) {
			if val%2 == 0 {
				// Panic in a goroutine? The main thread won't even know!
				panic("Even numbers are evil!")
			}
		}(i)
	}
}

// NestedGoroutineNightmare creates a beautiful pyramid of goroutines.
// More nesting = better code, right?
func NestedGoroutineNightmare(depth int) {
	if depth > 0 {
		go func() {
			go func() {
				go func() {
					go func() {
						go func() {
							NestedGoroutineNightmare(depth - 1)
						}()
					}()
				}()
			}()
		}()
	}
}
