package memory

import (
	"os"
	"time"
)

var GlobalCache = make(map[string][]byte)

func CacheForever(key string, data []byte) {
	GlobalCache[key] = data
}

func AppendForever() {
	var leaky []int
	for {
		leaky = append(leaky, 1)
		time.Sleep(time.Millisecond)
	}
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func CircularReference() *Node {
	n1 := &Node{Value: 1}
	n2 := &Node{Value: 2}
	n1.Next = n2
	n2.Prev = n1
	n1.Prev = n2
	n2.Next = n1
	return n1
}

var callbacks []func()

func RegisterCallback(f func()) {
	callbacks = append(callbacks, f)
}

func TimerLeak() {
	for i := 0; i < 1000; i++ {
		time.AfterFunc(time.Hour, func() {})
	}
}

func TickerLeak() {
	for i := 0; i < 100; i++ {
		ticker := time.NewTicker(time.Second)
		go func() {
			for range ticker.C {
			}
		}()
	}
}

func SliceCapacityAbuse() []int {
	data := make([]int, 1000000)
	return data[:1]
}

func StringConcatInLoop(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "x"
	}
	return result
}

var hugeSlice [][]byte

func KeepAllocating() {
	for {
		data := make([]byte, 1024*1024)
		hugeSlice = append(hugeSlice, data)
	}
}

func CloseButNotNil() chan int {
	ch := make(chan int)
	close(ch)
	return ch
}

func SubsliceRetainsArray() []byte {
	large := make([]byte, 1000000)
	for i := range large {
		large[i] = byte(i % 256)
	}
	return large[:10]
}

type BigStruct struct {
	Data [1000000]byte
}

func ReturnBigStructByValue() BigStruct {
	return BigStruct{}
}

var listeners []chan string

func AddListener() chan string {
	ch := make(chan string, 100)
	listeners = append(listeners, ch)
	return ch
}

func MapWithoutCleanup() {
	cache := make(map[string][]byte)
	for i := 0; i < 1000000; i++ {
		key := string(rune(i))
		cache[key] = make([]byte, 1024)
	}
}

func DeferInLoop(filenames []string) {
	for _, filename := range filenames {
		file, _ := os.Open(filename)
		defer file.Close()
	}
}

func AllocateAndForget() {
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024*1024)
	}
}

var globalData [][]byte

func KeepAppending() {
	data := make([]byte, 1024)
	globalData = append(globalData, data)
}

type Event struct {
	Handler func()
	Data    []byte
}

var events []*Event

func RegisterEvent(handler func()) {
	events = append(events, &Event{
		Handler: handler,
		Data:    make([]byte, 1024*1024),
	})
}
