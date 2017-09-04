package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int{
	return len(p)
}

func (p people) Swap(i, j int){
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool{
	return p[i][0] < p[j][0]
}

func main() {

	s := []string{"Zeno", "Jonh", "Al", "Jenny"}
	sort.Strings(s)
	fmt.Println(s)

	i := []int{7,4,8,2,9,19,12,32,3}
	sort.Ints(i)
	fmt.Println(i)


	studyPeople := people{"Zenoz", "Jonha", "Alc", "Jennyt"}
	sort.Sort(studyPeople)
	fmt.Println(studyPeople)
}
