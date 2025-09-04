package lib

import (
	"go.stellar.af/go-utils/slice"
	"go.stellar.af/stellar-ip-ranges/constants"
	"go.stellar.af/stellar-ip-ranges/types"
)

type Data map[string]any

type Condition struct {
	Scope string
	Type  string
	Data  string
}

type JSONCondition struct {
	Scope    string
	Type     string
	Response Data
}

func PopSliceValue[T comparable](slice []T, value T) []T {
	out := make([]T, 0, len(slice))
	for _, element := range slice {
		if element != value {
			out = append(out, element)
		}
	}
	return out
}

type ScopeManager struct {
	IPv4   bool
	IPv6   bool
	Dual   bool
	URL    bool
	Scopes []string
}

func FilterScopes(params []string) ScopeManager {
	v4 := false
	v6 := false
	dual := true
	url := false
	if slice.Contains(params, "json") {
		if (!slice.Contains(params, constants.IP4) && !slice.Contains(params, constants.IP6)) || len(params) == 1 {
			return ScopeManager{
				Dual: true,
				URL:  true,
			}
		}

	}
	if slice.Contains(params, constants.IP4) {
		v4 = true
		dual = false
	}
	if slice.Contains(params, constants.IP6) {
		v6 = true
		dual = false
	}
	if (slice.Contains(params, constants.IP4) && slice.Contains(params, constants.IP6)) || slice.Contains(params, constants.DUAL) {
		v4 = false
		v6 = false
		dual = true
	}
	if slice.Contains(params, constants.URL) {
		v4 = false
		v6 = false
		dual = false
		url = true
	}
	params = PopSliceValue(params, constants.IP4)
	params = PopSliceValue(params, constants.IP6)
	params = PopSliceValue(params, constants.URL)
	params = PopSliceValue(params, constants.DUAL)

	scopes := make([]string, 0, 2)
	if slice.Contains(params, constants.STELLAR) {
		scopes = append(scopes, constants.STELLAR)
	}
	if slice.Contains(params, constants.ORION) {
		scopes = append(scopes, constants.ORION)
	}
	return ScopeManager{
		IPv4:   v4,
		IPv6:   v6,
		Dual:   dual,
		URL:    url,
		Scopes: scopes,
	}
}

func MatchText(params []string) string {
	s := FilterScopes(params)
	matched := make([]types.List, 0, len(params))
	if s.URL {
		if slice.Contains(s.Scopes, constants.STELLAR) {
			matched = append(matched, constants.DOMAINS_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			matched = append(matched, constants.DOMAINS_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			matched = append(matched, constants.DOMAINS_ALL)
		}
	}
	if s.IPv4 {
		if slice.Contains(s.Scopes, constants.STELLAR) {
			matched = append(matched, constants.IP4_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			matched = append(matched, constants.IP4_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			matched = append(matched, constants.IP4_ALL)
		}
	}
	if s.IPv6 {
		if slice.Contains(s.Scopes, constants.STELLAR) {
			matched = append(matched, constants.IP6_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			matched = append(matched, constants.IP6_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			matched = append(matched, constants.IP6_ALL)
		}
	}
	if s.Dual {
		if slice.Contains(s.Scopes, constants.STELLAR) {
			matched = append(matched, constants.IP4_ST_CORP, constants.IP6_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			matched = append(matched, constants.IP4_ST_ORION, constants.IP6_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			matched = append(matched, constants.IP4_ALL, constants.IP6_ALL)
		}
	}
	merged := constants.Merge(types.List{}, matched...)
	return merged.Text()
}

func MatchJSON(params []string) Data {
	s := FilterScopes(params)
	out := Data{}
	if s.URL {
		all := []types.List{}
		if slice.Contains(s.Scopes, constants.STELLAR) {
			all = append(all, constants.DOMAINS_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			all = append(all, constants.DOMAINS_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			all = append(all, constants.DOMAINS_ALL)
		}
		out[constants.URL] = constants.Merge(types.List{}, all...)
	}
	if s.IPv4 {
		all := []types.List{}
		if slice.Contains(s.Scopes, constants.STELLAR) {
			all = append(all, constants.IP4_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			all = append(all, constants.IP4_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			all = append(all, constants.IP4_ALL)
		}
		out[constants.IP4] = constants.Merge(types.List{}, all...)
	}
	if s.IPv6 {
		all := []types.List{}
		if slice.Contains(s.Scopes, constants.STELLAR) {
			all = append(all, constants.IP6_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			all = append(all, constants.IP6_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			all = append(all, constants.IP6_ALL)
		}
		out[constants.IP6] = constants.Merge(types.List{}, all...)
	}
	if s.Dual {
		all4 := []types.List{}
		all6 := []types.List{}
		if slice.Contains(s.Scopes, constants.STELLAR) {
			all4 = append(all4, constants.IP4_ST_CORP)
			all6 = append(all6, constants.IP6_ST_CORP)
		}
		if slice.Contains(s.Scopes, constants.ORION) {
			all4 = append(all4, constants.IP4_ST_ORION)
			all6 = append(all6, constants.IP6_ST_ORION)
		}
		if len(s.Scopes) == 0 {
			all4 = append(all4, constants.IP4_ALL)
			all6 = append(all6, constants.IP6_ALL)
		}
		out[constants.IP4] = constants.Merge(types.List{}, all4...)
		out[constants.IP6] = constants.Merge(types.List{}, all6...)
	}
	if len(out) == 0 {
		return nil
	}
	return out
}



func MatchCrowdstrikeText(params []string) string {
	s := FilterScopes(params)
	matched := []types.List{}
	if s.URL {
		if len(s.Scopes) == 0 {
			matched = append(matched, constants.CS_DOMAINS)
		}
	}

	// IPv4
	if s.IPv4 {
		matched = append(matched, constants.CS_IP4)
	}

	// IPv6
	if s.IPv6 {
		matched = append(matched, constants.CS_IP6)
	}

	// Dual (both)
	if s.Dual {
		matched = append(matched, constants.CS_IP4, constants.CS_IP6)
	}

	merged := constants.Merge(types.List{}, matched...)
	return merged.Text()
}

func MatchCrowdstrikeJSON(params []string) Data {
	s := FilterScopes(params)
	out := Data{}

	if s.URL {
		out[constants.URL] = constants.CS_DOMAINS
	}
	if s.IPv4 {
		out[constants.IP4] = constants.CS_IP4
	}
	if s.IPv6 {
		out[constants.IP6] = constants.CS_IP6
	}
	if s.Dual {
		out[constants.IP4] = constants.CS_IP4
		out[constants.IP6] = constants.CS_IP6
	}

	if len(out) == 0 {
		return nil
	}
	return out
}
