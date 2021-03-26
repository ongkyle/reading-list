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
		response = NewResponse(200, items)
		fmt.Println(response)
		fmt.Println()
	} else {
		response = NewResponse(404, make([]Item, 0))
	}
	s.mu.RUnlock()
	fmt.Println()
	fmt.Println(response)
	fmt.Println()
	return response
}

func (s *DataService) Save(sessionOwner string, newItems []Item) error {
	var prevID int64
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
	return nil
}
