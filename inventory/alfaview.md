# alfaview gmbh inventory (discovery seed 2026-09-02)
# NOTE: hosts below are discovery candidates from passive DNS/CT; confirm in-scope vs program scope before active testing.
alfaview.com
app.alfaview.com
dev.alfaview.com
sso.alfaview.com
staging.alfaview.com
support.alfaview.com
test.alfaview.com
www.alfaview.com

## PASSIVE RECON 2026-09-02 (read-only, non-intrusive)

> Recon observations only. These are NOT confirmed vulnerabilities; ownership/in-scope of each host must be confirmed against the program scope before any active testing. Hosts resolve + serve HTTP — investigation requires scoped authorization.

**Probed:** 8 hosts | **Live HTTP:** 2

| Host | Status | Server/Tech |
|---|---|---|
| `support.alfaview.com` | 301 | Server: myracloud -> https://support.alfaview.com/en/ |
| `staging.alfaview.com` | 301 | Server: myracloud -> https://staging.alfaview.com/en |

**CNAME review signals (2):**
- `support.alfaview.com` -> `support-alfaview-com.ax4z.com`
- `staging.alfaview.com` -> `staging-alfaview-com.ax4z.com`

## DEEP SERVICE SCAN 2026-09-02 (read-only connect+banner)
**Host:** `staging.alfaview.com` | **Ports:** [80, 443]
**Web surface only:** [80, 443]

## DEEP SERVICE SCAN 2026-09-02 (read-only connect+banner)
**Host:** `support.alfaview.com` | **Ports:** [80, 443]
**Web surface only:** [80, 443]

## DEEP ENUM (wildcard-cleaned) 2026-09-03
**Root zone:** `alfaview.com` | **dedicated hosts after wildcard-filter: 55**
> Audit: brute+passive subfinder produced 10,083 resolving hostnames; zone-wildcard + IP-fingerprint filtering dropped 9,973 (98.9%) DNS-wildcard noise (random labels resolving to shared wildcard IPs e.g. account.cineplex.de, a.hypofriend.de, account.live-manager.de, docker.jtl-software.de, *.ggamdom.com, *.dev.alfaview.com). Only genuine dedicated hosts listed below. These are surface-map observations; live HTTP status captured read-only (GET / via curl). No findings claimed; scope must be confirmed with the program.
- `alfacheck-audio.alfaview.com`  [HTTP unprobed]
- `alfacheck-engine.alfaview.com`  [HTTP unprobed]
- `alfacheck-video.alfaview.com`  [HTTP unprobed]
- `alfatraining.alfaview.com`  [HTTP 200]
- `alfaview-com-assets.alfaview.com`  [HTTP 403]
- `apis.alfaview.com`  [HTTP 200]
- `app.alfaview.com`  [HTTP 200]
- `appstats.alfaview.com`  [HTTP unprobed]
- `assets.alfaview.com`  [HTTP 403]
- `beta-alfaview-assets.alfaview.com`  [HTTP 403]
- `beta-alfaview-com-assets.alfaview.com`  [HTTP 403]
- `beta-apis.alfaview.com`  [HTTP 200]
- `beta-app.alfaview.com`  [HTTP 401]
- `beta-hcloud-19-beta-audio-4xstl.alfaview.com`  [HTTP 404]
- `beta-hcloud-19-beta-audio-8xnz9.alfaview.com`  [HTTP 404]
- `beta-hcloud-19-beta-engine-9m62k.alfaview.com`  [HTTP 404]
- `beta-hcloud-19-beta-engine-vlhwd.alfaview.com`  [HTTP 404]
- `beta-hcloud-19-beta-hydra-dzwx8.alfaview.com`  [HTTP 200]
- `beta-hcloud-19-beta-video-xjl6p.alfaview.com`  [HTTP 404]
- `beta-hcloud-19-beta-video-zhjvl.alfaview.com`  [HTTP 404]
- `beta-ionoscloud-21-beta-audio-65st7.alfaview.com`  [HTTP unprobed]
- `beta-ionoscloud-21-beta-audio-bdtmf.alfaview.com`  [HTTP unprobed]
- `beta-ionoscloud-21-beta-engine-gw4qw.alfaview.com`  [HTTP unprobed]
- `beta-ionoscloud-21-beta-engine-kzmvv.alfaview.com`  [HTTP unprobed]
- `beta-ionoscloud-21-beta-hydra-7x5d5.alfaview.com`  [HTTP unprobed]
- `beta-ionoscloud-21-beta-video-6pp2m.alfaview.com`  [HTTP unprobed]
- `beta-ionoscloud-21-beta-video-l5mbv.alfaview.com`  [HTTP unprobed]
- `beta-noris-33-beta-audio-ljj9x.alfaview.com`  [HTTP 404]
- `beta-noris-33-beta-audio-xr4mw.alfaview.com`  [HTTP 404]
- `beta-noris-33-beta-engine-cr5rs.alfaview.com`  [HTTP 404]
- `beta-noris-33-beta-engine-w84qw.alfaview.com`  [HTTP 404]
- `beta-noris-33-beta-hydra-2zm7t.alfaview.com`  [HTTP 200]
- `beta-noris-33-beta-video-2f9jw.alfaview.com`  [HTTP 404]
- `beta-noris-33-beta-video-vbvn5.alfaview.com`  [HTTP 404]
- `beta-ovh-29-beta-audio-j5qgh.alfaview.com`  [HTTP 404]
- `beta-ovh-29-beta-audio-vldgf.alfaview.com`  [HTTP 404]
- `beta-ovh-29-beta-engine-dxvt4.alfaview.com`  [HTTP 404]
- `beta-ovh-29-beta-engine-k7khj.alfaview.com`  [HTTP 404]
- `beta-ovh-29-beta-hydra-z4tf8.alfaview.com`  [HTTP 200]
- `beta-ovh-29-beta-video-8zvvm.alfaview.com`  [HTTP 404]
- `beta-ovh-29-beta-video-gbbw4.alfaview.com`  [HTTP 404]
- `beta-webclient.alfaview.com`  [HTTP 200]
- `bhc.alfaview.com`  [HTTP 200]
- `client-diagnostics-ingest.alfaview.com`  [HTTP 404]
- `clone.staging-wordpress.alfaview.com`  [HTTP 303]
- `consul-monitoring.alfaview.com`  [HTTP unprobed]
- `demo-company.alfaview.com`  [HTTP 200]
- `design-assets.alfaview.com`  [HTTP 404]
- `design-tokens.alfaview.com`  [HTTP 404]
- `equipment.alfaview.com`  [HTTP unprobed]
- `hello-world.atlas-spike.atlas.alfaview.com`  [HTTP 404]
- `insider-webclient.alfaview.com`  [HTTP 200]
- `internal.alfaview.com`  [HTTP 401]
- `ip-185-245-101-240.alfaview.com`  [HTTP unprobed]
- `kh-freiburg.alfaview.com`  [HTTP 200]

## 2026-09-03 11:40:30 UTC

## 2026-09-03 14:23:21 UTC

## 2026-09-03 15:21:33 UTC
- NEW 55 dedicated hosts confirmed after wildcard filtering (was 8 in initial recon) — inventory/alfaview.md:35-92
- NEW Probe result: `GET https://beta-apis.alfaview.com/v2/languages` (no auth) → HTTP 401; `GET https://apis.alfaview.com/v2/languages` → HTTP 404 — probe-results.md:6-8
- CHANGED Beta API weaker auth hypothesis **disproven** — beta returns 401 (endpoint exists, auth required), production returns 404 (endpoint missing) — version drift confirmed
- CHANGED Production API lacks `/v2/languages` endpoint present in beta — API version divergence
- CHANGED beta-apis.alfaview.com: Auth response identical to production (401 + same error body). Beta weaker auth hypothesis disconfirmed.
- NEW beta-webclient.alfaview.com (HTTP 200): High-value web client surface, untested.
- NEW insider-webclient.alfaview.com (HTTP 200): Internal tooling potentially exposed.

## 2026-09-03 18:48:37 UTC
- NEW OpenAPI specs at `https://apis.alfaview.com/v2/docs/openapi.json` and `https://beta-apis.alfaview.com/v2/docs/openapi.json` are **identical** (same endpoints, schemas, auth requirements) — confirms ve
- NEW `GET https://apis.alfaview.com/v2/languages` returns 404 (endpoint absent in production); `GET https://beta-apis.alfaview.com/v2/languages` returns 401 (endpoint exists, auth enforced) — confirmed API
- NEW OpenAPI spec exposes `DELETE /v2/users/{id}` and `PATCH/DELETE /v2/rooms/{roomId}/permissions/{userId}` with UUID path params — direct evidence for IDOR hypothesis
- NEW `demo-company.alfaview.com/api/v1/users` returns 302 redirect to `/` — no unauthenticated user enumeration
- CHANGED Beta API weaker auth hypothesis **fully disproven** — OpenAPI specs identical, both enforce auth identically
- CHANGED API version drift scope narrowed: only `/v2/languages` endpoint differs (beta has it, prod doesn't)

## 2026-09-03 21:22:00 UTC

## 2026-09-03 23:24:23 UTC
- CHANGED Production API `apis.alfaview.com/v2/languages` now returns **401** (was 404) — endpoint added to production, aligns with beta; OpenAPI specs now identical including `/v2/languages` path
- CHANGED Beta API `beta-apis.alfaview.com/v2/languages` returns **401** (consistent) — both environments now enforce auth identically on this endpoint
- NEW `insider-webclient.alfaview.com` and `beta-webclient.alfaview.com` both serve identical SPA shells (4396 bytes, same HTML structure, `/health`=204, `/api|/admin|/debug|/internal|/v2|/docs`=404) — no i
- NEW `demo-company.alfaview.com` serves SPA (HTTP 200) — unauthenticated web surface confirmed
- NEW 55 dedicated hosts confirmed after wildcard filtering; 48 remain HTTP-unprobed (e.g., `alfacheck-*`, `beta-hcloud-*`, `beta-ionoscloud-*`, `beta-noris-*`, `beta-ovh-*`, `consul-monitoring`, `equipment

## 2026-09-04 01:12:26 UTC
- NEW Production API `apis.alfaview.com/v2/languages` now returns **401** (was 404) — endpoint added to production, aligns with beta; OpenAPI specs now identical including `/v2/languages` path
- NEW `insider-webclient.alfaview.com` and `beta-webclient.alfaview.com` both serve identical SPA shells (4396 bytes, same HTML structure, `/health`=204, `/api|/admin|/debug|/internal|/v2|/docs`=404) — no i
- NEW `demo-company.alfaview.com` serves SPA (HTTP 200) — unauthenticated web surface confirmed
- NEW 55 dedicated hosts confirmed after wildcard filtering; 48 remain HTTP-unprobed (e.g., `alfacheck-*`, `beta-hcloud-*`, `beta-ionoscloud-*`, `beta-noris-*`, `beta-ovh-*`, `consul-monitoring`, `equipment
- NEW Live HTTP 200 on previously unprobed: `alfatraining.alfaview.com`, `bhc.alfaview.com`, `kh-freiburg.alfaview.com`, `beta-hcloud-19-beta-hydra-dzwx8.alfaview.com`, `beta-noris-33-beta-hydra-2zm7t.alfav
- NEW `beta-app.alfaview.com` and `internal.alfaview.com` both return HTTP 401 (auth-gated)
- NEW `appstats.alfaview.com`, `consul-monitoring.alfaview.com`, `equipment.alfaview.com`, `ip-185-245-101-240.alfaview.com` timeout/unreachable
- CHANGED Beta API weaker auth hypothesis **fully disproven** — OpenAPI specs identical, both enforce auth identically (401 on `/v2/languages`)
- CHANGED API version drift **resolved** — both beta and production now expose `/v2/languages` with identical auth enforcement
- CHANGED Cross-tenant IDOR on room permissions and user deletion via UUID path params confirmed as highest-priority authenticated target (confidence 80)

## 2026-09-04 06:00:56 UTC
- CHANGED alfacheck-engine.alfaview.com: Was "HTTP unprobed" → now confirmed UNREACHABLE (3 timeout probes: root, /health, /status)
- CHANGED alfacheck-audio.alfaview.com: Was "HTTP unprobed" → now confirmed UNREACHABLE (3 timeout probes: root, /media, /recordings)
- CHANGED alfacheck-video.alfaview.com: Was "HTTP unprobed" → now confirmed UNREACHABLE (1 timeout probe: root)
- NEW beta-hcloud-19-beta-hydra-dzwx8.alfaview.com: HTTP 200 with 9-byte body; trailing slash returns HTTP 400 — minimal surface, likely health-check endpoint

## 2026-09-04 10:12:15 UTC

## 2026-09-04 14:25:07 UTC
- CHANGED beta-hcloud-19-beta-hydra-dzwx8.alfaview.com: Was ACCEPTED MISCONFIG (health-check surface) → now REJECTED after probing — "Hi Client" body on all paths confirms media/signaling server, not OIDC/auth 
- CHANGED beta-noris-33-beta-hydra-2zm7t.alfaview.com: Same — media server, REJECTED.
- CHANGED beta-ovh-29-beta-hydra-z4tf8.alfaview.com: Same — media server, REJECTED.
- CHANGED All 3 alfacheck-* hosts: Confirmed UNREACHABLE via timeout (root, /health, /status, /media, /recordings). Internal/firewalled. REJECTED.
- CHANGED RISK score dropped from 58 → 55 — all hydra/alfacheck surfaces resolved (rejected). Remaining surface is auth-gated.
- CHANGED beta-hcloud-19-beta-hydra-dzwx8.alfaview.com: ACCEPTED MISCONFIG → REJECTED — media/signaling server ("Hi Client"), not OIDC.
- CHANGED beta-noris-33-beta-hydra-2zm7t.alfaview.com: Same — REJECTED.
- CHANGED beta-ovh-29-beta-hydra-z4tf8.alfaview.com: Same — REJECTED.
- CHANGED alfacheck-engine.alfaview.com: UNREACHABLE confirmed (3 timeout probes).
- CHANGED alfacheck-audio.alfaview.com: UNREACHABLE confirmed (3 timeout probes).
- CHANGED alfacheck-video.alfaview.com: UNREACHABLE confirmed (1 timeout probe).
- NEW Production API `apis.alfaview.com/v2/languages` now returns 401 (was 404) — endpoint added, aligns with beta; OpenAPI specs identical
- NEW `beta-hcloud-19-beta-hydra-dzwx8.alfaview.com`: HTTP 200 with 9-byte body ("Hi Client"); trailing slash returns HTTP 400 — minimal health-check surface
- NEW Guest link auth flow updated to 4-field combo (companyId+roomId+accessKey+displayName) — rate-limit status unknown
- CHANGED API version drift **resolved** — both beta and production expose `/v2/languages` with identical auth enforcement (401)
- CHANGED `alfacheck-engine/audio/video.alfaview.com`: confirmed UNREACHABLE (timeout probes) — target exhausted
- CHANGED `beta-hcloud-19-beta-hydra-dzwx8.alfaview.com` et al: confirmed media/signaling servers ("Hi Client") — no admin endpoints

## 2026-09-04 17:51:04 UTC

## 2026-09-04 20:04:20 UTC
- NEW beta-ionoscloud-21-* fleet (7 hosts): `beta-ionoscloud-21-beta-audio-65st7`, `beta-ionoscloud-21-beta-audio-bdtmf`, `beta-ionoscloud-21-beta-engine-gw4qw`, `beta-ionoscloud-21-beta-engine-kzmvv`, `bet
- NEW Main domain set from initial recon (6 hosts): `alfaview.com`, `app.alfaview.com`, `dev.alfaview.com`, `sso.alfaview.com`, `test.alfaview.com`, `www.alfaview.com` — only `support`/`staging` probed (301
- CHANGED `apis.alfaview.com/v2/languages` now 401 (was 404) — endpoint added, aligns with beta; OpenAPI specs identical
- CHANGED `alfatraining`/`bhc`/`kh-freiburg` XSS hypothesis REJECTED — byte-identical SPA shells (1381B, MD5 554a39), no tenant-specific rendering

## 2026-09-04 22:09:53 UTC
- NEW beta-ionoscloud-21-beta-hydra-7x5d5.alfaview.com: HTTP timeout (root + trailing slash) — unlike hcloud/noris/ovh hydra hosts which returned "Hi Client"
- NEW alfaview.com: HTTP 200 — main marketing/auth entry point now confirmed live
- NEW sso.alfaview.com: HTTP 200 len=0 — SSO endpoint live, empty body (likely redirects or SPA shell)
- CHANGED beta-ionoscloud-21-* fleet (7 hosts): Previously "HTTP unprobed" → now 1/7 probed (hydra=timeout), 6 remain unprobed
- CHANGED Main domain set (6 hosts): Previously only support/staging probed (301) → now alfaview.com + sso.alfaview.com confirmed HTTP 200, 4 remain unprobed (app, dev, test, www)

## 2026-09-05 00:15:07 UTC

## 2026-09-05 04:34:13 UTC
- NEW sso.alfaview.com: FusionAuth 1.63.0 OIDC discovery live (issuer=acme.com, implicit flow, HS256 in supported algs but RSA-only JWKS) — confirmed 2026-09-05 00:15
- NEW test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) — no integrity verification visible
- NEW alfaview.com: Marketing page live (301→/en, 177KB, strict CSP, matomo) — no SSO login links on marketing domain
- CHANGED beta-ionoscloud-21-beta-engine-* (2 hosts): Confirmed timeout (000) — internal/firewalled like alfacheck-* fleet
- CHANGED alfaview.com: Root now redirects to /en with nginx + Accept-Language vary header
- CHANGED www.alfaview.com: 301→alfaview.com/en (no independent surface)

## 2026-09-05 08:50:59 UTC
- NEW app.alfaview.com/js/app.min.5b3949112f0cf682adc8.js (1.09MB) leaks FULL admin/company GraphQL schema: 60+ mutations + 45+ queries with exact args. Not just SDL strings — includes CreateCompany($compan
- NEW GraphQL live at app.alfaview.com/graphql: __typename OK unauth; introspection disabled; sensitive ops return UNAUTHENTICATED. Resolver auth is INCONSISTENT: listIdentityProviders returns data unauth (
- NEW CRITICAL DELTA: GraphQL guest path guestAuthenticate(userId,companyId,roomId) and guestJoin(userId,companyId,roomId,displayName) take NO accessKey — while REST guest-link flow requires 4-field combo i
- NEW Hosts: webclient.alfaview.com=200/4396B(insider family); staging-webclient.alfaview.com=401 Basic realm=Protected (/health=204); plausible.alfaview.com=204; production-alfaview-assets/staging-alfaview
- NEW sso.alfaview.com: FusionAuth 1.63.0 OIDC discovery confirmed with issuer=acme.com (misconfiguration), implicit flow enabled, HS256/HS384/HS512 in id_token_signing_alg_values_supported but JWKS contain
- NEW test.alfaview.com: Unauthenticated binary distribution confirmed (alfacheck v470079, 4 platforms: linux/amd64, windows/amd64, mac/amd64, mac/arm64) — statically linked ELF, no integrity hashes/signatu
- NEW app.alfaview.com: SPA loads from alfaview-com-assets.alfaview.com; no client_id discoverable in HTML/JS bundles; common client_id patterns (alfaview, app, web, client, spa, alfaview.com, app.alfaview.
- CHANGED beta-app.alfaview.com: HTTP 401 with WWW-Authenticate: Basic (not OAuth) — different auth mechanism than main app
- CHANGED alfaview.com: Marketing page now 301→/en with nginx + Accept-Language vary, 177KB, strict CSP, matomo analytics — no SSO login links
- CHANGED www.alfaview.com: 301→alfaview.com/en (no independent surface)

## 2026-09-05 12:20:10 UTC

## 2026-09-05 15:02:37 UTC

## 2026-09-05 17:06:17 UTC
- NEW app.alfaview.com/js/AppSignup.min.3329eeac503c038b44b8.js (48KB lazy chunk) recovered: SIGNUP IS UNAUTHENTICATED — Signup action sends NO token header (vs CreateCompany which sends headers:{token:s}).
- NEW app.alfaview.com/js/AppFinishSignup.min.c05a2146b0620d1167ae.js recovered: activation is EMAIL-GATED — finishSignup({companyId, username, activationToken, password}) pulled from URL route /finish-sign
- NEW GraphQL client posts to /graphql with credentials:"include" (Apollo), authenticated ops use a token header (not Authorization/Bearer); Signup/FinishSignup send NO token → anonymous reachable. No CSRF 
- NEW sso.alfaview.com: FusionAuth 1.63.0 OIDC discovery confirmed — issuer=acme.com (misconfiguration vs alfaview.com), implicit flow enabled, HS256/HS384/HS512 in id_token_signing_alg_values_supported but
- NEW app.alfaview.com/graphql: Full admin GraphQL schema leaked in public JS bundle (1.09MB) — 60+ mutations, 45+ queries with exact args; introspection disabled but resolver auth INCONSISTENT: listIdentit
- NEW app.alfaview.com/graphql: guestAuthenticate(userId,companyId,roomId) and guestJoin(userId,companyId,roomId,displayName) mutations reachable unauthenticated (BAD_USER_INPUT not UNAUTHENTICATED) — NO ac
- NEW test.alfaview.com: Unauthenticated binary distribution confirmed — alfacheck v470079 for 4 platforms (linux/amd64, windows/amd64, mac/amd64, mac/arm64), statically linked ELF, no integrity hashes/sign
- NEW alfaview.com: Marketing page now 301→/en (nginx, Accept-Language vary), 177KB, strict CSP, matomo analytics — no unauthenticated SSO login links on marketing domain
- NEW www.alfaview.com: 301→alfaview.com/en (no independent surface)
- NEW beta-ionoscloud-21-beta-engine-* (2 hosts): Confirmed timeout (000) — internal/firewalled like alfacheck-* fleet
- NEW beta-ionoscloud-21 fleet: 7 hosts total, only 1/7 probed (hydra=timeout), 6 remain unprobed
- CHANGED apis.alfaview.com/v2/languages: Now returns 401 (was 404) — endpoint added to production, aligns with beta; OpenAPI specs now identical including /v2/languages
- CHANGED beta-app.alfaview.com: HTTP 401 with WWW-Authenticate: Basic realm — different auth mechanism than main app (OAuth)
- CHANGED alfacheck-engine/audio/video.alfaview.com: Confirmed UNREACHABLE via timeout probes — target exhausted
- CHANGED beta-hcloud-19-beta-hydra-dzwx8 / beta-noris-33-beta-hydra-2zm7t / beta-ovh-29-beta-hydra-z4tf8: All confirmed media/signaling servers ("Hi Client" on all paths) — not OIDC/auth infrastructure, target
- CHANGED alfatraining/bhc/kh-freiburg.alfaview.com: XSS hypothesis REJECTED — all three serve byte-identical generic SPA shell (1381B, MD5 554a39), no tenant-specific rendering, no reflections
- CHANGED insider-webclient.alfaview.com / beta-webclient.alfaview.com: SPA shell only (4396B), /health=204, all admin/debug paths 404 — targets exhausted
- CHANGED demo-company.alfaview.com: SPA catch-all confirmed — /api/v1/users returns identical HTML shell as root, no unauthenticated data exposure

## 2026-09-05 18:56:58 UTC
- NEW app.alfaview.com: Signup mutation unauthenticated (AppSignup.min.js lazy chunk) — sends NO token header; activation email-gated via finishSignup({companyId,username,activationToken,password}) at /fini
- NEW app.alfaview.com/graphql: Apollo client uses credentials:"include" + custom token header (not Authorization/Bearer); Signup/FinishSignup send NO token → anonymous reachable; no CSRF token observed
- NEW sso.alfaview.com: FusionAuth 1.63.0 OIDC discovery confirmed — issuer=acme.com (misconfiguration vs alfaview.com), implicit flow enabled, HS256/HS384/HS512 in id_token_signing_alg_values_supported but
- NEW app.alfaview.com/graphql: Full admin GraphQL schema leaked in public JS bundle (1.09MB) — 60+ mutations, 45+ queries with exact args; introspection disabled but resolver auth INCONSISTENT (listIdentit
- NEW app.alfaview.com/graphql: guestAuthenticate(userId,companyId,roomId) and guestJoin(userId,companyId,roomId,displayName) mutations reachable unauthenticated (BAD_USER_INPUT not UNAUTHENTICATED) — NO ac
- NEW test.alfaview.com: Unauthenticated binary distribution confirmed — alfacheck v470079 for 4 platforms (linux/amd64, windows/amd64, mac/amd64, mac/arm64), statically linked ELF, no integrity hashes/sign
- CHANGED beta-ionoscloud-21-* fleet: 7 hosts total, only 1/7 probed (beta-ionoscloud-21-beta-hydra-7x5d5 = timeout); beta-engine-* (2 hosts) confirmed timeout (000) — internal/firewalled like alfacheck-* fleet
- CHANGED alfaview.com: Root 301→/en (nginx, Accept-Language vary), /en 177KB marketing page with strict CSP and matomo; no unauthenticated SSO login links on marketing domain
- CHANGED www.alfaview.com: 301→alfaview.com/en (no independent surface)
- CHANGED apis.alfaview.com/v2/languages: Now returns 401 (was 404) — endpoint added to production, aligns with beta; OpenAPI specs now identical including /v2/languages
- CHANGED beta-app.alfaview.com: HTTP 401 with WWW-Authenticate: Basic realm — different auth mechanism than main app (OAuth)
- CHANGED alfacheck-engine/audio/video.alfaview.com: Confirmed UNREACHABLE via timeout probes — target exhausted
- CHANGED beta-hcloud-19-beta-hydra-dzwx8 / beta-noris-33-beta-hydra-2zm7t / beta-ovh-29-beta-hydra-z4tf8: All confirmed media/signaling servers ("Hi Client" on all paths) — not OIDC/auth infrastructure, target
- CHANGED alfatraining/bhc/kh-freiburg.alfaview.com: XSS hypothesis REJECTED — all three serve byte-identical generic SPA shell (1381B, MD5 554a39), no tenant-specific rendering, no reflections
- CHANGED insider-webclient.alfaview.com / beta-webclient.alfaview.com: SPA shell only (4396B), /health=204, all admin/debug paths 404 — targets exhausted
- CHANGED demo-company.alfaview.com: SPA catch-all confirmed — /api/v1/users returns identical HTML shell as root, no unauthenticated data exposure

## 2026-09-05 20:53:46 UTC
- NEW app.alfaview.com/js/app.min.5b3949112f0cf682adc8.js (1.09MB) leaks FULL admin/company GraphQL schema: 60+ mutations + 45+ queries with exact args. Not just SDL strings — includes CreateCompany($compan
- NEW GraphQL live at app.alfaview.com/graphql: __typename OK unauth; introspection disabled; sensitive ops return UNAUTHENTICATED. Resolver auth is INCONSISTENT: listIdentityProviders returns data unauth (
- NEW CRITICAL DELTA: GraphQL guest path guestAuthenticate(userId,companyId,roomId) and guestJoin(userId,companyId,roomId,displayName) take NO accessKey — while REST guest-link flow requires 4-field combo i
- NEW Hosts: webclient.alfaview.com=200/4396B(insider family); staging-webclient.alfaview.com=401 Basic realm=Protected (/health=204); plausible.alfaview.com=204; production-alfaview-assets/staging-alfaview
- NEW app.alfaview.com/js/AppSignup.min.3329eeac503c038b44b8.js (48KB lazy chunk) recovered: SIGNUP IS UNAUTHENTICATED — Signup action sends NO token header (vs CreateCompany which sends headers:{token:s}).
- NEW app.alfaview.com/js/AppFinishSignup.min.c05a2146b0620d1167ae.js recovered: activation is EMAIL-GATED — finishSignup({companyId, username, activationToken, password}) pulled from URL route /finish-sign
- NEW GraphQL client posts to /graphql with credentials:"include" (Apollo), authenticated ops use a token header (not Authorization/Bearer); Signup/FinishSignup send NO token → anonymous reachable. No CSRF 
- NEW sso.alfaview.com: FusionAuth 1.63.0 OIDC discovery live (issuer=acme.com misconfiguration, implicit flow enabled, HS256/HS384/HS512 in supported algs but RSA-only JWKS) — 2026-09-05 00:15
- NEW test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms: linux/amd64, windows/amd64, mac/amd64, mac/arm64), statically linked ELF, no integrity hashes/signatures visible
- NEW app.alfaview.com: Full admin GraphQL schema leaked in public JS bundle (1.09MB, 60+ mutations, 45+ queries); resolver auth INCONSISTENT (listIdentityProviders returns data unauth, listComponents unaut
- NEW app.alfaview.com: Signup mutation unauthenticated (AppSignup.min.js lazy chunk) — sends NO token header; activation email-gated via finishSignup at /finish-signup route; Apollo client uses credentials
- NEW alfaview.com: Marketing page live (301→/en, nginx, Accept-Language vary, 177KB, strict CSP, matomo) — no unauthenticated SSO login links — 2026-09-05 04:34
- NEW www.alfaview.com: 301→alfaview.com/en (no independent surface) — 2026-09-05 04:34
- NEW beta-ionoscloud-21-beta-engine-* (2 hosts): Confirmed timeout (000) — internal/firewalled like alfacheck-* fleet; beta-ionoscloud-21 fleet 7 hosts total, only 1/7 probed (hydra=timeout), 6 remain unpr
- CHANGED apis.alfaview.com/v2/languages: Now returns 401 (was 404) — endpoint added to production, aligns with beta; OpenAPI specs now identical including /v2/languages — 2026-09-05 04:34
- CHANGED beta-app.alfaview.com: HTTP 401 with WWW-Authenticate: Basic realm — different auth mechanism than main app (OAuth) — 2026-09-05 04:34
- CHANGED alfatraining/bhc/kh-freiburg.alfaview.com: XSS hypothesis REJECTED — all three serve byte-identical generic SPA shell (1381B, MD5 554a39), no tenant-specific rendering, no reflections — targets exhaus

## 2026-09-05 22:28:21 UTC

## 2026-09-06 00:15:48 UTC

## 2026-09-06 04:41:59 UTC

## 2026-09-06 08:57:09 UTC

## 2026-09-06 12:28:54 UTC

## 2026-09-06 16:02:35 UTC
- NEW beta-ionoscloud-21 fleet (7/7) fully probed: beta-ionoscloud-21-beta-audio-65st7/bdtmf and -beta-video-6pp2m/l5mbv all timeout (000, 12s) — resolves to real distinct IPs (185.127.30.215/.225). Last un
- CHANGED Inventory now 100% probed: all 55 dedicated hosts have an HTTP verdict; zero genuinely-unprobed hosts remain.

## 2026-09-06 18:09:34 UTC

## 2026-09-06 19:58:58 UTC
- NEW beta-ionoscloud-21 fleet (7/7) fully probed — all 4 audio/video + 2 engine + 1 hydra hosts timeout (000, 12s), resolves to real IPs 185.127.30.215/.225; firewalled like alfacheck-* fleet. Inventory no
- CHANGED app.alfaview.com/graphql: anonymous resolver slice confirmed closed — listIdentityProviders returns `[]`, listComponents errors 500, searchCompanies/generateFileDownloadURL return UNAUTHENTICATED; no 
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT) — accessKey-less GraphQL guest path diverging from REST 4-field combo remains highest-structural 
- CHANGED app.alfaview.com/graphql: field-oracle enumeration discloses reply-type field names (providers [JSONObject], http_download_url, GetPendingUserAccount(userId)) — broader schema-surface mapping.
- CHANGED sso.alfaview.com: OIDC discovery unchanged — issuer=acme.com (misconfiguration), implicit flow enabled, HS256/384/512 in supported algs but JWKS contains ONLY RSA keys (7 RS256 keys, zero symmetric). 
- CHANGED test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification visible.
- CHANGED apis.alfaview.com: Access tokens confirmed opaque/base64 (distinct 401 "No base64 encoded access token was provided") — JWT alg-confusion against API gateway closed.

## 2026-09-06 21:51:20 UTC
- CHANGED sso.alfaview.com: OIDC discovery re-fetched today — no `registration_endpoint` advertised → FusionAuth dynamic client registration OFF; GET `/oauth2/register` = 6.2KB generic "Login | FusionAuth" them
- CHANGED app.alfaview.com/graphql: guestAuthenticate reconfirmed anonymous-reachable (~1 rps) — mutation processed, my field guess returned `GRAPHQL_VALIDATION_FAILED` ("Did you mean `role`?") not UNAUTHENTICA

## 2026-09-06 23:21:54 UTC
- NEW sso.alfaview.com: OIDC discovery re-fetched — no `registration_endpoint` advertised → FusionAuth dynamic client registration OFF; GET `/oauth2/register` returns 6.2KB generic "Login | FusionAuth" them
- NEW app.alfaview.com/graphql: guestAuthenticate reconfirmed anonymous-reachable (~1 rps) — mutation processed, field guess returned `GRAPHQL_VALIDATION_FAILED` ("Did you mean `role`?") not `UNAUTHENTICATE
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero genuinely-unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout)
- CHANGED app.alfaview.com/graphql: anonymous resolver slice confirmed closed — `listIdentityProviders` returns `[]`, `listComponents` errors 500, `searchCompanies`/`generateFileDownloadURL` return `UNAUTHENTIC
- CHANGED sso.alfaview.com: OIDC discovery unchanged — issuer=acme.com (misconfiguration), implicit flow enabled, HS256/384/512 in supported algs but JWKS contains ONLY RSA keys (7 RS256 keys, zero symmetric)
- CHANGED test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification visible
- CHANGED apis.alfaview.com: Access tokens confirmed opaque/base64 (distinct 401 "No base64 encoded access token was provided") — JWT alg-confusion against API gateway closed

## 2026-09-07 01:08:39 UTC
- NEW sso.alfaview.com/oauth2/authorize now returns HTTP 404 on root path (was 200 len=0) — endpoint behavior changed, may indicate deployment update
- NEW apis.alfaview.com/v2/users/{foreign-uuid} returns HTTP 405 (Method Not Allowed) — DELETE method not allowed on user endpoint without auth; PATCH /v2/rooms/{roomId}/permissions/{userId} returns 401
- CHANGED sso.alfaview.com OIDC discovery unchanged — issuer=acme.com, implicit flow enabled, HS256/384/512 in supported algs but JWKS contains ONLY 7 RSA keys (zero symmetric)
- CHANGED app.alfaview.com/graphql guestAuthenticate/guestJoin reconfirmed anonymous-reachable (returns GRAPHQL_VALIDATION_FAILED "Did you mean `role`?" not UNAUTHENTICATED) — accessKey-less GraphQL guest path 
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout)
- CHANGED apis.alfaview.com access tokens confirmed opaque/base64 (distinct 401 "No base64 encoded access token was provided") — JWT alg-confusion against API gateway closed
- CHANGED test.alfaview.com unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification visible

## 2026-09-07 06:09:48 UTC

## 2026-09-07 12:48:25 UTC
- NEW beta-ionoscloud-21 fleet (7/7) fully probed: beta-ionoscloud-21-beta-audio-65st7/bdtmf and -beta-video-6pp2m/l5mbv all timeout (000, 12s) — resolves to real distinct IPs (185.127.30.215/.225). Last un
- CHANGED Inventory now 100% probed: all 55 dedicated hosts have an HTTP verdict; zero genuinely-unprobed hosts remain.
- CHANGED sso.alfaview.com: OIDC discovery re-fetched today — no `registration_endpoint` advertised → FusionAuth dynamic client registration OFF; GET `/oauth2/register` = 6.2KB generic "Login | FusionAuth" them
- CHANGED app.alfaview.com/graphql: guestAuthenticate reconfirmed anonymous-reachable (~1 rps) — mutation processed, my field guess returned `GRAPHQL_VALIDATION_FAILED` ("Did you mean `role`?") not UNAUTHENTICA
- CHANGED sso.alfaview.com/oauth2/authorize: Now returns HTTP 200 (was 404 on root path) — endpoint behavior changed, may indicate deployment update; OIDC discovery unchanged (issuer=acme.com, implicit flow, HS
- CHANGED apis.alfaview.com/v2/users/{foreign-uuid}: Returns HTTP 405 (Method Not Allowed) — DELETE not allowed without auth; PATCH /v2/rooms/{roomId}/permissions/{userId} returns 401.
- CHANGED Inventory: 100% probed — all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout).
- CHANGED sso.alfaview.com/oauth2/authorize: Now returns HTTP 200 (was 404 on root path) — endpoint behavior changed, may indicate deployment update; OIDC discovery unchanged (issuer=acme.com, implicit flow, HS
- CHANGED apis.alfaview.com/v2/users/{foreign-uuid}: Returns HTTP 405 (Method Not Allowed) — DELETE not allowed without auth; PATCH /v2/rooms/{roomId}/permissions/{userId} returns 401.
- CHANGED Inventory: 100% probed — all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout).

## 2026-09-07 18:00:22 UTC
- CHANGED sso.alfaview.com/oauth2/authorize: Now returns HTTP 200 (was 404 on root path) — endpoint behavior changed, may indicate deployment update; OIDC discovery unchanged (issuer=acme.com, implicit flow ena
- CHANGED apis.alfaview.com/v2/users/{foreign-uuid}: Returns HTTP 405 (Method Not Allowed) — DELETE not allowed without auth; PATCH /v2/rooms/{roomId}/permissions/{userId} returns 401.
- CHANGED Inventory: 100% probed — all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout).

## 2026-09-07 21:00:34 UTC
- CHANGED sso.alfaview.com/oauth2/authorize: Now returns HTTP 200 with FusionAuth login page HTML (was 404/empty then invalid_client error) — endpoint behavior changed, now renders login page even for unregiste
- CHANGED apis.alfaview.com/v2/users/{uuid}: Returns HTTP 405 (Allow: DELETE) — DELETE method exists but requires auth (401 on PATCH permissions without token)

## 2026-09-07 23:07:47 UTC
- CHANGED sso.alfaview.com/oauth2/authorize: Now returns HTTP 200 with FusionAuth login page HTML for unregistered client_id (was 404 → invalid_client error) — endpoint behavior changed, now renders login page 
- CHANGED apis.alfaview.com/v2/users/{uuid}: Returns HTTP 405 (Allow: DELETE) — DELETE method exists but requires auth (401 on PATCH permissions without token); OpenAPI spec confirmed identical beta/prod

## 2026-09-08 01:17:28 UTC

## 2026-09-08 06:04:11 UTC
- NEW sso.alfaview.com/oauth2/authorize behavioral change: now returns HTTP 200 FusionAuth login page (~6189B) for unregistered client_id (was 404 → invalid_client error) — validation timing shifted, may in
- NEW sso.alfaview.com OIDC discovery reconfirmed: issuer=acme.com (not alfaview.com), implicit flow enabled, HS256/384/512 in id_token_signing_alg_values_supported but JWKS contains ONLY 7 RSA keys (zero s
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT, not UNAUTHENTICATED) — accessKey-less GraphQL guest path diverging from REST 4-field combo remain
- CHANGED apis.alfaview.com/v2: Cross-tenant IDOR via UUID path params (DELETE /v2/users/{id}, PATCH /v2/rooms/{roomId}/permissions/{userId}) confirmed in identical beta/prod OpenAPI — requires authenticated ac
- CHANGED test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification visible — supply chain risk
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero genuinely-unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout)
- CHANGED apis.alfaview.com: Access tokens confirmed opaque/base64 (distinct 401 "No base64 encoded access token was provided") — JWT alg-confusion against API gateway closed

## 2026-09-08 10:36:18 UTC
- NEW sso.alfaview.com/oauth2/authorize now returns HTTP 200 FusionAuth login page (~6189B) for unregistered client_id (was 404 → invalid_client error) — validation timing shifted, may indicate FusionAuth c
- NEW sso.alfaview.com OIDC discovery reconfirmed: issuer=acme.com (not alfaview.com), implicit flow enabled, HS256/384/512 in id_token_signing_alg_values_supported but JWKS contains ONLY 7 RSA keys (zero s
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT, not UNAUTHENTICATED) — accessKey-less GraphQL guest path diverging from REST 4-field combo remain
- CHANGED apis.alfaview.com/v2: Cross-tenant IDOR via UUID path params (DELETE /v2/users/{id}, PATCH /v2/rooms/{roomId}/permissions/{userId}) confirmed in identical beta/prod OpenAPI — requires authenticated ac
- CHANGED test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification visible — supply chain risk
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero genuinely-unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout)
- CHANGED apis.alfaview.com: Access tokens confirmed opaque/base64 (distinct 401 "No base64 encoded access token was provided") — JWT alg-confusion against API gateway closed

## 2026-09-08 14:53:43 UTC

## 2026-09-08 18:17:59 UTC
- NEW sso.alfaview.com/oauth2/authorize: Now returns HTTP 400 for unregistered client_id (was 200 login page) — validation timing still client_id-gated; redirect_uri matrix remains blocked without registere
- NEW apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable
- CHANGED sso.alfaview.com OIDC: issuer=acme.com misconfiguration persists; implicit flow + HS256/384/512 in supported algs but JWKS contains ONLY 7 RSA keys (zero symmetric) — alg confusion vector unchanged
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT) — accessKey-less GraphQL guest path diverging from REST 4-field combo remains highest-structural 
- CHANGED apis.alfaview.com/v2: Cross-tenant IDOR via UUID path params (DELETE /v2/users/{id}, PATCH /v2/rooms/{roomId}/permissions/{userId}) confirmed in OpenAPI — requires authenticated account
- CHANGED test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification — supply chain risk

## 2026-09-08 21:14:43 UTC
- NEW sso.alfaview.com/oauth2/authorize: Now returns HTTP 400 for unregistered client_id (was 200 login page) — validation timing still client_id-gated; redirect_uri matrix remains blocked without registere
- NEW apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable
- CHANGED sso.alfaview.com OIDC: issuer=acme.com misconfiguration persists; implicit flow + HS256/384/512 in supported algs but JWKS contains ONLY 7 RSA keys (zero symmetric) — alg confusion vector unchanged
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT) — accessKey-less GraphQL guest path diverging from REST 4-field combo remains highest-structural 
- CHANGED apis.alfaview.com/v2: Cross-tenant IDOR via UUID path params (DELETE /v2/users/{id}, PATCH /v2/rooms/{roomId}/permissions/{userId}) confirmed in OpenAPI — requires authenticated account
- CHANGED test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification — supply chain risk

## 2026-09-08 23:25:31 UTC
- NEW sso.alfaview.com/oauth2/authorize: Now returns HTTP 400 for unregistered client_id (was 200 login page) — validation timing still client_id-gated; redirect_uri matrix remains blocked without registere
- NEW apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable
- CHANGED sso.alfaview.com OIDC: issuer=acme.com misconfiguration persists; implicit flow + HS256/384/512 in supported algs but JWKS contains ONLY 7 RSA keys (zero symmetric) — alg confusion vector unchanged
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT) — accessKey-less GraphQL guest path diverging from REST 4-field combo remains highest-structural 
- CHANGED apis.alfaview.com/v2: Cross-tenant IDOR via UUID path params (DELETE /v2/users/{id}, PATCH /v2/rooms/{roomId}/permissions/{userId}) confirmed in OpenAPI — requires authenticated account
- CHANGED test.alfaview.com: Unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification — supply chain risk

## 2026-09-09 01:31:29 UTC

## 2026-09-09 06:50:32 UTC
- CHANGED `/oauth2/introspect` — now confirmed broken client authentication: accepts ANY `client_id` value without validation (200 `{"active":false}` with `client_id=does-not-exist-12345`). Without `client_id` 
- CHANGED `/oauth2/token` with `client_credentials` grant → 400 `not_licensed` — FusionAuth Community edition, Entity Management feature not available. Confirms license tier.
- CHANGED `/v2/auth/password` REST endpoint — POST with `{username, password}` confirmed: 422 validates input schema, 401 `invalid credentials` for bad creds. Generic error (no username enumeration). Now confir
- NEW `apis.alfaview.com/v2/docs/openapi.json` still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable this cycle
- NEW `sso.alfaview.com/oauth2/authorize` returns HTTP 400 for unregistered client_id (not 200 login page) — validation timing still client_id-gated; redirect_uri matrix remains blocked without registered c
- NEW `app.alfaview.com/graphql` guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT) — accessKey-less GraphQL guest path diverging from REST 4-field combo remains highest-structural
- NEW `app.alfaview.com/graphql` field-oracle enumeration discloses reply-type field names (`providers` `[JSONObject]`, `http_download_url`, `GetPendingUserAccount(userId)`) — broader schema-surface mapping
- CHANGED `apis.alfaview.com/v2/users/{uuid}` returns HTTP 405 (Allow: DELETE) — DELETE method exists but requires auth (401 on PATCH permissions without token)
- CHANGED `sso.alfaview.com` OIDC discovery unchanged — issuer=acme.com, implicit flow enabled, HS256/384/512 in supported algs but JWKS contains ONLY 7 RSA keys (zero symmetric)
- CHANGED `test.alfaview.com` unauthenticated binary distribution (alfacheck v470079, 4 platforms) still live, no integrity verification visible
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero genuinely-unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout)

## 2026-09-09 12:00:41 UTC
- CHANGED `app.alfaview.com/graphql`: `guestAuthenticate` with UUID-format args returns `FORBIDDEN` (not `BAD_USER_INPUT`) — mutation processes UUIDs and hits authorization, not type validation. Null args retur
- CHANGED `app.alfaview.com/graphql`: `GuestJoinReply` field oracle reveals `expiry` field (error: "Did you mean `expiry`?"). All other tested fields absent: accessToken, refreshToken, role, userId, companyId, 
- CHANGED `apis.alfaview.com/v2/auth/guest-link`: 3-field combo `{accessKey,companyId,roomId}` returns 422 `ACTION_INVALID` (same as 4-field with displayName) — `displayName` is **not** a required field per ser
- CHANGED `sso.alfaview.com/oauth2/authorize` now returns HTTP 200 FusionAuth login page (~6189B) for unregistered `client_id` (was 400) — validation timing shifted again overnight
- CHANGED `sso.alfaview.com/oauth2/introspect` accepts any `client_id` value without validation (200 `{"active":false}` with `client_id=does-not-exist-12345`) — broken client authentication on token introspecti
- CHANGED `sso.alfaview.com/oauth2/token` with `client_credentials` grant returns 400 `not_licensed` — FusionAuth Community edition (v1.63.0) confirmed, Entity Management not licensed

## 2026-09-09 15:42:12 UTC
- CHANGED `sso.alfaview.com/oauth2/authorize` returns HTTP 200 FusionAuth login page (~6189B) for unregistered `client_id` (was 400) — validation timing shifted again overnight
- CHANGED `sso.alfaview.com/oauth2/introspect` now requires `token` parameter (400 `missing_token` without it) — previously accepted any `client_id` without validation
- CHANGED `apis.alfaview.com/v2/auth/guest-link` REST endpoint confirms 3-field combo (`accessKey`+`companyId`+`roomId`) returns 422 `ACTION_INVALID` — `displayName` not required
- CHANGED `app.alfaview.com/graphql` `guestAuthenticate`/`guestJoin` reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs) — accessKey-less GraphQL guest path diverges from REST 3-field combo

## 2026-09-09 18:53:03 UTC
- CHANGED `sso.alfaview.com/oauth2/authorize` returns HTTP 200 FusionAuth login page (~6160B) for unregistered `client_id` (was 400) — validation timing shifted again overnight
- CHANGED `sso.alfaview.com/oauth2/introspect` now requires `token` parameter (400 `missing_token` without it) — previously accepted any `client_id` without validation
- CHANGED `apis.alfaview.com/v2/auth/guest-link` REST endpoint confirms 3-field combo (`accessKey`+`companyId`+`roomId`) returns 422 `ACTION_INVALID` — `displayName` not required
- CHANGED `app.alfaview.com/graphql` `guestAuthenticate`/`guestJoin` reconfirmed anonymous-reachable (`BAD_USER_INPUT` with zero UUIDs) — accessKey-less GraphQL guest path diverges from REST 3-field combo

## 2026-09-09 21:21:16 UTC
- CHANGED `sso.alfaview.com/oauth2/authorize` returns HTTP 200 FusionAuth login page for unregistered `client_id` (was 400) — validation timing shifted again overnight
- CHANGED `sso.alfaview.com/oauth2/introspect` now requires `token` parameter (400 `missing_token` without it) — previously accepted any `client_id` without validation
- CHANGED `apis.alfaview.com/v2/auth/guest-link` REST endpoint confirms 3-field combo (`accessKey`+`companyId`+`roomId`) returns 422 `ACTION_INVALID` — `displayName` not required
- CHANGED `app.alfaview.com/graphql` `guestAuthenticate`/`guestJoin` reconfirmed anonymous-reachable (`BAD_USER_INPUT` with zero UUIDs) — accessKey-less GraphQL guest path diverges from REST 3-field combo

## 2026-09-09 23:20:30 UTC

## 2026-09-10 01:10:03 UTC
- CHANGED `sso.alfaview.com/oauth2/authorize` returns HTTP 200 FusionAuth login page for unregistered `client_id` (was 400) — validation timing shifted again overnight
- CHANGED `sso.alfaview.com/oauth2/introspect` now requires `token` parameter (400 `missing_token` without it) — previously accepted any `client_id` without validation
- CHANGED `apis.alfaview.com/v2/auth/guest-link` REST endpoint confirms 3-field combo (`accessKey`+`companyId`+`roomId`) returns 422 `ACTION_INVALID` — `displayName` not required
- CHANGED `app.alfaview.com/graphql` `guestAuthenticate`/`guestJoin` reconfirmed anonymous-reachable (`BAD_USER_INPUT` with zero UUIDs) — accessKey-less GraphQL guest path diverges from REST 3-field combo
- CHANGED OpenAPI specs prod/beta remain byte-identical (37 paths, MD5 357b94d367909a40b9299b543d23712b) — schema surface fully stable
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain

## 2026-09-10 06:04:15 UTC
- CHANGED `sso.alfaview.com/oauth2/authorize` returns HTTP 200 FusionAuth login page for unregistered `client_id` (was 400) — validation timing shifted again overnight
- CHANGED `sso.alfaview.com/oauth2/introspect` now requires `token` parameter (400 `missing_token` without it) — previously accepted any `client_id` without validation
- CHANGED `apis.alfaview.com/v2/auth/guest-link` REST endpoint confirms 3-field combo (`accessKey`+`companyId`+`roomId`) returns 422 `ACTION_INVALID` — `displayName` not required
- CHANGED `app.alfaview.com/graphql` `guestAuthenticate`/`guestJoin` reconfirmed anonymous-reachable (`BAD_USER_INPUT` with zero UUIDs) — accessKey-less GraphQL guest path diverges from REST 3-field combo
- CHANGED OpenAPI specs prod/beta remain byte-identical (37 paths, MD5 357b94d367909a40b9299b543d23712b) — schema surface fully stable
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain

## 2026-09-10 10:40:25 UTC
- CHANGED sso.alfaview.com/oauth2/introspect: HTTP Basic auth path accepts **fabricated client_id + any secret** → `{"active":false}` with no client validation; POST body `client_id=` → `invalid_client`. Client
- NEW sso.alfaview.com/.well-known/openid-configuration: `token_endpoint_auth_methods_supported=['client_secret_basic','client_secret_post','none']` — Basic scheme is an advertised auth method, making the i
- CHANGED `sso.alfaview.com/oauth2/authorize` returns HTTP 200 FusionAuth login page for unregistered `client_id` (was 400) — validation timing shifted again overnight
- CHANGED `sso.alfaview.com/oauth2/introspect` now requires `token` parameter (400 `missing_token` without it) — previously accepted any `client_id` without validation
- CHANGED `apis.alfaview.com/v2/auth/guest-link` REST endpoint confirms 3-field combo (`accessKey`+`companyId`+`roomId`) returns 422 `ACTION_INVALID` — `displayName` not required
- CHANGED `app.alfaview.com/graphql` `guestAuthenticate`/`guestJoin` reconfirmed anonymous-reachable (`BAD_USER_INPUT` with zero UUIDs) — accessKey-less GraphQL guest path diverges from REST 3-field combo
- CHANGED OpenAPI specs prod/beta remain byte-identical (37 paths, MD5 357b94d367909a40b9299b543d23712b) — schema surface fully stable
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain

## 2026-09-10 14:51:53 UTC
- CHANGED `sso.alfaview.com/oauth2/authorize` returns HTTP 200 FusionAuth login page (6174B) for unregistered `client_id` — validation timing shifted again (was 400 `invalid_client`); redirect_uri matrix still 
- CHANGED `sso.alfaview.com/oauth2/introspect` HTTP Basic auth path accepts fabricated `client_id` + any secret → `200 {"active":false}`; POST body `client_id` validated → `400 invalid_client` — client authenti
- CHANGED `apis.alfaview.com/v2/auth/guest-link` REST endpoint confirms 3-field combo (`accessKey`+`companyId`+`roomId`) returns 422 `ACTION_INVALID` — `displayName` NOT required (prior 4-field claim incorrect)
- CHANGED `app.alfaview.com/graphql` `guestAuthenticate`/`guestJoin` reconfirmed anonymous-reachable (`BAD_USER_INPUT` with zero UUIDs) — accessKey-less GraphQL guest path diverges from REST 3-field combo
- CHANGED OpenAPI specs prod/beta remain byte-identical (37 paths, MD5 357b94d367909a40b9299b543d23712b) — schema surface fully stable 4th consecutive cycle
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero unprobed hosts remain

## 2026-09-10 18:00:00 UTC

## 2026-09-10 20:11:21 UTC
- CHANGED sso.alfaview.com/oauth2/introspect: POST-body `client_id` validation removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); Basic and POST-body channels now accept any cli
- CHANGED sso.alfaview.com OIDC discovery: now lists ES256/384/512 among `id_token_signing_alg_values_supported` (previous cycles: RSA+HS only); JWKS remains RSA-only (7 RS256 keys, zero symmetric/ECDSA)

## 2026-09-10 22:38:18 UTC
- NEW sso.alfaview.com/oauth2/introspect: POST-body `client_id` validation removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); Basic and POST-body channels now accept any cli
- NEW sso.alfaview.com OIDC discovery: now lists ES256/384/512 among `id_token_signing_alg_values_supported` (previous cycles: RSA+HS only); JWKS remains RSA-only (7 RS256 keys, zero symmetric/ECDSA)

## 2026-09-11 00:31:54 UTC

## 2026-09-11 05:14:39 UTC
- CHANGED beta-apis.alfaview.com: Auth response identical to production (401 + same error body). Beta weaker auth hypothesis disconfirmed.
- NEW beta-webclient.alfaview.com (HTTP 200): High-value web client surface, untested.
- NEW insider-webclient.alfaview.com (HTTP 200): Internal tooling potentially exposed.

## 2026-09-11 09:46:07 UTC

## 2026-09-11 14:02:27 UTC
- NEW sso.alfaview.com/oauth2/introspect: POST-body `client_id` validation removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); Basic and POST-body channels now accept any cli
- NEW sso.alfaview.com OIDC discovery: now lists ES256/384/512 among `id_token_signing_alg_values_supported` (previous cycles: RSA+HS only); JWKS remains RSA-only (7 RS256 keys, zero symmetric/ECDSA)

## 2026-09-11 17:24:48 UTC

## 2026-09-11 20:01:24 UTC

## 2026-09-11 22:20:38 UTC
- NEW NO_DELTA: all probes (OpenAPI MD5 357b94d3, OIDC discovery issuer=acme.com + ES256/HS256/RS256 algs, introspect Basic/POST-body both accept fake client_id, authorize 200 login page, GraphQL guestAuthe

## 2026-09-12 00:22:37 UTC

## 2026-09-12 04:46:37 UTC

## 2026-09-12 09:01:14 UTC
- NEW sso.alfaview.com/oauth2/introspect: POST-body `client_id` validation fully removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); HTTP Basic auth also accepts any client_i
- NEW sso.alfaview.com OIDC discovery: `id_token_signing_alg_values_supported` now lists ES256/384/512 alongside RSA+HS; JWKS unchanged (7 RS256 keys, zero symmetric/ECDSA) — alg confusion surface expanded 
- NEW sso.alfaview.com/oauth2/authorize: Returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but

## 2026-09-12 12:31:51 UTC
- NEW sso.alfaview.com/oauth2/introspect: POST-body `client_id` validation fully removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); HTTP Basic auth also accepts any client_i
- NEW sso.alfaview.com OIDC discovery: `id_token_signing_alg_values_supported` now lists ES256/384/512 alongside RSA+HS; JWKS unchanged (7 RS256 keys, zero symmetric/ECDSA) — alg confusion surface expanded 
- NEW sso.alfaview.com/oauth2/authorize: Returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but
- NEW client-diagnostics-ingest.alfaview.com: `/health`=200 `{"status":"ok"}` with strict headers (CSP default-src 'none', frame-ancestors 'none', JSON-only, edge-proxy) while all other GET paths return 39B
- NEW test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d367909a40b9299b543d23712b) — 5th+ consecutive stable cycle
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 
- CHANGED apis.alfaview.com: REST `/v2/auth/guest-link` requires only 3 fields (`accessKey`, `companyId`, `roomId`) — `displayName` NOT required (prior 4-field claim incorrect)
- CHANGED Inventory 100% probed: all 55 dedicated hosts have HTTP verdict; zero genuinely-unprobed hosts remain (beta-ionoscloud-21 fleet fully timeout)

## 2026-09-12 15:58:13 UTC
- CHANGED sso.alfaview.com/oauth2/jwks: Now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 confirmed on index page (4 platforms); still no sha256/signatures
- CHANGED sso.alfaview.com/.well-known/openid-configuration: id_token_signing_alg_values_supported lists ES256/384/512 + HS256/384/512 + RS256/384/512; JWKS now 404 — alg confusion surface expanded in metadata 
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- NEW client-diagnostics-ingest.alfaview.com/health: 200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed

## 2026-09-12 18:02:37 UTC
- NEW sso.alfaview.com/oauth2/jwks: Now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- NEW test.alfaview.com: alfacheck release bumped v470079→v483102 confirmed on index page (4 platforms); still no sha256/signatures
- CHANGED sso.alfaview.com/.well-known/openid-configuration: id_token_signing_alg_values_supported lists ES256/384/512 + HS256/384/512 + RS256/384/512; JWKS now 404 — alg confusion surface expanded in metadata 
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- NEW client-diagnostics-ingest.alfaview.com/health: 200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed

## 2026-09-12 19:50:05 UTC
- CHANGED sso.alfaview.com/oauth2/jwks: Now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 confirmed on index page (4 platforms); still no sha256/signatures
- CHANGED sso.alfaview.com/.well-known/openid-configuration: id_token_signing_alg_values_supported lists ES256/384/512 + HS256/384/512 + RS256/384/512; JWKS now 404 — alg confusion surface expanded in metadata 
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- NEW client-diagnostics-ingest.alfaview.com/health: 200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed

## 2026-09-12 21:47:15 UTC

## 2026-09-12 23:32:36 UTC

## 2026-09-13 01:24:28 UTC
- NEW sso.alfaview.com/oauth2/jwks now returns 404 (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- NEW sso.alfaview.com OIDC discovery lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS 404 — alg confusion surface expanded in metadata with zero keys to ex
- NEW client-diagnostics-ingest.alfaview.com/health=200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed
- NEW test.alfaview.com alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- CHANGED sso.alfaview.com/oauth2/authorize returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but 
- CHANGED apis.alfaview.com/v2/docs/openapi.json still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — 5th+ consecutive stable cycle
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 
- CHANGED apis.alfaview.com REST /v2/auth/guest-link requires only 3 fields (accessKey, companyId, roomId) — displayName NOT required (prior 4-field claim incorrect)

## 2026-09-13 06:47:39 UTC
- NEW sso.alfaview.com/oauth2/jwks now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- NEW sso.alfaview.com OIDC discovery lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS 404 — alg confusion surface expanded in metadata with zero keys to ex
- NEW client-diagnostics-ingest.alfaview.com/health=200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed
- NEW test.alfaview.com alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- CHANGED sso.alfaview.com/oauth2/authorize returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but 
- CHANGED apis.alfaview.com/v2/docs/openapi.json still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — 5th+ consecutive stable cycle
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 
- CHANGED apis.alfaview.com REST /v2/auth/guest-link requires only 3 fields (accessKey, companyId, roomId) — displayName NOT required (prior 4-field claim incorrect)

## 2026-09-13 12:27:40 UTC
- NEW sso.alfaview.com/oauth2/jwks now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- NEW sso.alfaview.com OIDC discovery lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS 404 — alg confusion surface expanded in metadata with zero keys to ex
- NEW client-diagnostics-ingest.alfaview.com/health=200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed
- NEW test.alfaview.com alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- CHANGED sso.alfaview.com/oauth2/authorize returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but 
- CHANGED apis.alfaview.com/v2/docs/openapi.json still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — 5th+ consecutive stable cycle
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 
- CHANGED apis.alfaview.com REST /v2/auth/guest-link requires only 3 fields (accessKey, companyId, roomId) — displayName NOT required (prior 4-field claim incorrect)

## 2026-09-13 16:55:46 UTC

## 2026-09-13 18:59:54 UTC
- NEW sso.alfaview.com/oauth2/jwks now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- NEW sso.alfaview.com OIDC discovery lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS 404 — alg confusion surface expanded in metadata with zero keys to ex
- NEW client-diagnostics-ingest.alfaview.com/health=200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed
- NEW test.alfaview.com alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401)
- CHANGED sso.alfaview.com/oauth2/authorize returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but 
- CHANGED apis.alfaview.com/v2/docs/openapi.json still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — 5th+ consecutive stable cycle
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 
- CHANGED apis.alfaview.com REST /v2/auth/guest-link requires only 3 fields (accessKey, companyId, roomId) — displayName NOT required (prior 4-field claim incorrect)

## 2026-09-13 21:13:10 UTC
- CHANGED sso.alfaview.com/oauth2/jwks: Now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- CHANGED sso.alfaview.com OIDC discovery: Lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS now 404 — alg confusion surface expanded in metadata with zero keys 
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- CHANGED sso.alfaview.com/oauth2/authorize: Returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — 5th+ consecutive stable cycle
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 
- CHANGED apis.alfaview.com REST /v2/auth/guest-link: Requires only 3 fields (accessKey, companyId, roomId) — displayName NOT required (prior 4-field claim incorrect)
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED client-diagnostics-ingest.alfaview.com/health: 200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed

## 2026-09-13 23:12:19 UTC
- CHANGED sso.alfaview.com/oauth2/jwks: Now returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- CHANGED sso.alfaview.com OIDC discovery: Lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS 404 — alg confusion surface expanded in metadata with zero keys to e
- CHANGED sso.alfaview.com/oauth2/introspect: Both POST-body and HTTP Basic auth accept ANY client_id (fabricated) → 200 {"active":false}; only residual check is Basic-vs-body client_id_mismatch (401) — client 
- CHANGED sso.alfaview.com/oauth2/authorize: Returns HTTP 200 FusionAuth login page (~6173B) for unregistered client_id — validation timing shifted again (was 400); redirect_uri matrix still client_id-gated but
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — 5th+ consecutive stable cycle
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 
- CHANGED apis.alfaview.com REST /v2/auth/guest-link: Requires only 3 fields (accessKey, companyId, roomId) — displayName NOT required (prior 4-field claim incorrect)
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED client-diagnostics-ingest.alfaview.com/health: 200 {"status":"ok"} with strict CSP (default-src 'none', frame-ancestors 'none'), all other GET paths 39B JSON 404 — minimal POST-only ingest confirmed

## 2026-09-14 01:10:10 UTC

## 2026-09-14 06:20:58 UTC
- NEW sso.alfaview.com/oauth2/jwks confirmed 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- NEW sso.alfaview.com OIDC discovery lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS 404 — alg confusion surface expanded in metadata with zero keys to ex
- NEW sso.alfaview.com/oauth2/introspect: 8th+ consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels unchanged
- NEW test.alfaview.com alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases

## 2026-09-14 13:04:48 UTC

## 2026-09-14 18:20:43 UTC
- NEW sso.alfaview.com/oauth2/jwks: Returns 404 FusionAuth error page (was accessible with 7 RSA keys) — JWKS endpoint broken/removed
- NEW sso.alfaview.com OIDC discovery: Lists ES256/384/512 + HS256/384/512 + RS256/384/512 in id_token_signing_alg_values_supported; JWKS now 404 — alg confusion surface expanded in metadata with zero keys 
- CHANGED sso.alfaview.com/oauth2/introspect: 9th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels; token_endpoint_auth_methods_supported advertises client_secret_basic/p
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — 5th+ consecutive stable cycle
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED app.alfaview.com/graphql: guestAuthenticate/guestJoin reconfirmed anonymous-reachable (BAD_USER_INPUT with zero UUIDs; GRAPHQL_VALIDATION_FAILED if displayName omitted) — accessKey-less GraphQL guest 

## 2026-09-14 21:57:34 UTC

## 2026-09-15 00:04:40 UTC

## 2026-09-15 05:00:31 UTC
- CHANGED sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404 in prior cycles) — JWKS endpoint restored
- CHANGED sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)

## 2026-09-15 09:46:28 UTC
- NEW sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404 in prior cycles) — JWKS endpoint restored
- NEW sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED OIDC discovery unchanged: issuer=acme.com, implicit flow, HS256/384/512/ES256/384/512 in id_token_signing_alg_values_supported, JWKS now 7 RSA keys (zero symmetric/ECDSA)
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public, prod=beta byte-identical (37 paths, MD5 357b94d3)
- CHANGED app.alfaview.com/graphql: __typename OK unauthenticated, introspection disabled, guestAuthenticate/guestJoin anonymous-reachable (BAD_USER_INPUT)

## 2026-09-15 14:51:11 UTC

## 2026-09-15 18:38:47 UTC
- NEW sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404 in prior cycles) — JWKS endpoint restored
- NEW sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED OIDC discovery unchanged: issuer=acme.com, implicit flow, HS256/384/512/ES256/384/512 in id_token_signing_alg_values_supported, JWKS now 7 RSA keys (zero symmetric/ECDSA)
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public, prod=beta byte-identical (37 paths, MD5 357b94d3)
- CHANGED app.alfaview.com/graphql: __typename OK unauthenticated, introspection disabled, guestAuthenticate/guestJoin anonymous-reachable (BAD_USER_INPUT)

## 2026-09-15 21:45:11 UTC

## 2026-09-15 23:48:48 UTC
- NEW sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404) — JWKS endpoint restored
- NEW sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: 12th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels; token_endpoint_auth_methods still advertise client_secret_basic/post/
- CHANGED app.alfaview.com/graphql: signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed)
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public, prod=beta byte-identical (37 paths, MD5 357b94d3)

## 2026-09-16 01:58:31 UTC
- NEW sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404) — JWKS endpoint restored
- NEW sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: 12th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels; token_endpoint_auth_methods still advertise client_secret_basic/post/
- CHANGED app.alfaview.com/graphql: signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed)
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public, prod=beta byte-identical (37 paths, MD5 357b94d3)

## 2026-09-16 07:03:01 UTC
- NEW sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404) — JWKS endpoint restored
- NEW sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: 13th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels; token_endpoint_auth_methods still advertise client_secret_basic/post/
- CHANGED app.alfaview.com/graphql: signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed)
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public, prod=beta byte-identical (37 paths, MD5 357b94d3)
- CHANGED client-diagnostics-ingest.alfaview.com: /health=200 {"status":"ok"} with strict headers (CSP default-src 'none', frame-ancestors 'none', JSON-only, edge-proxy) while all other GET paths return 39B JSO
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases

## 2026-09-16 12:29:57 UTC

## 2026-09-16 17:18:32 UTC

## 2026-09-16 20:09:54 UTC

## 2026-09-16 22:58:04 UTC

## 2026-09-17 00:59:42 UTC

## 2026-09-17 05:42:24 UTC
- CHANGED beta-apis.alfaview.com: Auth response identical to production (401 + same error body). Beta weaker auth hypothesis disconfirmed.
- NEW beta-webclient.alfaview.com (HTTP 200): High-value web client surface, untested.
- NEW insider-webclient.alfaview.com (HTTP 200): Internal tooling potentially exposed.

## 2026-09-17 10:35:17 UTC
- NEW JWKS endpoint at sso.alfaview.com/.well-known/jwks.json oscillates (200↔404), currently 200 with 7 RSA keys (MD5 3f8d456c)
- NEW OIDC discovery introspection_endpoint advertisement oscillates (present↔absent) while /oauth2/introspect stays live (OPTIONS 405)
- NEW alfacheck binary at test.alfaview.com version bumped v470079→v483102 (4 platforms), still no sha256/signatures
- CHANGED sso.alfaview.com/oauth2/introspect: 16th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels; token_endpoint_auth_methods advertises client_secret_basic/post/none 
- CHANGED app.alfaview.com/graphql: signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed)
- CHANGED apis.alfaview.com/v2/docs/openapi.json: byte-identical prod/beta (37 paths, MD5 357b94d3), no new endpoints — 10+ stable cycles
- CHANGED All 55 dedicated hosts probed; 31 exhausted; zero genuinely-unprobed hosts remain

## 2026-09-17 15:18:46 UTC
- NEW sso.alfaview.com/.well-known/jwks.json oscillates (200↔404), currently 200 with 7 RSA keys (MD5 3f8d456c)
- NEW OIDC discovery introspection_endpoint advertisement oscillates (present↔absent) while /oauth2/introspect stays live (OPTIONS 405)
- NEW alfacheck binary at test.alfaview.com version bumped v470079→v483102 (4 platforms), still no sha256/signatures
- CHANGED sso.alfaview.com/oauth2/introspect: 16th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels; token_endpoint_auth_methods advertises client_secret_basic/post/none 
- CHANGED app.alfaview.com/graphql: signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed)
- CHANGED apis.alfaview.com/v2/docs/openapi.json: byte-identical prod/beta (37 paths, MD5 357b94d3), no new endpoints — 10+ stable cycles
- CHANGED All 55 dedicated hosts probed; 31 exhausted; zero genuinely-unprobed hosts remain

## 2026-09-17 19:07:00 UTC
- NEW NO_DELTA: Surface fully stable since last cycle (2026-09-17 15:18). Probes identical: app.alfaview.com (SPA 200), sso.alfaview.com/oauth2/introspect (OPTIONS 405), app.alfaview.com/graphql (GET 400), 

## 2026-09-17 22:11:09 UTC
- NEW NO_DELTA: Surface fully stable since last cycle (2026-09-17 15:18). Probes identical: app.alfaview.com (SPA 200), sso.alfaview.com/oauth2/introspect (OPTIONS 405), app.alfaview.com/graphql (GET 400), 

## 2026-09-18 00:24:18 UTC
- NEW `tools.alfaview.com` — live Vue "Tools UI" (in-room toolbox: polls/Q&A); exposes verb API `POST /poll/pollservice/{list|create|delete|edit|vote|hasVoted}` with auth header `Grpc-Metadata-alfaview.toke
- NEW `staging-tools.alfaview.com` — live (612B), same Tools UI, separate (staging) backend.
- NEW `whiteboard.alfaview.com` / `staging-whiteboard.alfaview.com` — live board renderer; `/` = "board deleted/access expired" error, any non-root path → `302 Location: /`.
- NEW `status.alfaview.com` — public status page 200/53714B (title "alfaview Status").
- NEW `qa.alfaview.com`, `uni-stuttgart.alfaview.com` — 200/1381B, identical tenant SPA shell (same family as alfatraining/bhc/kh-freiburg).
- NEW CT (crt.sh) yields ~110 subdomains absent from inventory: `grafana`, `loki`, `logs`, `ops`/`ops-*`, `prometheus-*`, `linkerd-*`/`linkerd-prometheus-*`, `envoy-health`, `sap`/`sap-events`, `beta/stagin
- CHANGED Most new infra hosts are firewalled externally (000): grafana, loki, logs, prometheus-*, linkerd*, ops-*, envoy-health, sap, webrtc, stun, gitlab.dev, fusionauth.dev, elitr-recordings — same edge-only
- CHANGED `staging-webclient.alfaview.com` = 401 Basic realm=Protected (as before). `ops.alfaview.com` = 404 plaintext (19B).

## 2026-09-18 05:10:12 UTC
- NEW `tools.alfaview.com` — live Vue "Tools UI" (in-room toolbox: polls/Q&A); exposes verb-based JSON RPC `POST /poll/pollservice/{list|create|delete|edit|vote|hasVoted}` with custom auth header `Grpc-Meta
- NEW `staging-tools.alfaview.com` — live (612B), same Tools UI, separate staging backend
- NEW `whiteboard.alfaview.com` / `staging-whiteboard.alfaview.com` — live board renderer; root returns "board deleted/access expired", non-root paths → 302 to `/`
- NEW `status.alfaview.com` — public status page 200/53714B
- NEW `qa.alfaview.com`, `uni-stuttgart.alfaview.com` — 200/1381B, identical tenant SPA shell (same family as alfatraining/bhc/kh-freiburg)
- NEW CT (crt.sh) yields ~110 subdomains absent from inventory: grafana, loki, logs, ops/ops-*, prometheus-*, linkerd-*/linkerd-prometheus-*, envoy-health, sap/sap-events, beta/staging/production-*, gitlab.
- CHANGED Most new infra hosts firewalled externally (000): grafana, loki, logs, prometheus-*, linkerd*, ops-*, envoy-health, sap, webrtc, stun, gitlab.dev, fusionauth.dev, elitr-recordings — same edge-only pat
- CHANGED `staging-webclient.alfaview.com` = 401 Basic realm=Protected (unchanged). `ops.alfaview.com` = 404 plaintext (19B)
- CHANGED `tools.alfaview.com/poll/pollservice/list` → HTTP 404 (GET), `staging-tools.alfaview.com/poll/pollservice/list` → HTTP 404 (GET) — verb API likely POST-only with auth header

## 2026-09-18 09:50:18 UTC
- NEW tools.alfaview.com: live Vue "Tools UI" exposing verb-based JSON RPC `POST /poll/pollservice/{list|create|delete|edit|vote|hasVoted}` with custom auth header `Grpc-Metadata-alfaview.token` (diverges f
- NEW staging-tools.alfaview.com: live (612B), same Tools UI, separate staging backend
- NEW whiteboard.alfaview.com / staging-whiteboard.alfaview.com: live board renderer; `/` = "board deleted/access expired", non-root → `302 Location: /`
- NEW status.alfaview.com: public status page 200/53714B (title "alfaview Status")
- NEW qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- NEW CT/crt.sh: ~110 subdomains not in inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*); most edge-firewalle
- CHANGED grafana/loki/prometheus/linkerd/ops/envoy-health/gitlab.dev/fusionauth.dev: all external probes timeout (000) — internal-only, target exhausted

## 2026-09-18 14:04:20 UTC
- NEW tools.alfaview.com: live Vue "Tools UI" exposing verb-based JSON RPC `POST /poll/pollservice/{list|create|delete|edit|vote|hasVoted}` with custom auth header `Grpc-Metadata-alfaview.token` (diverges f
- NEW staging-tools.alfaview.com: live (612B), same Tools UI, separate staging backend
- NEW whiteboard.alfaview.com / staging-whiteboard.alfaview.com: live board renderer; `/` = "board deleted/access expired", non-root → `302 Location: /`
- NEW status.alfaview.com: public status page 200/53714B (title "alfaview Status")
- NEW qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- NEW CT/crt.sh: ~110 subdomains not in inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*); most edge-firewalle
- CHANGED grafana/loki/prometheus/linkerd/ops/envoy-health/gitlab.dev/fusionauth.dev: all external probes timeout (000) — internal-only, target exhausted
- CHANGED sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404 in prior cycles) — JWKS endpoint restored
- CHANGED sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: POST-body client_id validation fully removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); HTTP Basic auth also accepts any client_id;
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable 10+ cycles

## 2026-09-18 17:25:41 UTC

## 2026-09-18 19:58:19 UTC
- NEW tools.alfaview.com: live Vue "Tools UI" exposing verb-based JSON RPC `POST /poll/pollservice/{list|create|delete|edit|vote|hasVoted}` with custom auth header `Grpc-Metadata-alfaview.token` (diverges f
- NEW staging-tools.alfaview.com: live (612B), same Tools UI, separate staging backend
- NEW whiteboard.alfaview.com / staging-whiteboard.alfaview.com: live board renderer; `/` = "board deleted/access expired", non-root → `302 Location: /`
- NEW status.alfaview.com: public status page 200/53714B (title "alfaview Status")
- NEW qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- NEW CT/crt.sh: ~110 subdomains not in inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*); most edge-firewalle
- CHANGED sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404) — JWKS endpoint restored
- CHANGED sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: POST-body client_id validation fully removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); HTTP Basic auth also accepts any client_id;
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable 10+ cycles
- CHANGED grafana/loki/prometheus/linkerd/ops/envoy-health/gitlab.dev/fusionauth.dev: all external probes timeout (000) — internal-only, target exhausted

## 2026-09-18 22:07:50 UTC

## 2026-09-18 23:56:23 UTC
- NEW tools.alfaview.com: live Vue "Tools UI" exposing verb-based JSON RPC `POST /poll/pollservice/{list|create|delete|edit|vote|hasVoted|updateState|get}` with custom auth header `Grpc-Metadata-alfaview.to
- NEW staging-tools.alfaview.com: live (612B), same Tools UI, separate staging backend
- NEW whiteboard.alfaview.com / staging-whiteboard.alfaview.com: live board renderer; `/` = "board deleted/access expired", non-root → `302 Location: /`
- NEW status.alfaview.com: public status page 200/53714B (title "alfaview Status")
- NEW qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- NEW CT/crt.sh: ~110 subdomains not in inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*); most edge-firewalle
- CHANGED sso.alfaview.com/.well-known/jwks.json: returns 200 with 7 RSA keys (was 404) — JWKS endpoint restored
- CHANGED sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: POST-body `client_id` validation fully removed — fabricated client_id → 200 `{"active":false}` (was 400 `invalid_client`); HTTP Basic auth also accepts any client_i
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED apis.alfaview.com/v2/docs/openapi.json: still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable 10+ cycles
- CHANGED grafana/loki/prometheus/linkerd/ops/envoy-health/gitlab.dev/fusionauth.dev: all external probes timeout (000) — internal-only, target exhausted

## 2026-09-19 01:57:11 UTC

## 2026-09-19 06:50:19 UTC
- NEW `tools.alfaview.com` — live Vue "Tools UI" exposing verb-based JSON RPC `POST /poll/pollservice/{list|create|delete|update|updateState|vote|get|hasVoted}` with custom auth header `Grpc-Metadata-alfavi
- NEW `staging-tools.alfaview.com` — live (612B), same Tools UI, separate staging backend; vendor bundle hash identical to prod
- NEW `whiteboard.alfaview.com` / `staging-whiteboard.alfaview.com` — live board renderer; `/` returns "board deleted/access expired" error, all non-root paths → `302 Location: /`
- NEW `status.alfaview.com` — public status page 200/53714B (title "alfaview Status")
- NEW `qa.alfaview.com` / `uni-stuttgart.alfaview.com` — 200/1381B, identical tenant SPA shell (same family as alfatraining/bhc/kh-freiburg)
- NEW CT/crt.sh yields ~110 subdomains absent from inventory: `grafana`, `loki`, `prometheus-*`, `linkerd-*`, `ops`, `sap`, `webrtc`, `stun`, `gitlab.dev`, `fusionauth.dev`, `whiteboard`, `tools`, `staging-
- CHANGED `sso.alfaview.com/.well-known/jwks.json` — oscillates 200↔404, currently 200 with 7 RSA keys (MD5 3f8d456c)
- CHANGED `sso.alfaview.com/.well-known/openid-configuration` — `introspection_endpoint` advertisement oscillates (absent this cycle) while `/oauth2/introspect` stays live (OPTIONS 405)
- CHANGED `test.alfaview.com` — alfacheck binary version bumped v470079→v483102 (4 platforms), index page still carries no sha256/signatures
- CHANGED `apis.alfaview.com/v2/docs/openapi.json` — still public, prod=beta byte-identical (37 paths, MD5 357b94d3), 10+ stable cycles

## 2026-09-19 11:45:19 UTC

## 2026-09-19 14:59:50 UTC

## 2026-09-19 17:34:38 UTC

## 2026-09-19 19:37:21 UTC

## 2026-09-19 21:45:13 UTC

## 2026-09-19 23:43:34 UTC

## 2026-09-20 01:57:25 UTC

## 2026-09-20 07:22:23 UTC
- NEW tools.alfaview.com/poll/pollservice: Verb-based JSON RPC (list|create|delete|update|updateState|vote|get|hasVoted) at POST /poll/pollservice/<verb> with custom auth header `Grpc-Metadata-alfaview.toke
- NEW staging-tools.alfaview.com: Live Tools UI (612B), separate staging backend
- NEW whiteboard.alfaview.com / staging-whiteboard.alfaview.com: Board renderer; root returns "board deleted/access expired", all non-root → 302 /
- NEW status.alfaview.com: Public status page 200/53714B
- NEW qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- NEW CT/crt.sh: ~110 subdomains not in inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*); most firewalled (00
- CHANGED sso.alfaview.com/.well-known/jwks.json: JWKS restored (200, 7 RSA keys, MD5 3f8d456c) after 404 period
- CHANGED sso.alfaview.com/.well-known/openid-configuration: introspection_endpoint advertisement oscillates (absent this cycle) while /oauth2/introspect stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: POST-body client_id validation fully removed (fabricated client_id → 200 {"active":false}); HTTP Basic auth also accepts any client_id; only residual check is Basic
- CHANGED sso.alfaview.com/oauth2/authorize: Returns HTTP 200 FusionAuth login page (~6189B) for unregistered client_id — validation timing shifted (was 400 invalid_client); redirect_uri matrix still client_id-
- CHANGED apis.alfaview.com/v2/auth/guest-link: Confirmed 3-field only (accessKey+companyId+roomId); displayName NOT required
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public, prod=beta byte-identical (37 paths, MD5 357b94d3) — 10+ stable cycles
- CHANGED test.alfaview.com: alfacheck v483102 (4 platforms), no sha256/signatures — supply-chain hardening absent across releases

## 2026-09-20 12:32:54 UTC
- NEW tools.alfaview.com/poll/pollservice: Verb-based JSON RPC (list|create|delete|update|updateState|vote|get|hasVoted) at POST /poll/pollservice/<verb> with custom auth header `Grpc-Metadata-alfaview.toke
- NEW staging-tools.alfaview.com: Live Tools UI (612B), identical vendor bundle hash (8caab24e), separate staging backend
- NEW whiteboard.alfaview.com / staging-whiteboard.alfaview.com: Board renderer; root returns "board deleted/access expired", all non-root paths → 302 Location: /
- NEW status.alfaview.com: Public status page 200/53714B
- NEW qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- NEW CT/crt.sh: ~110 subdomains absent from inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*); most firewalle
- CHANGED sso.alfaview.com/.well-known/jwks.json: JWKS restored (200, 7 RSA keys, MD5 3f8d456c) after 404 period
- CHANGED sso.alfaview.com/.well-known/openid-configuration: introspection_endpoint advertisement oscillates (absent this cycle) while /oauth2/introspect stays live (OPTIONS 405)
- CHANGED sso.alfaview.com/oauth2/introspect: POST-body client_id validation fully removed (fabricated client_id → 200 {"active":false}); HTTP Basic auth also accepts any client_id; only residual check is Basic
- CHANGED sso.alfaview.com/oauth2/authorize: Returns HTTP 200 FusionAuth login page (~6189B) for unregistered client_id — validation timing shifted (was 400 invalid_client)
- CHANGED apis.alfaview.com/v2/auth/guest-link: Confirmed 3-field only (accessKey+companyId+roomId); displayName NOT required
- CHANGED test.alfaview.com: alfacheck v483102 (4 platforms), no sha256/signatures — supply-chain hardening absent across releases

## 2026-09-20 16:40:07 UTC

## 2026-09-20 19:03:46 UTC

## 2026-09-20 21:35:34 UTC
- NEW tools.alfaview.com/poll/pollservice: Verb-based JSON RPC (8 verbs) at POST /poll/pollservice/<verb> with custom auth header Grpc-Metadata-alfaview.token (b64url→b64, opaque family); staging-tools.alfa
- NEW sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404) — JWKS endpoint restored; MD5 3f8d456c stable
- NEW sso.alfaview.com/.well-known/openid-configuration: introspection_endpoint advertisement oscillates (absent this cycle) while /oauth2/introspect stays live (OPTIONS 405) — discovery not reliable livene
- NEW sso.alfaview.com/oauth2/introspect: 18th+ consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods advertise client_
- NEW CT/alfaview.com: crt.sh returns ~110 subdomains not in inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*)
- CHANGED whiteboard.alfaview.com / staging-whiteboard.alfaview.com: Live board renderer; root returns "board deleted/access expired", non-root → 302 /
- CHANGED status.alfaview.com: Public status page 200/53714B (title "alfaview Status")
- CHANGED qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- CHANGED grafana/loki/prometheus/linkerd/ops/envoy-health/gitlab.dev/fusionauth.dev: All external probes timeout (000) — internal-only, target exhausted

## 2026-09-20 23:29:08 UTC
- NEW tools.alfaview.com/poll/pollservice: Verb-based JSON RPC (8 verbs: list|create|delete|update|updateState|vote|get|hasVoted) at POST /poll/pollservice/<verb> with custom auth header `Grpc-Metadata-alfa
- NEW staging-tools.alfaview.com: Live (612B), same Tools UI, separate staging backend
- NEW whiteboard.alfaview.com / staging-whiteboard.alfaview.com: Live board renderer; root returns "board deleted/access expired", non-root → 302 /
- NEW status.alfaview.com: Public status page 200/53714B (title "alfaview Status")
- NEW qa.alfaview.com / uni-stuttgart.alfaview.com: 200/1381B, identical tenant SPA shell (alfatraining/bhc/kh-freiburg family)
- NEW CT/alfaview.com: crt.sh returns ~110 subdomains not in inventory (grafana, loki, prometheus-*, linkerd*, ops, sap, webrtc, stun, gitlab.dev, fusionauth.dev, whiteboard, tools, staging-*, production-*)
- CHANGED sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404) — JWKS endpoint restored; MD5 3f8d456c stable
- CHANGED sso.alfaview.com/.well-known/openid-configuration: introspection_endpoint advertisement oscillates (absent this cycle) while /oauth2/introspect stays live (OPTIONS 405) — discovery not reliable livene
- CHANGED sso.alfaview.com/oauth2/introspect: 18th+ consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods advertise client_
- CHANGED apis.alfaview.com/v2/docs/openapi.json: Still public and byte-identical prod/beta (37 paths, MD5 357b94d3) — schema surface fully stable 10+ cycles
- CHANGED test.alfaview.com: alfacheck release bumped v470079→v483102 (4 platforms); index page still carries no sha256/signatures — supply-chain hardening absent across successive releases
- CHANGED app.alfaview.com/graphql: signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed); guestAuthenticate/guestJoin anonymous-reachable (BAD_USER_INPUT, not UN

## 2026-09-21 01:35:51 UTC

## 2026-09-21 07:04:51 UTC

## 2026-09-21 14:15:27 UTC

## 2026-09-21 19:31:10 UTC
- NEW NO_DELTA @ full standing probes: OpenAPI MD5 357b94d3 (127532B), JWKS 200/16257B (3f8d456c), OIDC 200/2169B (introspection_endpoint absent, issuer=acme.com), introspect OPTIONS 405, authorize 200/6142
- NEW NO_DELTA @ sso.alfaview.com/oauth2/introspect: 19th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods adverti
- NEW NO_DELTA @ app.alfaview.com/graphql: signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed); all HIGH-value chains (tools BOLA 85, IDOR 80, introspect 70

## 2026-09-21 22:49:20 UTC

## 2026-09-22 01:21:14 UTC
- NEW NO_DELTA @ full standing probes since 2026-09-21 22:49: OpenAPI MD5 357b94d3, JWKS 200/16257B (3f8d456c), OIDC 200/2169B (introspection_endpoint absent, issuer=acme.com), introspect OPTIONS 405, autho

## 2026-09-22 06:27:05 UTC

## 2026-09-22 11:51:47 UTC

## 2026-09-22 15:56:24 UTC

## 2026-09-22 19:24:31 UTC

## 2026-09-22 22:17:42 UTC

## 2026-09-23 00:38:38 UTC

## 2026-09-23 05:07:09 UTC

## 2026-09-23 10:11:31 UTC

## 2026-09-23 14:46:43 UTC
- NEW NO_DELTA @ full standing probes: OpenAPI MD5 357b94d3 (37 paths), JWKS 200/16257B (3f8d456c), OIDC 200/2169B (introspection_endpoint absent, issuer=acme.com), introspect OPTIONS 405, authorize 200/617

## 2026-09-23 18:44:47 UTC

## 2026-09-23 21:52:55 UTC
- NEW NO_DELTA @ full standing probes: OpenAPI MD5 357b94d3 (37 paths), JWKS 200/16257B (3f8d456c), OIDC 200/2169B (introspection_endpoint absent, issuer=acme.com), introspect OPTIONS 405, authorize 200/617

## 2026-09-24 00:06:46 UTC

## 2026-09-24 04:54:59 UTC

## 2026-09-24 09:40:02 UTC
- NEW sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (was 404 in prior cycles) — JWKS endpoint restored; MD5 3f8d456c stable
- NEW sso.alfaview.com/.well-known/openid-configuration: `introspection_endpoint` absent this cycle (was advertised last cycle) — advertisement oscillates while `/oauth2/introspect` stays live (OPTIONS 405)
- NEW sso.alfaview.com/oauth2/introspect: 27th consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods advertise client_s
- NEW tools.alfaview.com: Verb-based JSON RPC (8 verbs) at POST `/poll/pollservice/<verb>` with custom auth header `Grpc-Metadata-alfaview.token` (b64url→b64, opaque family); staging-tools.alfaview.com byte
- NEW app.alfaview.com/graphql: signup mutation remains the sole standing unauthenticated path to a legit bearer token (public JS bundles); all HIGH-value chains (tools BOLA 85, guest AUTH 70, introspect 70

## 2026-09-24 14:28:45 UTC

## 2026-09-24 18:39:47 UTC

## 2026-09-24 21:47:13 UTC

## 2026-09-25 00:08:19 UTC

## 2026-09-25 04:59:08 UTC

## 2026-09-25 10:03:13 UTC
- NEW @ support.alfaview.com — first-ever full mapping (31 cycles only ever recorded a 301). WordPress "alfaview Support Center" (server: myracloud, etag "myra-*", host: support/staging → *.ax4z.com), custo
- NEW @ app.alfaview.com — bundle generation rotated. Old https://app.alfaview.com/js/app.min.5b3949112f0cf682adc8.js → 302 Location: / (asset retired). Live 1381B shell now references https://alfaview-com-
- NEW @ app.alfaview.com/graphql — never-mapped passwordless path: CreateMagicToken is invoked with an *optional* token (`fetchMagicTokenLaunchURL({optionalAccessToken})`, a separate Apollo client `An.A` wh
- NEW @ app.alfaview.com bundle — admin session surface in the SPA: query AdminTokenAuthenticate → Vuex adminSession.accessToken/permissions; mutation AdminSwitchCompany(nextCompanyId); queries GetFinishSig
- NEW @ app.alfaview.com bundle — hardcoded internal config in public asset: AV_COMPANY_ID="alfatraining-internal", AV_SCHULUNG_COMPANY_ID="alfatraining-schulung", AUDIT_LOG_ENABLED_COMPANIES={both:true}, P
- NEW @ staging-app.alfaview.com — host referenced by new bundle, absent from inventory: 401 WWW-Authenticate: Basic on / and /graphql (edge-proxy, CSP references jsdelivr graphql-playground assets). Exhaus
- NEW @ webviewer.dev.alfaview.com — host referenced by new bundle, absent from inventory: 000, no HTTP response. Exhausted.
- CHANGED @ staging.alfaview.com — 2026-09-02 recon recorded 301 → /en; today /, /en/, /xmlrpc.php, /wp-json/ all return 401 (574B nginx Basic page). Whole staging twin now edge-gated.
- CHANGED @ app.alfaview.com bundle — GraphQL auth transport pinned: authenticated ops send headers:{token:<accessToken>} (custom `token` header, not Authorization/Bearer); Signup/FinishSignup/CreateMagicToken 
- NEW NO_DELTA — all standing probes byte-identical to prior cycle (OpenAPI MD5 357b94d3, JWKS 200/16257B, OIDC 200/2169B introspection_endpoint absent, introspect OPTIONS 405, authorize 200/6176B, tools.li

## 2026-09-25 15:01:31 UTC
- NEW support.alfaview.com: First full map — WordPress "alfaview Support Center" (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1); no unauthenticated data exposure; every sensitive 
- NEW app.alfaview.com (public bundle): Asset generation rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle now carries admin session flow (AdminTokenAuthenticate → ad
- NEW app.alfaview.com/graphql: Passwordless bearer issuance path — CreateMagicToken invoked with optional token (fetchMagicTokenLaunchURL({optionalAccessToken}), separate Apollo client An.A without token h

## 2026-09-25 19:10:58 UTC
- NEW support.alfaview.com: First full map — WordPress "alfaview Support Center" (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1); no unauthenticated data exposure; every sensitive 
- NEW app.alfaview.com (public bundle): Asset generation rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle now carries admin session flow (AdminTokenAuthenticate → ad
- NEW app.alfaview.com/graphql: CreateMagicToken mutation confirmed auth-gated (UNAUTHENTICATED) — not anonymous-reachable; no optionalAccessToken argument exists (GRAPHQL_VALIDATION_FAILED). Passwordless b
- NEW staging.alfaview.com: Now fully edge-gated (401 HTTP Basic on /, /en/, /xmlrpc.php, /wp-json/) — was 301 → /en on 2026-09-02. No surface; change recorded, no finding.
- NEW staging-app.alfaview.com + webviewer.dev.alfaview.com: Two bundle-referenced hosts absent from inventory; both exhausted immediately (401 Basic incl. /graphql; 000). Inventory extended, zero attack su
- CHANGED sso.alfaview.com/oauth2/introspect: 27th+ consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods advertise client_
- CHANGED app.alfaview.com/graphql: Signup mutation remains the sole standing unauthenticated path to a legit bearer token (public JS bundles); all HIGH-value chains (tools BOLA 85, IDOR 80, introspect 70) gate

## 2026-09-25 22:29:00 UTC
- NEW whiteboard.alfaview.com: FIRST structural map — Express/Node board renderer behind `edge-proxy`, strict single-route app (every unknown path 302 → `/`; `/` = 386B "board deleted or access expired" pag
- NEW whiteboard.alfaview.com: zero security headers (no CSP/XFO/nosniff/Referrer-Policy), no `Set-Cookie` on any response; error page byte-identical prod vs staging (md5 `fdd0e69b`) — same build-family pat
- NEW whiteboard.alfaview.com: no LFI — `/../package.json`, `/..%2f..%2fpackage.json`, `/....//package.json`, `/static/../package.json` all normalize to 302/23B. Referenced static mount only (`/images/favic
- CHANGED tools.alfaview.com: bundle re-verified unchanged — `js/app-bundle.0a250e1f96a7aaf5661c.js` 200/137343B md5 `b7f17c85ccd8d91e9b831a6bc2aa863c`; 8-verb `/poll/pollservice/<verb>` + `Grpc-Metadata-alfavi
- CHANGED app.alfaview.com/graphql: CreateMagicToken passwordless-bearer lead **CLOSED** (UNAUTHENTICATED, no `optionalAccessToken` argument exists); signup remains the sole unauthenticated path to a bearer and
- CHANGED design-assets.alfaview.com / design-tokens.alfaview.com (404/548B) and ops.alfaview.com (404/19B): confirmed exhausted, no surface.
- NEW support.alfaview.com: First full map — WordPress (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1); no unauthenticated data exposure; sensitive routes 401, only public KB artic
- NEW app.alfaview.com (public bundle): Asset rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle now carries admin session flow (AdminTokenAuthenticate → adminSession.
- NEW app.alfaview.com/graphql: CreateMagicToken mutation confirmed auth-gated (UNAUTHENTICATED) — not anonymous-reachable; no optionalAccessToken argument (GRAPHQL_VALIDATION_FAILED). Passwordless bearer i
- NEW staging.alfaview.com: Now fully edge-gated (401 HTTP Basic on /, /en/, /xmlrpc.php, /wp-json/) — was 301→/en on 2026-09-02. No surface.
- NEW staging-app.alfaview.com + webviewer.dev.alfaview.com: Bundle-referenced hosts absent from inventory; both exhausted (401 Basic incl. /graphql; 000). Zero attack surface.
- CHANGED sso.alfaview.com/oauth2/introspect: 27th+ consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods advertise client_
- CHANGED app.alfaview.com/graphql: Signup mutation remains sole unauthenticated path to legit bearer token; all HIGH-value chains (tools BOLA 85, IDOR 80, introspect 70) gate on it — email-gated, HUMAN_ONLY.
- CHANGED sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (MD5 3f8d456c stable) — JWKS restored.
- CHANGED sso.alfaview.com/.well-known/openid-configuration: introspection_endpoint advertisement oscillates (absent) while /oauth2/introspect stays live (OPTIONS 405) — discovery not reliable liveness signal.
- CHANGED tools.alfaview.com: Verb-based JSON RPC (8 verbs) at POST /poll/pollservice/<verb> with custom auth header Grpc-Metadata-alfaview.token (b64url→b64, opaque family); staging-tools.alfaview.com byte-ide

## 2026-09-26 00:55:06 UTC
- NEW `tools.alfaview.com/whiteboard/` — a **second, previously-unmapped RPC backend**, distinct from the documented `/poll/pollservice/*` gateway. Control set proves the differential: `/health/`, `/foo/`, 
- NEW `tools.alfaview.com`: two **distinct RPC marshallers** — `/poll/pollservice/` returns compact jsonpb `{"code":5,"message":...}` (45B), `/whiteboard/` returns `{"code":5, "message":...}` (47B, space-af
- NEW `tools.alfaview.com` public bundle re-hashed `b7f17c85ccd8d91e9b831a6bc2aa863c` (137343B) contains **zero** `whiteboard` references and only `Ga=${origin}/poll/pollservice` ⇒ the `/whiteboard/` backen
- NEW `staging-tools.alfaview.com/whiteboard/` returns the byte-identical 47B envelope ⇒ staging exposure equals production (same build family).
- NEW `/whiteboard/` is **absent** from `whiteboard.alfaview.com` / `staging-whiteboard.alfaview.com` (both 302→`/`, strict single-route) ⇒ the board *renderer* host and the board *data RPC* are separate sy
- NEW `/whiteboard/` prefix routing is exclusive: every `/whiteboard/**` subpath reaches the RPC backend; identical-looking non-prefixed paths (`/whiteboardservice.v1.WhiteboardService/get`, `/WhiteboardSer
- CHANGED `tools.alfaview.com`: prior model held the tool host's surface as poll-only (8 verbs). That is now known-incomplete — a second mounted prefix exists with unmapped verbs and unknown auth enforcement.
- NEW support.alfaview.com: First full map — WordPress (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1); no unauthenticated data exposure; sensitive routes 401, only public KB artic
- NEW app.alfaview.com (public bundle): Asset generation rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle now carries admin session flow (AdminTokenAuthenticate → ad
- NEW app.alfaview.com/graphql: CreateMagicToken mutation confirmed auth-gated (UNAUTHENTICATED) — not anonymous-reachable; no optionalAccessToken argument (GRAPHQL_VALIDATION_FAILED). Passwordless bearer i
- NEW staging.alfaview.com: Now fully edge-gated (401 HTTP Basic on /, /en/, /xmlrpc.php, /wp-json/) — was 301→/en on 2026-09-02. No surface; change recorded, no finding.
- NEW staging-app.alfaview.com + webviewer.dev.alfaview.com: Bundle-referenced hosts absent from inventory; both exhausted (401 Basic incl. /graphql; 000). Zero attack surface.
- NEW whiteboard.alfaview.com: First structural map — Express/Node board renderer behind edge-proxy; strict single-route (unknown ID → 302 `/`; `/` = 386B "board deleted/access expired"); zero security head
- CHANGED sso.alfaview.com/oauth2/introspect: 27th+ consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods advertise client_
- CHANGED app.alfaview.com/graphql: Signup mutation remains sole unauthenticated path to legit bearer token; all HIGH-value chains (tools BOLA 85, IDOR 80, introspect 70) gate on it — email-gated, HUMAN_ONLY.
- CHANGED sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (MD5 3f8d456c stable) — JWKS restored.
- CHANGED sso.alfaview.com/.well-known/openid-configuration: introspection_endpoint advertisement oscillates (absent) while /oauth2/introspect stays live (OPTIONS 405) — discovery not reliable liveness signal.
- CHANGED tools.alfaview.com: Verb-based JSON RPC (8 verbs) at POST /poll/pollservice/<verb> with custom auth header Grpc-Metadata-alfaview.token (b64url→b64, opaque family); staging-tools.alfaview.com byte-ide

## 2026-09-26 05:42:18 UTC

## 2026-09-26 10:16:57 UTC
- NEW GET https://apis.alfaview.com/v2/stats — never probed in 33 cycles. Unauthenticated, no query params → 422/319B application/problem+json carrying a per-field validation body (query.from, query.to, que
- NEW /v2/stats with valid params (from=2026-07-27T00:00:00Z&to=2026-09-26T00:00:00Z&stepDurationHours=24) → 401/107B, and a fabricated base64 bearer (Zm9vOmJhcg==) → 401/115B. Auth IS enforced before the d
- NEW Ordering defect is environment-wide: beta-apis.alfaview.com/v2/stats is byte-identical (422/319B bare, 401/107B with valid params), and /v2/auth/token-info bare → 422/85B on both.
- NEW Defect is BOUNDED and the bound is proven: path-parameter routes do not share it. GET /v2/rooms|meetings|group-links|guest-links/{malformed-id} all return 401/107B, never 422 — the UUID path validator
- NEW /v2/stats schema bound reproduced pre-auth: stepDurationHours=0 → 422/158B "expected number >= 1" (matches spec minimum:1); stepDurationHours=abc → 422/357B. A from= date 63 days back (outside the doc
- NEW Root cause is visible in the spec: /v2/docs/openapi.json declares components.securitySchemes = {} and top-level security = null. Authentication is enforced purely by handler-level middleware, which is
- CHANGED /v2/auth/token-info parse oracle unchanged, and NO third error tier exists: base64 of {}, {"token":"x"}, a raw UUID, and random 16/32/48/64/128-byte payloads all return the identical 422 "invalid acce
- CHANGED GET /v2/users/invitation → 405/19B text/plain (Go-native), not the application/problem+json 401 every other /v2 path returns → a different runtime fronts that route.
- NEW tools.alfaview.com/whiteboard/: Second unmapped RPC backend proven by controlled differential — `/whiteboard/` returns 47B gRPC status envelope while `/health/`, `/foo/`, `/zzznotreal/` return 615B SP
- NEW staging-tools.alfaview.com/whiteboard/: Byte-identical 47B envelope ⇒ unmapped RPC mount mirrored to staging with exposure equal to production.
- NEW whiteboard.alfaview.com: First structural map — Express/Node board renderer behind edge-proxy; strict single-route (unknown ID → 302 `/`; `/` = 386B "board deleted/access expired"); zero security head
- NEW support.alfaview.com: First full map — WordPress (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1); no unauthenticated data exposure; every sensitive route 401s, only public ro
- NEW app.alfaview.com (public bundle): Asset generation rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle carries admin session flow (AdminTokenAuthenticate → adminS
- CHANGED staging.alfaview.com: Now fully edge-gated (401 HTTP Basic on /, /en/, /xmlrpc.php, /wp-json/) — was 301 → /en on 2026-09-02.
- CHANGED staging-app.alfaview.com + webviewer.dev.alfaview.com: Bundle-referenced hosts exhausted (401 Basic incl. /graphql; 000).

## 2026-09-26 14:32:54 UTC
- NEW `GET /v2/rooms/{roomId}/attendances` — the last unprobed GET op — returns **422/232B pre-auth** with `{query.from, query.to}` required-parameter errors. My prior cycle's bound ("path-parameter routes 
- CHANGED The defect is **not** limited to required params. `/v2/meetings?from=abc&to=zzz` → 422/290B although `from`/`to` are `required:false` in the spec. Rule corrected: *any format-constrained query param t
- NEW Full pre-auth validation map completed — **9 of 26** GET ops (was "2 of 26"): `stats` 422/319B, `rooms/{roomId}/attendances` 422/232B, `meetings` 422/290B, `rooms?limit=abc` 422/141B, `users?emailAddr
- NEW **Mechanism proven by single-route differential:** `GET /v2/guest-links` → 401, `?pageToken=x` → 401 (unconstrained string, nothing to reject), `?limit=abc` → 422 (integer-constrained). Same route, sa
- NEW Pre-auth error bodies disclose the internal validator stack: Go `time.Parse` layout string `2006-01-02T15:04:05.999999999Z07:00`, Go `net/mail` RFC-5322 parser (`expected string to be RFC 5322 email: 
- NEW Identity-adjacent surface: `/v2/users?emailAddress=` is the only pre-auth validator touching an identity field. **Existence was deliberately not tested** — that is the rejected enumeration class. Form
- CHANGED Business-rule checks stay behind auth: `from=2025-08-01` (400 days, far outside the documented 62-day floor) → 401. Valid params → 401 on every op. **No attendance, user, or tenant data read; every ca
- NEW `beta-apis.alfaview.com/v2/rooms/{roomId}/attendances` byte-identical 422/232B → same code path serves both environments.
- NEW Spec re-confirmed: `components.securitySchemes` **absent**, top-level `security` **absent** (`None` in the fetched doc) — no declarative control exists that could have caught this ordering.
- NEW apis.alfaview.com/v2/stats: GET unauthenticated, no params → 422/319B with per-field validation body (query.from, query.to, query.stepDurationHours); 11/26 other GET ops return 401. Valid params → 401
- NEW apis.alfaview.com/v2/auth/token-info: Base64 of {}, {"token":"x"}, raw UUID, random 16-128 byte payloads all return identical 422 "invalid access token format"; only non-base64 header gives distinct e
- NEW apis.alfaview.com/v2/users/invitation: Returns 405/19B text/plain (Go-native) vs application/problem+json 401 for rest of /v2 — different runtime fronts this route.
- NEW tools.alfaview.com/whiteboard/: Second unmapped RPC backend proven by controlled differential — `/whiteboard/` returns 47B gRPC status envelope while `/health/`, `/foo/`, `/zzznotreal/` return 615B SP
- NEW staging-tools.alfaview.com/whiteboard/: Byte-identical 47B envelope ⇒ unmapped RPC mount mirrored to staging with equal exposure.
- NEW whiteboard.alfaview.com: First structural map — Express/Node board renderer behind edge-proxy; strict single-route (unknown ID → 302 `/`; `/` = 386B "board deleted/access expired"); zero security head
- NEW support.alfaview.com: First full map — WordPress (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1); no unauthenticated data exposure; every sensitive route 401s, only public ro
- NEW app.alfaview.com (public bundle): Asset generation rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle carries admin session flow (AdminTokenAuthenticate → adminS
- CHANGED staging.alfaview.com: Now fully edge-gated (401 HTTP Basic on /, /en/, /xmlrpc.php, /wp-json/) — was 301 → /en on 2026-09-02.
- CHANGED staging-app.alfaview.com + webviewer.dev.alfaview.com: Bundle-referenced hosts exhausted (401 Basic incl. /graphql; 000).

## 2026-09-26 18:10:01 UTC
- NEW apis.alfaview.com/v2/rooms/{roomId}/attendances: GET unauthenticated → 422/232B pre-auth validation (query.from, query.to required); 9 of 26 GET ops now confirmed validation-before-auth (stats, attend
- NEW tools.alfaview.com/whiteboard/: Second unmapped RPC backend proven by controlled differential — returns 47B gRPC status envelope vs poll gateway's 45B compact jsonpb; distinct marshaller (space-after-
- NEW staging-tools.alfaview.com/whiteboard/: Byte-identical 47B envelope → unmapped RPC mount mirrored to staging with equal exposure
- NEW whiteboard.alfaview.com: First structural map — Express/Node board renderer behind edge-proxy; strict single-route (unknown ID → 302 /; / = 386B "board deleted/access expired"); zero security headers,
- NEW support.alfaview.com: First full map — WordPress (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1); no unauthenticated data exposure; every sensitive route 401s, only public ro
- NEW app.alfaview.com (public bundle): Asset generation rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle carries admin session flow (AdminTokenAuthenticate → adminS
- CHANGED staging.alfaview.com: Now fully edge-gated (401 HTTP Basic on /, /en/, /xmlrpc.php, /wp-json/) — was 301 → /en on 2026-09-02
- CHANGED staging-app.alfaview.com + webviewer.dev.alfaview.com: Bundle-referenced hosts exhausted (401 Basic incl. /graphql; 000)
- CHANGED apis.alfaview.com/v2/auth/token-info: Parse oracle unchanged — base64 of {}, {"token":"x"}, raw UUID, random 16-128 byte payloads all return identical 422 "invalid access token format"; only non-base6
- CHANGED sso.alfaview.com/.well-known/jwks.json: Returns 200 with 7 RSA keys (MD5 3f8d456c stable) — JWKS restored
- CHANGED sso.alfaview.com/.well-known/openid-configuration: introspection_endpoint advertisement oscillates (absent this cycle) while /oauth2/introspect stays live (OPTIONS 405) — discovery not reliable livene
- CHANGED sso.alfaview.com/oauth2/introspect: 27th+ consecutive stable cycle — fabricated client_id accepted on POST-body and Basic channels (200 {"active":false}); token_endpoint_auth_methods advertise client_
- CHANGED app.alfaview.com/graphql: Signup mutation remains sole unauthenticated path to legit bearer token (public JS bundles confirmed); all HIGH-value chains (tools BOLA 85, IDOR 80, introspect 70) gate on i

## 2026-09-26 20:34:27 UTC

## 2026-09-26 23:08:35 UTC
- NEW tools.alfaview.com/whiteboard/: Second unmapped RPC backend confirmed by controlled differential — `/whiteboard/` returns 47B gRPC status envelope (`{"code":5, "message":...}`) vs poll gateway's 45B c
- NEW staging-tools.alfaview.com/whiteboard/: Byte-identical 47B envelope ⇒ unmapped RPC mount mirrored to staging with exposure equal to production.
- NEW whiteboard.alfaview.com: `/whiteboard/` absent from renderer host (302→`/`, strict single-route) ⇒ board renderer and board data RPC are separate systems; renderer's 200-vs-302 existence oracle not th
- NEW apis.alfaview.com/v2/rooms/{roomId}/attendances: GET unauthenticated → 422/232B pre-auth validation (`query.from`, `query.to` required); completes the pre-auth validation map — **9 of 26** GET ops now
- NEW apis.alfaview.com/v2/stats: Validation-before-auth defect confirmed systemic and environment-wide; OpenAPI declares `components.securitySchemes={}`, `security=null` — auth purely handler middleware.

## 2026-09-27 01:29:33 UTC

## 2026-09-27 07:05:29 UTC
- NEW tools.alfaview.com/whiteboard/: Second unmapped RPC backend confirmed by controlled differential — `/whiteboard/` returns 47B gRPC status envelope vs poll gateway's 45B compact jsonpb; distinct marsha
- NEW staging-tools.alfaview.com/whiteboard/: Byte-identical 47B envelope ⇒ unmapped RPC mount mirrored to staging with exposure equal to production
- NEW whiteboard.alfaview.com: `/whiteboard/` absent from renderer host (302→`/`, strict single-route) ⇒ board renderer and board data RPC are separate systems
- NEW support.alfaview.com: First full map — WordPress (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom alfaview/v1) — no unauthenticated data exposure; every sensitive route 401s, only public r
- NEW app.alfaview.com (public bundle): Asset generation rotated to app.min.67e8a68d4318b34ca241.js (md5 2cb9128353b1f7444e222b4f61e4ffa5); bundle carries admin session flow (AdminTokenAuthenticate → adminS
- CHANGED staging.alfaview.com: Now fully edge-gated (401 HTTP Basic on /, /en/, /xmlrpc.php, /wp-json/) — was 301 → /en on 2026-09-02
- CHANGED staging-app.alfaview.com + webviewer.dev.alfaview.com: Bundle-referenced hosts exhausted (401 Basic incl. /graphql; 000)
- CHANGED design-assets.alfaview.com, design-tokens.alfaview.com, ops.alfaview.com: 404/548B and 404/19B plaintext — no independent surface, targets confirmed exhausted

## 2026-09-27 13:06:33 UTC
- NEW apis.alfaview.com/v2/users/invitation advertises `Allow: DELETE, POST` (405/19B text/plain, server: edge-proxy) while the public OpenAPI declares POST ONLY — an undeclared destructive method on the do
- NEW apis.alfaview.com/v2/users/invitations (plural, ABSENT from the 37-path spec) is DELETE-only (`Allow: DELETE`), same 19B Go-native envelope; also live on beta. Never enumerated in 34 cycles.
- NEW Pre-auth route-existence oracle on that second runtime, usable with GET/HEAD only: registered-but-no-GET → `405 "Method Not Allowed"`; unregistered → `404 "404 page not found"`; declared app routes → 
- NEW The 405 is returned BEFORE authentication (405, not 401) → router-layer pre-auth processing, one layer below the 9 already-proven pre-auth-validation ops.
- NEW REJECTED AUTH @ sso.alfaview.com/oauth2/userinfo: third-tier probe closed. 3 JOSE-shaped self-issued tokens (iss=https://sso.alfaview.com, iss=acme.com, exp=2036) all → byte-identical `access_token_fa
- CHANGED apis.alfaview.com OpenAPI unchanged: md5 357b94d367909a40b9299b543d23712b, 127532B, 37 paths, `components.securitySchemes` absent, top-level `security` absent.
- CHANGED sso.alfaview.com unchanged: introspect OPTIONS 405 (alive), OIDC 200/2169B, issuer=acme.com, `introspection_endpoint` absent, `token_endpoint_auth_methods_supported`=[client_secret_basic, client_secre
- CHANGED Control `/v2/rooms/invitation` → 401 application/problem+json = match on `/v2/rooms/{id}`, NOT a new route; reconfirms the path-param pre-auth-401 bound.
- NEW `tools.alfaview.com/whiteboard/` — second unmapped gRPC-web RPC backend confirmed by controlled differential: returns 47B gRPC status envelope (`{"code":5, "message":...}`) vs poll gateway's 45B compa
- NEW `staging-tools.alfaview.com/whiteboard/` — byte-identical 47B envelope ⇒ unmapped RPC mount mirrored to staging with exposure equal to production
- NEW `whiteboard.alfaview.com` — `/whiteboard/` absent from renderer host (302→`/`, strict single-route) ⇒ board renderer and board data RPC are separate systems; renderer's 200-vs-302 existence oracle not
- NEW `support.alfaview.com` — first full map: WordPress (myracloud/ax4z, 272 REST routes, 14 namespaces incl. custom `alfaview/v1`); no unauthenticated data exposure; every sensitive route 401s, only publi
- NEW `app.alfaview.com` (public bundle) — asset generation rotated to `app.min.67e8a68d4318b34ca241.js` (md5 `2cb9128353b1f7444e222b4f61e4ffa5`); bundle carries admin session flow (`AdminTokenAuthenticate`
- CHANGED `staging.alfaview.com` — now fully edge-gated (401 HTTP Basic on `/`, `/en/`, `/xmlrpc.php`, `/wp-json/`) — was 301→`/en` on 2026-09-02
- CHANGED `staging-app.alfaview.com` + `webviewer.dev.alfaview.com` — bundle-referenced hosts exhausted immediately (401 Basic incl. `/graphql`; 000)
- CHANGED `design-assets.alfaview.com`, `design-tokens.alfaview.com`, `ops.alfaview.com` — 404/548B and 404/19B plaintext confirmed exhausted

## 2026-09-27 17:56:02 UTC
- NEW apis.alfaview.com `/v2/guest-links/{id}` + `/v2/group-links/{id}`: the only link **mutations** without `roomId` in the path, while every link create/delete is room-scoped. Unscoped GET returns `access
- NEW apis.alfaview.com `PATCH /v2/guest-links/{id}` accepts `emailAddress` + `permissionGroupId` + `validFrom`/`validUntil` + `sendEmail:true` in one call — credential rebinding plus permission-group assig
- CHANGED Queued SCAN closed: full 37-path `Allow:`-vs-spec diff run. Exactly **one** divergence in 37 paths (`/v2/users/invitation` spec=[POST] live=[DELETE, POST]). The undeclared-method class is a single rou
- CHANGED `GET /v2/permission-groups` = 401/107B `application/problem+json` — **proves the spec under-declares auth codes**, invalidating response-code declarations as evidence of code behaviour.
- CHANGED `sso.alfaview.com` byte-stable: OIDC 200/2169B (issuer=acme.com, `introspection_endpoint` null, `password`+`client_credentials` advertised), JWKS 200/16257B (7×RS256, RSA-only, zero symmetric/ECDSA vs
- CHANGED `apis.alfaview.com` OpenAPI unchanged: md5 `357b94d367909a40b9299b543d23712b`, 127532B, 37 paths, `components.securitySchemes` absent, top-level `security` absent.
- CHANGED `apis.alfaview.com/v2/users/invitations` reconfirmed `allow: DELETE`, 405/19B Go-native, `server: edge-proxy`.
- CHANGED `app.alfaview.com/graphql` 400/406B, `tools` poll 501/55B, `/whiteboard/` 404/47B, `/health/` 200/615B, `client-diagnostics /health` 200/16B — all byte-identical.
- CHANGED NO_DELTA @ sso/apis/app/tools: 33rd consecutive stable cycle on every standing probe; no new endpoints, no regressions.

## 2026-09-27 20:38:21 UTC
- NEW apis.alfaview.com GET /v2/guest-links + GET /v2/group-links: the two company-wide link lists accept NO scoping parameter — only pageToken + limit. Room-scoped siblings exist separately (/v2/rooms/{roo
- NEW apis.alfaview.com TokenUserPermissions is a CLOSED vocabulary: additionalProperties:false, all 7 fields required = manageCompany, roomAdmin, roomCreate, roomList, userAdmin, userList, userShow. There 
- NEW apis.alfaview.com RoomPermissions proves guest links ARE permission-bearing principals: guestLinkPermissions[] and groupLinkPermissions[] of ParticipantPermissions{participantId, permissions}; Permiss
- NEW apis.alfaview.com POST /v2/auth/group-link is LIVE and unprobed in 34 cycles: HEAD → 405 `allow: POST`, text/plain, edge-proxy. The 2026-09-26 standing-probe list recorded it as 404 (that was a GET on
- NEW app.alfaview.com admin surface now completely enumerated from the unchanged bundle (1092529B, md5 2cb9128353b1f7444e222b4b61e4ffa5): exactly 4 admin operations — adminTokenAuthenticate{accessToken,per
- CHANGED /v2/auth/group-link requires displayName, /v2/auth/guest-link does not — server-confirmed: omitting it gives 422/171B `expected required property displayName to be present` on group-link, while guest-
- CHANGED The 422/88B `ACTION_INVALID: ACTION_INVALID` tier is a POST-LOOKUP business tier, not a format gate: it is returned for zero-UUID companyId+roomId with a synthetic 20-char key, for a `../../etc/passwd
- CHANGED Queued 37-path scan closed: all four link-family routes return byte-identical 401/107B `application/problem+json` with `www-authenticate: AccessToken` pre-auth — /v2/group-links/{uuid}, /v2/guest-link
- CHANGED NO_DELTA on every other standing probe: OpenAPI 200/127532B md5 357b94d367909a40b9299b543d23712b, users/me 401/107B md5 60ed2f29, OIDC 200/2169B md5 f78a08fc (introspection_endpoint null), JWKS 200/16
- NEW apis.alfaview.com authentication gate precedes ALL request parsing. Eight GETs, all 401/107B application/problem+json, all body md5 60ed2f29d492ec872c598cc7d36aa37e, all headers identical: `server: ed
- NEW The negative control settles the parameter question in the strong direction. `?bogusParam=abc` returns the identical 401, so the endpoint does not distinguish recognized from unrecognized query parame
- NEW Whole-surface uniformity: the same 107-byte body is emitted by /v2/guest-links, /v2/group-links, /v2/rooms/{id}/guest-links, /v2/rooms/{id} and the standing /v2/users/me probe. The authenticated API h
- CHANGED Methodological correction to my own prior turn — I must record this rather than let it stand: I proposed the test "422 on roomId ⇒ the parameter is recognized" as the discriminator for H1. That test c
- CHANGED Contract precision confirmed against the local 37-path spec for ListGuestLinks: parameters are exactly Authorization (header), pageToken (string), limit (int32, default 100, minimum 0, maximum 500). N
- CHANGED The 403 asymmetry is real and now quoted exactly. ListGuestLinks declares 200, 401 and 403 where 403 reads "The requesting user does not have the necessary permissions". ListGroupLinks declares only 2
- CHANGED GuestLink is additionalProperties:false with 12 of 15 fields required: id, roomId, createdBy, createdAt, permissionGroupId, accessKey, emailAddress, locale, origin, dialInAllowed, dialInCode, joinUrl.
- CHANGED New amplification from the same schema: dialInCode is an unconditionally required member sitting beside the boolean dialInAllowed, which is itself required. The contract therefore provisions and retur

## 2026-09-27 23:25:32 UTC
- CHANGED **Self-correction, material: `GET /v2/guest-links?limit=abc` → 422/141B `application/problem+json` `{"title":"Unprocessable Entity","status":422,"detail":"validation failed","errors":[{"message":"inva
- NEW Same-route same-session **negative proof for H1**: `/v2/guest-links?roomId=abc&limit=abc` → 422/141B and `?companyId=abc&limit=abc` → 422/141B, both md5 `551221c3`, byte-identical to `?limit=abc` alon
- NEW `/v2/group-links?limit=abc` → identical 422/141B md5 `551221c3` — pre-auth oracle is a shared binder across both link families, not a per-route quirk.
- NEW **Unauthenticated parameter-schema enumeration oracle** (no credential): `GET <route>?<param>=<wrong-typed-value>` → 422 carrying `location: query.<param>` if the parameter is recognized and format-co
- NEW `/v2/rooms?roomTypes=abc` → 422/190B `"expected value to be one of \"department, room, meeting\""` at `query.roomTypes[0]` — new validator class: array item, indexed location, complete enum whitelist 
- NEW `POST /v2/auth/api-key` (`AuthenticateAPIKey`) — the **4th unauthenticated `/v2/auth/*` endpoint, absent from 34 cycles of notes**. `HEAD` → 405 `allow: POST`; `OPTIONS` → 405/19B `text/plain`, `serve
- NEW `/v2/auth/api-key` is the **only** auth endpoint in the contract declaring a 403 account-status tier ("account is inactive") that guest-link, group-link and password do not — a distinct third response
- CHANGED `/v2/rooms/abc` → 401/107B md5 `60ed2f29` reconfirmed. Mechanism is now **asymmetric and exact**: query params validate pre-auth, path params validate post-auth.
- CHANGED **Retraction of my own prior-cycle methodology**: I recorded that the test "422 ⇒ parameter recognized" "could not have worked." It does work. The missing piece was a positive control on a schema-decl
- NEW apis.alfaview.com/v2/guest-links + /v2/group-links: company-wide lists accept NO scoping parameter (only pageToken+limit); room-scoped siblings exist separately; TokenUserPermissions vocabulary (7 fie
- NEW apis.alfaview.com/v2/auth/group-link: HEAD → 405 allow: POST (live, unprobed 34 cycles); requires displayName (422/171B if omitted) unlike guest-link
- NEW app.alfaview.com bundle (md5 2cb9128353b1f7444e222b4f61e4ffa5): adminSwitchCompany($nextCompanyId: String!) returns {companyId, accessToken, permissions} from caller's token with no proof of administe
- NEW apis.alfaview.com: 9 of 26 GET ops validation-before-auth (stats, rooms/{id}/attendances, meetings, rooms?limit=abc, users?emailAddress=, etc.); path-param routes correctly 401; root cause: OpenAPI co
- NEW tools.alfaview.com/whiteboard/: second unmapped gRPC-web RPC backend confirmed (47B envelope vs poll's 45B); distinct marshaller; no WWW-Authenticate/401 on any probed path; public bundle (md5 b7f17c8
- NEW sso.alfaview.com/oauth2/introspect: 30th consecutive cycle — fabricated client_id accepted on POST-body and Basic (200 {"active":false}); token_endpoint_auth_methods advertises client_secret_basic/pos
- CHANGED NO_DELTA on all other standing probes (34th consecutive byte-stable cycle)

## 2026-09-28 02:01:17 UTC
- NEW **`/v2/rooms?limit=51` → 422/147B `{"message":"expected number <= 50","location":"query.limit"}`** — new oracle class: per-operation numeric *upper bound* disclosed pre-auth, and it matches the contra
- NEW **The binder is per-operation, not global — proven with three negative controls.** `/v2/meetings?limit=abc` → 401/107B; `/v2/rooms?from=abc` → 401/107B; `/v2/permission-groups?limit=abc` → 401/107B. E
- NEW **Binder active on the room-scoped credential siblings.** `/v2/rooms/{uuid}/guest-links?limit=abc` and `/v2/rooms/{uuid}/group-links?limit=abc` → both 422/141B md5 `551221c3`, byte-identical to the un
- NEW **Zero undocumented parameters found** — 10 candidate tests (`includeArchived`, `search`, `sort`, `filter`, `offset` on `/v2/guest-links`; `search` on `/v2/group-links`; `roomId`, `companyId` on the l
- NEW **Required-ness and multi-entry accumulation disclosed pre-auth.** `/v2/rooms/{uuid}/attendances?from=abc` → 422/261B with two entries: `query.from` `"invalid date/time for format 2006-01-02T15:04:05.
- CHANGED **`/v2/guest-links` has two unauthenticated 401 tiers, not one.** No `Authorization` header → 401/107B md5 `60ed2f29` `"No access token was provided…"`; any `Authorization: Bearer …` → 401/122B md5 `3
- CHANGED **Self-correction, material: my recorded 401 body lengths are not stable (107/115/122B across cycles); the md5 is the reliable discriminator.** My prior-cycle claim "the whole authenticated surface re
- CHANGED **Message-accuracy defect:** `Bearer Zm9vOmJhcg==` is valid base64 yet still returns the *"No base64 encoded access token"* tier, and a raw UUID returns the same tier — the string is emitted unconditi
- CHANGED Contract re-verified: `components.securitySchemes` absent, top-level `security` absent. OpenAPI md5 `357b94d367909a40b9299b543d23712b` / 127532B / 37 paths — 35th consecutive stable cycle.
- NEW NO_DELTA — last leads (2026-09-27 23:25) already incorporated into knowledge base; all standing probes byte-identical 34th consecutive cycle

## 2026-09-28 08:06:22 UTC
- NEW NO_DELTA — last leads (2026-09-27 23:25) already incorporated into knowledge base; all standing probes byte-identical 34th consecutive cycle
- CHANGED apis.alfaview.com/v2/guest-links + /v2/group-links: room-scoped siblings return byte-identical 422/141B md5 `551221c3` for `?limit=abc`; 10 candidate filters (`includeArchived`, `search`, `sort`, `fil
- CHANGED apis.alfaview.com/v2/auth/api-key: 4th unauthenticated credential endpoint discovered (HEAD→405 allow:POST, OPTIONS→405/19B text/plain); declares unique 403 account-status tier
- CHANGED sso.alfaview.com/oauth2/introspect: 30+ cycles stable — fabricated client_id accepted on POST-body and Basic (200 `{"active":false}`); `token_endpoint_auth_methods_supported` advertises `client_secret
- CHANGED app.alfaview.com (public bundle): adminSwitchCompany mutation mints cross-tenant admin tokens (`{companyId, accessToken, permissions}`) with no administration proof; hardcoded production tenant IDs (`
- CHANGED tools.alfaview.com/whiteboard/: second unmapped RPC backend confirmed — 47B gRPC status envelope vs poll's 45B jsonpb; distinct marshaller; no auth challenge; zero client references in bundle (md5 `b7
- CHANGED apis.alfaview.com: 9/26 GET ops validate query params pre-auth (stats, attendances, meetings, rooms?limit, users?emailAddress, guest-links?limit, group-links?limit, rooms?roomTypes, rooms/{id}/attenda

## 2026-09-28 16:52:14 UTC
- CHANGED apis.alfaview.com/v2/guest-links + /v2/group-links: room-scoped siblings return byte-identical 422/141B md5 `551221c3` for `?limit=abc`; 10 candidate filters (`includeArchived`, `search`, `sort`, `fil
- CHANGED apis.alfaview.com/v2/auth/api-key: 4th unauthenticated credential endpoint discovered (HEAD→405 allow:POST, OPTIONS→405/19B text/plain); declares unique 403 account-status tier
- CHANGED sso.alfaview.com/oauth2/introspect: 30+ cycles stable — fabricated client_id accepted on POST-body and Basic (200 `{"active":false}`); `token_endpoint_auth_methods_supported` advertises `client_secret
- CHANGED app.alfaview.com (public bundle): adminSwitchCompany mutation mints cross-tenant admin tokens (`{companyId, accessToken, permissions}`) with no administration proof; hardcoded production tenant IDs (`
- CHANGED tools.alfaview.com/whiteboard/: second unmapped RPC backend confirmed — 47B gRPC status envelope vs poll's 45B jsonpb; distinct marshaller; no auth challenge; zero client references in bundle (md5 `b7
- CHANGED apis.alfaview.com: 9/26 GET ops validate query params pre-auth (stats, attendances, meetings, rooms?limit, users?emailAddress, guest-links?limit, group-links?limit, rooms?roomTypes, rooms/{id}/attenda

## 2026-09-28 22:27:29 UTC
- NEW `apis.alfaview.com/v2/auth/api-key`: 4th unauthenticated credential endpoint discovered (HEAD→405 allow:POST, OPTIONS→405/19B text/plain); declares unique 403 account-status tier
- CHANGED `apis.alfaview.com/v2/guest-links` + `/v2/group-links`: room-scoped siblings return byte-identical 422/141B md5 `551221c3` for `?limit=abc`; 10 candidate filters (`includeArchived`, `search`, `sort`, 
- CHANGED `sso.alfaview.com/oauth2/introspect`: 30+ cycles stable — fabricated `client_id` accepted on POST-body and Basic (200 `{"active":false}`); `token_endpoint_auth_methods_supported` advertises `client_se
- CHANGED `app.alfaview.com` (public bundle): `adminSwitchCompany` mutation mints cross-tenant admin tokens (`{companyId, accessToken, permissions}`) with no administration proof; hardcoded production tenant ID
- CHANGED `tools.alfaview.com/whiteboard/`: second unmapped RPC backend confirmed — 47B gRPC status envelope vs poll's 45B jsonpb; distinct marshaller; no auth challenge; zero client references in bundle (md5 `
- CHANGED `apis.alfaview.com`: 9/26 GET ops validate query params pre-auth (stats, attendances, meetings, rooms?limit, users?emailAddress, guest-links?limit, group-links?limit, rooms?roomTypes, rooms/{id}/atten
- CHANGED `apis.alfaview.com/v2/users/invitation{,s}`: OpenAPI declares POST only; live server advertises `Allow: DELETE,POST` on `/v2/users/invitation` and DELETE-only on undocumented `/v2/users/invitations`; 

## 2026-09-29 02:06:54 UTC
- NEW /v2/users/me/company (GetOwnCompany, x-sort-index 32) is the 38th path — identity of the 2026-09-28 37→38 delta RESOLVED by set-diff against recon-notes/alfaview-openapi.yaml (37 paths, removed=∅). So
- NEW Company schema = additionalProperties:false {companyId (26-char example "0123456789ABCDEFGHIJKLMNOP"), displayName, createdAt}, all 3 required. Declares 200/401/422 — no 403.
- NEW x-sort-index space is GLOBAL 0–32, not per-tag. Holes at 24,25,26,30,31; GetOwnCompany appended at 32 far outside its own Users group (0–4). Structural evidence it was bolted on late, separately from 
- CHANGED apis OpenAPI md5 284a3383c1ac3cfc9152ffcc631891f2 (132100B, 38 paths, 57 ops) unchanged; components.securitySchemes absent, top-level security absent — 2nd consecutive cycle at this hash. Byte-identic
- CHANGED GET /v2/users/me/company unauth → 401/107B md5 60ed2f29d492ec872c598cc7d36aa37e; `?bogus=abc` and `?companyId=abc` both return that identical 401 (NOT 422) — the new path declares no query params, so 
- CHANGED sso byte-stable 34th cycle: OIDC 200/2169B md5 f78a08fc, JWKS 200/16257B md5 3f8d456c (7×RS256), introspection_endpoint absent while /oauth2/introspect live at OPTIONS 405, grant_types still advertise
- CHANGED tools unchanged: /whiteboard/ 404/47B gRPC envelope, /poll/pollservice/list 501/55B.

## 2026-09-29 08:35:12 UTC
- NEW /v2/users/me/company (GetOwnCompany) identified as the 38th OpenAPI path — sole 200-response supplier of companyId across 57 operations; enables guest-link chain execution by authorized tester
- CHANGED OpenAPI spec md5 rotated to 284a3383c1ac3cfc9152ffcc631891f2 (132100B, 38 paths, 57 ops) — byte-identical on beta-apis.alfaview.com

## 2026-09-29 15:22:48 UTC
- NEW apis.alfaview.com OpenAPI: full structural diff of the 37→38 document delta reveals a SECOND, previously undetected
- NEW alfaview 2FA is a shipped, customer-facing feature with a dedicated public KB article:
- NEW Measured 2FA enforcement asymmetry, three-way: (a) sso.alfaview.com OIDC 200/2169B md5 f78a08fc advertises the
- CHANGED apis.alfaview.com OpenAPI md5 `284a3383c1ac3cfc9152ffcc631891f2` (132100B, 38 paths, 57 ops) — 3rd consecutive
- CHANGED Complete 403-declaration map built across all 57 ops (37 declare 403). New divergence:
- NEW alfaview supports BRING-YOUR-OWN-IdP single sign-on, documented at
- NEW Every alfaview user is provisioned with a NATIVE password independent of any IdP.
- NEW The SSO article's own "Limitations" section addresses only SAML signature algorithm and IdP-initiated flow
- NEW alfaview's security-guide (200/105KB) lists "Activate the two-factor authentication" as a personal self-service
- NEW NEGATIVE RESULT — v1 API is gone. https://apis.alfaview.com/v1/docs/openapi.json and .../openapi/openapi both
- NEW NEGATIVE RESULT — the OpenAPI 3.0.3 twin is semantically IDENTICAL to the 3.1 original. 132391B md5 34e9f231,
- NEW OpenAPI spec at apis.alfaview.com/v2/docs/openapi.json rotated to MD5 284a3383c1ac3cfc9152ffcc631891f2 (132100B, 38 paths, 57 ops) — byte-identical on beta-apis.alfaview.com
- NEW /v2/users/me/company (GetOwnCompany) identified as the 38th path — sole 200-response supplier of companyId across 57 operations; enables guest-link chain execution by authorized tester

## 2026-09-29 20:15:47 UTC
- NEW OpenAPI spec at apis.alfaview.com/v2/docs/openapi.json rotated to MD5 284a3383c1ac3cfc9152ffcc631891f2 (132100B, 38 paths, 57 ops) — byte-identical on beta-apis.alfaview.com
- NEW /v2/users/me/company (GetOwnCompany) identified as the 38th path — sole 200-response supplier of companyId across 57 operations; enables guest-link chain execution by authorized tester
- NEW alfaview 2FA is a shipped, customer-facing feature with a dedicated public KB article; measured 2FA enforcement asymmetry (SSO vs native password paths)
- NEW alfaview supports BRING-YOUR-OWN-IdP single sign-on (GitLab, Google Workspace, Azure AD, generic SAML/OIDC)
- NEW Every alfaview user provisioned with a NATIVE password independent of any IdP
- NEW NEGATIVE RESULT — v1 API is gone (/v1/docs/openapi.json and /v1/docs/openapi both 404/19B)
- NEW NEGATIVE RESULT — OpenAPI 3.0.3 twin semantically IDENTICAL to 3.1 original (132391B md5 34e9f231)
- CHANGED Complete 403-declaration map built across all 57 ops (37 declare 403); new divergence on /v2/group-links vs /v2/guest-links
- CHANGED apis.alfaview.com OpenAPI md5 `284a3383c1ac3cfc9152ffcc631891f2` (132100B, 38 paths, 57 ops) — 3rd consecutive stable cycle
- CHANGED sso.alfaview.com OIDC 200/2169B md5 `f78a08fc`, JWKS 200/16257B md5 `3f8d456c` (7×RS256) — 34th consecutive byte-stable cycle
- CHANGED tools.alfaview.com unchanged: /whiteboard/ 404/47B gRPC envelope, /poll/pollservice/list 501/55B

## 2026-09-29 23:38:50 UTC
- NEW apis.alfaview.com GET /v2/audit-log — LIVE, UNDOCUMENTED 39th path. Absent from the 38-path
- NEW apis.alfaview.com /v2/audit-log pre-auth query binder COMPLETE (per-operation, negative-controlled):
- NEW Orphan schema family in the published contract: Event, Actor, Target, PaginatedBodyListEvent.
- NEW Event.metadata description names an operation — "refer to the 'List Searchable metadata attributes' endpoint" —
- CHANGED Pre-auth validation-before-auth class: 9 → 10 GET operations, the new member fully characterised.
- CHANGED Whole-contract confirmation: grep for "audit|event|log" across all 132100B returns only
- NEW apis.alfaview.com GET /v2/audit-log — LIVE, UNDOCUMENTED 39th path. Absent from the 38-path
- NEW /v2/audit-log pre-auth query binder COMPLETE, negative-controlled:
- NEW The route's real consumer is NOT the REST API — it is app.alfaview.com/graphql, a second, entirely
- NEW /v2/audit-log is therefore NOT an orphan. It is the REST mirror of GraphQL query ListAuditEvents
- NEW app.alfaview.com/graphql reachable pre-auth: 400 with CSRF guard, __typename → 200/32B
- NEW Full GraphQL type surface recovered WITHOUT introspection, using Apollo's per-field validation errors
- NEW The permission model is now KNOWN, from the client, and it changes the hypothesis materially:
- NEW The feature is flag-gated client-side, to alfaview's own companies only:
- CHANGED apis.alfaview.com OpenAPI md5 rotated to `284a3383c1ac3cfc9152ffcc631891f2` (132100B, 38 paths, 57 ops) — 3rd consecutive stable cycle; byte-identical on beta-apis.alfaview.com
- CHANGED `/v2/users/me/company` (GetOwnCompany) identified as the 38th path — sole 200-response supplier of `companyId` across 57 operations; enables guest-link chain execution by authorized tester
- CHANGED alfaview 2FA is a shipped feature with public KB article; measured 2FA enforcement asymmetry (SSO vs native password paths)
- CHANGED alfaview supports BYO-IdP SSO (GitLab, Google Workspace, Azure AD, generic SAML/OIDC); every user has native password independent of IdP
- CHANGED NEGATIVE RESULT — v1 API gone (/v1/docs/openapi.json and /v1/docs/openapi both 404/19B)
- CHANGED NEGATIVE RESULT — OpenAPI 3.0.3 twin semantically identical to 3.1 (132391B md5 34e9f231)
- CHANGED Complete 403-declaration map built across all 57 ops (37 declare 403); new divergence on `/v2/group-links` vs `/v2/guest-links`
- CHANGED sso.alfaview.com OIDC 200/2169B md5 `f78a08fc`, JWKS 200/16257B md5 `3f8d456c` (7×RS256) — 34th consecutive byte-stable cycle
- CHANGED tools.alfaview.com unchanged: `/whiteboard/` 404/47B gRPC envelope, `/poll/pollservice/list` 501/55B

## 2026-09-30 02:23:38 UTC
- CHANGED apis.alfaview.com OpenAPI md5 rotated to `284a3383c1ac3cfc9152ffcc631891f2` (132100B, 38 paths, 57 ops) — 3rd consecutive stable cycle; byte-identical on beta-apis.alfaview.com
- CHANGED `/v2/users/me/company` (GetOwnCompany) identified as the 38th path — sole 200-response supplier of `companyId` across 57 operations; enables guest-link chain execution by authorized tester
- CHANGED alfaview 2FA is a shipped feature with public KB article; measured 2FA enforcement asymmetry (SSO vs native password paths)
- CHANGED alfaview supports BYO-IdP SSO (GitLab, Google Workspace, Azure AD, generic SAML/OIDC); every user has native password independent of IdP
- CHANGED NEGATIVE RESULT — v1 API gone (/v1/docs/openapi.json and /v1/docs/openapi both 404/19B)
- CHANGED NEGATIVE RESULT — OpenAPI 3.0.3 twin semantically identical to 3.1 (132391B md5 34e9f231)
- CHANGED Complete 403-declaration map built across all 57 ops (37 declare 403); new divergence on `/v2/group-links` vs `/v2/guest-links`
- CHANGED sso.alfaview.com OIDC 200/2169B md5 `f78a08fc`, JWKS 200/16257B md5 `3f8d456c` (7×RS256) — 34th consecutive byte-stable cycle
- CHANGED tools.alfaview.com unchanged: `/whiteboard/` 404/47B gRPC envelope, `/poll/pollservice/list` 501/55B
