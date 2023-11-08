package server

import (
	"testing"

	"github.com/armnerd/zcache"
)

func TestServer(t *testing.T) {
	server := zcache.NewServer(
		zcache.WithCleanSeq(5),
		zcache.WithLandSeq(5),
	)
	server.Run()
}
