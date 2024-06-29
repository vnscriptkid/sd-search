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
		matrix:    [][]int{}, // Value of matrix represents frequency of term in document
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

func main() {
	tdm := NewTermDocumentMatrix()

	// Example documents
	documents := map[string]string{
		"doc1": "apple banana apple orange mango",
		"doc2": "banana orange banana",
		"doc3": "apple orange",
	}

	// Add documents to the Term-Document Matrix
	for docName, content := range documents {
		tdm.AddDocument(docName, content)
	}

	// Print the Term-Document Matrix
	fmt.Println("Terms:", tdm.terms)
	fmt.Println("Documents:", tdm.documents)
	fmt.Println("Frequency Map:", tdm.freqMap)
	fmt.Println("Term-Document Matrix:")
	for i := range tdm.terms {
		fmt.Printf("%-10s", tdm.terms[i])
		for j := range tdm.documents {
			fmt.Printf("%-4d", tdm.matrix[i][j])
		}
		fmt.Println()
	}
}
