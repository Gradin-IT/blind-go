package model

import "sort"

type Result struct {
	HunterName string
	BlindName  string
}
type ResultSorter struct {
	results []Result
	by      func(p1, p2 *Result) bool // Closure used in the Less method.
}

type By func(p1, p2 *Result) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(results []Result) {
	ps := &ResultSorter{
		results: results,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *ResultSorter) Len() int {
	return len(s.results)
}

// Swap is part of sort.Interface.
func (s *ResultSorter) Swap(i, j int) {
	s.results[i], s.results[j] = s.results[j], s.results[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *ResultSorter) Less(i, j int) bool {
	return s.by(&s.results[i], &s.results[j])
}

// Closures that order the Planet structure.
var Name = func(p1, p2 *Result) bool {
	return p1.HunterName < p2.HunterName
}
