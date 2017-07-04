package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"io"
)

func heartBeat(conn net.Conn, interval time.Duration) {
	fmt.Println("starting heart beat ...")
	for {
		msg := fmt.Sprintf("I am alive ... %v", time.Now())
		fmt.Println(msg)
		//writer.Write([]byte(msg))
		conn.Write([]byte(msg))
		time.Sleep(interval*time.Second)
	}
}

func Listen(conn net.Conn) error {
	fmt.Println("starting reader")
	response := make([]byte, 100)
	for {
		for index := range response {
			response[index] = 0
		}

		n, err := conn.Read(response)
		if err == io.EOF {
			fmt.Println("Connection closed")
			return err
		} else if err != nil {
			fmt.Println(err.Error())
			return err
		}
		//fmt.Println(response)
		fmt.Println(string(response[:n]))
	}
}


func KeyboardInput(conn net.Conn) error {
	reader := bufio.NewReader(os.Stdin)
	// read from stdin
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	// request - send something
	_, err := conn.Write([]byte(strings.Trim(text, "\n")))
	if err == io.EOF {
		fmt.Println("Connection closed")
		return err
	} else if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func channel_hold() {
	c := make(chan string)
	for {
		log := <- c
		fmt.Println(log)
	}
}

func main() {
	// connect
	conn, err := net.DialTimeout("tcp", "localhost:2999", 3*time.Second)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Printf("Connected. %v %v\n", conn.LocalAddr(), conn.RemoteAddr())

	// heart beat
	go heartBeat(conn, 10)
	go Listen(conn)

	KeyboardInput(conn)
	channel_hold()

}
