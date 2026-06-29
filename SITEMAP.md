# API Sitemap — stellar-ip-ranges

All routes are registered in [api/index.go](api/index.go) and dispatched through Echo. Path segments after the prefix act as **filter tokens** (any order, slash-separated) parsed by `FilterScopes` in [lib/match.go](lib/match.go).

The Vercel rewrite in [vercel.json](vercel.json) funnels every request to `/api/index.go`, so all prefix matching happens inside the Echo router.

---

## Route prefixes

| Prefix | Handler | Source |
|---|---|---|
| `/rapid7*` | `Rapid7Handler` | [lib/handlers.go:80](lib/handlers.go#L80) |
| `/crowdstrike*` | `CrowdstrikeHandler` | [lib/handlers.go:53](lib/handlers.go#L53) |
| `/geofeed*` | `GeofeedHandler` | [lib/handlers.go:38](lib/handlers.go#L38) |
| `/*` (everything else) | `BaseHandler` — Stellar/Orion data | [lib/handlers.go:11](lib/handlers.go#L11) |

---

## Filter tokens

Defined in [constants/path.go](constants/path.go). Tokens can appear in any order, separated by `/`.

| Token | Effect |
|---|---|
| `json` | Return JSON instead of plain text |
| `ipv4` | IPv4 ranges only |
| `ipv6` | IPv6 ranges only |
| `dual` | Both IPv4 and IPv6 (default when neither is specified) |
| `url` | Return domains instead of IPs |
| `stellar` | Scope filter — Stellar Corp only *(BaseHandler only)* |
| `orion` | Scope filter — Stellar Orion only *(BaseHandler only)* |

---

## Effective endpoints

### Base (`/*`) — Stellar + Orion ranges

| URL | Returns |
|---|---|
| `/` | All IPv4 + IPv6 (text) |
| `/json` | All IPv4 + IPv6 (JSON) |
| `/ipv4` | IPv4 only (text) |
| `/ipv6` | IPv6 only (text) |
| `/dual` | Both families (text) |
| `/url` | Domains list (text) |
| `/stellar` | Stellar Corp scope, both families |
| `/orion` | Orion scope, both families |
| `/stellar/ipv6/json` | Example — Stellar Corp IPv6 as JSON |
| `/orion/url/json` | Example — Orion domains as JSON |

Any combination of tokens is valid.

### Crowdstrike (`/crowdstrike*`)

Scope filters (`stellar`, `orion`) are ignored — Crowdstrike data is a single dataset.

| URL | Returns |
|---|---|
| `/crowdstrike` | CS IPv4 + IPv6 (text) |
| `/crowdstrike/json` | CS IPv4 + IPv6 (JSON) |
| `/crowdstrike/ipv4` | CS IPv4 only |
| `/crowdstrike/ipv6` | CS IPv6 only |
| `/crowdstrike/dual` | Both families |
| `/crowdstrike/url` | CS domains |
| Append `/json` to any of the above | JSON variant |

### Rapid7 (`/rapid7*`)

Mirrors Crowdstrike.

| URL | Returns |
|---|---|
| `/rapid7` | R7 IPv4 + IPv6 (text) |
| `/rapid7/json` | R7 IPv4 + IPv6 (JSON) |
| `/rapid7/ipv4` | R7 IPv4 only |
| `/rapid7/ipv6` | R7 IPv6 only |
| `/rapid7/dual` | Both families |
| `/rapid7/url` | R7 domains |
| Append `/json` to any of the above | JSON variant |

### Geofeed (`/geofeed*`)

See [lib/handlers.go:38](lib/handlers.go#L38). Does not accept filter tokens — only suffix variants.

| URL | Returns |
|---|---|
| `/geofeed` | CSV body served as `text/plain` |
| `/geofeed.csv` | CSV download (`stellar-geofeed.csv`) |
| `/geofeed.txt` | Same CSV as `.txt` download |

---

## Response contract

- **Text** — newline-joined CIDRs/domains via `List.Text()`
- **JSON** — `{ "ipv4": [...], "ipv6": [...], "url": [...] }`; only keys for requested types are included ([lib/match.go:146](lib/match.go#L146))
- **No match** — `400` with `"no matching parameters"` (text) or `{"error": "no matching parameters"}` (JSON)
- **CORS** — `*` for all origins, all standard methods ([api/index.go:17-21](api/index.go#L17-L21))
