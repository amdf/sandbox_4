package main

import "testing"

func TestAb(t *testing.T) {

	var str string
	var b bool
	var errors int
	tr := NewTrie()
	tr.Insert("ab")
	tr.Insert("az")
	tr.Insert("faz")
	tr.Insert("braz")
	tr.Insert("godaz")

	// str = tr.SearchMore("az")
	// // tr.DeleteWord("az")
	// b = tr.SearchWord("az")
	// tr.Insert("az")
	// b = tr.SearchWord("az")

	search := []string{
		"ab",
		"az",
		"faz",
		"strab",
		"straz",
		"mafaz",
		"kez",
	}

	for _, s := range search {

		str = tr.Search(s)
		if str == s {
			errors++
		}
	}

	str = tr.SearchMore("z")
	if "z" == str || !b {
		t.Error()
	}
}
