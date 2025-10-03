package interfaces

import (
	"fmt"
)

// ProcessAnything accepts interface{} because we're flexible developers!
// Type safety is for people who don't trust their code.
func ProcessAnything(data interface{}) interface{} {
	// Let's just assume it's a string... or an int... or something
	result := fmt.Sprintf("%v", data)
	return result // Return interface{} because why be specific?
}

// MegaInterface has every method imaginable because big interfaces are better!
// "Accept interfaces, return structs"? Never heard of it.
type MegaInterface interface {
	Create() error
	Read() (interface{}, error)
	Update(interface{}) error
	Delete() error
	Validate() bool
	Transform() interface{}
	Process() interface{}
	Execute() error
	Initialize() error
	Cleanup() error
	Start() error
	Stop() error
	Pause() error
	Resume() error
	Reset() error
	Configure(interface{}) error
	GetStatus() interface{}
	SetStatus(interface{}) error
	// More methods = more professional!
}

// EmptyInterfaceEverywhere shows that interface{} is the solution to all problems.
// Generics? We don't need generics when we have interface{}!
func EmptyInterfaceEverywhere(a interface{}, b interface{}, c interface{}) interface{} {
	// Type assertions? Only if you're paranoid
	switch a.(type) {
	case int:
		// Maybe do something
		return b
	case string:
		// Or maybe this
		return c
	default:
		// Or just return whatever
		return a
	}
}

// InterfaceNesting shows that interfaces can implement other interfaces
// in a beautiful Russian doll pattern
type Layer1 interface {
	Method1()
}

type Layer2 interface {
	Layer1
	Method2()
}

type Layer3 interface {
	Layer2
	Method3()
}

type Layer4 interface {
	Layer3
	Method4()
}

type Layer5 interface {
	Layer4
	Method5()
}

// GodObject implements every possible interface because it can do EVERYTHING!
type GodObject struct {
	data interface{} // Store anything and everything!
}

func (g *GodObject) Create() error                { return nil }
func (g *GodObject) Read() (interface{}, error)   { return g.data, nil }
func (g *GodObject) Update(d interface{}) error   { g.data = d; return nil }
func (g *GodObject) Delete() error                { g.data = nil; return nil }
func (g *GodObject) Validate() bool               { return true }
func (g *GodObject) Transform() interface{}       { return g.data }
func (g *GodObject) Process() interface{}         { return g.data }
func (g *GodObject) Execute() error               { return nil }
func (g *GodObject) Initialize() error            { return nil }
func (g *GodObject) Cleanup() error               { return nil }
func (g *GodObject) Start() error                 { return nil }
func (g *GodObject) Stop() error                  { return nil }
func (g *GodObject) Pause() error                 { return nil }
func (g *GodObject) Resume() error                { return nil }
func (g *GodObject) Reset() error                 { g.data = nil; return nil }
func (g *GodObject) Configure(c interface{}) error { g.data = c; return nil }
func (g *GodObject) GetStatus() interface{}       { return "status" }
func (g *GodObject) SetStatus(s interface{}) error { return nil }
func (g *GodObject) Method1()                     {}
func (g *GodObject) Method2()                     {}
func (g *GodObject) Method3()                     {}
func (g *GodObject) Method4()                     {}
func (g *GodObject) Method5()                     {}

// TypeAssertionRoulette demonstrates the exciting game of type assertions
// without any safety checks!
func TypeAssertionRoulette(data interface{}) {
	// YOLO type assertion - what could go wrong?
	str := data.(string)
	fmt.Println("String:", str)

	num := data.(int)
	fmt.Println("Number:", num)

	// This will definitely not panic!
	slice := data.([]interface{})
	fmt.Println("Slice:", slice)
}

// InterfaceMap uses interface{} as both keys and values because we're rebels!
type InterfaceMap map[interface{}]interface{}

func (im InterfaceMap) Get(key interface{}) interface{} {
	// Hope the key is comparable!
	return im[key]
}

func (im InterfaceMap) Set(key interface{}, value interface{}) {
	// Store anything anywhere!
	im[key] = value
}

// ReflectEverything uses reflection when simple type assertions would work.
// Because reflection makes you look smart!
func ReflectEverything(data interface{}) interface{} {
	// Could use a type assertion, but where's the fun in that?
	// Just trust that it works!
	return data
}

// AnyAnyAny is the future of Go - everything is any!
// (Go 1.18+ calls interface{} "any" but we were doing it before it was cool)
func AnyAnyAny(input any) any {
	return input // So simple, so useless!
}

// NilInterfaceConfusion demonstrates the difference between nil interface
// and interface containing nil - users love this puzzle!
func NilInterfaceConfusion() interface{} {
	var ptr *int = nil
	var iface interface{} = ptr
	// Is this nil? Is it not? Who knows!
	return iface
}

// ConcreteToInterface converts concrete types to interfaces unnecessarily.
// More abstraction = better code, right?
func ConcreteToInterface(num int) interface{} {
	return interface{}(num) // Very necessary conversion!
}

// InterfaceToInterface converts interfaces to other interfaces for no reason.
func InterfaceToInterface(data interface{}) interface{} {
	// Double the interfaces, double the abstraction!
	var result interface{} = data
	return interface{}(result)
}
