//
// An echo service that receive messages from client and send back exactly the same one
//

package main

import (
	"net"
	"fmt"
	"os"
	//"time"
	"io"
	"time"
)

func Listen(listener net.Listener, c chan string) {
	c <- fmt.Sprintf("listener goroutine in %v", os.Getpid())
	// Listen for incoming connection
	conn, err := listener.Accept()

	// a timeout error
	// if err, ok := err.(*net.OpError); ok && err.Timeout() {}

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("connected localAddr=%v remoteAddr=%v\n", conn.LocalAddr(), conn.RemoteAddr())
	defer conn.Close()

	// send welcome message
	conn.Write([]byte(fmt.Sprintf("welcome %v", conn.RemoteAddr())))

	content := make([]byte, 100)
	for {
		// clean input array for new request
		for c := range content {
			content[c] = 0
		}
		// read request
		n, err := conn.Read(content)
		if err == io.EOF {
			fmt.Printf("Connection closed. remoteAddr=%v\n", conn.RemoteAddr())
			break
		} else if err != nil {
			fmt.Println(err.Error())
			//fmt.Println("error reading request: " + err.Error())
		} else {
			fmt.Println(string(content[:n]))
			// send response - the same content
			conn.Write(content)
		}
	}
	c <- "Connection closed"
}

func HanleConnection(conn net.Conn, c chan string) {
	fmt.Printf("connected localAddr=%v remoteAddr=%v\n", conn.LocalAddr(), conn.RemoteAddr())
	defer conn.Close()

	// send welcome message
	conn.Write([]byte(fmt.Sprintf("welcome %v", conn.RemoteAddr())))

	content := make([]byte, 100)
	for {
		// clean input array for new request
		for c := range content {
			content[c] = 0
		}
		// read request
		n, err := conn.Read(content)
		if err == io.EOF {
			fmt.Printf("Connection closed. remoteAddr=%v\n", conn.RemoteAddr())
			break
		} else if err != nil {
			fmt.Println(err.Error())
			//fmt.Println("error reading request: " + err.Error())
		} else {
			fmt.Println(string(content[:n]))
			// send response - the same content
			conn.Write(content)
		}
	}
	c <- "Connection closed"
}

func main() {
	fmt.Printf("Main thread in pid %v\n", os.Getpid())
	c := make(chan string)
	listener, err := net.Listen("tcp", ":2999")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Server started")

	// exp: sleep some time and accept connections
	time.Sleep(5*time.Second)


	// print messages in channel
	go func(c chan string) {
		for {
			fmt.Println(<-c)
		}
	} (c)

	// Listen for incoming connection
	for {
		conn, err := listener.Accept()
		// a timeout error
		// if err, ok := err.(*net.OpError); ok && err.Timeout() {}
		if err != nil {
			fmt.Println(err.Error())
		}
		go HanleConnection(conn, c)
	}

	// read from channel forever
	for {
		fmt.Println(<-c)
	}
}
