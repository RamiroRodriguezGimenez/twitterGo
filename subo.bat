git add .
git commit -m "utlimo commit"
git push
 go build -o bootstrap main.go
del lambda-handler.zip
tar.exe -a -cf lambda-handler.zip bootstrap