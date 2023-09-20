package types

import (
	"strings"
)

type List []string

func (l List) Text() []byte {
	stringBlock := strings.Join(l, "\n")
	return []byte(stringBlock)
}
