package inventory

import (
	"fmt"
	"path"
)

func Path(elements ...string) string {
	return fmt.Sprintf("/%s", path.Join(elements...))
}
