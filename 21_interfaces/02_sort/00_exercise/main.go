package main

import (
	"sort"
	"fmt"
)

type People []string

func (p People) Len() int		{ return len(p) }
func (p People) Swap(i, j int)		{ p[i], p[j] = p[j], p[i] }
func (p People) Less(i, j int)bool	{ return p[i] < p[j] }

func main() {
	studyGroup := People{"Zeno", "John", "AI", "Jenny"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}