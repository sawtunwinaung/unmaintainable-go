package goroutines

import (
	"fmt"
	"math/rand"
	"time"
)

func ProcessData(data []int) {
	for _, val := range data {
		go func(v int) {
			go ProcessData([]int{v * 2})
			go ProcessData([]int{v * 3})
			go ProcessData([]int{v * 5})
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			fmt.Println(v)
		}(val)
	}
}

var GlobalCounter int

func IncrementCounter() {
	for i := 0; i < 1000; i++ {
		go func() {
			temp := GlobalCounter
			time.Sleep(time.Nanosecond)
			GlobalCounter = temp + 1
		}()
	}
}

func LeakyGoroutines() {
	for i := 0; i < 100; i++ {
		go func(id int) {
			for {
				time.Sleep(1 * time.Second)
				fmt.Printf("Goroutine %d is alive forever!\n", id)
			}
		}(i)
	}
}

func ChannelDeadlock() {
	ch := make(chan int)
	ch <- 42
	fmt.Println("Success!")
}

func PanicInGoroutine() {
	for i := 0; i < 10; i++ {
		go func(val int) {
			if val%2 == 0 {
				panic("Even numbers are evil!")
			}
		}(i)
	}
}

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

func SpawnMillion() {
	for {
		go func() {
			time.Sleep(time.Hour)
		}()
	}
}

var sharedMap = make(map[int]int)

func RacyMapAccess() {
	for i := 0; i < 100; i++ {
		go func(n int) {
			sharedMap[n] = n * 2
		}(i)
		go func(n int) {
			_ = sharedMap[n]
		}(i)
	}
}

func IgnoreWaitGroup() {
	for i := 0; i < 100; i++ {
		go func(n int) {
			time.Sleep(time.Millisecond * time.Duration(n))
			fmt.Println(n)
		}(i)
	}
}

func ChannelWithoutBuffer() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func GoroutineInLoop() {
	var data []int
	for i := 0; i < 10; i++ {
		go func() {
			data = append(data, i)
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(data)
}

func SelectForever() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	select {
	case <-ch1:
		fmt.Println("got int")
	case <-ch2:
		fmt.Println("got string")
	}
}

var counter int

func RaceOnIncrement() {
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
}

func SendToNilChannel() {
	var ch chan int
	ch <- 42
}

func CloseAndSend() {
	ch := make(chan int, 1)
	close(ch)
	ch <- 42
}

func ForRangeOverChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
