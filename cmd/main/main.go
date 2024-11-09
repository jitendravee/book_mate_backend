package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"github.com/jitendra/book_mate_backend/internals/config"
)

var client *firestore.Client
var ctx context.Context

func main() {
	// Initialize Firestore client and context
	ctx = context.Background()
	client = config.FireBaseConfig()
	defer config.CloseFirestoreClient() // Properly close the Firestore client

	// Set up router
	router := mux.NewRouter()

	// Define POST route for creating a Firestore document
	router.HandleFunc("/api/create-document", createDocumentHandler).Methods("POST")

	// Start server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// DocumentData represents the data structure of the document to be saved
type DocumentData struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// createDocumentHandler handles the POST request to create a Firestore document
func createDocumentHandler(w http.ResponseWriter, r *http.Request) {
	var data DocumentData

	// Parse JSON body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Create a document in Firestore
	docRef := client.Collection("test_collection").NewDoc() // Generates a new doc with random ID
	_, err = docRef.Set(ctx, map[string]interface{}{
		"title":       data.Title,
		"description": data.Description,
		"created_at":  time.Now(),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create document: %v", err), http.StatusInternalServerError)
		return
	}

	// Return a success response with the document ID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Document created successfully",
		"doc_id":  docRef.ID,
	})
}
