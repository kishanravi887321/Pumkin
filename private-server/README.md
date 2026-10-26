# Private FastAPI server

This service listens on `127.0.0.1:8179` by default. The `/server-secret`
route is available only when both conditions are met:

1. The request source IP is listed in `ALLOWED_CALLER_IPS`.
2. The caller sends the configured shared token:

```http
X-Internal-Token: change-this-shared-token
```

Set the same strong `PRIVATE_SERVER_AUTH_TOKEN` in both the Python and Go
service environments. `ALLOWED_CALLER_IPS` accepts comma-separated IP
addresses, for example `127.0.0.1,10.0.0.5`. Change both the allowlisted IP
and token for production as needed.

The server checks the actual TCP peer address and does not trust
`X-Forwarded-For` or other client-controlled URL/host headers. Requests from
all other IP addresses, URLs, or routes are rejected.

## Run

```powershell
python -m venv .venv
.venv\Scripts\Activate.ps1
pip install -r requirements.txt
Copy-Item .env.example .env
python run.py
```

The existing JavaScript service should call:

```js
fetch("http://127.0.0.1:8179/server-secret", {
  headers: { "X-Internal-Token": "change-this-shared-token" }
});
```

The Go backend calls `PRIVATE__SERVER__URL` and sends
`PRIVATE_SERVER_AUTH_TOKEN` as `X-Internal-Token`.