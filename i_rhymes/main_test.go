package main

import "testing"

func TestAb(t *testing.T) {

	tr := NewTrie()
	tr.Insert("ve")
	tr.Insert("neeevzzv")

	tr.SearchWord("v")          //neeevzzv
	tr.SearchWord("vnveneezve") //ve
	tr.SearchWord("evzz")
	tr.SearchWord("ezenne") //ve
	tr.SearchWord("nneenn")
	//t.Error()
}
