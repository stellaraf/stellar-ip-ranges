package types_test

import (
	"net/http"
	"strings"
	"testing"

	"go.stellar.af/stellar-ip-ranges/constants"

	"github.com/stretchr/testify/assert"
	"go.stellar.af/stellar-ip-ranges/lib"
)



func Test_Rapid7Handler_JSON_Default(t *testing.T) {
	ctx, rec := newEchoCtx(http.MethodGet, "/rapid7/json", "")
	ctx.SetParamNames("*")
	ctx.SetParamValues("rapid7/json")

	err := lib.Rapid7Handler(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	body := rec.Body.String()
	assert.Contains(t, body, `"ipv4"`)
	assert.Contains(t, body, `"ipv6"`)
	assert.Contains(t, body, `"url"`)
}

func Test_Rapid7Handler_Text_IPv4(t *testing.T) {
	ctx, rec := newEchoCtx(http.MethodGet, "/rapid7/ipv4", "")
	ctx.SetParamNames("*")
	ctx.SetParamValues("rapid7/ipv4")

	err := lib.Rapid7Handler(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	body := rec.Body.String()
	assert.NotEmpty(t, body)
	firstLine := strings.Split(body, "\n")[0]
	assert.Contains(t, firstLine, ".") // crude sanity check: IPv4 dotted quad
}


func Test_MatchRapid7Text_IPv4(t *testing.T) {
	out := lib.MatchRapid7Text([]string{"ipv4"})
	assert.Equal(t, constants.R7_IP4.Text(), out)
}

func Test_MatchRapid7Text_IPv6(t *testing.T) {
	out := lib.MatchRapid7Text([]string{"ipv6"})
	assert.Equal(t, constants.R7_IP6.Text(), out)
}

func Test_MatchRapid7Text_URL(t *testing.T) {
	out := lib.MatchRapid7Text([]string{"url"})
	assert.Equal(t, constants.R7_DOMAINS.Text(), out)
}

func Test_MatchRapid7Text_Dual(t *testing.T) {
	// Dual == IPv4 + IPv6 joined with newlines
	expected := constants.R7_IP4.Text() + "\n" + constants.R7_IP6.Text()
	out := lib.MatchRapid7Text([]string{"dual"})
	assert.Equal(t, expected, out)
}

// JSON matchers

func Test_MatchRapid7JSON_DefaultWithJSON(t *testing.T) {
	out := lib.MatchRapid7JSON([]string{"json"})
	assert.NotNil(t, out)

	// URL key present + non-empty
	urls, ok := out[constants.URL]
	assert.True(t, ok)
	assert.NotEmpty(t, urls)

	// IPv4 + IPv6 present + non-empty
	ip4, ok := out[constants.IP4]
	assert.True(t, ok)
	assert.NotEmpty(t, ip4)

	ip6, ok := out[constants.IP6]
	assert.True(t, ok)
	assert.NotEmpty(t, ip6)
}

func Test_MatchRapid7JSON_OnlyIPv4(t *testing.T) {
	out := lib.MatchRapid7JSON([]string{"json", "ipv4"})
	assert.NotNil(t, out)

	_, hasURL := out[constants.URL]
	assert.False(t, hasURL)

	_, hasIPv6 := out[constants.IP6]
	assert.False(t, hasIPv6)

	ip4, hasIPv4 := out[constants.IP4]
	assert.True(t, hasIPv4)
	assert.Equal(t, constants.R7_IP4, ip4)
}
