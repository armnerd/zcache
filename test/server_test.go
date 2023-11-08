package server

import (
	"testing"
	"time"

	"github.com/armnerd/zcache"
)

func TestServer(t *testing.T) {
	server := zcache.NewServer(
		zcache.WithCleanSeq(5*time.Second),
		zcache.WithLandSeq(5*time.Second),
	)
	server.Run()
}
