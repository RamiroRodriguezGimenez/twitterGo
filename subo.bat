git add .
git commit -m "utlimo commit"
git push
go build main.go
del main.zip
tar.exe -a -cf main.zip main.exe