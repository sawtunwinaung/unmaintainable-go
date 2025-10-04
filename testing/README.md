# Testing: How to Make Code Untestable

Global state, hard-coded dependencies, side effects, and non-deterministic behavior. Everything that makes testing impossible.

## What's Inside

**GetUser** - Direct database access via global `DB` variable. Can't test without real database. No dependency injection.

**ProcessOrder** - Hard-coded HTTP call to production API. Tests hit real API. No mocking possible.

**GetCurrentTime** - Returns current time. Tests are time-dependent. Non-deterministic. Fails at midnight.

**RandomNumber** - Returns random value. Non-deterministic. Tests can't assert expected value.

**ReadConfig** - Reads from `/etc/app/config.json`. Tests require filesystem setup. Hard-coded path.

**SaveToFile** - Writes to `/tmp/output.txt`. Tests create real files. Hard-coded path. Side effects.

**ComplexBusinessLogic** - Database + HTTP + file I/O all mixed together. No separation of concerns. Impossible to unit test.

**NotifyUser** - Global `emailer` variable. Database access. Can't test without sending real emails.

**ProcessWithSideEffects** - Reads and writes file. State persists across test runs. Tests interfere with each other.

**IncrementGlobal** - Modifies global counter. Tests can't run in parallel. Order-dependent.

**DependsOnEnvironment** - Reads environment variable. Tests must set ENV vars. Affects other tests.

**CallsOtherService** - Creates HTTP client internally. Hard-coded URL. No timeout configuration. Can't mock.

**TimeBasedLogic** - Behavior changes based on time of day. Tests fail outside business hours.

**RandomBehavior** - Non-deterministic logic. Different result each run. Can't write assertions.

**NewService** - Constructor does I/O and connects to database. Can't instantiate in tests without real dependencies.

**DoWork** - Database, HTTP, and logging all together. No interfaces. All dependencies hard-coded.

**ConcurrentOperation** - Race condition with random sleep. Non-deterministic timing. Flaky tests.

**PublicWrapper** - Calls private function. Can't test private logic in isolation. No way to inject behavior.

**LotsOfParameters** - Ten parameters. Tests require setting up all parameters. Hard to read test cases.

**ReturnsMultipleThings** - Five return values. Tests must handle all of them. Error-prone assertions.

## Why This Can't Be Tested

1. **Global state** - Can't isolate tests, can't run in parallel
2. **Hard-coded dependencies** - Database, HTTP, filesystem all real
3. **No interfaces** - Can't mock or stub
4. **Side effects** - Tests modify shared state
5. **Non-deterministic** - Random, time-based, concurrent
6. **Constructor I/O** - Can't create instances without real resources
7. **Mixed concerns** - Business logic entangled with infrastructure

## What You Should Do Instead

- Inject dependencies through constructors or parameters
- Use interfaces for external dependencies
- Make time injectable (accept `time.Time` parameter)
- Use deterministic random with seeded `rand.New()`
- Separate business logic from I/O
- Return errors instead of panicking
- Keep functions pure when possible
- Use table-driven tests
- Mock external services
- Use in-memory implementations for testing

## Do Not

Write code like this if you ever want to write tests. Or if you want anyone else to write tests.
