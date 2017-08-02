package utils

import "fmt"

func channel_hold() {
	c := make(chan string)
	for {
		log := <- c
		fmt.Println(log)
	}
}
