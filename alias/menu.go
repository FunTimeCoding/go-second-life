package alias

type Menu string

func (m Menu) String() string {
	return string(m)
}
