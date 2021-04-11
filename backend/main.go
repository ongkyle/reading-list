/*
	Author: Kyle Ong
	Date: 10/25/2018

	backend server for reading-list application
*/
package main

import (
	//standard libary
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"net/http"

	. "ongkyle.com/reading-list/backend/datasource"
	. "ongkyle.com/reading-list/backend/services"
	. "ongkyle.com/reading-list/common"
)

type Message struct {
	SessionID, HttpMethod string
	Body                  []Item
}

func handleConnection(conn net.Conn, dataService *DataService) {
	defer conn.Close()
	decoder := gob.NewDecoder(conn)
	message := &Message{}
	decoder.Decode(message)
	fmt.Printf("Received : %+v", message)
	fmt.Println()
	encoder := gob.NewEncoder(conn)
	response := Response{}
	if message.HttpMethod == "GET" {
		response = dataService.Get(message.SessionID)
		fmt.Println()
		fmt.Println(response)
		fmt.Println()
	} else if message.HttpMethod == "SAVE" {
		response = dataService.Save(message.SessionID, message.Body)
		if response.Code != http.StatusCreated {
			log.Fatal(response)
		}
		response = dataService.Get(message.SessionID)
	}
	encoder.Encode(response)
}

func main() {
	dataService := NewDataService(Items)
	portNum := ParseListenPort("8888")
	fmt.Println("Starting backend server...")
	ln, err := net.Listen("tcp", ":"+portNum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("listening on port:", portNum)
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			log.Fatal("Setup error", err)
		}
		handleConnection(conn, dataService) // a goroutine handles conn so that the loop can accept other connections
	}
}
