package service

import (
	// standard lib
	"fmt"
	"net/http"
	"testing"

	// user defined
	"ongkyle.com/reading-list/backend/datasource"
	"ongkyle.com/reading-list/common"
)

func TestNewDataService(t *testing.T) {
	t.Parallel()
	dataService := NewDataService(datasource.Items)
	if _, exists := dataService.items["test"]; !exists {
		message := fmt.Sprintf("Data Service does not have key 'test'")
		t.Error(message)
	}
	testItems := dataService.items["test"]
	for _, item := range testItems {
		if item.SessionID != "test" {
			message := fmt.Sprintf("Expected: test. got: %s", item.SessionID)
			t.Error(message)
		}
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	dataService := NewDataService(datasource.Items)
	testResponse := dataService.Get("test")
	for _, item := range testResponse.Data {
		if item.SessionID != "test" {
			message := fmt.Sprintf("Expected: test. got: %s", item.SessionID)
			t.Error(message)
		}
	}
}

func TestSave(t *testing.T) {
	t.Parallel()
	dataService := NewDataService(datasource.Items)
	saveItem := common.Item{
		SessionID: "saved_sessID",
		Title:     "saved_title",
		Completed: true}
	saveItems := []common.Item{
		saveItem,
	}
	response := dataService.Save("test", saveItems)
	if response.Code != http.StatusCreated {
		t.Error(response)
	}
}
