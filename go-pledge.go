package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hallo Welt…")
	for i := 0; i < 10; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println("pledge?…")
	}
}
