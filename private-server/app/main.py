import os
import hmac
import logging
from ipaddress import ip_address
from pathlib import Path

from dotenv import load_dotenv
from fastapi import FastAPI, Header, HTTPException, Request, status

load_dotenv(Path(__file__).resolve().parent.parent / ".env")

app = FastAPI(title="Private Server", version="1.0.0")
logger = logging.getLogger("private-server")

PRIVATE_SERVER_AUTH_TOKEN = os.getenv("PRIVATE_SERVER_AUTH_TOKEN", "")
SERVER_SECRET = os.getenv("SERVER_SECRET", "")
ALLOWED_CALLER_IPS = frozenset(
    ip_address(value.strip())
    for value in os.getenv("ALLOWED_CALLER_IPS", "").split(",")
    if value.strip()
)

def is_allowed_ip(host: str | None) -> bool:
    if not host:
        return False
    try:
        return ip_address(host) in ALLOWED_CALLER_IPS
    except ValueError:
        return False


@app.get("/health")
async def health() -> dict[str, str]:
    return {"status": "ok"}


@app.get("/server-secret")
async def server_secret(
    request: Request,
    x_internal_token: str | None = Header(default=None),
) -> dict[str, str]:
    caller_host = request.client.host if request.client else None
    caller_ip = caller_host or "unknown"
    ip_allowed = is_allowed_ip(caller_host)
    token_valid = (
        bool(PRIVATE_SERVER_AUTH_TOKEN)
        and bool(x_internal_token)
        and hmac.compare_digest(x_internal_token, PRIVATE_SERVER_AUTH_TOKEN)
    )

    if not ip_allowed:
        logger.warning("UNVERIFIED /server-secret request from non-whitelisted IP %s", caller_ip)
    elif not token_valid:
        logger.warning("UNVERIFIED /server-secret request from whitelisted IP %s: invalid token", caller_ip)

    if not ip_allowed or not token_valid:
        raise HTTPException(
            status_code=status.HTTP_403_FORBIDDEN,
            detail="Only the authorized local service may access this route",
        )

    logger.info("VERIFIED /server-secret request from whitelisted IP %s", caller_ip)
    return {
        "message": "Server secret retrieved successfully",
        "key": SERVER_SECRET,
    }