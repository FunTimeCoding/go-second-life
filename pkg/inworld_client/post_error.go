package inworld_client

func (c *Client) postError(
	locator string,
	text string,
) {
	c.logger.Agent("post error (%s): %s", locator, text)
}
