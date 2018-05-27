package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

// func (s string) Len() int           { return len(s) }
// func (s string) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s string) Less(i, j int) bool { return s[i] > s[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println(studyGroup)
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.StringSlice(s).Sort()
	sort.Sort(sort.Reverse(sort.StringSlice(s)))

	fmt.Println(s)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.IntSlice(n).Sort()
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
