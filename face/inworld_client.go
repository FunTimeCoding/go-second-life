package face

import (
	"github.com/funtimecoding/go-library/pkg/web/web_client/web_response"
	"github.com/funtimecoding/go-second-life/command"
)

type InworldClient interface {
	IsInworldClient()

	Post(
		locator string,
		o *command.Command,
	) (bool, *web_response.Response)
}
