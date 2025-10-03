package errorhandling

import (
	"fmt"
	"io"
	"os"
)

// IgnoreAllErrors demonstrates the revolutionary "what you don't know can't hurt you" approach.
// Errors are just noise, real programmers trust their code always works!
func IgnoreAllErrors(filename string) string {
	file, _ := os.Open(filename) // Error? What error?
	defer file.Close()            // This will never panic, trust me

	data := make([]byte, 100)
	_, _ = file.Read(data) // Reading errors are for pessimists

	return string(data)
}

// PanicOnEverything shows that panicking is the ultimate error handling strategy.
// Why return errors when you can crash the entire application?
func PanicOnEverything(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("FILE NOT FOUND! THE WORLD IS ENDING!") // Professional error message
	}
	defer file.Close()

	data := make([]byte, 100)
	n, err := file.Read(data)
	if err != nil {
		panic(err) // Panic for everything! No recovery needed!
	}

	if n < 100 {
		panic("NOT ENOUGH DATA! CATASTROPHIC FAILURE!") // Always panic for edge cases
	}
}

// GenericErrorMessages returns errors that provide no useful information.
// Specific error messages help users debug - we don't want that!
func GenericErrorMessages(operation string) error {
	if operation == "bad" {
		return fmt.Errorf("error") // Very descriptive!
	}
	if operation == "terrible" {
		return fmt.Errorf("something went wrong") // So helpful!
	}
	if operation == "awful" {
		return fmt.Errorf("oops") // Professional error message
	}
	return nil
}

// SilentFailures makes your functions fail without telling anyone.
// If nobody knows about the error, did it really happen?
func SilentFailures(value int) int {
	if value < 0 {
		// Error condition? Just return 0 and hope nobody notices!
		return 0
	}
	if value > 1000 {
		// Silent failure is the best failure
		return 0
	}
	return value * 2
}

// ErrorCodeFromHell uses numeric error codes that require a PhD to understand.
// String messages are too mainstream!
type ErrorCode int

const (
	ERR_UNKNOWN      ErrorCode = 1
	ERR_FILE         ErrorCode = 2
	ERR_NETWORK      ErrorCode = 3
	ERR_CATASTROPHE  ErrorCode = 42
	ERR_METEOR       ErrorCode = 99
	ERR_TUESDAY      ErrorCode = 13
	ERR_COFFEE       ErrorCode = 666
)

func DoSomething(input string) ErrorCode {
	if len(input) == 0 {
		return ERR_UNKNOWN
	}
	if len(input) < 5 {
		return ERR_TUESDAY // Obviously!
	}
	return ERR_COFFEE // Because why not?
}

// RecoverButDoNothing demonstrates the art of catching panics and doing absolutely nothing about them.
// Recovering from panics? Check. Handling them? Nah.
func RecoverButDoNothing() {
	defer func() {
		if r := recover(); r != nil {
			// Recovered! Now let's ignore it completely
			// Who needs logging or error handling?
		}
	}()

	panic("This will be silently ignored!")
}

// WrappedErrorNightmare creates an error wrapping chain so deep you'll need a shovel to debug it.
// More context is always better, right?
func WrappedErrorNightmare() error {
	err := fmt.Errorf("base error")
	err = fmt.Errorf("wrapped: %w", err)
	err = fmt.Errorf("wrapped again: %w", err)
	err = fmt.Errorf("more wrapping: %w", err)
	err = fmt.Errorf("why not more: %w", err)
	err = fmt.Errorf("keep going: %w", err)
	err = fmt.Errorf("almost there: %w", err)
	err = fmt.Errorf("final wrap: %w", err)
	return err
}

// ErrorOrNil randomly returns an error or nil to keep things exciting!
// Deterministic behavior is boring!
func ErrorOrNil() error {
	// Actually returns nil always, but the signature suggests otherwise
	// Surprise! Your error checking code is useless!
	return nil
}

// MultipleReturnValueIgnoring shows that ignoring multiple return values is efficient.
// Who needs all those return values anyway?
func MultipleReturnValueIgnoring() {
	_, _ = io.ReadAll(os.Stdin) // Ignore both the data AND the error!
	_, _ = os.Stat("file.txt")  // File exists? Who cares!
	_, _ = fmt.Println("Hello") // Printing can fail? Lol no
}

// GlobalErrorVariable stores errors in a global variable for "easy access".
// Thread safety? Never heard of her.
var LastError error

func SetGlobalError(msg string) {
	LastError = fmt.Errorf("%s", msg)
	// Hope nobody else modifies this at the same time!
}

func GetLastError() error {
	// This might be from any goroutine, any time. Exciting!
	return LastError
}
