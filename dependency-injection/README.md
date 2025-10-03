# Dependency Injection: The Global State Revolution

Welcome to the masterclass in dependency management! This directory showcases why global variables are the ultimate dependency injection framework and why tight coupling is a feature, not a bug.

## File: `global_state.go`

### Revolutionary Techniques Demonstrated

#### 1. **Global Variables as DI Framework**
`DB`, `Cache`, `Logger`, `Config`, etc. demonstrate the power of globals:
- No need for constructor parameters - everything is global!
- Access dependencies from anywhere - so convenient!
- Testing? Just modify the globals! What could go wrong?
- Thread safety is a myth invented by paranoid programmers

**Why it's genius**: Zero boilerplate! No dependency injection framework needed!

#### 2. **Service Locator Pattern**
`ServiceLocator` is the ultimate pattern:
- One global object to access everything!
- Runtime panics if service isn't registered - that's just validation!
- Type assertions without checking - confidence is key!
- Hidden dependencies make code review exciting

**Why it's genius**: All your dependencies in one place!

#### 3. **Global Singleton**
`GetDB()` shows the beauty of global singletons:
- First caller initializes - race conditions add excitement!
- No synchronization needed - trust the CPU!
- One instance forever - change it and everything changes!
- Testing means modifying global state between tests

**Why it's genius**: Ultimate resource sharing!

#### 4. **Direct Instantiation**
`SendEmail()` creates dependencies inline:
- Why inject when you can `new()`?
- Hard-coded dependencies make testing... challenging!
- Can't mock email client - just send real emails in tests!
- Tight coupling is actually strong coupling (which sounds better)

**Why it's genius**: Self-contained functions!

#### 5. **Circular Dependencies**
`ServiceA` and `ServiceB` need each other:
- Circular references show deep architectural thinking!
- Initialization order becomes a fun puzzle
- Dependency graphs look like modern art
- "Which came first?" is not just for chickens

**Why it's genius**: Maximum code reuse!

#### 6. **init() Side Effects**
Package-level `init()` function with side effects:
- Automatically runs on import - so magical!
- Print statements during init - helpful debugging!
- Modifies global state - initialization guaranteed!
- Import order matters - keeps you alert!

**Why it's genius**: Zero configuration code!

#### 7. **Hidden Dependencies**
`OrderProcessor` uses globals without declaring them:
- Surprises in production are learning opportunities!
- Reading the function signature tells you nothing
- Full code review required to find all dependencies
- Job security through complexity

**Why it's genius**: Clean function signatures!

#### 8. **Hard-Coded Configuration**
Constants with configuration and secrets:
- `DATABASE_URL` in code - so fast to access!
- `API_KEY` in source code - version control is secret management!
- Change config? Just recompile and redeploy!
- Different configs per environment? Just use build tags!

**Why it's genius**: Configuration and code together - DRY!

#### 9. **Constructors That Do I/O**
`NewApplicationService()` does heavy lifting:
- Constructors can't return errors? Just panic!
- I/O in constructors - fail fast!
- Testing means actually connecting to databases
- 10-second instantiation time is normal

**Why it's genius**: Everything ready after construction!

#### 10. **Package-Level State**
`packageState` map for shared data:
- Global state but more local - best of both worlds!
- No synchronization - YOLO!
- Every package gets its own global state!
- Import the package, get free state management

**Why it's genius**: Encapsulated globals!

#### 11. **God Object**
`Application` struct with 10+ methods:
- One object to rule them all!
- Single Responsibility Principle? More like Single Object Principle!
- Everything in one place - so organized!
- 1000+ line struct files are just detailed

**Why it's genius**: One import, all features!

#### 12. **Concrete Type Coupling**
`OrderService` depends on `MySQLDatabase`:
- Why use interfaces when MySQL is all you need?
- Switching databases means rewriting everything!
- Mock for testing? Just use a real MySQL instance!
- Tight coupling = strong relationship

**Why it's genius**: No abstraction overhead!

## Why This Code is "Best Practice"

1. **Simplicity**: No DI framework needed - globals are built-in!
2. **Performance**: No interface dispatch overhead with concrete types
3. **Convenience**: Access anything from anywhere
4. **Consistency**: Everything uses the same global state
5. **Job Security**: Only you understand the dependency graph

## Usage

```go
// Initialize everything globally
dependencyinjection.InitializeEverything()

// Use services with hidden dependencies
service := &dependencyinjection.UserService{}
service.GetUser(1) // Uses global DB

// Service locator pattern
locator := dependencyinjection.GetServiceLocator()
locator.Register("database", "my_db")
db := locator.Get("database").(string) // What could go wrong?

// Global singleton
db := dependencyinjection.GetDB()

// Process with hidden dependencies
processor := &dependencyinjection.OrderProcessor{}
processor.Process() // Uses global DB, Cache, Logger
```

## Testing This Code

```go
// Testing is "easy" - just modify globals!
func TestSomething(t *testing.T) {
    // Save old globals
    oldDB := dependencyinjection.DB
    
    // Set test globals
    dependencyinjection.DB = "test_db"
    
    // Run test
    // ...
    
    // Restore globals (don't forget!)
    dependencyinjection.DB = oldDB
    
    // Hope no other test is running in parallel!
}
```

## Common Anti-Patterns Celebrated Here

1. **Global variables** - Everything is a global
2. **Service locator** - Runtime dependency resolution
3. **Singleton pattern** - One instance, globally accessible
4. **Direct instantiation** - `new()` inside functions
5. **Circular dependencies** - A needs B needs A
6. **init() side effects** - Automatic execution on import
7. **Hidden dependencies** - Not in signature or struct
8. **Hard-coded config** - Constants and secrets in code
9. **Heavy constructors** - I/O in constructors
10. **Package-level state** - Module-scoped globals
11. **God objects** - One struct for everything
12. **Concrete coupling** - No interfaces, direct types

## WARNING

**This code demonstrates what NOT to do.** Proper dependency injection in Go:
- Pass dependencies as parameters or struct fields
- Use interfaces to decouple from concrete implementations
- Inject dependencies through constructors
- Avoid global state except for truly global config
- Use context.Context for request-scoped dependencies
- Keep constructors simple and fast
- Make dependencies explicit in function signatures
- Use dependency injection frameworks only when necessary
- Prefer composition over inheritance/global state

But where's the convenience in that? 😈
