package main

import "testing"

func TestAb(t *testing.T) {

	var str string

	tr := NewTrie()

	tr.Insert("iiii")
	tr.Insert("iiiii")
	tr.Insert("iiiiiiii")
	tr.Insert("iiiiiiiii")
	tr.Insert("iii")
	tr.Insert("ii")
	tr.Insert("i")
	tr.Insert("iiiiiii")
	tr.Insert("iiiiii")
	tr.Insert("iiiiiiiiii")

	str = tr.Search("iiiiii")
	str = tr.Search("iiii")
	str = tr.Search("iiiiiiiiii")
	str = tr.Search("iiiiiiii")
	str = tr.Search("ii")
	str = tr.Search("iiiiiii")
	str = tr.Search("iiiiiiii")
	str = tr.Search("iiiiiiii")
	str = tr.Search("i")
	str = tr.Search("i")
	str = tr.Search("iiiiii")
	str = tr.Search("i")
	str = tr.Search("i")
	str = tr.Search("iiiii")
	str = tr.Search("iii")
	str = tr.Search("iii")
	str = tr.Search("iiiiiiiiii")
	str = tr.Search("iiii")
	str = tr.Search("iiiiiiiiii")
	str = tr.Search("iii")

	if "" == str {
		t.Error("")
	}
}
