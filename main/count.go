package main

import (
	"fmt"
	"log"
	"sort"
)

// CounterEntry represents an association between a token and
// its count.
type CounterEntry struct {
	Word  string
	Count int
}

// CounterList is a list of CounterEntries. It implements
// sort.Interface (with methods Swap, Len, Less below), so we're
// able to order CounterEntries.
type CounterList []CounterEntry

func (c CounterList) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CounterList) Len() int { return len(c) }

// Sorting is by Counts, not Tokens.
func (c CounterList) Less(i, j int) bool { return c[i].Count < c[j].Count }

// Counter is a frequency table, mapping Tokens to their
// numbers of occurrence. It's loosely modeled after
// python's collections.Counter.
type Counter struct {
	Data map[string]int
}

// Update bumps the count for Token up.
func (c Counter) Update(token Token) {
	if count, ok := c.Data[token.Datum]; ok {
		c.Data[token.Datum] = count + 1
	} else {
		c.Data[token.Datum] = 1
	}
}

// Top returns the top n most frequent CounterEntries in the Counter.
// N must be an integer greater than or equal to 0 and less than
// the size of the Counter. If n is 0, all entries will be returned.
// The order parameter takes two values: 'asc' and 'desc' and is used
// to determine the sorting order: ascending or descending. It returns
// an ordered CounterList.
func (c Counter) Top(n int, order string) CounterList {
	p := make(CounterList, len(c.Data))
	i := 0
	for k, v := range c.Data {
		p[i] = CounterEntry{k, v}
		i++
	}
	switch order {
	case "asc":
		sort.Sort(p)
	case "desc":
		sort.Sort(sort.Reverse(p))
	default:
		log.Fatal(fmt.Sprintf(
			"error: unrecognized order argument '%s'", order))
	}

	if n > 0 && n < len(p) {
		return p[:n]
	}
	return p
}

func (c Counter) Lookup(word string) int {
	if count, ok := c.Data[word]; ok {
		return count
	}
	return 0
}

// NewCounter is a Counter factory. It takes a pointer to a list of
// Tokens, populates a Counter with it, and returns the Counter.
func NewCounter(tokens []Token) Counter {
	counter := Counter{}
	counter.Data = make(map[string]int)

	count := 0
	last := ""
	for _, token := range tokens {
		if token.Datum == last {
			count++
		} else {
			count = 1
		}
		last = token.Datum
		counter.Update(token)
	}
	return counter
}

