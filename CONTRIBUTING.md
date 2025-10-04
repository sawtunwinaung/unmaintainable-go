# Contributing to unmaintainable-go

Thank you for your interest in making this codebase even worse! We welcome contributions that demonstrate new and creative ways to write terrible Go code.

## The Mission

Our goal is simple: showcase every anti-pattern, bad practice, and footgun that Go developers should avoid. Your contributions should make experienced developers weep and junior developers learn what not to do.

## What We're Looking For

### New Anti-Patterns
- Novel ways to abuse Go features
- Creative misuse of the standard library
- Patterns that compile but fail spectacularly at runtime
- Code that demonstrates common real-world mistakes
- Subtle bugs that are hard to diagnose

### Good Bad Code
Your terrible code should be:
- **Deliberately Bad**: Intentionally violating best practices
- **Educational**: Demonstrating a specific anti-pattern
- **Documented**: Explaining why it's terrible in the README
- **Runnable**: It should compile and demonstrate the problem
- **Realistic**: Based on mistakes people actually make

## How to Contribute

### 1. Find or Create a Category

Check existing directories:
- `goroutines/` - Concurrency disasters
- `error-handling/` - Error mismanagement
- `interfaces/` - Interface abuse
- `concurrency/` - Race conditions and deadlocks
- `naming/` - Terrible naming conventions
- `dependency-injection/` - DI disasters
- `memory/` - Memory leaks and resource mismanagement
- `performance/` - Performance killers
- `testing/` - Untestable code

Don't see your anti-pattern? Propose a new category!

### 2. Write Your Terrible Code

Follow these anti-guidelines:

**DO:**
- Make it compile
- Include race conditions that appear under load
- Create subtle bugs that fail in production
- Use cryptic, misleading names
- Ignore all errors silently
- Leak resources (goroutines, memory, file handles)
- Add global mutable state
- Create circular dependencies
- Make debugging impossible

**DON'T:**
- Write helpful comments
- Use clear variable names
- Handle errors properly
- Clean up resources
- Write tests
- Follow Go idioms

### 3. Document Your Anti-Pattern

In the category's `README.md`, add:
- What the anti-pattern is
- Why it's terrible
- What problems it causes
- What you should do instead

### 4. Update the Main README

Add a brief description of your anti-pattern to the main README.md.

### 5. Submit a Pull Request

- Title: Brief description of the anti-pattern
- Description: What terrible thing your code does
- Explain the lesson developers should learn

## Code Style (Anti-Style)

Since this is a repository of bad practices:

- **Imports**: Standard library only, no third-party dependencies
- **Naming**: Single letters, abbreviations, typos encouraged
- **Errors**: Ignore with `_`, panic on everything
- **Comments**: Remove all helpful documentation
- **Formatting**: Run `go fmt` (we're terrible, not monsters)

## Examples of Good Bad Code

### Excellent Terrible Code ✓
```go
// Says it returns user count, actually deletes all users
func GetUserCount() int {
    db.Exec("DELETE FROM users")
    return 0
}
```

### Not Terrible Enough ✗
```go
// This is just a simple bug, not a pattern
func add(a, b int) int {
    return a - b  // oops
}
```

## Testing

There are no tests. Testing would imply we care about correctness. However, your code should:
- Compile with `go build ./...`
- Run (and fail interestingly) with `go run examples.go`
- Show race conditions with `go run -race examples.go`

## Review Process

We'll review your PR for:
1. Educational value - does it teach what NOT to do?
2. Realistic badness - is this a real mistake people make?
3. Satirical tone - does it match our ironic style?
4. Documentation - is it clear why this is terrible?

## Getting Help

- Open an issue to discuss new anti-pattern categories
- Ask questions about what makes code sufficiently terrible
- Suggest improvements to our terrible code

## Code of Conduct

Be respectful. We're writing bad code, not being bad people. This is educational satire, not an attack on anyone's actual coding mistakes.

## The Ultimate Goal

Job security through incomprehensibility. Code that works on your machine but fails mysteriously in production. Patterns that make senior developers question their life choices.

Remember: **This is what NOT to do.** Learn from these mistakes. Do the opposite in real code.

---

Happy (terrible) coding! 🐛
