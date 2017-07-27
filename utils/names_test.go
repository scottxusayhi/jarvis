package utils

import (
	"testing"
	"fmt"
)

func TestRandomDataCenter(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(UnknownDataCenter())
	}
}

func TestRandomRack(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(UnknownRack())
	}
}

func TestRandomSlot(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(UnknownSlot())
	}
}



