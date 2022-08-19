package main

import "testing"

func TestNorm(t *testing.T) {
	var s1, s2 word

	s1.set("task")
	s2.set("flask")

	if 3 != rhymeValue(s1, s2) {
		t.Error("3 != task flask")
	}

	s1.set("decide")
	s2.set("code")

	if 2 != rhymeValue(s1, s2) {
		t.Error("2 != decide code")
	}

	s1.set("id")
	s2.set("void")

	if 2 != rhymeValue(s1, s2) {
		t.Error("2 != id void")
	}

	s1.set("code")
	s2.set("forces")

	if 0 != rhymeValue(s1, s2) {
		t.Error("0 != code forces")
	}
}
