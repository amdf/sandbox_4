package main

import "testing"

func TestAb(t *testing.T) {

	var str string
	var b bool
	var errors int
	tr := NewTrie()
	tr.Insert("ab")
	tr.Insert("az")
	tr.Insert("fb")
	tr.Insert("fz")

	b = tr.SearchWord("az")
	tr.DeleteWord("az")
	b = tr.SearchWord("az")
	tr.Insert("az")
	b = tr.SearchWord("az")

	search := []string{
		"ab",
		"az",
		"fab",
		"faz",
		"bb",
		"bz",
	}

	for _, s := range search {

		b = tr.SearchWord(s)
		str = tr.SearchMore(s[:len(s)-1])
		if str == s {
			errors++
		}
	}

	str = tr.SearchMore("z")
	if "z" == str || !b {
		t.Error()
	}
}
