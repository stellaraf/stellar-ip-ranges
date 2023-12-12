package types

import (
	"strings"
)

type List []string

func (l List) Text() string {
	return strings.Join(l, "\n")
}
