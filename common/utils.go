/* 
	Author: Kyle Ong
	Date: 10/25/2018

	utils for frontend server

	todo
	- [ ] remove duplicated utils on the front end and backend (pretty sure this is really hard)
*/
package common

import (
	"flag"
	"bytes"
)

func ParseListenPort(defaultVal string) string {
	portPtr := flag.String("listen", defaultVal, "Server Port")
	return *portPtr
}

func ParseBackendHost() string {
	backendHostPtr := flag.String("backend", "8888", "Backend Server Port")
	var buffer bytes.Buffer;
	buffer.WriteString("localhost:")
	buffer.WriteString(*backendHostPtr)
	return buffer.String()
}
