package main

import "testing"

func TestRhymes(t *testing.T) {
	var s1, s2 word

	s1.set("task")
	s2.set("flask")

	if v, _ := rhymeValue(s1, s2); v != 3 {
		t.Error("3 != task flask")
	}

	s1.set("decide")
	s2.set("code")

	if v, _ := rhymeValue(s1, s2); v != 2 {
		t.Error("2 != decide code")
	}

	s1.set("id")
	s2.set("void")

	if v, _ := rhymeValue(s1, s2); v != 2 {
		t.Error("2 != id void")
	}

	s1.set("code")
	s2.set("forces")

	if v, _ := rhymeValue(s1, s2); v != 0 {
		t.Error("0 != code forces")
	}
}

func TestEqual(t *testing.T) {
	var s1, s2 word

	s1.set("flaskflask")
	s2.set("flaskflask")

	if v, eq := rhymeValue(s1, s2); v != 10 || eq != true {
		t.Error("equals1")
	}

	s1.set("laskflask")
	s2.set("flaskflask")

	if v, eq := rhymeValue(s1, s2); v != 9 || eq != false {
		t.Error("equals2")
	}
}
