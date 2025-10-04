# Interfaces: Type Safety is Optional

`interface{}` for everything. God interfaces. Type assertions without checks. The erasure of all type safety.

## What's Inside

**ProcessAnything** - Accepts `interface{}`, returns `interface{}`. No type information anywhere. Runtime panics when you guess wrong.

**EmptyInterfaceEverywhere** - Three `interface{}` parameters, returns `interface{}`. Type assertions without comma-ok. Panic roulette.

**MegaInterface** - 18+ methods in one interface. Violates interface segregation. Implementing this is a full-time job.

**Layer1 through Layer5** - Five levels of nested interface composition. Each adds one method. Enterprise-grade over-engineering.

**GodObject** - Implements every interface in the codebase. 25+ methods on one struct. Single Responsibility Principle is dead.

**TypeAssertionRoulette** - Type assertions without checking. `str := data.(string)` followed by `num := data.(int)` on the same value. One will panic.

**InterfaceMap** - Map with `interface{}` keys and values. Non-comparable keys panic at runtime. Type safety is a memory.

**NilInterfaceConfusion** - Interface containing nil vs nil interface. `i != nil` but value is nil. Classic Go gotcha.

**ConcreteToInterface** - Unnecessary conversions to `interface{}`. Abstractions for the sake of abstractions.

**InterfaceToInterface** - Converting `interface{}` to `interface{}`. Does nothing. Looks smart.

## Why This Destroys Type Safety

1. **Runtime panics** - Type assertions fail at runtime instead of compile time
2. **No IDE support** - Auto-complete can't help with `interface{}`
3. **Zero documentation** - Function signatures tell you nothing
4. **Impossible refactoring** - No compiler to catch breaking changes
5. **Testing nightmare** - Can't mock what you can't type

## What You Should Do Instead

- Keep interfaces small (1-3 methods)
- Accept interfaces, return concrete types
- Define interfaces at point of use
- Use comma-ok idiom: `val, ok := i.(Type)`
- Avoid `interface{}` unless you truly need any type
- Use generics (Go 1.18+) for type-safe generic code
- Never create god interfaces

## Do Not

Use these patterns unless you want runtime panics instead of compile-time safety.
