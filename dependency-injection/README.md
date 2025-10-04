# Dependency Injection: Global Variables as Architecture

Global state everywhere. Service locator pattern. Circular dependencies. Hard-coded configuration. Every DI anti-pattern.

## What's Inside

**DB, Cache, Logger, Config** - All global variables. Access from anywhere, modify from anywhere. No dependency injection needed.

**InitializeEverything** - One massive function to initialize all globals. Separation of concerns is overrated.

**UserService** - Empty struct, uses global dependencies. Function signature tells you nothing about what it needs.

**ServiceLocator** - Runtime dependency resolution with type assertions. Hidden dependencies, runtime panics if service not registered.

**GetDB** - Singleton with race condition in initialization. First caller wins. No synchronization.

**SendEmail** - Creates dependencies inline. Hard-coded email server. Can't mock in tests. Tight coupling as design.

**ServiceA/ServiceB** - Circular dependencies. A needs B, B needs A. Initialization order puzzle.

**init()** - Package-level init with side effects. Runs on import. Prints to stdout during initialization. Import order matters.

**OrderProcessor** - Hidden dependencies on global DB, Cache, Logger. Reading the struct tells you nothing.

**DATABASE_URL, API_KEY** - Hard-coded config and secrets in code. Version control as secret management. Different configs require recompilation.

**NewApplicationService** - Constructor that does I/O. Can't return errors. Takes 10 seconds to instantiate. Testing requires real connections.

**packageState** - Package-level map accessed without synchronization. Global state but more local. Race conditions included.

**Application** - God object with 10+ methods. Single Responsibility Principle violated. Everything in one struct.

**OrderService** - Depends on concrete `*MySQLDatabase`. Can't swap database. Can't mock in tests. Tight coupling to implementation.

## Why This Destroys Maintainability

1. **Impossible to test** - Global state makes unit testing a nightmare
2. **Hidden dependencies** - Can't tell what a function needs from its signature
3. **Race conditions** - Global state accessed from multiple goroutines
4. **Tight coupling** - Can't swap implementations
5. **Initialization order** - init() side effects depend on import order
6. **No mocking** - Hard-coded dependencies can't be replaced

## What You Should Do Instead

- Pass dependencies as parameters or struct fields
- Use interfaces to decouple from concrete implementations
- Inject dependencies through constructors
- Avoid global state except for truly global config
- Use `context.Context` for request-scoped dependencies
- Keep constructors simple and fast
- Make dependencies explicit in function signatures
- Use dependency injection only when needed
- Prefer composition over global state

## Do Not

Use these patterns unless you want untestable, tightly-coupled code.
