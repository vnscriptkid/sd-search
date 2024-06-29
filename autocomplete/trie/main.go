package main

import (
	"fmt"
	"sort"
)

// TrieNode represents a single node in the Trie.
type TrieNode struct {
	children   map[rune]*TrieNode
	isEnd      bool
	popularity int
	word       string
}

// Trie represents the Trie structure.
type Trie struct {
	root *TrieNode
}

// NewTrieNode creates a new TrieNode.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
	}
}

// NewTrie creates a new Trie.
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// Insert inserts a word into the Trie.
func (t *Trie) Insert(word string, popularity int) {
	node := t.root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.popularity = popularity
	node.word = word
}

// Autocomplete returns a list of words with the given prefix, ordered by popularity.
func (t *Trie) Autocomplete(prefix string) []string {
	node := t.root
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return []string{}
		}
		node = node.children[ch]
	}

	var results []TrieNode
	t.dfs(node, &results)

	// Sort results by popularity in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].popularity > results[j].popularity
	})

	var words []string
	for _, result := range results {
		words = append(words, result.word)
	}

	return words
}

// dfs performs a depth-first search from the given node.
func (t *Trie) dfs(node *TrieNode, results *[]TrieNode) {
	if node.isEnd {
		*results = append(*results, *node)
	}
	for _, child := range node.children {
		t.dfs(child, results)
	}
}

func main() {
	trie := NewTrie()
	trie.Insert("hello", 5)
	trie.Insert("hell", 10)
	trie.Insert("heaven", 7)
	trie.Insert("heavy", 2)

	results := trie.Autocomplete("hea")
	for _, result := range results {
		fmt.Println(result)
	}
}
