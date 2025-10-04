# unmaintainable-go

A comprehensive guide to writing Go code that will make your colleagues quit. Every anti-pattern you should never use, demonstrated with precision.

## What You'll Find Here

This repository is a deliberate exercise in writing the worst possible Go code while remaining syntactically valid. If you're looking for best practices, you're in the wrong place. If you want to see exactly what not to do, welcome.

### The Collection

**goroutines/** - Recursive spawning without bounds. Race conditions as features. Deadlocks on demand. Memory leaks guaranteed. No synchronization, no context, no cleanup. Watch your server die.

**error-handling/** - The blank identifier as error handling strategy. Panics instead of returns. Generic error messages like "oops". Silent failures that corrupt data. Global error state shared across goroutines.

**interfaces/** - `interface{}` for everything. Type safety is optional. God interfaces with 18+ methods. Type assertions without checks. Runtime panics instead of compile-time safety.

**concurrency/** - Unsynchronized shared state. Concurrent map access without locks. Double-checked locking that doesn't work. WaitGroups used incorrectly. Sleep as synchronization.

**naming/** - Single letter variables everywhere. Functions that lie about what they do. Permanent typos in the API. Variables shadowing built-in types. Hungarian notation in a language with type inference.

**dependency-injection/** - Global variables as dependency injection. Service locator pattern. Circular dependencies. Hard-coded secrets in source code. init() with side effects. Hidden dependencies everywhere.

**memory/** - Memory leaks everywhere. Unbounded caches, circular references, timer leaks. Global slices that grow forever. Defers in loops. Watch your process get OOM killed.

**performance/** - Quadratic algorithms, excessive allocations, operations in hot paths. String concatenation in loops. Regex compiled every time. Reflection everywhere. Death by a thousand cuts.

**testing/** - Code designed to be untestable. Global state, hard-coded dependencies, side effects. Non-deterministic behavior. Time-based logic. Makes writing tests impossible.

## The Point

Every pattern here violates Go idioms deliberately. This is educational material showing you what happens when you ignore best practices. The code compiles. The code runs. The code will destroy your production environment.

## How to Use This

Study what not to do. Run the examples with `-race` to see the carnage. Watch goroutines leak. See race conditions in action. Experience deadlocks. Then do the opposite in your real code.

```bash
# See the disasters in action
go run examples.go

# See race conditions exposed
go run -race examples.go

# Build everything (it compiles, unfortunately)
go build ./...

# Tests (there are none, testing would imply care)
go test ./...
```

## The Anti-Patterns

- Ignoring errors with `_`
- Panic instead of error returns
- No goroutine cleanup or cancellation
- Race conditions on shared state
- Deadlocks with circular channel dependencies
- `interface{}` everywhere
- Global mutable state
- Hard-coded configuration
- Cryptic naming conventions
- Type assertions without checks
- Concurrent map access without synchronization
- Sleep as a synchronization primitive

## What You Should Actually Do

The opposite of everything in this repository. Use proper error handling, synchronize goroutines, protect shared state, use interfaces sparingly, inject dependencies properly, and name things descriptively.

## Fair Warning

Do not use any of these patterns in production code. Do not use them in personal projects. Do not use them in code samples. This is a museum of bad decisions. Look, learn what not to do, then close the tab.

## Contributing

Want to make this codebase even worse? We welcome contributions! Check out [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on adding new anti-patterns, writing terrible code, and helping others learn what not to do.

## License

MIT - Because even terrible code deserves to be free.
