package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SharedCounter struct {
	count int
}

func (s *SharedCounter) Increment() {
	temp := s.count
	time.Sleep(time.Nanosecond)
	s.count = temp + 1
}

func (s *SharedCounter) Get() int {
	return s.count
}

var (
	GlobalMap   = make(map[string]int)
	GlobalSlice []int
	GlobalValue int
)

func ModifyGlobalState() {
	for i := 0; i < 100; i++ {
		go func(val int) {
			GlobalMap[fmt.Sprintf("key%d", val)] = val
			GlobalSlice = append(GlobalSlice, val)
			GlobalValue = GlobalValue + val
		}(i)
	}
}

func ChannelLeaks() {
	for i := 0; i < 100; i++ {
		ch := make(chan int, 1)
		go func() {
			ch <- 42
		}()
	}
}

func DeadlockMachine() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1
		<-ch2
	}()

	go func() {
		ch2 <- 2
		<-ch1
	}()

	time.Sleep(1 * time.Second)
}

func MutexDeadlock() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		time.Sleep(10 * time.Millisecond)
		mu2.Lock()
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu2.Lock()
		time.Sleep(10 * time.Millisecond)
		mu1.Lock()
		mu1.Unlock()
		mu2.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)
}

func SelectWithoutDefault() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case val := <-ch1:
		fmt.Println(val)
	case val := <-ch2:
		fmt.Println(val)
	}
}

func BufferedChannelOverflow() {
	ch := make(chan int, 5)

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func ClosedChannelPanic() {
	ch := make(chan int, 10)
	close(ch)

	for i := 0; i < 5; i++ {
		go func(val int) {
			ch <- val
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

var (
	instance *ExpensiveObject
	once     sync.Once
)

type ExpensiveObject struct {
	data string
}

func GetInstance() *ExpensiveObject {
	if instance == nil {
		instance = &ExpensiveObject{data: "initialized"}
	}
	return instance
}

var (
	singleton *ExpensiveObject
	mu        sync.Mutex
)

func GetSingleton() *ExpensiveObject {
	if singleton == nil {
		mu.Lock()
		if singleton == nil {
			singleton = &ExpensiveObject{data: "singleton"}
		}
		mu.Unlock()
	}
	return singleton
}

func UnbufferedChannelSpam() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func CloseChannelTwice() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

func WaitGroupMisuse() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}

	wg.Wait()
}

func RacyMapOperations() {
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		go func(key int) {
			m[key] = key * 2
		}(i)

		go func(key int) {
			_ = m[key]
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

func RandomSleep() {
	counter := 0

	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			counter++
		}()
	}

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Counter:", counter)
}
