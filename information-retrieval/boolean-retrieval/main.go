package main

import (
	"fmt"
	"strings"
)

type TermDocumentMatrix struct {
	terms     []string               // List of terms
	documents []string               // List of document names
	matrix    [][]int                // Term-Document Matrix
	freqMap   map[string]map[int]int // Frequency map: term -> document -> frequency
}

func NewTermDocumentMatrix() *TermDocumentMatrix {
	return &TermDocumentMatrix{
		terms:     []string{},
		documents: []string{},
		matrix:    [][]int{},
		freqMap:   make(map[string]map[int]int),
	}
}

func (tdm *TermDocumentMatrix) AddDocument(docName string, content string) {
	// Tokenize document content
	terms := strings.Fields(content)

	// Track document
	tdm.documents = append(tdm.documents, docName)

	// Update term frequency map
	for _, term := range terms {
		if _, ok := tdm.freqMap[term]; !ok {
			tdm.terms = append(tdm.terms, term)
			tdm.freqMap[term] = make(map[int]int)
		}
		tdm.freqMap[term][len(tdm.documents)-1]++
	}

	// Update term-document matrix
	tdm.updateMatrix()
}

func (tdm *TermDocumentMatrix) updateMatrix() {
	// Initialize or resize matrix
	tdm.matrix = make([][]int, len(tdm.terms))
	for i := range tdm.matrix {
		tdm.matrix[i] = make([]int, len(tdm.documents))
	}

	// Populate matrix with frequencies
	for termIndex, term := range tdm.terms {
		for docIndex, _ := range tdm.documents {
			tdm.matrix[termIndex][docIndex] = tdm.freqMap[term][docIndex]
		}
	}
}

func (tdm *TermDocumentMatrix) BooleanQuery(query string) []string {
	// Parse the query
	query = strings.ToLower(query)
	parts := strings.Fields(query)

	// Separate terms and operators
	var terms []string
	var operators []string
	for _, part := range parts {
		if part == "and" || part == "or" {
			operators = append(operators, part)
		} else {
			terms = append(terms, part)
		}
	}

	// Find documents matching the boolean query
	var result []string
	for docIndex, _ := range tdm.documents {
		if tdm.matchesQuery(terms, operators, docIndex) {
			result = append(result, tdm.documents[docIndex])
		}
	}

	return result
}

func (tdm *TermDocumentMatrix) matchesQuery(terms []string, operators []string, docIndex int) bool {
	// Initialize boolean evaluation
	result := true

	for i, term := range terms {
		isPresent := tdm.termPresent(term, docIndex)

		if i > 0 {
			op := operators[i-1]
			switch op {
			case "and":
				result = result && isPresent
			case "or":
				result = result || isPresent
			}
		} else {
			// Default behavior for the first term
			result = result && isPresent
		}

		// term {x} is {present} in the document {y}
		fmt.Printf("term %s is `%t` in the document %s\n", term, isPresent, tdm.documents[docIndex])
	}
	fmt.Printf("> Result: %t\n", result)
	return result
}

func (tdm *TermDocumentMatrix) termPresent(term string, docIndex int) bool {
	// Check if the term is present in the document
	for termIndex, t := range tdm.terms {
		if t == term && tdm.matrix[termIndex][docIndex] > 0 {
			return true
		}
	}
	return false
}

func main() {
	tdm := NewTermDocumentMatrix()

	// Example documents
	documents := map[string]string{
		"doc1": "apple banana apple orange avocado",
		"doc2": "banana orange banana",
		"doc3": "apple orange mango",
	}

	// Add documents to the Term-Document Matrix
	for docName, content := range documents {
		tdm.AddDocument(docName, content)
	}

	// Perform boolean retrieval queries
	fmt.Println("Boolean Retrieval Results:")
	query1 := "apple AND orange"
	fmt.Printf("\nQuery: %s\n\n", query1)
	fmt.Println("Results:", tdm.BooleanQuery(query1))

	query2 := "banana OR orange"
	fmt.Printf("\nQuery: %s\n\n", query2)
	fmt.Println("Results:", tdm.BooleanQuery(query2))

	query3 := "mango OR avocado"
	fmt.Printf("\nQuery: %s\n\n", query3)
	fmt.Println("Results:", tdm.BooleanQuery(query3))
}
