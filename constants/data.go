package constants

import "github.com/stellaraf/stellar-ip-ranges/types"

func Merge(init types.List, rest ...types.List) types.List {
	merged := init
	for _, list := range rest {
		merged = append(merged, list...)
	}
	return merged
}

var IP4_ALL types.List = types.List{
	"199.34.92.0/22",
	"216.250.230.0/23",
}
var IP6_ALL types.List = types.List{
	"2604:c0c0::/32",
}

var IP4_ST_CORP types.List = types.List{
	"199.34.92.144/29",
	"199.34.94.120/30",
	"216.250.230.192/30",
	"216.250.231.64/29",
}

var IP6_ST_CORP types.List = types.List{
	"2604:c0c0:100e::/48",
	"2604:c0c0:1828::/48",
	"2604:c0c0:3001::/48",
	"2604:c0c0:500e::/48",
	"2604:c0c0:800e::/48",
}

var IP4_ST_ORION types.List = types.List{
	"199.34.92.64/26",
	"199.34.94.32/27",
	"216.250.230.0/27",
	"216.250.231.0/27",
}

var IP6_ST_ORION types.List = types.List{
	"2604:c0c0:15af::/48",
	"2604:c0c0:35af::/48",
	"2604:c0c0:55af::/48",
	"2604:c0c0:85af::/48",
}

var DOMAINS_ST_CORP = types.List{
	"stellar.tech",
	"stellar.af",
	"*.stellar.tech",
	"*.stellar.af",
}

var DOMAINS_ST_ORION = types.List{
	"chi01.orion.cloud",
	"gvl01.orion.cloud",
	"hnl01.orion.cloud",
	"phx01.orion.cloud",
	"orion.cloud",
	"*.chi01.orion.cloud",
	"*.gvl01.orion.cloud",
	"*.hnl01.orion.cloud",
	"*.phx01.orion.cloud",
	"*.orion.cloud",
}

var DOMAINS_ALL types.List = Merge(DOMAINS_ST_ORION, DOMAINS_ST_CORP)

var IP_DUAL types.List = Merge(IP4_ALL, IP6_ALL)

var IP_DUAL_ST_CORP types.List = Merge(IP4_ST_CORP, IP6_ST_CORP)

var IP_DUAL_ST_ORION types.List = Merge(IP4_ST_ORION, IP6_ST_ORION)

var IP_DUAL_ST_CORP_ORION types.List = Merge(IP4_ST_CORP, IP6_ST_CORP, IP4_ST_ORION, IP6_ST_ORION)
