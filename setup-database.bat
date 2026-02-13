@echo off
echo ========================================
echo   SI-AZISAH Database Setup
echo ========================================
echo.

set /p DB_USER="Enter MySQL username (default: root): "
if "%DB_USER%"=="" set DB_USER=root

set /p DB_PASS="Enter MySQL password (press Enter if no password): "

echo.
echo Creating database...
mysql -u %DB_USER% -p%DB_PASS% -e "CREATE DATABASE IF NOT EXISTS siazisah;"

if errorlevel 1 (
    echo ERROR: Failed to create database!
    echo Please check your MySQL credentials.
    pause
    exit /b 1
)

echo.
echo Importing schema...
mysql -u %DB_USER% -p%DB_PASS% siazisah < database\schema.sql

if errorlevel 1 (
    echo ERROR: Failed to import schema!
    pause
    exit /b 1
)

echo.
echo Do you want to import sample data? (Y/N)
set /p IMPORT_SAMPLE=
if /i "%IMPORT_SAMPLE%"=="Y" (
    echo Importing sample data...
    mysql -u %DB_USER% -p%DB_PASS% siazisah < database\sample_data.sql
    echo Sample data imported!
)

echo.
echo ========================================
echo   Database setup completed!
echo ========================================
echo.
echo Next steps:
echo 1. Generate password for superadmin:
echo    cd backend
echo    go run cmd/generate_password.go admin123
echo.
echo 2. Update password in database:
echo    mysql -u %DB_USER% -p%DB_PASS% siazisah
echo    UPDATE users SET password = 'your_hash' WHERE username = 'superadmin';
echo.
echo 3. Start backend server:
echo    start-backend.bat
echo.

pause
