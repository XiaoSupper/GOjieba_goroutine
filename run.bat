cls
@echo off
color 06
:MENU
ECHO ѡ�����
ECHO 1 go run 
ECHO 2 go build
ECHO 3 main.exe
ECHO 4 ��    ��
echo ������ѡ����ţ�
set /p ID=
if "%id%"=="1" goto cmd1
if "%id%"=="2" goto cmd2
if "%id%"=="3" goto cmd3
exit

:cmd1
time /T
go run main.go -GOjieba config.json 
time /T
goto MENU

:cmd2
time /T
go build main.go 
time /T
GOTO MENU

:cmd3
time /T
start "GOjieba" main.exe
time /T
GOTO MENU