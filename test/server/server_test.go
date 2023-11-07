package server

import (
	"testing"

	"github.com/armnerd/zcache"
)

func TestServer(t *testing.T) {
	server := zcache.NewServer()
	server.Run()
}
