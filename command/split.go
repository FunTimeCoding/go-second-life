package command

func Split(
	c *Command,
	maximumSize int,
) []Command {
	var original Command
	original.Q = append(original.Q, c.Q...)
	original.O = c.O
	var commands []Command

	for {
		var command Command
		var split bool

		if !split && len(original.Q) > 0 {
			var lastIndex int
			var sliceEmptied bool

			for i, message := range original.Q {
				command.Q = append(command.Q, message)
				size := command.Size()

				if size == maximumSize {
					split = true

					break
				} else if size > maximumSize {
					if command.RemoveLastQueuedMessage() {
						sliceEmptied = true
					}

					split = true

					break
				}

				lastIndex = i
			}

			if sliceEmptied {
				original.Q = nil
			} else {
				original.Q = original.Q[lastIndex+1:]
			}
		}

		if !split && original.O != "" {
			command.O = original.O
			size := command.Size()
			split = size >= maximumSize

			if size == maximumSize {
				break
			} else if size > maximumSize {
				command.O = ""

				break
			}

			original.O = ""
		}

		if split {
			commands = append(commands, command)
		}

		if len(original.Q) == 0 && original.O == "" {
			if !split {
				commands = append(commands, command)
			}

			break
		}
	}

	return commands
}
