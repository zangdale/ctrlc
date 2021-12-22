@echo off
rem run this script as admin

net stop go-svc-ctrlc
sc delete go-svc-ctrlc
