package chat

import "fmt"

func Link(
	channel int,
	message string,
	display string,
) string {
	return fmt.Sprintf(
		"[secondlife:///app/chat/%d/%s %s]",
		channel,
		message,
		display,
	)
}
