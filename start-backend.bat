@echo off
echo ========================================
echo   SI-AZISAH Backend Server
echo ========================================
echo.

echo Checking Go installation...
go version
if errorlevel 1 (
    echo ERROR: Go is not installed!
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo.
echo Starting backend server...
echo Server will run on http://localhost:8080
echo Press Ctrl+C to stop the server
echo.

cd backend
go run cmd/main.go

pause
