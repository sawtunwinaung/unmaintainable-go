# Error Handling: Ignore, Panic, or Lie

Every terrible way to handle errors in Go. From ignoring them completely to panicking on everything.

## What's Inside

**IgnoreAllErrors** - `file, _ := os.Open(filename)` then `defer file.Close()`. Nil pointer dereference guaranteed. Blank identifier as error handling strategy.

**PanicOnEverything** - Every error becomes a panic. Stack traces instead of graceful degradation. Crashes instead of error returns.

**GenericErrorMessages** - `return fmt.Errorf("error")` and `fmt.Errorf("oops")`. Zero context, zero debugging information. Good luck.

**SilentFailures** - Returns 0 for all error cases without returning an error. Function signature lies about its behavior. Corruption instead of failure.

**DoSomething** - Returns cryptic error codes. `ERR_TUESDAY`, `ERR_COFFEE`, `ERR_METEOR`. Numeric codes requiring documentation nobody wrote.

**RecoverButDoNothing** - Recovers from panics then ignores them. No logging, no error handling. Silent failure with extra steps.

**WrappedErrorNightmare** - Eight levels of error wrapping for a simple error. Error messages longer than the code. Stack overflow when printing.

**LastError** - Global variable to store errors. Race conditions between goroutines. Thread-unsafe error state.

**ChainOfPanics** - Panic in a defer that's recovering from a panic. Double panic chaos.

**WrongErrorCheck** - Panics when error is nil. Backwards error checking logic.

**IgnoreErrorInLoop** - Ignores errors in a loop. If one file fails, corrupt all results silently.

## Why This Destroys Your Application

1. **Nil pointer crashes** - Ignoring errors leads to nil dereferences
2. **No debugging info** - Generic errors provide zero context
3. **Silent data corruption** - Functions that fail silently
4. **Unrecoverable crashes** - Panic instead of returning errors
5. **Race conditions** - Global error state accessed from multiple goroutines

## What You Should Do Instead

- Always check errors: `if err != nil { return err }`
- Return descriptive errors with context
- Use `fmt.Errorf("operation failed: %w", err)` for wrapping
- Never panic in library code (only for programmer errors)
- Don't use global error state
- Recover from panics only when you can actually handle them

## Do Not

Use any of these patterns unless you enjoy production incidents.
