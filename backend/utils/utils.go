/* 
	Author: Kyle Ong
	Date: 10/25/2018

	utilities for backend server
*/

package utils

import (
	"bytes"
	"strings"
)

func ParseListenPort(args []string) string {
	portNum := "8090"
	for idx, ele := range args {
		if ele == "--listen" && idx < (len(args)-1) {
			portNum = args[idx+1]
		}
	}
	return portNum
}

func ParseBackendHost(args []string) string {
	backend := "localhost:8090"
	for idx, ele := range args {
		if ele == "--backend" && idx < len(args)-1 {
			backend = args[idx+1]
		}
	}
	if len(strings.Split(backend, ":")) == 1 {
		var buffer bytes.Buffer
		buffer.WriteString("localhost")
		buffer.WriteString(backend)
		backend = buffer.String()
	}
	return backend
}
