# Performance: Death by a Thousand Allocations

Inefficient algorithms, excessive allocations, and operations in hot paths. Watch your CPU burn.

## What's Inside

**StringConcatInLoop** - String concatenation with `+` in a loop. Creates new string on every iteration. O(n²) time and allocations.

**CompileRegexEveryTime** - Compiles regex on every call. Regex compilation is expensive. Should compile once and reuse.

**MD5InLoop** - Hashes in a loop without reusing hash object. Allocates hasher every iteration.

**ReflectInHotPath** - Uses reflection in hot path. Reflection is slow. Type assertions are faster.

**JSONMarshalInLoop** - JSON encoding in a loop. Each marshal allocates buffers. Should reuse encoder.

**NoPreallocation** - Appends to slice without preallocating capacity. Slice grows and reallocates multiple times.

**SortInLoop** - Sorts slices inside loop without checking if already sorted. Redundant work.

**DeferInTightLoop** - Uses defer in tight loop. Defer has overhead. Stack of deferred calls grows.

**ConvertIntToString** - Uses `fmt.Sprintf` for int-to-string. `strconv.Itoa` is faster. Sprintf parses format string every time.

**TimeNowInLoop** - Calls `time.Now()` in tight loop. System call every iteration. Cache time outside loop.

**InterfaceConversion** - Converts concrete types to interfaces. Boxing allocates. Keep concrete types when possible.

**DeepCopy** - Deep copies map via JSON marshal/unmarshal. Extremely slow. Manual copy is faster.

**CheckContains** - Linear search on every call. O(n) lookups. Use map for O(1).

**RemoveFromSlice** - Removes from slice by creating new slice. Allocates and copies. In-place removal is faster.

**ReverseString** - Reverses string by concatenating one character at a time. Quadratic allocations.

**FilterSlice** - Filters without preallocating result slice. Multiple reallocations as slice grows.

**ParseIntsSlowly** - Parses ints without preallocating result. Each append may reallocate.

**Contains** - Creates map on every call just to check membership. Map allocation overhead.

**BufferedChannelVsUnbuffered** - Uses unbuffered channel. Each send blocks until receive. Buffered channels reduce context switches.

**GetOrCompute** - Cache without mutex. Race conditions. Concurrent map access.

**NestedLoops** - O(n³) algorithm. Cubic time complexity. Could be optimized.

## Why This Kills Performance

1. **Excessive allocations** - GC pressure, memory bandwidth waste
2. **Wrong data structures** - Linear search instead of hash lookup
3. **Redundant work** - Recompiling regex, sorting sorted data
4. **Reflection overhead** - Runtime type inspection instead of compile-time
5. **Blocking operations** - Unbuffered channels, synchronous I/O
6. **String allocations** - Immutable strings copied repeatedly

## What You Should Do Instead

- Use `strings.Builder` for string concatenation
- Compile regex once, use many times
- Preallocate slices when size is known
- Use `strconv` instead of `fmt` for conversions
- Cache expensive computations
- Use maps for O(1) lookups instead of O(n) linear search
- Avoid reflection in hot paths
- Use buffered channels to reduce blocking
- Profile with pprof before optimizing
- Benchmark with testing.B

## Do Not

Ship any of this if you care about performance. Your users will notice.
