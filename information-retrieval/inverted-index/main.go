package main

import (
	"fmt"
	"strings"
)

// InvertedIndex represents the inverted index structure
type InvertedIndex struct {
	index     map[string][]int
	documents []string
}

// buildIndex creates an inverted index from a slice of documents
func buildIndex(documents []string) *InvertedIndex {
	index := make(map[string][]int)
	for docID, document := range documents {
		terms := extractTerms(document)
		for _, term := range terms {
			if _, ok := index[term]; !ok {
				index[term] = make([]int, 0)
			}
			index[term] = append(index[term], docID)
		}
	}
	return &InvertedIndex{
		index:     index,
		documents: documents,
	}
}

// extractTerms extracts terms from a document (simple tokenizer)
func extractTerms(document string) []string {
	return strings.Fields(strings.ToLower(document))
}

// addDocument adds a new document to the index
func (idx *InvertedIndex) addDocument(document string) {
	idx.documents = append(idx.documents, document)
	docID := len(idx.documents) - 1
	terms := extractTerms(document)
	for _, term := range terms {
		if _, ok := idx.index[term]; !ok {
			idx.index[term] = make([]int, 0)
		}
		idx.index[term] = append(idx.index[term], docID)
	}
}

// queryIndex retrieves documents containing a specific term
func (idx *InvertedIndex) queryIndex(term string) []string {
	if docIDs, ok := idx.index[term]; ok {
		results := make([]string, 0)
		for _, docID := range docIDs {
			results = append(results, idx.documents[docID])
		}
		return results
	}
	return nil
}

func main() {
	// Initialize index with initial documents
	initialDocuments := []string{
		"This is the first document.",
		"Second document is here.",
		"And this is the third one.",
		"Is this the first document?",
	}
	index := buildIndex(initialDocuments)

	// Add a new document
	newDocument := "A new document added dynamically."
	index.addDocument(newDocument)

	fmt.Printf("index: %v\n\n", index.index)

	// Query example
	searchTerm := "document"
	results := index.queryIndex(searchTerm)
	fmt.Printf("Documents containing '%s':\n", searchTerm)
	for i, result := range results {
		fmt.Printf("%d: %s\n", i+1, result)
	}
}
