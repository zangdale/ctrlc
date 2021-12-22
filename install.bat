@echo off
rem run this script as admin

if not exist ctrlc.exe (
    echo Build the ctrlc server before installing by running "go build"
    goto :exit
)

sc create go-svc-ctrlc binpath= "%CD%\ctrlc.exe" start= auto DisplayName= "go-svc-ctrlc-desplay"
sc description go-svc-ctrlc "go-svc-ctrlc-description"
net start go-svc-ctrlc
sc query go-svc-ctrlc

echo Check ctrlc-log.log

:exit
