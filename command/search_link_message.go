package command

import "strings"

func (c *Command) SearchLinkMessage(
	number int,
	search []string,
) int {
	for messageIndex, message := range c.L {
		if message.N == number {
			arguments := strings.Split(message.T, ",")

			if len(search) > len(arguments) {
				continue
			}

			match := true

			for argumentIndex, argument := range search {
				if arguments[argumentIndex] != argument {
					match = false

					break
				}
			}

			if match {
				return messageIndex
			}
		}
	}

	return -1
}
