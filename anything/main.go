package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()

	fmt.Println(now.Format("02.01.2006 15:04:05"))
}