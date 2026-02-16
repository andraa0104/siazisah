-- Update password superadmin ke admin123
-- Hash: $2a$14$vSC1dcQRbgAFj.M8kV9gTOpTNRCis14naHk6SsUYJOdR99oGnOaJG

UPDATE users 
SET password = '$2a$14$vSC1dcQRbgAFj.M8kV9gTOpTNRCis14naHk6SsUYJOdR99oGnOaJG' 
WHERE username = 'superadmin';

-- Verify update
SELECT id, username, role, is_active FROM users WHERE username = 'superadmin';
