package utils

import (
	"testing"
	"fmt"
)

func TestRandomDataCenter(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomDataCenter())
	}
}

func TestRandomRack(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomRack())
	}
}

func TestRandomSlot(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandomSlot())
	}
}



