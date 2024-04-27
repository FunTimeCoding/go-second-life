package inworld_client

import (
	"github.com/funtimecoding/go-library/pkg/face"
	secondLife "github.com/funtimecoding/go-second-life/face"
)

type Client struct {
	clock  face.Clock
	logger secondLife.AgentLogger
}
