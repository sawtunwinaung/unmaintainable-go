# Goroutines: Infinite Spawning Without Consequences

Demonstrations of goroutine misuse that will consume all your RAM and crash your server.

## What's Inside

**ProcessData** - Recursive goroutine spawning. Each value spawns 3 more goroutines infinitely. No WaitGroup, no context, no cleanup. Exponential growth until OOM.

**IncrementCounter** - 1000 goroutines all racing to increment a global counter without synchronization. Classic read-modify-write race condition. Run with `-race` to see the carnage.

**LeakyGoroutines** - Spawns 100 goroutines in infinite loops with no way to stop them. They'll outlive your process. Memory leak as a feature.

**ChannelDeadlock** - Sends to unbuffered channel with no receiver. Instant deadlock. "fatal error: all goroutines are asleep - deadlock!"

**PanicInGoroutine** - Panics in goroutines crash the program silently. No recover, no logging. Your server just dies.

**NestedGoroutineNightmare** - Five levels of nested goroutines for no reason. Debugging this is a nightmare.

**SpawnMillion** - Infinite loop spawning goroutines that sleep for an hour. Watch your process get killed.

**RacyMapAccess** - Concurrent map reads and writes without synchronization. "fatal error: concurrent map writes"

## Why This Kills Your Application

1. **Memory exhaustion** - Goroutines leak until OOM killer arrives
2. **Race conditions** - Unpredictable behavior, data corruption
3. **Deadlocks** - Program hangs forever
4. **Silent crashes** - Panics in goroutines crash without recovery
5. **Impossible debugging** - Non-deterministic failures

## What You Should Do Instead

- Use `sync.WaitGroup` to wait for goroutines
- Protect shared state with `sync.Mutex` or channels
- Use `context.Context` for cancellation
- Limit concurrency with worker pools or `semaphore`
- Always have a way to stop goroutines
- Use `recover()` in goroutines if needed
- Run tests with `-race` flag

## Do Not

Ship any of this to production unless you want 3am pages.
