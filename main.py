"""Tax Calculation API - Main Application."""
import os
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles
from fastapi.responses import FileResponse
from handlers import router
import uvicorn


# Create FastAPI application
app = FastAPI(
    title="Tax Calculation API",
    description="API for calculating sales tax based on location",
    version="1.0.0"
)

# Configure CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include API routes
app.include_router(router)

# Serve static files
app.mount("/static", StaticFiles(directory="static"), name="static")


@app.get("/")
async def serve_index():
    """Serve the main index.html file."""
    return FileResponse("static/index.html")


def main():
    """Run the application."""
    port = int(os.getenv("PORT", "8080"))
    
    print(f"Server starting on port {port}")
    print(f"Web UI available at http://localhost:{port}")
    print(f"API endpoints at http://localhost:{port}/api/v1/")
    print(f"API documentation at http://localhost:{port}/docs")
    
    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=port,
        reload=True
    )


if __name__ == "__main__":
    main()
