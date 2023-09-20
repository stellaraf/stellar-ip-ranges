package types_test

import (
	"testing"

	"github.com/stellaraf/stellar-ip-ranges/types"
	"github.com/stretchr/testify/assert"
)

func Test_List(t *testing.T) {
	t.Run("base", func(t *testing.T) {
		list := types.List{
			"192.0.2.1",
			"192.0.2.2",
			"192.0.2.3",
			"192.0.2.4",
		}
		expected := []byte(`192.0.2.1
192.0.2.2
192.0.2.3
192.0.2.4`)
		result := list.Text()
		assert.Equal(t, expected, result)
	})
}
