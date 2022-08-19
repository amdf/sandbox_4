package main

import (
	"bytes"
	"testing"
)

func TestMax(t *testing.T) {
	var s word

	s.set("id")

	if 1 != s.maxrhyme() {
		t.Error("TestMax wrong")
	}

	s.set("abc")

	if 2 != s.maxrhyme() {
		t.Error("TestMax wrong")
	}

	s.set("aaaaabbbbb")

	if 9 != s.maxrhyme() {
		t.Error("TestMax wrong")
	}
}

func TestNorm1(t *testing.T) {
	var buf, outbuf bytes.Buffer
	buf.WriteString(
		"3\n" +
			"task\n" +
			"decide\n" +
			"id\n" +
			"6\n" +
			"flask\n" +
			"code\n" +
			"void\n" +
			"forces\n" +
			"id\n" +
			"ask\n")

	processing(&buf, &outbuf)

	str := outbuf.String()

	if 0 == len(str) {
		t.Error("TestNorm1 wrong")
	}

}

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

	s1.set("id")
	s2.set("id")

	if v, eq := rhymeValue(s1, s2); v != 2 || eq != true {
		t.Error("id equals")
	}

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
