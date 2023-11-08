package server

import (
	"testing"
	"time"

	"github.com/armnerd/zcache"
)

func TestServer(t *testing.T) {
	server := zcache.NewServer(
		zcache.WithCleanSeq(60*time.Second),
		zcache.WithLandSeq(60*time.Second),
	)
	server.Run()
}
