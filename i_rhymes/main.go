package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var debug bool

func reverse(s string) string {
	var byte strings.Builder
	byte.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		byte.WriteByte(s[i])
	}
	return byte.String()
}

//Node represent each character
type Node struct {
	Char     byte
	Children [27]*Node
	Prev     *Node
	Word     bool
}

func (node Node) HasChildren() bool {
	for i := 0; i < 27; i++ {
		if nil != node.Children[i] {
			return true
		}
	}
	return false
}
func (node Node) CountChildren() (result int) {
	for i := 0; i < 27; i++ {
		if nil != node.Children[i] {
			result++
		}
	}
	return
}

func NewNode(char byte) *Node {
	node := &Node{Char: char}
	for i := 0; i < 27; i++ {
		node.Children[i] = nil
	}
	return node
}

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	root := NewNode('\x00')
	return &Trie{RootNode: root}
}

/// Insert inserts a word to the trie
func (t *Trie) Insert(s string) {
	current := t.RootNode

	word := reverse(s)
	//word := s
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'

		if current.Children[index] == nil {
			current.Children[index] = NewNode(word[i])
			current.Children[index].Prev = current
		}
		current = current.Children[index]
	}
	current.Word = true
}

func (t *Trie) SearchWord(s string) bool {
	word := reverse(s)
	//word := s
	current := t.RootNode
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'

		if current == nil || current.Children[index] == nil {
			return false
		}
		current = current.Children[index]
	}
	return current.Word
}

func (t *Trie) DeleteWord(s string) {
	word := reverse(s)
	//word := s
	current := t.RootNode

	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'

		if current == nil || current.Children[index] == nil {
			return
		}
		current = current.Children[index]
	}

	i := len(word) - 1

	allow := true

	for {
		current = current.Prev
		if nil == current {
			break
		}
		count := current.CountChildren()
		if i >= 0 {
			if current.Children[word[i]-'a'].Word {
				if allow {
					allow = false
				} else {
					break
				}
			}
			current.Children[word[i]-'a'] = nil

		} else {
			break
		}

		if count > 1 {
			break
		}
		i--
	}
}

func (t *Trie) Search(word string) (result string) {
	result = t.SearchMore(word)
	if result == word {
		t.DeleteWord(word)
		result = t.SearchMore(word)
		t.Insert(word)
		if "" == result {
			result = word
			for i := 0; (result == word) && i < len(dict_s); i++ {
				result = dict_s[i]
			}

		}
	}

	return
}

func (t *Trie) SearchMore(s string) (result string) {

	word := reverse(s)
	//word := s

	current := t.RootNode

	var b strings.Builder

	treeEnd := false
	i := 0
	for !treeEnd {
		index := byte(0)

		if i < len(word) {
			index = word[i] - 'a'
		}

		found := false

		if nil != current.Children[index] {

			b.WriteByte('a' + index)
			current = current.Children[index]

			found = true

		}

		if !found {
			for j := byte(0); j < 27; j++ {
				if nil != current.Children[j] {

					b.WriteByte('a' + j)
					current = current.Children[j]

					found = true
					break

				}
			}
		}

		treeEnd = !found
		i++
	}

	result = reverse(b.String())

	return
}

var dictSize, wordCount int64
var threads int

//uniqwords int

var dict_s []string

func processing(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)

	if !sc.Scan() {
		return
	}

	tr := NewTrie()
	dictSize, _ = strconv.ParseInt(sc.Text(), 10, 64)

	for i := 0; i < int(dictSize); i++ {
		if !sc.Scan() {
			break
		}
		s := sc.Text()
		tr.Insert(s)
		dict_s = append(dict_s, s)
	}

	if !sc.Scan() {
		return
	}
	wordCount, _ = strconv.ParseInt(sc.Text(), 10, 64)

	uniqw := make(map[string]struct{})
	words := make([]string, wordCount)

	for i := 0; i < int(wordCount); i++ {
		if !sc.Scan() {
			break
		}
		s := sc.Text()

		words[i] = s
		uniqw[s] = struct{}{}
	}

	res := make([]string, wordCount)
	for i := range words {
		res[i] = tr.Search(words[i])
	}

	for i := range res {
		fmt.Fprintln(w, res[i])
	}
}

func main() {
	t := time.Now()
	if len(os.Args) > 1 {
		debug = true
	}
	processing(os.Stdin, os.Stdout)

	if len(os.Args) > 1 {
		fmt.Println()
		fmt.Println(time.Since(t))
		fmt.Printf("threads: %v\n", threads)
		fmt.Printf("dictSize: %v\n", dictSize)
		fmt.Printf("wordCount: %v\n", wordCount)

	}
}
