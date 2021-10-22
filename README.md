Для запуска необходимо установить:   
go get github.com/lxn/walk  
go get -d  github.com/akavel/rsrc

Далее необходимо создать *.manifest и скомпилировать его  
Команда компиляции:
*\go\bin\rsrc.exe -manifest .\GUI_Interface\GUI.manifest -o rsrc.syso

После этого собираем файл:
go build

/*Для сборки файла без системного окна (cmd)*/  
go build -ldflags="-H windowsgui"

Далее просто вызов получившегося исполняемого файла: *.exe
