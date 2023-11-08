package main

import (
	"github.com/armnerd/zcache"
)

func main() {
	client := zcache.NewClient()
	client.Run()
}
