package types_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go.stellar.af/stellar-ip-ranges/constants"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.stellar.af/stellar-ip-ranges/lib"
)

func newEchoCtx(method, target string, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, target, strings.NewReader(body))
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func Test_CrowdstrikeHandler_JSON_Default(t *testing.T) {
	ctx, rec := newEchoCtx(http.MethodGet, "/crowdstrike/json", "")
	ctx.SetParamNames("*")
	ctx.SetParamValues("crowdstrike/json")

	err := lib.CrowdstrikeHandler(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	body := rec.Body.String()
	assert.Contains(t, body, `"ipv4"`)
	assert.Contains(t, body, `"ipv6"`)
	assert.Contains(t, body, `"url"`)
}

func Test_CrowdstrikeHandler_Text_IPv4(t *testing.T) {
	ctx, rec := newEchoCtx(http.MethodGet, "/crowdstrike/ipv4", "")
	ctx.SetParamNames("*")
	ctx.SetParamValues("crowdstrike/ipv4")

	err := lib.CrowdstrikeHandler(ctx)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	body := rec.Body.String()
	assert.NotEmpty(t, body)
	firstLine := strings.Split(body, "\n")[0]
	assert.Contains(t, firstLine, ".") // crude sanity check: IPv4 dotted quad
}


func Test_MatchCrowdstrikeText_IPv4(t *testing.T) {
	out := lib.MatchCrowdstrikeText([]string{"ipv4"})
	assert.Equal(t, constants.CS_IP4.Text(), out)
}

func Test_MatchCrowdstrikeText_IPv6(t *testing.T) {
	out := lib.MatchCrowdstrikeText([]string{"ipv6"})
	assert.Equal(t, constants.CS_IP6.Text(), out)
}

func Test_MatchCrowdstrikeText_URL(t *testing.T) {
	out := lib.MatchCrowdstrikeText([]string{"url"})
	assert.Equal(t, constants.CS_DOMAINS.Text(), out)
}

func Test_MatchCrowdstrikeText_Dual(t *testing.T) {
	// Dual == IPv4 + IPv6 joined with newlines
	expected := constants.CS_IP4.Text() + "\n" + constants.CS_IP6.Text()
	out := lib.MatchCrowdstrikeText([]string{"dual"})
	assert.Equal(t, expected, out)
}

// JSON matchers

func Test_MatchCrowdstrikeJSON_DefaultWithJSON(t *testing.T) {
	out := lib.MatchCrowdstrikeJSON([]string{"json"})
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

func Test_MatchCrowdstrikeJSON_OnlyIPv4(t *testing.T) {
	out := lib.MatchCrowdstrikeJSON([]string{"json", "ipv4"})
	assert.NotNil(t, out)

	_, hasURL := out[constants.URL]
	assert.False(t, hasURL)

	_, hasIPv6 := out[constants.IP6]
	assert.False(t, hasIPv6)

	ip4, hasIPv4 := out[constants.IP4]
	assert.True(t, hasIPv4)
	assert.Equal(t, constants.CS_IP4, ip4)
}
