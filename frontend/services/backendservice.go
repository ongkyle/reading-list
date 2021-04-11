
/* 
	Author: Kyle Ong
	Date: 10/25/2018

	backendservice for frontend server
*/
package backendservice

import (
	// standard
	"encoding/gob"
	"fmt"
	"log"
	"net" // user-defined

	. "ongkyle.com/reading-list/common"
)

type Service interface {
	Get(owner string) []Item
	Save(owner string, newItems []Item) error
}

type Message struct {
	SessionID, HttpMethod string
	Body                  []Item
}

type DataService struct {
	HostName string
}

func NewDataService(hostName string) *DataService {
	return &DataService{
		HostName: hostName}
}

func (s *DataService) Get(sessionOwner string) Response {
	message := Message{
		SessionID:  sessionOwner,
		HttpMethod: "GET",
	}

	conn, err := net.Dial("tcp", s.HostName)
	if err != nil {
		log.Fatal("Connection error", err)
	} else {
		fmt.Println("Connection Established!")
	}

	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)

	decoder := gob.NewDecoder(conn)
	response := &Response{}
	decoder.Decode(response)

	log.Println(response)

	return *response
}

func (s *DataService) Save(sessionOwner string, newItems []Item) Response {
	message := Message{
		SessionID:  sessionOwner,
		HttpMethod: "SAVE",
		Body:       newItems,
	}

	conn, err := net.Dial("tcp", s.HostName)
	if err != nil {
		log.Fatal("Connection error", err)
	} else {
		fmt.Println("Connection Established")
	}

	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)

	decoder := gob.NewDecoder(conn)
	response := &Response{}
	decoder.Decode(response)

	log.Println(response)

	return *response
}
