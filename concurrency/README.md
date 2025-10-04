# Concurrency: Unsynchronized Chaos

Race conditions, deadlocks, and concurrent map access without protection. Every way to misuse concurrency primitives.

## What's Inside

**SharedCounter** - Read-modify-write without mutex. Classic race condition. Counter incremented from multiple goroutines, final value is random.

**GlobalMap, GlobalSlice, GlobalValue** - Concurrent modification of non-thread-safe data structures. "fatal error: concurrent map writes" guaranteed.

**ModifyGlobalState** - 100 goroutines modifying global map, slice, and int without synchronization. Everything that can race, will race.

**ChannelLeaks** - Creates 100 goroutines that send to channels then leak. Goroutines blocked forever. Memory leak as a service.

**DeadlockMachine** - Circular dependency with channels. Both goroutines waiting for each other. Perfect deadlock.

**MutexDeadlock** - Two mutexes locked in opposite order. Classic deadlock scenario. Program hangs forever.

**SelectWithoutDefault** - Select on empty channels without default case. Blocks forever waiting for data that never arrives.

**BufferedChannelOverflow** - Sends more items than buffer size without receiver. Blocks forever on 6th item.

**ClosedChannelPanic** - Sends to closed channel from multiple goroutines. "send on closed channel" panic.

**CloseChannelTwice** - Closes a channel twice. "close of closed channel" panic.

**GetInstance** - Check-then-act singleton without synchronization. Multiple instances created. Race condition in initialization.

**GetSingleton** - Double-checked locking that doesn't work. First check is racy. Still creates multiple instances.

**WaitGroupMisuse** - Calls `Add` inside goroutine after `Wait`. Race condition. Sometimes returns before goroutines start.

**RacyMapOperations** - Concurrent map reads and writes without `sync.RWMutex` or `sync.Map`. Runtime will crash your program.

**RandomSleep** - Uses `time.Sleep` to "fix" race conditions. Doesn't fix anything. Creates heisenbugs.

## Why This Crashes Your Application

1. **Race conditions** - Non-deterministic behavior, data corruption
2. **Deadlocks** - Program hangs forever
3. **Runtime panics** - Concurrent map access, sending to closed channels
4. **Memory leaks** - Goroutines that never terminate
5. **Heisenbugs** - Bugs that disappear when you try to debug them

## What You Should Do Instead

- Use `sync.Mutex` or `sync.RWMutex` for shared state
- Use `sync/atomic` for simple counters
- Use `sync.Once` for one-time initialization
- Call `WaitGroup.Add` before starting goroutines
- Never send to closed channels, never close channels twice
- Use `sync.Map` for concurrent map access
- Use proper synchronization, never use sleep
- Run tests with `-race` flag always

## Do Not

Ship this unless you want your application to crash randomly in production.
