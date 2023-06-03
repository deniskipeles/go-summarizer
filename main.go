package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/jaytaylor/html2text"
	"github.com/james-bowman/nlp"
)

// RequestBody represents the JSON payload for text summarization request.
type RequestBody struct {
	Text string `json:"text"`
}

// ResponseBody represents the JSON response containing the summarized text.
type ResponseBody struct {
	Summary string `json:"summary"`
}

// summarizeText performs text summarization using the TextRank algorithm.
func summarizeText(text string) string {
	// Convert HTML to plain text
	plainText, err := html2text.FromString(text)
	if err != nil {
		log.Println("Error converting HTML to plain text:", err)
		return ""
	}

	// Tokenize the text
	tokenizer := nlp.NewSentenceTokenizer(nil)
	sentences := tokenizer.Tokenize(strings.NewReader(plainText))

	// Build a TextRank model
	model := nlp.NewTextRank()
	model.StopWords = nlp.DefaultStopWords()

	// Add the sentences to the model
	for _, sentence := range sentences {
		model.AddSentence(sentence)
	}

	// Calculate the TextRank scores
	model.Rank()

	// Get the top-ranked sentences
	topSentences := model.RankedSentences(5) // Adjust the number of sentences as needed

	// Combine the top sentences to form the summary
	var summary strings.Builder
	for _, sentence := range topSentences {
		summary.WriteString(sentence.Value())
		summary.WriteString(" ")
	}

	return summary.String()
}

// summarizeHandler handles the text summarization request.
func summarizeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Summarize the text
	summary := summarizeText(reqBody.Text)

	// Create the response body
	resBody := ResponseBody{
		Summary: summary,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resBody)
	if err != nil {
		log.Println("Error encoding response body:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// Set up the HTTP server
	http.HandleFunc("/summary", summarizeHandler)
	http.HandleFunc("/", helloHandler)

	// Start the server
	fmt.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
