package inworld_client

func (c *Client) proxyError(
	locator string,
	status int,
) {
	c.logger.Agent("proxy error %d %s", status, locator)
}
