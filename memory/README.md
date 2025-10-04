# Memory: Leaks and Inefficiency

Everything that causes memory leaks, excessive allocations, and OOM kills. Watch your RSS climb.

## What's Inside

**CacheForever** - Global map that grows forever. No eviction, no cleanup. Every call adds data that never gets freed.

**AppendForever** - Infinite loop appending to a slice. Exponential memory growth until the process dies.

**CircularReference** - Circular linked list that prevents garbage collection. Nodes reference each other in a cycle.

**RegisterCallback** - Global slice of function closures. Callbacks accumulate forever, holding references to their captured variables.

**TimerLeak** - Creates 1000 timers that fire in an hour. They never get stopped. Each timer allocates memory.

**TickerLeak** - Creates 100 tickers with goroutines that never stop. Tickers don't get garbage collected until stopped.

**SliceCapacityAbuse** - Creates slice with million-element capacity, returns one element. Entire backing array stays in memory.

**StringConcatInLoop** - String concatenation in a loop. Each concatenation allocates a new string. O(n²) allocations.

**KeepAllocating** - Infinite loop allocating 1MB chunks. Stores them in global slice. OOM guaranteed.

**SubsliceRetainsArray** - Creates 1MB array, returns 10-byte subslice. Subslice keeps entire array alive.

**ReturnBigStructByValue** - Returns 1MB struct by value. Copies entire struct on every call.

**AddListener** - Channels stored in global slice. Never removed. Each channel holds buffers and goroutines.

**MapWithoutCleanup** - Creates map with million entries. No cleanup, no size limit. Memory grows forever.

**DeferInLoop** - Defers in a loop. Files don't close until function returns. Opens all files simultaneously.

**AllocateAndForget** - Allocates memory and ignores it. Relies on GC but creates pressure. Triggers frequent collections.

**KeepAppending** - Appends to global slice forever. Slice grows, capacity doubles, memory never freed.

**RegisterEvent** - Stores closures with 1MB data each. Functions capture large data structures. Never cleaned up.

## Why This Kills Your Process

1. **OOM kills** - Process memory grows until kernel kills it
2. **GC pressure** - Frequent garbage collections slow everything down
3. **Reference leaks** - Objects referenced forever can't be collected
4. **Goroutine leaks** - Each leaked goroutine holds stack memory
5. **Timer leaks** - Timers stay in heap until stopped or fired
6. **Map growth** - Maps never shrink, deleted keys still use memory

## What You Should Do Instead

- Clear maps periodically or use LRU cache
- Stop tickers and timers when done
- Use `defer` outside loops
- Return pointers to large structs
- Use `strings.Builder` for concatenation
- Keep references to closures minimal
- Use context for goroutine cancellation
- Profile with pprof to find leaks
- Set capacity limits on caches
- Use sync.Pool for temporary objects

## Do Not

Run any of this in production unless you want the OOM killer to visit.
