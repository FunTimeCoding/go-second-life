package inworld_client

import (
	"github.com/funtimecoding/go-library/pkg/face"
	secondLife "github.com/funtimecoding/go-second-life/face"
)

type Client struct {
	web    face.WebClient
	logger secondLife.AgentLogger
}
