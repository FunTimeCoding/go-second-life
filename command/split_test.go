package command

import (
	"github.com/funtimecoding/go-library/assert"
	"github.com/funtimecoding/go-second-life/link_message"
	"testing"
)

func TestSplitEmpty(t *testing.T) {
	assert.Any(
		t,
		[]Command{{}},
		Split(&Command{}, 10),
	)
}

func TestSplitNoSplit(t *testing.T) {
	assert.Any(
		t,
		[]Command{{L: []link_message.Message{{T: "a"}}}},
		Split(
			&Command{L: []link_message.Message{{T: "a"}}},
			100,
		),
	)
}

func TestSplitNoSplitQueued(t *testing.T) {
	assert.Any(
		t,
		[]Command{{Q: []link_message.Message{{T: "a"}}}},
		Split(
			&Command{Q: []link_message.Message{{T: "a"}}},
			100,
		),
	)
}

func TestSplit(t *testing.T) {
	assert.Any(
		t,
		[]Command{
			{L: []link_message.Message{{T: "a"}}},
			{L: []link_message.Message{{T: "b"}}},
		},
		Split(
			&Command{L: []link_message.Message{{T: "a"}, {T: "b"}}},
			50,
		),
	)
}

func TestSplitQueued(t *testing.T) {
	assert.Any(
		t,
		[]Command{
			{L: []link_message.Message{{T: "a1234567890123456789"}}},
			{Q: []link_message.Message{{T: "b1234567890123456789"}}},
		},
		Split(
			&Command{
				L: []link_message.Message{{T: "a1234567890123456789"}},
				Q: []link_message.Message{{T: "b1234567890123456789"}},
			},
			65,
		),
	)
}
