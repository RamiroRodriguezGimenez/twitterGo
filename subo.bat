git add .
git commit -m "utlimo commit"
git push
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bootstrap main.go
del lambda-handler.zip
tar.exe -a -cf lambda-handler.zip bootstrap