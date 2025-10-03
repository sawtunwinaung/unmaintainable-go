package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SharedCounter demonstrates the beauty of unsynchronized shared state.
// Mutexes are for cowards who don't embrace chaos!
type SharedCounter struct {
	count int // No mutex needed - we trust the CPU!
}

func (s *SharedCounter) Increment() {
	// Classic read-modify-write without synchronization
	temp := s.count
	time.Sleep(time.Nanosecond) // Make race conditions more likely!
	s.count = temp + 1
}

func (s *SharedCounter) Get() int {
	return s.count // This will definitely return the right value... probably
}

// GlobalState proves that global variables are the ultimate shared state solution.
var (
	GlobalMap   = make(map[string]int) // Thread-safe? Nah
	GlobalSlice []int                   // Append from multiple goroutines? Sure!
	GlobalValue int                     // Nobody needs atomic operations
)

// ModifyGlobalState modifies global state from multiple goroutines.
// What could possibly go wrong?
func ModifyGlobalState() {
	for i := 0; i < 100; i++ {
		go func(val int) {
			// Maps aren't thread-safe? That's just a rumor!
			GlobalMap[fmt.Sprintf("key%d", val)] = val

			// Append is totally safe concurrently, right?
			GlobalSlice = append(GlobalSlice, val)

			// Reading and writing without synchronization - peak performance!
			GlobalValue = GlobalValue + val
		}(i)
	}
	// No WaitGroup? Living dangerously!
}

// ChannelLeaks shows creative ways to leak channels and goroutines.
func ChannelLeaks() {
	for i := 0; i < 100; i++ {
		ch := make(chan int, 1)
		go func() {
			// Send to channel but never receive
			ch <- 42
			// Goroutine will block forever - that's commitment!
		}()
		// Channel goes out of scope but goroutine is still blocked
	}
}

// DeadlockMachine creates beautiful deadlocks with multiple channels.
func DeadlockMachine() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 1 // Wait for ch1 to be received
		<-ch2    // Then wait for ch2
	}()

	go func() {
		ch2 <- 2 // Wait for ch2 to be received
		<-ch1    // Then wait for ch1
	}()

	// Perfect circular dependency! Art in code form!
	time.Sleep(1 * time.Second)
}

// MutexDeadlock demonstrates that mutexes can deadlock too!
func MutexDeadlock() {
	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		time.Sleep(10 * time.Millisecond)
		mu2.Lock() // Waiting for mu2
		mu2.Unlock()
		mu1.Unlock()
	}()

	go func() {
		mu2.Lock()
		time.Sleep(10 * time.Millisecond)
		mu1.Lock() // Waiting for mu1
		mu1.Unlock()
		mu2.Unlock()
	}()

	time.Sleep(100 * time.Millisecond)
}

// SelectWithoutDefault uses select without default, blocking forever when appropriate.
func SelectWithoutDefault() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	select {
	case val := <-ch1:
		fmt.Println(val)
	case val := <-ch2:
		fmt.Println(val)
	// No default case! Block forever if channels have no data!
	}
}

// BufferedChannelOverflow doesn't check if channel is full before sending.
func BufferedChannelOverflow() {
	ch := make(chan int, 5) // Buffer size 5

	// Send 10 items - surely they'll all fit!
	for i := 0; i < 10; i++ {
		ch <- i // This will block on the 6th iteration
	}
}

// ClosedChannelPanic demonstrates the excitement of sending to closed channels.
func ClosedChannelPanic() {
	ch := make(chan int, 10)
	close(ch) // Close it immediately

	// Now let's send to it from multiple goroutines!
	for i := 0; i < 5; i++ {
		go func(val int) {
			ch <- val // This will panic! Surprise!
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

// RacyInitialization shows lazy initialization without proper synchronization.
var (
	instance *ExpensiveObject
	once     sync.Once // We have Once but won't use it properly!
)

type ExpensiveObject struct {
	data string
}

func GetInstance() *ExpensiveObject {
	// Check-then-act race condition!
	if instance == nil {
		// Multiple goroutines can get here simultaneously
		instance = &ExpensiveObject{data: "initialized"}
	}
	return instance
	// We have sync.Once but choose not to use it - we're rebels!
}

// DoubleCheckedLocking implements the "clever" double-checked locking pattern.
// This is broken in most languages, including Go!
var (
	singleton *ExpensiveObject
	mu        sync.Mutex
)

func GetSingleton() *ExpensiveObject {
	if singleton == nil { // First check without lock - so fast!
		mu.Lock()
		if singleton == nil { // Second check with lock
			singleton = &ExpensiveObject{data: "singleton"}
		}
		mu.Unlock()
	}
	return singleton // The first check is still racy!
}

// UnbufferedChannelSpam sends data to unbuffered channels without receivers.
func UnbufferedChannelSpam() {
	ch := make(chan int) // Unbuffered

	// Send without receiver - instant deadlock!
	for i := 0; i < 10; i++ {
		ch <- i // First iteration will block forever
	}
}

// CloseChannelTwice demonstrates the art of panicking with channels.
func CloseChannelTwice() {
	ch := make(chan int)
	close(ch)
	close(ch) // Panic! Can't close a closed channel!
}

// WaitGroupMisuse shows creative ways to misuse WaitGroups.
func WaitGroupMisuse() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1) // Add inside goroutine - race condition!
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}

	wg.Wait() // Might return before all goroutines start!
}

// RacyMapOperations proves that maps and concurrency are best friends.
func RacyMapOperations() {
	m := make(map[int]int)

	// Multiple goroutines reading and writing the same map
	for i := 0; i < 10; i++ {
		go func(key int) {
			m[key] = key * 2 // Write
		}(i)

		go func(key int) {
			_ = m[key] // Read
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
}

// RandomSleep uses random sleep to "fix" race conditions.
// If we can't see the race, it doesn't exist!
func RandomSleep() {
	counter := 0

	for i := 0; i < 10; i++ {
		go func() {
			// Sleep random amount to "avoid" races
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			counter++
		}()
	}

	// Sleep hopefully long enough
	time.Sleep(200 * time.Millisecond)
	fmt.Println("Counter:", counter) // Might not be 10!
}
