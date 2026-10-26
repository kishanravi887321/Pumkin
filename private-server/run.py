import os

import uvicorn


if __name__ == "__main__":
    uvicorn.run(
        "app.main:app",
        host=os.getenv("PRIVATE_SERVER_HOST", "127.0.0.1"),
        port=int(os.getenv("PRIVATE_SERVER_PORT", "8179")),
        reload=os.getenv("APP_ENV", "development") == "development",
    )