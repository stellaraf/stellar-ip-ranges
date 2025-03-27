package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.stellar.af/stellar-ip-ranges/types"
)

func Test_GeofeedEntry(t *testing.T) {
	t.Parallel()
	g := &types.GeofeedEntry{
		Name:        "Orion GVL01 IPv4",
		Prefix:      "216.250.231.0/24",
		CountryCode: "US",
		RegionCode:  "SC",
		City:        "Greenville",
		PostalCode:  "29607",
	}
	result := g.CSV()
	exp := `# Orion GVL01 IPv4
216.250.231.0/24,US,SC,Greenville,29607,`

	require.Equal(t, exp, result)
}

func Test_Geofeed(t *testing.T) {
	t.Parallel()
	e1 := types.GeofeedEntry{
		Name:        "Orion GVL01 IPv4",
		Prefix:      "216.250.231.0/24",
		CountryCode: "US",
		RegionCode:  "SC",
		City:        "Greenville",
		PostalCode:  "29607",
	}
	e2 := types.GeofeedEntry{
		Name:        "Orion PHX01 IPv6",
		Prefix:      "2604:c0c0:1000::/36",
		CountryCode: "US",
		RegionCode:  "AZ",
		City:        "Phoenix",
		PostalCode:  "85004",
	}
	g := types.Geofeed{e1, e2}
	result := g.CSV("AS14525 Geofeed")
	exp := `# AS14525 Geofeed
# Orion GVL01 IPv4
216.250.231.0/24,US,SC,Greenville,29607,
# Orion PHX01 IPv6
2604:c0c0:1000::/36,US,AZ,Phoenix,85004,
`

	require.Equal(t, exp, result)
}
