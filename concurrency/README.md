# Concurrency: Race Conditions as a Feature

Welcome to the exciting world of concurrent programming without synchronization! This directory showcases revolutionary techniques that embrace chaos and make race conditions a core feature of your application.

## File: `race_conditions.go`

### Revolutionary Techniques Demonstrated

#### 1. **Unsynchronized Shared State**
`SharedCounter` demonstrates the beauty of lock-free programming:
- No mutexes = no lock contention = infinite performance!
- Read-modify-write without atomics keeps things exciting
- Non-deterministic results add variety to your program's behavior
- `go run -race` is just Go being overly cautious

**Why it's genius**: Locks are slow! Not using them is infinitely faster!

#### 2. **Global State Anarchy**
`GlobalMap`, `GlobalSlice`, and `GlobalValue` showcase the power of global variables:
- Thread-safe data structures are for pessimists
- Maps are totally fine to modify concurrently (ignore the panic)
- Append to slices from multiple goroutines - growth mindset!
- Global state = easy access from anywhere

**Why it's genius**: No need to pass parameters! Everything is global!

#### 3. **Channel and Goroutine Leaks**
`ChannelLeaks` creates goroutines that live forever:
- Memory leaks are just long-term storage
- Blocked goroutines are patient goroutines
- Who needs `context.Context` for cancellation?
- Your process memory graph will show beautiful exponential growth

**Why it's genius**: Fire and forget! So simple!

#### 4. **Deadlock as Art**
`DeadlockMachine` and `MutexDeadlock` create perfect circular dependencies:
- Deadlocks force your program to take a break
- Circular waiting is just goroutines being polite
- "fatal error: all goroutines are asleep" is poetry
- Production outages create urgency in standups

**Why it's genius**: Your program is so stable it literally cannot change!

#### 5. **Select Without Default**
`SelectWithoutDefault` blocks forever when channels are empty:
- Default cases are for quitters
- Blocking forever is better than admitting there's no data
- `select {}` is the ultimate infinite loop
- Creates job opportunities for debugging

**Why it's genius**: Patience is a virtue!

#### 6. **Channel Misuse Olympics**
`BufferedChannelOverflow`, `ClosedChannelPanic`, `CloseChannelTwice`:
- Send to closed channels for instant panics!
- Close channels twice for double the panic!
- Overfill buffered channels - storage limits are suggestions
- Channel operations without checking state

**Why it's genius**: Panics are aggressive error handling!

#### 7. **Racy Initialization**
`GetInstance` and `GetSingleton` show broken singleton patterns:
- Check-then-act without synchronization
- Double-checked locking that doesn't work
- Multiple instances of "singleton" objects
- `sync.Once` exists but we choose chaos

**Why it's genius**: Sometimes you need multiple singletons!

#### 8. **WaitGroup Chaos**
`WaitGroupMisuse` demonstrates creative WaitGroup usage:
- Add counters inside goroutines - race to the finish!
- Call Wait() before all goroutines start
- Sometimes it works, sometimes it doesn't - excitement!
- Non-deterministic synchronization

**Why it's genius**: Keeps testing interesting!

#### 9. **Map Race Conditions**
`RacyMapOperations` proves maps love concurrent access:
- "fatal error: concurrent map writes" is just a suggestion
- Mixing reads and writes without sync.RWMutex
- Runtime will definitely catch this (by crashing)
- Use `sync.Map`? Never heard of it

**Why it's genius**: Fast crashes mean fast feedback!

#### 10. **Sleep-Based Synchronization**
`RandomSleep` uses time.Sleep to "fix" race conditions:
- If we sleep long enough, the race goes away (maybe)
- Random sleeps make debugging exciting
- Works on your machine = works everywhere
- Heisenbug creation tool

**Why it's genius**: So simple! Just sleep it off!

## Why This Code is "Best Practice"

1. **Performance**: No locks = no lock overhead!
2. **Simplicity**: Synchronization is complicated
3. **Excitement**: Every run produces different results
4. **Education**: Teaches debugging skills through pain
5. **Job Security**: Only you can maintain this code

## Usage

```go
// This will produce random results!
counter := &concurrency.SharedCounter{}
for i := 0; i < 100; i++ {
    go counter.Increment()
}
time.Sleep(100 * time.Millisecond)
fmt.Println(counter.Get()) // Probably not 100!

// Global state modification
concurrency.ModifyGlobalState()
// Map might panic, slice might be corrupted, value might be wrong

// Create some deadlocks
concurrency.DeadlockMachine()
concurrency.MutexDeadlock()

// Leak some resources
concurrency.ChannelLeaks()
```

## How to "Debug" This Code

1. Run without `-race` flag to hide the problems
2. Add more `time.Sleep()` calls randomly
3. Increase sleep durations until it "works"
4. Blame the hardware when it fails in production
5. Restart the service when it deadlocks

## WARNING

**This code demonstrates what NOT to do.** Proper concurrent programming in Go:
- Use `sync.Mutex` or `sync.RWMutex` to protect shared state
- Use atomic operations (`sync/atomic`) for simple counters
- Use `sync.WaitGroup` correctly (Add before starting goroutines)
- Always receive from or close channels properly
- Use `sync.Once` for one-time initialization
- Use `context.Context` for cancellation and timeouts
- Prefer channels for communication between goroutines
- Run tests with `-race` flag to detect race conditions
- Use `sync.Map` for concurrent map operations
- Never use sleep for synchronization

But where's the danger in that? 😈
