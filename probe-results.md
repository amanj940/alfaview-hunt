
## 2026-09-03 11:40:30 UTC


## 2026-09-03 14:23:25 UTC
https://beta-apis.alfaview.com/v2/languages -> HTTP 401
https://beta-apis.alfaview.com/v2/languages` -> HTTP 404
https://apis.alfaview.com/v2/languages` -> HTTP 404

## 2026-09-03 15:21:47 UTC
https://beta-apis.alfaview.com/v2/languages` -> HTTP 404
https://apis.alfaview.com/v2/languages` -> HTTP 404
https://beta-apis.alfaview.com/openapi.json -> HTTP 404
https://apis.alfaview.com/openapi.json -> HTTP 404
https://beta-apis.alfaview.com/openapi.json` -> HTTP 404
https://apis.alfaview.com/openapi.json` -> HTTP 404
https://demo-company.alfaview.com/ -> 200 len=1381
https://demo-company.alfaview.com/api/v1/users -> 200 len=1381
https://demo-company.alfaview.com/api/v1/users` -> 200 len=1381

## 2026-09-03 18:48:54 UTC
https://apis.alfaview.com/v2/docs/openapi.json` -> HTTP 404
https://beta-apis.alfaview.com/v2/docs/openapi.json` -> HTTP 404
https://apis.alfaview.com/v2/languages` -> HTTP 404
https://beta-apis.alfaview.com/v2/languages` -> HTTP 404
https://beta-apis.alfaview.com/v2/docs/openapi.json -> 200 len=?
https://apis.alfaview.com/v2/docs/openapi.json -> 200 len=?
https://beta-apis.alfaview.com/v2/debug` -> HTTP 404
https://beta-apis.alfaview.com/v2/admin` -> HTTP 404
https://beta-apis.alfaview.com/v2/internal` -> HTTP 404
https://beta-apis.alfaview.com/v2/test` -> HTTP 404
https://beta-apis.alfaview.com/v2/health` -> HTTP 404
https://insider-webclient.alfaview.com/ -> 200 len=4396

## 2026-09-03 21:22:16 UTC
https://insider-webclient.alfaview.com/ -> 200 len=4396
https://insider-webclient.alfaview.com/api/* -> HTTP 404
https://beta-apis.alfaview.com/v2/languages -> HTTP 401
https://insider-webclient.alfaview.com/` -> HTTP 404
https://insider-webclient.alfaview.com/api` -> HTTP 404
https://insider-webclient.alfaview.com/admin` -> HTTP 404
https://insider-webclient.alfaview.com/debug` -> HTTP 404
https://insider-webclient.alfaview.com/internal` -> HTTP 404
https://insider-webclient.alfaview.com/health` -> HTTP 404
https://insider-webclient.alfaview.com/api/health -> HTTP 404
https://insider-webclient.alfaview.com/api/config -> HTTP 404

## 2026-09-03 23:25:03 UTC
https://alfacheck-audio.alfaview.com/` -> ERR The read operation timed out
https://alfacheck-engine.alfaview.com/` -> ERR The read operation timed out
https://alfacheck-video.alfaview.com/` -> ERR The read operation timed out

## 2026-09-04 01:14:01 UTC
https://beta-hcloud-19-beta-hydra-dzwx8.alfaview.com/ -> 200 len=9
https://beta-hcloud-19-beta-hydra-dzwx8.alfaview.com/` -> HTTP 400
https://alfacheck-engine.alfaview.com/ -> ERR The read operation timed out
https://alfacheck-engine.alfaview.com/health -> ERR The read operation timed out
https://alfacheck-engine.alfaview.com/status -> ERR The read operation timed out
https://alfacheck-audio.alfaview.com/ -> ERR The read operation timed out
https://alfacheck-audio.alfaview.com/media -> ERR The read operation timed out
https://alfacheck-audio.alfaview.com/recordings -> ERR The read operation timed out
https://alfacheck-video.alfaview.com/ -> ERR The read operation timed out

## 2026-09-04 06:01:02 UTC
https://apis.alfaview.com/v2/docs/openapi.json` -> HTTP 404
https://beta-hcloud-19-beta-hydra-dzwx8.alfaview.com/ -> 200 len=9
https://beta-hcloud-19-beta-hydra-dzwx8.alfaview.com/.well-known/openid-configuration -> 200 len=9
https://beta-hcloud-19-beta-hydra-dzwx8.alfaview.com/health -> 200 len=9

## 2026-09-04 10:12:17 UTC
https://apis.alfaview.com/v2/auth/guest-link` -> HTTP 404

## 2026-09-04 14:25:13 UTC
https://apis.alfaview.com/v2/auth/guest-link` -> HTTP 404
https://alfatraining.alfaview.com/ -> 200 len=1381
https://bhc.alfaview.com/ -> 200 len=1381
https://kh-freiburg.alfaview.com/ -> 200 len=1381

## 2026-09-04 17:51:06 UTC
https://apis.alfaview.com/v2/auth/guest-link` -> HTTP 404

## 2026-09-04 20:04:50 UTC
https://beta-ionoscloud-21-beta-hydra-7x5d5.alfaview.com/ -> ERR <urlopen error timed out>
https://alfaview.com/ -> 200 len=?
https://sso.alfaview.com/ -> 200 len=0
https://beta-ionoscloud-21-beta-hydra-7x5d5.alfaview.com/` -> ERR <urlopen error timed out>

## 2026-09-04 22:10:15 UTC
https://alfaview.com/ -> 200 len=?
https://sso.alfaview.com/ -> 200 len=0
https://alfaview.com/@evil.com -> HTTP 404
https://beta-ionoscloud-21-beta-engine-gw4qw.alfaview.com/ -> ERR <urlopen error timed out>
https://alfaview.com/` -> HTTP 404
https://sso.alfaview.com/authorize?client_id=test&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 404

## 2026-09-05 00:15:17 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://sso.alfaview.com/oauth2/authorize?client_id=<valid_id>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test -> HTTP 400
https://sso.alfaview.com/oauth2/token -> HTTP 405
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/js/app.min.5b3949112f0cf682adc8.js` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://sso.alfaview.com/oauth2/register` -> HTTP 404

## 2026-09-05 04:34:25 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/js/app.min.5b3949112f0cf682adc8.js -> 200 len=1090319
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://sso.alfaview.com/.well-known/openid-configuration -> 200 len=0
https://sso.alfaview.com/.well-known/jwks.json -> 200 len=0
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/js/app.min.5b3949112f0cf682adc8.js` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400

## 2026-09-05 08:51:14 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://sso.alfaview.com/.well-known/openid-configuration -> 200 len=0
https://sso.alfaview.com/.well-known/jwks.json -> 200 len=0
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://app.alfaview.com/graphql` -> 200 len=1381
https://apis.alfaview.com/v2/auth/guest-link` -> HTTP 404

## 2026-09-05 12:20:14 UTC
https://apis.alfaview.com/v2/auth/password -> HTTP 405
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/rooms/{victimRoomId -> HTTP 401

## 2026-09-05 15:02:37 UTC


## 2026-09-05 17:06:28 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://app.alfaview.com -> 200 len=1381

## 2026-09-05 18:57:10 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://test.alfaview.com/` -> HTTP 404

## 2026-09-05 20:54:05 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://sso.alfaview.com/oauth2/register` -> HTTP 404
https://app.alfaview.com/graphql` -> 200 len=1381
https://apis.alfaview.com/v2/auth/guest-link` -> HTTP 404
https://apis.alfaview.com/v2/auth/password -> HTTP 405
https://apis.alfaview.com/v2/users/me -> HTTP 401

## 2026-09-05 22:28:31 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400

## 2026-09-06 00:15:59 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400

## 2026-09-06 04:42:09 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400

## 2026-09-06 08:57:21 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://apis.alfaview.com/v2/rooms/{victimRoomId -> HTTP 401

## 2026-09-06 12:29:06 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://apis.alfaview.com/v2/rooms/{victimRoomId -> HTTP 401

## 2026-09-06 16:02:46 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://apis.alfaview.com/v2/rooms/{victimRoomId -> HTTP 401

## 2026-09-06 18:09:34 UTC


## 2026-09-06 19:59:08 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400

## 2026-09-06 21:51:30 UTC
https://sso.alfaview.com/oauth2/authorize` -> HTTP 404
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://test.alfaview.com/ -> 200 len=494
https://sso.alfaview.com/oauth2/authorize?client_id=<x>&redirect_uri=https://evil.com&response_type=code -> HTTP 400
https://apis.alfaview.com/v2/rooms/{victimRoomId -> HTTP 401

## 2026-09-06 23:22:06 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401
https://sso.alfaview.com/oauth2/authorize?client_id=<x>&redirect_uri=https://evil.com&response_type=code -> HTTP 400

## 2026-09-07 01:08:49 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-07 06:09:58 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-07 12:48:44 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401
https://apis.alfaview.com/v2/rooms/{victimRoomId -> HTTP 401
https://test.alfaview.com/ -> 200 len=494
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123` -> HTTP 400
https://sso.alfaview.com/oauth2/authorize?client_id=<x>&redirect_uri=https://evil.com&response_type=code -> HTTP 400

## 2026-09-07 18:00:33 UTC
https://sso.alfaview.com/oauth2/authorize -> 200 len=0
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-07 21:00:42 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-07 23:07:56 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-08 01:17:36 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-08 06:04:20 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-08 10:36:26 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-08 14:53:52 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-08 18:18:08 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-08 21:14:52 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-08 23:25:39 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-09 01:31:37 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-09 06:50:41 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-09 12:00:50 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-09 15:42:21 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-09 18:53:12 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-09 21:21:25 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-09 23:20:30 UTC


## 2026-09-10 01:10:12 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-10 06:04:23 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-10 10:40:33 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/me -> HTTP 401
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-room-id -> HTTP 401

## 2026-09-10 14:51:59 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-10 18:00:06 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-10 20:11:27 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-10 22:38:24 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-11 00:31:59 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-11 05:14:58 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://beta-apis.alfaview.com/v2/languages -> HTTP 401
https://beta-apis.alfaview.com/v2/languages` -> HTTP 404
https://apis.alfaview.com/v2/languages` -> HTTP 404
https://demo-company.alfaview.com/ -> 200 len=1381
https://demo-company.alfaview.com/api/v1/users -> 200 len=1381
https://demo-company.alfaview.com/api/v1/users` -> 200 len=1381
https://insider-webclient.alfaview.com/ -> 200 len=4396
https://insider-webclient.alfaview.com/api/health -> HTTP 404

## 2026-09-11 09:46:07 UTC


## 2026-09-11 14:02:32 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-11 17:24:54 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-11 20:01:30 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-11 22:20:43 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 00:22:42 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 04:46:43 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 09:01:19 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 12:31:57 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 15:58:19 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 18:02:44 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 19:50:10 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 21:47:21 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-12 23:32:42 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-13 01:24:34 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-13 06:47:45 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-13 12:27:46 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-13 16:55:46 UTC


## 2026-09-13 19:00:00 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-13 21:13:16 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-13 23:12:25 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-14 01:10:15 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-14 06:21:04 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-14 13:04:53 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-14 18:20:49 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-14 21:57:40 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-15 00:04:46 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-15 05:00:37 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-15 09:46:34 UTC
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-15 14:51:14 UTC
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405

## 2026-09-15 18:38:53 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/authorize?client_id=<found>&redirect_uri=https://evil.com&response_type=code&scope=openid&state=test123 -> HTTP 400
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405

## 2026-09-15 21:45:16 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-15 23:48:55 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-16 01:58:37 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-16 07:03:07 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-16 12:30:02 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-16 17:18:38 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-16 20:09:56 UTC
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-16 22:58:04 UTC


## 2026-09-17 00:59:42 UTC


## 2026-09-17 05:42:34 UTC
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/ -> 200 len=1381
https://app.alfaview.com/graphql -> HTTP 400
https://beta-apis.alfaview.com/v2/languages -> HTTP 401
https://beta-apis.alfaview.com/v2/languages` -> HTTP 404
https://apis.alfaview.com/v2/languages` -> HTTP 404

## 2026-09-17 10:35:23 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-17 15:18:52 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-17 19:07:06 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-17 22:11:15 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405

## 2026-09-18 00:24:27 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{other-tenant-user-uuid -> HTTP 405
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-18 05:10:18 UTC
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400

## 2026-09-18 09:50:24 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501

## 2026-09-18 14:04:29 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-18 17:25:49 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-18 19:58:28 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-18 22:08:01 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615

## 2026-09-18 23:56:31 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615

## 2026-09-19 01:57:24 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/int -> HTTP 404
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-19 06:50:32 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-19 11:45:31 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-19 15:00:03 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-19 17:34:51 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-19 19:37:33 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://staging-tools.alfaview.com/poll/pollservice/list -> HTTP 501

## 2026-09-19 21:45:24 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-19 23:43:44 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-20 01:57:34 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-20 07:22:34 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-20 12:33:04 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404

## 2026-09-20 16:40:17 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-20 19:03:57 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-20 21:35:44 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-20 23:29:18 UTC
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-21 01:36:02 UTC
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-21 07:05:02 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-21 14:15:38 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-21 19:31:21 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-21 22:49:31 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-22 01:21:25 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-22 06:27:16 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-22 11:51:57 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-22 15:56:35 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-22 19:24:42 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-22 22:17:53 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-23 00:38:50 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-23 05:07:19 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-23 10:11:31 UTC


## 2026-09-23 14:46:54 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-23 18:44:58 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-23 21:53:05 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-24 00:06:56 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-24 04:55:09 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-24 09:40:12 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-24 14:28:56 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-24 18:39:58 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-24 21:47:23 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-25 00:08:30 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-25 04:59:19 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-25 10:03:26 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://app.alfaview.com/js/app.min.5b3949112f0cf682adc8.js -> 200 len=1381
https://alfaview-com-assets.alfaview.com/production/alfaview-com-frontend/js/app.min.67e8a68d4318b34ca241.js -> 200 len=1092529

## 2026-09-25 15:01:49 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://alfaview-com-assets.alfaview.com/production/alfaview-com-frontend/js/app.min.67e8a68d4318b34ca241.js` -> HTTP 403
https://sso.alfaview.com/.well-known/openid-configuration` -> HTTP 404
https://sso.alfaview.com/.well-known/jwks.json` -> HTTP 404
https://sso.alfaview.com/oauth2/authorize?client_id=probe-invalid-20260925&redirect_uri=https%3A%2F%2Fexample.invalid%2Fcb&response_type=code` -> 200 len=0
https://tools.alfaview.com/` -> 200 len=615

## 2026-09-25 19:11:11 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-roomId -> HTTP 401
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-25 22:29:17 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-roomId -> HTTP 401
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://tools.alfaview.com/` -> 200 len=615
https://tools.alfaview.com/poll/pollservice/list` -> HTTP 404
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615

## 2026-09-26 00:55:24 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-roomId -> HTTP 401
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://tools.alfaview.com/` -> 200 len=615
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js` -> 200 len=615

## 2026-09-26 05:42:35 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/js/app-bundle.0a250e1f96a7aaf5661c.js -> 200 len=137343
https://tools.alfaview.com/poll/pollservice/list -> HTTP 501
https://app.alfaview.com/graphql -> HTTP 400
https://apis.alfaview.com/v2/users/{foreign-uuid -> HTTP 405
https://apis.alfaview.com/v2/rooms/{foreign-roomId -> HTTP 401
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/users/me` -> HTTP 405
https://apis.alfaview.com/v2/auth/group-link` -> HTTP 404

## 2026-09-26 10:17:13 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://app.alfaview.com/graphql -> HTTP 400
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404
https://apis.alfaview.com/v2/stats -> HTTP 422
https://apis.alfaview.com/v2/auth/token-info -> HTTP 422
https://apis.alfaview.com/v2/rooms/not-a-uuid -> HTTP 401
https://sso.alfaview.com/.well-known/openid-configuration -> 200 len=0
https://app.alfaview.com/js/app.min.67e8a68d4318b34ca241.js -> 200 len=1092529

## 2026-09-26 14:33:02 UTC
https://apis.alfaview.com/v2/guest-links?limit=abc` -> HTTP 422
https://apis.alfaview.com/v2/guest-links?pageToken=x` -> HTTP 401
https://apis.alfaview.com/v2/rooms/{any-uuid -> HTTP 401
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://sso.alfaview.com/.well-known/openid-configuration` -> HTTP 404
https://app.alfaview.com/graphql` -> 200 len=1381

## 2026-09-26 18:10:18 UTC
https://apis.alfaview.com/v2/rooms/00000000-0000-4000-8000-000000000001/attendances -> HTTP 422
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404
https://app.alfaview.com/` -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://sso.alfaview.com/.well-known/openid-configuration` -> HTTP 404
https://apis.alfaview.com/v2/guest-links?limit=1&limit=abc` -> HTTP 401

## 2026-09-26 20:34:39 UTC
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://apis.alfaview.com/v2/rooms/00000000-0000-4000-8000-000000000001/attendances -> HTTP 422
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404

## 2026-09-26 23:08:46 UTC
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://apis.alfaview.com/v2/rooms/00000000-0000-4000-8000-000000000001/attendances -> HTTP 422
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404

## 2026-09-27 01:29:44 UTC
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://apis.alfaview.com/v2/rooms/00000000-0000-4000-8000-000000000001/attendances -> HTTP 422
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404

## 2026-09-27 07:05:46 UTC
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://apis.alfaview.com/v2/rooms/00000000-0000-4000-8000-000000000001/attendances -> HTTP 422
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404
https://sso.alfaview.com/.well-known/openid-configuration -> 200 len=0
https://sso.alfaview.com/oauth2/userinfo -> HTTP 401
https://sso.alfaview.com/oauth2/userinfo` -> HTTP 404
https://sso.alfaview.com` -> ERR <urlopen error [Errno -2] Name or service not know

## 2026-09-27 13:06:51 UTC
https://apis.alfaview.com/v2/rooms/00000000-0000-4000-8000-000000000001/attendances -> HTTP 422
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404
https://sso.alfaview.com -> 200 len=0
https://apis.alfaview.com/v2/users/invitations` -> HTTP 405
https://apis.alfaview.com/v2/users/invitation` -> HTTP 405
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404

## 2026-09-27 17:56:18 UTC
https://apis.alfaview.com/v2/users/invitation -> HTTP 405
https://apis.alfaview.com/v2/users/invitations -> HTTP 405
https://apis.alfaview.com/v2/users/invitation/{valid-uuid -> HTTP 404
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404
https://apis.alfaview.com/v2/guest-links/00000000-0000-0000-0000-000000000000 -> HTTP 401
https://sso.alfaview.com/.well-known/openid-configuration -> 200 len=0

## 2026-09-27 20:38:39 UTC
https://app.alfaview.com/ -> 200 len=1381
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://tools.alfaview.com/whiteboard/ -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/ -> HTTP 404
https://apis.alfaview.com/v2/guest-links/00000000-0000-0000-0000-000000000000 -> HTTP 401
https://apis.alfaview.com/v2/guest-links/{owned-uuid -> HTTP 401
https://tools.alfaview.com/whiteboard/` -> HTTP 404
https://tools.alfaview.com/whiteboard/List` -> HTTP 404
https://staging-tools.alfaview.com/whiteboard/` -> HTTP 404
https://apis.alfaview.com/v2/guest-links?roomId=00000000-0000-0000-0000-000000000000 -> HTTP 401
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/.well-known/openid-configuration -> 200 len=0

## 2026-09-27 23:25:45 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://apis.alfaview.com/v2/guest-links?roomId=abc&limit=abc` -> HTTP 422
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/guest-links?limit=1` -> HTTP 422

## 2026-09-28 02:01:33 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/guest-links?limit=1` -> HTTP 422
https://apis.alfaview.com/v2/auth/api-key` -> HTTP 404
https://sso.alfaview.com/.well-known/openid-configuration` -> HTTP 404
https://sso.alfaview.com/.well-known/jwks.json` -> HTTP 404

## 2026-09-28 08:06:40 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://apis.alfaview.com/v2/zzznotreal` -> HTTP 404
https://apis.alfaview.com/v2/languages` -> HTTP 404
https://apis.alfaview.com/v2/users/invitations` -> HTTP 405
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/guest-links?limit=1` -> HTTP 422
https://sso.alfaview.com/.well-known/openid-configuration` -> HTTP 404

## 2026-09-28 16:52:32 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/rooms/<own-room-id>/participants` -> HTTP 404
https://sso.alfaview.com/.well-known/openid-configuration` -> HTTP 404
https://sso.alfaview.com/.well-known/jwks.json` -> HTTP 404
https://apis.alfaview.com/v2/rooms/{id -> HTTP 401

## 2026-09-28 22:27:47 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://sso.alfaview.com/oauth2/introspect` -> HTTP 404
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/users/me/company` -> HTTP 404
https://apis.alfaview.com/v2/guest-links/<own-guest-link-id>` -> HTTP 401
https://apis.alfaview.com/v2/docs/openapi.json` -> HTTP 404
https://sso.alfaview.com/.well-known/openid-configuration` -> HTTP 404
https://sso.alfaview.com/.well-known/jwks.json` -> HTTP 404

## 2026-09-29 02:07:04 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/users/me/company` -> HTTP 404

## 2026-09-29 08:35:24 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/users/me/company` -> HTTP 404
https://apis.alfaview.com/v2/guest-links?limit=1` -> HTTP 422
https://app.alfaview.com/graphql` -> 200 len=1381

## 2026-09-29 15:23:00 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
https://apis.alfaview.com/v2/auth/token-info` -> HTTP 404
https://apis.alfaview.com/v2/docs/openapi.json -> 200 len=?
https://apis.alfaview.com/v1/docs/openapi.json -> HTTP 404

## 2026-09-29 20:15:54 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405

## 2026-09-29 23:38:57 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405

## 2026-09-30 02:23:46 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405

## 2026-09-30 08:48:34 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405

## 2026-09-30 15:31:56 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405

## 2026-09-30 20:25:28 UTC
https://apis.alfaview.com/v2/guest-links -> HTTP 401
https://apis.alfaview.com/v2/group-links -> HTTP 401
https://apis.alfaview.com/v2/auth/guest-link -> HTTP 405
https://app.alfaview.com/graphql -> HTTP 400
https://sso.alfaview.com/oauth2/introspect -> HTTP 405
