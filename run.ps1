
Start-Process -WindowStyle hidden -FilePath .\ctrlc.exe
$LastExitCode
ps | findstr ctrlc
