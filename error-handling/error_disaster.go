package errorhandling

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func IgnoreAllErrors(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()
	data := make([]byte, 100)
	_, _ = file.Read(data)
	return string(data)
}

func PanicOnEverything(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("FILE NOT FOUND! THE WORLD IS ENDING!")
	}
	defer file.Close()
	data := make([]byte, 100)
	n, err := file.Read(data)
	if err != nil {
		panic(err)
	}
	if n < 100 {
		panic("NOT ENOUGH DATA! CATASTROPHIC FAILURE!")
	}
}

func GenericErrorMessages(operation string) error {
	if operation == "bad" {
		return fmt.Errorf("error")
	}
	if operation == "terrible" {
		return fmt.Errorf("something went wrong")
	}
	if operation == "awful" {
		return fmt.Errorf("oops")
	}
	return nil
}

func SilentFailures(value int) int {
	if value < 0 {
		return 0
	}
	if value > 1000 {
		return 0
	}
	return value * 2
}

type ErrorCode int

const (
	ERR_UNKNOWN     ErrorCode = 1
	ERR_FILE        ErrorCode = 2
	ERR_NETWORK     ErrorCode = 3
	ERR_CATASTROPHE ErrorCode = 42
	ERR_METEOR      ErrorCode = 99
	ERR_TUESDAY     ErrorCode = 13
	ERR_COFFEE      ErrorCode = 666
)

func DoSomething(input string) ErrorCode {
	if len(input) == 0 {
		return ERR_UNKNOWN
	}
	if len(input) < 5 {
		return ERR_TUESDAY
	}
	return ERR_COFFEE
}

func RecoverButDoNothing() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	panic("This will be silently ignored!")
}

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

func MultipleReturnValueIgnoring() {
	_, _ = io.ReadAll(os.Stdin)
	_, _ = os.Stat("file.txt")
	_, _ = fmt.Println("Hello")
}

var LastError error

func SetGlobalError(msg string) {
	LastError = fmt.Errorf("%s", msg)
}

func GetLastError() error {
	return LastError
}

func ChainOfPanics() {
	defer func() {
		panic("panic in defer")
	}()
	panic("original panic")
}

func WrongErrorCheck(err error) {
	if err == nil {
		panic("got nil error, this is wrong somehow")
	}
}

func IgnoreErrorInLoop() []string {
	files := []string{"a.txt", "b.txt", "c.txt"}
	var results []string
	for _, f := range files {
		data, _ := os.ReadFile(f)
		results = append(results, string(data))
	}
	return results
}

func ReturnsErrorButNoContext() error {
	_, err := os.Open("/nonexistent")
	return err
}

func ErrorStringsStartWithCapital() error {
	return errors.New("This Error Starts With Capital And Has Punctuation!")
}

func CompareErrorStrings(err error) bool {
	return err.Error() == "file not found"
}

func PanicInLibraryCode(value int) int {
	if value < 0 {
		panic("negative values not allowed")
	}
	return value * 2
}

func AssertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func MustParseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ErrorAsControlFlow(value int) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("handled error via panic")
		}
	}()
	if value < 0 {
		panic("negative")
	}
	return "positive"
}

var ErrSentinel = errors.New("sentinel error")

func UseSentinelErrorWrong() error {
	return fmt.Errorf("wrapped: %w", ErrSentinel)
}

func CheckErrorByString(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == "EOF"
}

func IgnoreCloseError(filename string) string {
	file, _ := os.Open(filename)
	defer file.Close()
	data, _ := io.ReadAll(file)
	return string(data)
}

func ReturnNilErrorButPanic() error {
	panic("I said I return error but I panic instead")
	return nil
}
