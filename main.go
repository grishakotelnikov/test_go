package main

import (
	"fmt"
	"sort"
	"strconv"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}
func IntSliceToString(t []int) string {
	var result string
	for i := 0; i < len(t); i++ {
		result += strconv.Itoa(t[i])
	}
	return result
}

func MergeSlices(f []float32, i []int32) []int {
	var result []int
	for i := 0; i < len(f); i++ {
		result = append(result, int(f[i]))
	}
	for j := 0; j < len(i); j++ {
		result = append(result, int(i[j]))
	}

	return result

}

type Pair struct {
	key   int
	value string
}
type SortBy []Pair

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].key < a[j].key }

func GetMapValuesSortedByKey(m1 map[int]string) (result []string) {
	p := make(SortBy, len(m1))
	i := 0
	for k, v := range m1 {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)

	for _, v := range p {
		result = append(result, v.value)
	}
	return
}

func main() {
	var input = map[int]string{
		10: "Зима",
		30: "Лето",
		20: "Весна",
		40: "Осень",
	}
	for _, v := range GetMapValuesSortedByKey(input) {
		fmt.Println(v)
	}
}
