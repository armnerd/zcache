package client

import (
	"testing"

	"github.com/armnerd/zcache"
)

func TestClient(t *testing.T) {
	client := zcache.NewClient()
	client.Run()
}
