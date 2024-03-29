package input

import "github.com/funtimecoding/go-second-life/key"

type Input struct {
	User key.Key

	// Avatar sitting on the object of the input
	AvatarSitting key.Key

	// Attach point of the object of the input
	AttachPoint int

	Menu          string
	ButtonClicked string
	TextEntered   string
	ObjectTouched string

	TimedOut bool
}
