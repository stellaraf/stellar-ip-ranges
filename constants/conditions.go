package constants

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stellaraf/stellar-ip-ranges/types"
)

type Condition struct {
	Scope string
	Type  string
	Data  types.List
}

type JSONCondition struct {
	Scope    string
	Type     string
	Response fiber.Map
}

func (c *Condition) Match(scope, _type string) bool {
	return scope == c.Scope && _type == c.Type

}

func (c *JSONCondition) Match(scope, _type string) bool {
	return scope == c.Scope && _type == c.Type
}

var TextConditions = []Condition{
	{Scope: STELLAR, Type: IP4, Data: IP4_ST_CORP},
	{Scope: STELLAR, Type: IP6, Data: IP6_ST_CORP},
	{Scope: STELLAR, Type: URL, Data: DOMAINS_ST_CORP},
	{Scope: STELLAR, Type: DUAL, Data: IP_DUAL_ST_CORP},
	{Scope: STELLAR, Type: GLOBAL, Data: IP_DUAL_ST_CORP},
	{Scope: ORION, Type: IP4, Data: IP4_ST_ORION},
	{Scope: ORION, Type: IP6, Data: IP6_ST_ORION},
	{Scope: ORION, Type: URL, Data: DOMAINS_ST_ORION},
	{Scope: ORION, Type: DUAL, Data: IP_DUAL_ST_ORION},
	{Scope: ORION, Type: GLOBAL, Data: IP_DUAL_ST_ORION},
	{Scope: IP4, Type: STELLAR, Data: IP4_ST_CORP},
	{Scope: IP4, Type: GLOBAL, Data: IP4_ALL},
	{Scope: IP4, Type: ORION, Data: IP4_ST_ORION},
	{Scope: IP4, Type: DUAL, Data: IP4_ALL},
	{Scope: IP6, Type: STELLAR, Data: IP6_ST_CORP},
	{Scope: IP6, Type: ORION, Data: IP6_ST_ORION},
	{Scope: IP6, Type: GLOBAL, Data: IP6_ALL},
	{Scope: IP6, Type: DUAL, Data: IP6_ALL},
	{Scope: URL, Type: STELLAR, Data: DOMAINS_ST_CORP},
	{Scope: URL, Type: ORION, Data: DOMAINS_ST_ORION},
	{Scope: URL, Type: GLOBAL, Data: DOMAINS_ALL},
	{Scope: URL, Type: DUAL, Data: DOMAINS_ALL},
	{Scope: DUAL, Type: IP4, Data: IP4_ALL},
	{Scope: DUAL, Type: IP6, Data: IP6_ALL},
	{Scope: DUAL, Type: STELLAR, Data: IP_DUAL_ST_CORP},
	{Scope: DUAL, Type: ORION, Data: IP_DUAL_ST_ORION},
	{Scope: DUAL, Type: GLOBAL, Data: IP_DUAL},
	{Scope: GLOBAL, Type: IP4, Data: IP4_ALL},
	{Scope: GLOBAL, Type: IP6, Data: IP6_ALL},
	{Scope: GLOBAL, Type: URL, Data: DOMAINS_ALL},
	{Scope: GLOBAL, Type: DUAL, Data: IP_DUAL},
	{Scope: GLOBAL, Type: STELLAR, Data: IP_DUAL_ST_CORP},
	{Scope: GLOBAL, Type: STELLAR, Data: IP_DUAL_ST_CORP},
	{Scope: STELLAR, Type: ORION, Data: IP_DUAL_ST_CORP_ORION},
	{Scope: ORION, Type: STELLAR, Data: IP_DUAL_ST_CORP_ORION},
}

var JSONConditions []JSONCondition = []JSONCondition{
	{Scope: STELLAR, Type: IP4, Response: fiber.Map{IP4: IP4_ST_CORP}},
	{Scope: STELLAR, Type: IP6, Response: fiber.Map{IP6: IP6_ST_CORP}},
	{Scope: STELLAR, Type: URL, Response: fiber.Map{URL: DOMAINS_ST_CORP}},
	{Scope: STELLAR, Type: DUAL, Response: fiber.Map{IP4: IP4_ST_CORP, IP6: IP6_ST_CORP}},
	{Scope: STELLAR, Type: GLOBAL, Response: fiber.Map{IP4: IP4_ST_CORP, IP6: IP4_ST_CORP, URL: DOMAINS_ST_CORP}},
	{Scope: ORION, Type: IP4, Response: fiber.Map{IP4: IP4_ST_ORION}},
	{Scope: ORION, Type: IP6, Response: fiber.Map{IP6: IP6_ST_ORION}},
	{Scope: ORION, Type: URL, Response: fiber.Map{URL: DOMAINS_ST_ORION}},
	{Scope: ORION, Type: DUAL, Response: fiber.Map{IP4: IP4_ST_ORION, IP6: IP6_ST_ORION}},
	{Scope: ORION, Type: GLOBAL, Response: fiber.Map{IP4: IP4_ST_ORION, IP6: IP4_ST_ORION, URL: DOMAINS_ST_ORION}},
	{Scope: IP4, Type: STELLAR, Response: fiber.Map{IP4: IP4_ST_CORP}},
	{Scope: IP4, Type: ORION, Response: fiber.Map{IP4: IP4_ST_ORION}},
	{Scope: IP4, Type: GLOBAL, Response: fiber.Map{IP4: IP4_ALL}},
	{Scope: IP4, Type: DUAL, Response: fiber.Map{IP4: IP4_ALL}},
	{Scope: IP6, Type: STELLAR, Response: fiber.Map{IP6: IP6_ST_CORP}},
	{Scope: IP6, Type: ORION, Response: fiber.Map{IP6: IP6_ST_ORION}},
	{Scope: IP6, Type: GLOBAL, Response: fiber.Map{IP6: IP6_ALL}},
	{Scope: IP6, Type: DUAL, Response: fiber.Map{IP6: IP6_ALL}},
	{Scope: URL, Type: STELLAR, Response: fiber.Map{URL: DOMAINS_ST_CORP}},
	{Scope: URL, Type: ORION, Response: fiber.Map{URL: DOMAINS_ST_ORION}},
	{Scope: URL, Type: GLOBAL, Response: fiber.Map{URL: DOMAINS_ALL}},
	{Scope: DUAL, Type: IP4, Response: fiber.Map{IP4: IP4_ALL}},
	{Scope: DUAL, Type: IP6, Response: fiber.Map{IP6: IP6_ALL}},
	{Scope: DUAL, Type: STELLAR, Response: fiber.Map{IP4: IP4_ST_CORP, IP6: IP6_ST_CORP}},
	{Scope: DUAL, Type: ORION, Response: fiber.Map{IP4: IP4_ST_ORION, IP6: IP6_ST_ORION}},
	{Scope: DUAL, Type: GLOBAL, Response: fiber.Map{IP4: IP4_ALL, IP6: IP6_ALL}},
	{Scope: GLOBAL, Type: IP4, Response: fiber.Map{IP4: IP4_ALL}},
	{Scope: GLOBAL, Type: IP6, Response: fiber.Map{IP6: IP6_ALL}},
	{Scope: GLOBAL, Type: URL, Response: fiber.Map{URL: DOMAINS_ALL}},
	{Scope: GLOBAL, Type: DUAL, Response: fiber.Map{IP4: IP4_ALL, IP6: IP6_ALL, URL: DOMAINS_ALL}},
	{Scope: GLOBAL, Type: STELLAR, Response: fiber.Map{IP4: IP4_ST_CORP, IP6: IP4_ST_CORP, URL: DOMAINS_ST_CORP}},
	{Scope: GLOBAL, Type: ORION, Response: fiber.Map{IP4: IP4_ST_ORION, IP6: IP4_ST_ORION, URL: DOMAINS_ST_ORION}},
}
