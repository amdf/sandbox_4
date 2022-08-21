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
	//this is a single letter stored for example letter a,b,c,d,etc
	Char string
	//store all children  of a node
	//that is from a-z
	//a slice of Nodes(and each child will also have 26 children)
	Children [26]*Node
}

/// NewNode this will be used to initialize a new node with 26 children
///each child should first be initialized to nil
func NewNode(char string) *Node {
	node := &Node{Char: char}
	for i := 0; i < 26; i++ {
		node.Children[i] = nil
	}
	return node
}

// Trie  is our actual tree that will hold all of our nodes
//the Root node will be nil
type Trie struct {
	RootNode *Node
}

// NewTrie Creates a new trie with a root('constructor')
func NewTrie() *Trie {
	//we will not use this node so
	//it can be anything
	root := NewNode("\000")
	return &Trie{RootNode: root}
}

/// Insert inserts a word to the trie
func (t *Trie) Insert(word string) {
	///this will keep track of our current node
	///when transversing our tree
	///it should always start at the top of our tree
	///i.e our root
	current := t.RootNode

	strippedWord := reverse(word)
	for i := 0; i < len(strippedWord); i++ {
		//from the ascii table a represent decimal number 97
		//from the ascii table b represent decimal number 98
		//from the ascii table c represent decimal number 99
		/// and so on so basically if we were to say  98-97=1 which means the index of b is 1 and for  c is 99-97
		///that what is happening below (we are taking the decimal representation of a character and subtracting decimal representation of a)
		index := strippedWord[i] - 'a'
		///check if current already has a node created at our current node
		//if not create the node
		if current.Children[index] == nil {
			current.Children[index] = NewNode(string(strippedWord[i]))
		}
		current = current.Children[index]
		//since we want to support autocomplete
	}

}

func (t *Trie) SearchWord(word string) (result string) {
	if debug {
		fmt.Println("Searching", word)
	}
	strippedWord := reverse(word)

	current := t.RootNode

	var b strings.Builder

	treeEnd := false
	i := 0
	for !treeEnd {
		index := byte(0)
		if i < len(strippedWord) {
			index = strippedWord[i] - 'a'
			i++
		}

		found := false

		if nil == current.Children[index] {
			for j := byte(0); j < 26; j++ {
				if nil != current.Children[j] {

					b.WriteByte('a' + j)
					current = current.Children[j]

					found = true
					break
				}
			}
		} else {
			b.WriteByte('a' + index)
			current = current.Children[index]

			found = true
		}

		treeEnd = !found
	}

	result = reverse(b.String())

	if debug && result == word {
		fmt.Println("Searching", word, " got ", result)
	}

	for i := 0; (result == word) && i < len(dict_s); i++ {
		result = dict_s[i]
	}

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
		res[i] = tr.SearchWord(words[i])
	}

	// uniqwords = len(uniqw)

	// res := parallel(words, result)

	for i := range res {
		fmt.Fprintln(w, res[i])
	}
}

// func parallel(words []word, result map[word]string) []string {
// 	var mut sync.RWMutex
// 	var mut2 sync.Mutex

// 	r := make([]string, len(words))

// 	proc := func(wg *sync.WaitGroup, part []word, res []string) {

// 		for i := range part {
// 			mut.RLock()
// 			s, ok := result[part[i]]
// 			mut.RUnlock()
// 			if !ok {
// 				s = maxRhymeWord(part[i])
// 				mut.Lock()
// 				result[part[i]] = s
// 				mut.Unlock()
// 			}
// 			mut2.Lock()
// 			res[i] = s
// 			mut2.Unlock()
// 		}
// 		wg.Done()
// 	}
// 	var wg sync.WaitGroup

// 	packsize := len(words) / 4

// 	for endOffset := len(words); endOffset > 0; endOffset -= packsize {
// 		beginOffset := endOffset - packsize
// 		if beginOffset < 0 {
// 			beginOffset = 0
// 		}
// 		wg.Add(1)
// 		threads++
// 		go proc(&wg, words[beginOffset:endOffset], r[beginOffset:endOffset])
// 	}

// 	wg.Wait()

// 	return r
// }

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
		// fmt.Printf("uniqwords: %v\n", uniqwords)
		// fmt.Printf("result size: %v\n", len(result))
	}
}
