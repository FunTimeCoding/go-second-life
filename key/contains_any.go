package key

func ContainsAny(
	elements []Key,
	search []Key,
) bool {
	for _, v := range search {
		if Contains(elements, v) {
			return true
		}
	}

	return false
}
