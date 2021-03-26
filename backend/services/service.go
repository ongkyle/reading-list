/*
	Author: Kyle Ong
	Date: 10/25/2018

	backend service for reading list application
*/
package service

import (
	// standard
	"fmt"
	"sync"
	"net/http"

	// user-defined
	. "ongkyle.com/reading-list/common"
)

type Service interface {
	Get(owner string) Response
	Save(owner string, newItems []Item) error
}

type DataService struct {
	items map[string][]Item
	mu    sync.RWMutex
}

func NewDataService(source map[string][]Item) *DataService {
	return &DataService{
		items: source}
}

func (s *DataService) Get(sessionOwner string) Response {
	var response Response
	fmt.Println(s.items)
	s.mu.RLock()
	if _, exists := s.items["test"]; exists {
		testTasks := s.items["test"]
		delete(s.items, "test")
		for i := 0; i < len(testTasks); i++ {
			testTasks[i].SessionID = sessionOwner
			s.items[sessionOwner] = append(s.items[sessionOwner], testTasks[i])
		}
	}

	if _, exists := s.items[sessionOwner]; exists {
		items := s.items[sessionOwner]
		fmt.Println(items)
		fmt.Println()
		response = NewResponse(http.StatusOK, items)
		fmt.Println(response)
		fmt.Println()
	} else {
		response = NewResponse(http.StatusNotFound, EmptyItems())
	}
	s.mu.RUnlock()
	fmt.Println()
	fmt.Println(response)
	fmt.Println()
	return response
}

func (s *DataService) Save(sessionOwner string, newItems []Item) Response {
	var prevID int64
	var response Response = NewResponse(http.StatusInternalServerError, EmptyItems())
	for i := range newItems {
		if newItems[i].ID == 0 {
			newItems[i].ID = prevID
			prevID++
		}
	}

	s.mu.Lock()
	s.items[sessionOwner] = newItems
	s.mu.Unlock()
	fmt.Println(s.items)
	response = NewResponse(http.StatusCreated, EmptyItems())
	return response
}
