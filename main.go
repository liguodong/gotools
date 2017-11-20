package main

import (
	"fmt"
	"gotools/tools"
	"time"
)

func main() {
	fmt.Println("running")
	tools.SendMail(1, 1, 10 * time.Second)
	for true {}
}