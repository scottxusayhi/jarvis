package utils

import (
	"testing"
	"fmt"
	"crypto/rand"
	"time"
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

func TestMisc(t *testing.T) {
	fmt.Println(UnknownDataCenter())
	fmt.Println(UnknownRack())
	fmt.Println(UnknownSlot())

	fmt.Println(UnknownDataCenter())
	fmt.Println(UnknownRack())
	fmt.Println(UnknownSlot())
}

func TestRand(t *testing.T) {
    n := 5
    b := make([]byte, n)
    if _, err := rand.Read(b); err != nil {
        panic(err)
    }
    s := fmt.Sprintf("%X", b)
    fmt.Println(s)
}

func TestLongLongAgo(t *testing.T) {
	fmt.Println(time.Time{})
	//fmt.Println(LongLongAgo())
}







