package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/ariel17/project-euler/problem1"
	"github.com/ariel17/project-euler/problem2"
	"github.com/ariel17/project-euler/problem3"
	"github.com/ariel17/project-euler/problem4"
)

var (
	problems = map[int]func(){
		1: problem1.Solve,
		2: problem2.Solve,
		3: problem3.Solve,
		4: problem4.Solve,
	}
)

func main() {
	problem := flag.Int("problem", 0, "The problem number to solve.")
	show := flag.Bool("show-availables", false, "Show solved problem numbers.")
	flag.Parse()

	if *show {
		fmt.Printf("> Available problems: %v", problemList())
		return
	}

	if problem == nil || (problem != nil && *problem <= 0) {
		return
	}

	problems[*problem]()
}

func problemList() []int {
	l := []int{}
	for k := range problems {
		l = append(l, k)
	}
	sort.Ints(l)
	return l
}
