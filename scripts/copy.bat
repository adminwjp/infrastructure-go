@echo
rem del E:\software\go\src\github.com\adminwjp /s /h /e /y
cd e:
rem https://www.yisu.com/zixun/404689.html
rem md software\go\src\github.com\adminwjp\infrastructure-go
xcopy   E:\work\utility\Utility-for-go\infrastructure E:\software\go\src\github.com\adminwjp\infrastructure-go  /s /h /e  /y

rem for /r "E:\work\utility\Utility-for-go\infrastructure" %f in (*.*,*) do @xcopy "%f" "E:\software\go\src\github.com\adminwjp\infrastructure-go"

rem xcopy ..\infrastructure E:\software\go\src\github.com\adminwjp\infrastructure-go
rem ren infrastructure-go infrastructure-go
rem ren oldDir newDir
rem ren oldFile newFile
rem copy file newFile
rem copyx dir new_dir
rem del dir