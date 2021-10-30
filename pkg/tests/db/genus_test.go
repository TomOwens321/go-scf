package database

import	"testing"
import  database "tomo/go-scf/pkg/db"

func TestGenusName(t *testing.T) {
	testName := "testGenus"
	genus := database.Genus{Name: testName}
	if genus.Name != testName {
		t.Errorf("Failed to create a genus with name %s", testName)
	}
}

func TestGenusDescription(t *testing.T) {
	testDesc := "Lovely test"
	genus := database.Genus{Description: testDesc}
	if genus.Description != testDesc {
		t.Errorf("The description does not match")
	}
}