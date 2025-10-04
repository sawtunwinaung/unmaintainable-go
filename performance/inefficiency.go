package performance

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func StringConcatInLoop(items []string) string {
	result := ""
	for _, item := range items {
		result = result + item + ","
	}
	return result
}

func CompileRegexEveryTime(input string) bool {
	re, _ := regexp.Compile(`[a-z]+`)
	return re.MatchString(input)
}

func MD5InLoop(data []string) []string {
	var results []string
	for _, d := range data {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(d)))
		results = append(results, hash)
	}
	return results
}

func ReflectInHotPath(value interface{}) string {
	v := reflect.ValueOf(value)
	return v.Type().String()
}

func JSONMarshalInLoop(items []map[string]interface{}) []string {
	var results []string
	for _, item := range items {
		data, _ := json.Marshal(item)
		results = append(results, string(data))
	}
	return results
}

func NoPreallocation(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}
	return result
}

func SortInLoop(data [][]int) {
	for i := 0; i < len(data); i++ {
		sort.Ints(data[i])
	}
}

func DeferInTightLoop() int {
	sum := 0
	for i := 0; i < 10000; i++ {
		defer func(n int) {
			sum += n
		}(i)
	}
	return sum
}

func StringSplit(data string) []string {
	parts := strings.Split(data, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func ConvertIntToString(nums []int) []string {
	var results []string
	for _, n := range nums {
		results = append(results, fmt.Sprintf("%d", n))
	}
	return results
}

func MapLookupInLoop(data []string, lookup map[string]int) int {
	sum := 0
	for _, key := range data {
		if val, ok := lookup[key]; ok {
			sum += val
		}
	}
	return sum
}

func TimeNowInLoop() {
	for i := 0; i < 10000; i++ {
		_ = time.Now()
	}
}

func InterfaceConversion(items []int) []interface{} {
	result := make([]interface{}, len(items))
	for i, v := range items {
		result[i] = v
	}
	return result
}

func DeepCopy(m map[string]interface{}) map[string]interface{} {
	data, _ := json.Marshal(m)
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	return result
}

func CheckContains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func RemoveFromSlice(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func ReverseString(s string) string {
	result := ""
	for _, c := range s {
		result = string(c) + result
	}
	return result
}

func CountWords(text string) map[string]int {
	words := strings.Fields(text)
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func FilterSlice(items []int, predicate func(int) bool) []int {
	var result []int
	for _, item := range items {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func ParseIntsSlowly(strings []string) []int {
	var ints []int
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

var globalRegex = regexp.MustCompile(`[a-z]+`)

func UseCompiledRegex(input string) bool {
	return globalRegex.MatchString(input)
}

func CopyMap(m map[string]int) map[string]int {
	result := make(map[string]int, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func Contains(slice []string, item string) bool {
	set := make(map[string]bool)
	for _, s := range slice {
		set[s] = true
	}
	return set[item]
}

func BufferedChannelVsUnbuffered() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

type Cache struct {
	data map[string]interface{}
}

func (c *Cache) GetOrCompute(key string, compute func() interface{}) interface{} {
	if val, ok := c.data[key]; ok {
		return val
	}
	val := compute()
	c.data[key] = val
	return val
}

func NestedLoops(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				sum++
			}
		}
	}
	return sum
}
