"""Health check endpoint."""

from fastapi import APIRouter, Depends

from middleware.auth import verify_token

router = APIRouter(tags=["health"])


@router.get("/health")
def health(_: str = Depends(verify_token)):
    """
    Health check endpoint.

    Returns the service status and identifier.
    Requires Bearer token authentication.
    """
    return {"status": "ok", "service": "python"}
