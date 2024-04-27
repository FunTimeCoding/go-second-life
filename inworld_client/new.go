package inworld_client

import (
	"github.com/funtimecoding/go-library/pkg/face"
	secondLife "github.com/funtimecoding/go-second-life/face"
)

func New(
	c face.Clock,
	l secondLife.AgentLogger,
) *Client {
	return &Client{clock: c, logger: l}
}
