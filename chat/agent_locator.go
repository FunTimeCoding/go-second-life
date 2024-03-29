package chat

import (
	"fmt"
	"github.com/funtimecoding/go-second-life/key"
)

func AgentLocator(agent key.Key) string {
	return fmt.Sprintf("secondlife:///app/agent/%s/about", agent)
}
