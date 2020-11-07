package main

import (
	"fmt"
	"time"

	"github.com/ssargent/play/services/wiki/server"
)

func main() {
	fmt.Printf("Wiki %d\n", time.Now().Unix())
	server.Init()
}
