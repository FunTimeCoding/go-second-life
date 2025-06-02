package inworld_client

import (
	"github.com/funtimecoding/go-library/pkg/face"
	secondLife "github.com/funtimecoding/go-second-life/pkg/face"
)

func New(
	w face.WebClient,
	l secondLife.AgentLogger,
) *Client {
	return &Client{web: w, logger: l}
}
