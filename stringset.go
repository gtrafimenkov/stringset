// SPDX-License-Identifier: MIT-0

package stringset

import (
	"fmt"
	"sort"
	"strings"
)

// StringSet describes a set of strings
type StringSet interface {

	// HasItem checks if the set has the given item
	HasItem(value string) bool

	// Add adds single item into the set
	Add(item string)

	// AddItems adds items into the set
	AddItems(items ...string)

	// Remove removes a single item from the set
	Remove(item string)

	// Size returns size of the set
	Size() int

	// Items returns list of all items in the set
	Items() []string
}

// New creates a new StringSet add adds provided items into it
func New(items ...string) StringSet {
	ss := stringSetImpl{}
	ss.AddItems(items...)
	return &ss
}

type stringSetImpl map[string]bool

// HasItem checks if the set has the given item
func (ss *stringSetImpl) HasItem(value string) bool {
	_, ok := (*ss)[value]
	return ok
}

// Add adds single item into the set
func (ss *stringSetImpl) Add(item string) {
	(*ss)[item] = true
}

// AddItems adds items into the set
func (ss *stringSetImpl) AddItems(items ...string) {
	for _, item := range items {
		(*ss)[item] = true
	}
}

// Remove removes a single item from the set
func (ss *stringSetImpl) Remove(item string) {
	delete(*ss, item)
}

// Size returns size of the set
func (ss *stringSetImpl) Size() int {
	return len(*ss)
}

// Items returns list of all items in the set
func (ss *stringSetImpl) Items() []string {
	items := make([]string, len(*ss), len(*ss))
	i := 0
	for item := range *ss {
		items[i] = item
		i++
	}
	return items
}

func (ss *stringSetImpl) String() string {
	items := (*ss).Items()
	sort.Strings(items)
	quotedItems := make([]string, len(items), len(items))
	for i := range items {
		quotedItems[i] = fmt.Sprintf("%q", items[i])
	}
	return fmt.Sprintf("[%s]", strings.Join(quotedItems, ", "))
}
