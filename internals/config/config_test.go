package config

import (
	"testing"
)

func TestFireBaseConfig(t *testing.T) {
	client := FireBaseConfig()

	_, err := client.Collection("test_collection").Doc("test_document").Create(ctx, map[string]interface{}{
		"sampleField": "samleVale",
		"hello":       "jkl",
	})
	if err != nil {
		t.Fatalf("Failed to create test document: %v", err)
	}
	doc, err := client.Collection("test_collection").Doc("test_document").Get(ctx)
	if err != nil {
		t.Errorf("Failed to retrieve document: %v", err)
	} else {
		t.Logf("Document data: %v", doc.Data())
	}

	CloseFirestoreClient()
}
