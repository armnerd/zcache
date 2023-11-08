package main

import (
	"github.com/armnerd/zcache"
)

func main() {
	client := zcache.NewClient(
		zcache.WithAddr("127.0.0.1"),
		zcache.WithPort(8999),
	)
	client.Run()
}
