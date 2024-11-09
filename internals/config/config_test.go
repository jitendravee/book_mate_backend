package config

import (
	"testing"
)

func TestFireBaseConfig(t *testing.T) {
	fireBaseConfig() // Initialize Firestore client

	// Ensure the test document exists
	hello, err := client.Collection("test_collection").Doc("test_document").Set(ctx, map[string]interface{}{
		"sampleField": "sampleVale",
		"hello":       "jkjl",
	})
	if err != nil {
		t.Fatalf("Failed to create test document: %v", err)
	}
	t.Fatalf("hkjhfkjs: %v", hello)
	// Now attempt to retrieve it
	doc, err := client.Collection("test_collection").Doc("test_document").Get(ctx)
	if err != nil {
		t.Errorf("Failed to retrieve document: %v", err)
	} else {
		t.Logf("Document data: %v", doc.Data())
	}

	CloseFirestoreClient()
}
