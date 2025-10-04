package main

import (
	"fmt"
	"time"

	"github.com/malewicz1337/unmaintainable-go/concurrency"
	"github.com/malewicz1337/unmaintainable-go/dependency-injection"
	errorhandling "github.com/malewicz1337/unmaintainable-go/error-handling"
	"github.com/malewicz1337/unmaintainable-go/goroutines"
	"github.com/malewicz1337/unmaintainable-go/interfaces"
	"github.com/malewicz1337/unmaintainable-go/naming"
)

func main() {
	fmt.Println("=== Unmaintainable Go Examples ===")
	fmt.Println("DO NOT use any of these patterns in production")
	fmt.Println()

	fmt.Println("1. Error Handling Disasters:")
	data := errorhandling.IgnoreAllErrors("nonexistent.txt")
	fmt.Printf("Read data (probably nil pointer crash): %v\n", data)

	fmt.Println("\n2. Silent Failures:")
	result := errorhandling.SilentFailures(-100)
	fmt.Printf("Silent failure returns: %d (should be error)\n", result)

	fmt.Println("\n3. Generic Error Messages:")
	err := errorhandling.GenericErrorMessages("bad")
	fmt.Printf("Error: %v (zero context)\n", err)

	fmt.Println("\n4. Race Conditions:")
	counter := &concurrency.SharedCounter{}
	for i := 0; i < 10; i++ {
		go counter.Increment()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Counter value (racy): %d (probably not 10)\n", counter.Get())

	fmt.Println("\n5. Global State Modification:")
	concurrency.ModifyGlobalState()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Global state modified (might crash with concurrent map writes)")

	fmt.Println("\n6. Interface Abuse:")
	result2 := interfaces.ProcessAnything("hello")
	fmt.Printf("Result type unknown: %v\n", result2)

	fmt.Println("\n7. Type Assertion Without Checks:")
	fmt.Println("This will panic:")

	fmt.Println("\n8. Dependency Injection via Globals:")
	dependencyinjection.InitializeEverything()
	service := &dependencyinjection.UserService{}
	user, _ := service.GetUser(1)
	fmt.Printf("User: %s (using global DB)\n", user)

	fmt.Println("\n9. Terrible Naming:")
	badResult := naming.SingleLetterFunc(1, 2, "test", true, []int{1, 2, 3})
	fmt.Printf("Single letter variables everywhere: %d\n", badResult)

	fmt.Println("\n10. Goroutine Hell:")
	fmt.Println("Don't run this - it will spawn infinite goroutines")
	fmt.Println("Uncomment at your own risk:")
	fmt.Println("// goroutines.ProcessData([]int{1, 2, 3})")

	fmt.Println("\n11. Increment Counter Race:")
	goroutines.IncrementCounter()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Global counter (race condition): %d\n", goroutines.GlobalCounter)

	fmt.Println("\n=== All examples executed (that don't crash) ===")
	fmt.Println("Run with -race to see the race conditions")
}
