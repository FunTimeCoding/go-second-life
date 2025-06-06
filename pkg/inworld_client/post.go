package inworld_client

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/web_client/web_response"
	"github.com/funtimecoding/go-second-life/pkg/command"
	secondLifeWeb "github.com/funtimecoding/go-second-life/pkg/web"
	"net/http"
	"strings"
)

// Post
// In false cases the post or read failed, or the object became unreachable
//
// Do not capture known errors
func (c *Client) Post(
	locator string,
	o *command.Command,
) (bool, *web_response.Response) {
	result, e := c.web.Post(locator, o)

	if e != nil {
		text := e.Error()

		if strings.Contains(text, "i/o timeout") {
			c.postError(locator, text)
		} else if strings.Contains(text, "connection reset by peer") {
			c.postError(locator, text)
		} else if strings.Contains(text, "connection refused") {
			c.postError(locator, text)
		} else {
			c.logger.Sentry(e)
		}

		return false, result
	}

	// Avoid confusion: These are not errors happening in the inworld object
	// They are reported by the Second Life reverse proxy
	// See: http://wiki.secondlife.com/wiki/Http_response
	switch s := result.Response.StatusCode; s {
	case http.StatusOK:
		return true, result
	case http.StatusNotFound:
		// The object locator was released
		c.proxyError(locator, s)
	case http.StatusBadGateway:
		c.proxyError(locator, s)
	default:
		b := web.ReadString(result.Response)

		if s == http.StatusServiceUnavailable &&
			b == secondLifeWeb.ServiceUnavailable {
			c.proxyError(locator, s)
		} else if s == http.StatusInternalServerError &&
			b == secondLifeWeb.InternalServerError {
			c.proxyError(locator, s)
		} else {
			c.logger.Sentryf(
				"object replied bad status (%d) (%s): %s",
				s,
				locator,
				b,
			)
		}
	}

	return false, result
}
