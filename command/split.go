package command

func Split(
	c *Command,
	maximumSize int,
) []Command {
	var original Command
	original.L = append(original.L, c.L...)
	original.O = c.O
	var commands []Command

	for {
		var command Command
		var split bool

		if len(original.L) > 0 {
			var lastIndex int
			var sliceEmptied bool

			for i, message := range original.L {
				command.L = append(command.L, message)
				size := command.Size()

				if size == maximumSize {
					split = true

					break
				} else if size > maximumSize {
					if command.RemoveLast() {
						sliceEmptied = true
					}

					split = true

					break
				}

				lastIndex = i
			}

			if sliceEmptied {
				original.L = nil
			} else {
				original.L = original.L[lastIndex+1:]
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

		if len(original.L) == 0 && original.O == "" {
			if !split {
				commands = append(commands, command)
			}

			break
		}
	}

	return commands
}
