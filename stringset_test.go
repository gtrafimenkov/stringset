// SPDX-License-Identifier: MIT-0

package stringset

import (
	"fmt"
	"testing"
)

func TestNewEmpty(t *testing.T) {
	ss := New()
	if ss.Size() != 0 {
		t.Errorf("New string set should be of size 0; got %v", ss.Size())
	}
}

func TestNewFilled(t *testing.T) {
	ss := New("foo", "bar")
	if ss.Size() != 2 {
		t.Errorf("New string set %v should be of size 2; got %v", ss, ss.Size())
	}
}

func TestAdd(t *testing.T) {
	ss := New()
	ss.Add("foo")
	ss.Add("foo")
	if ss.Size() != 1 || ss.Items()[0] != "foo" {
		t.Errorf("string set should contain only one item foo; got %v", ss)
	}

	ss.Add("bar")
	if ss.Size() != 2 {
		t.Errorf("string set should contain only two items; got %v", ss.Size())
	}

	for _, lookingFor := range []string{"foo", "bar"} {
		found := false
		for _, item := range ss.Items() {
			if item == lookingFor {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("item %q not found in %v", lookingFor, ss)
		}
	}
}

func TestHasItem(t *testing.T) {
	ss := New("foo", "bar")

	for _, lookingFor := range []string{"foo", "bar"} {
		if !ss.HasItem(lookingFor) {
			t.Errorf("Item %q not found in %v", lookingFor, ss)
		}
	}

	for _, lookingFor := range []string{"baz", ""} {
		if ss.HasItem("baz") {
			t.Errorf("Item %q should not be in the set %v, but it is", lookingFor, ss)
		}
	}
}

func TestRemove(t *testing.T) {
	ss := New("foo", "bar")
	ss.Remove("bar")
	if ss.Size() != 1 {
		t.Errorf("Size of the set %v shold be %d, but it is %d", ss, 1, ss.Size())
	}
	ss.Remove("foo")
	if ss.Size() != 0 {
		t.Errorf("Size of the set %v shold be %d, but it is %d", ss, 0, ss.Size())
	}
	ss.Remove("baz")
	if ss.Size() != 0 {
		t.Errorf("Size of the set %v shold be %d, but it is %d", ss, 0, ss.Size())
	}
}

func TestConvertionToString(t *testing.T) {
	tests := []struct {
		ss       StringSet
		expected string
	}{
		{New(), `[]`},
		{New("foo"), `["foo"]`},
		{New("foo", "bar"), `["bar", "foo"]`},
	}
	for _, test := range tests {
		if fmt.Sprint(test.ss) != test.expected {
			t.Errorf("Expecting %v, got %v", test.expected, test.ss)
		}
	}
}
