GOOS=windows go build -o ./bin/windows/dotnet-tpl.exe &&

GOOS=linux go build -o ./bin/linux/dotnet-tpl &&

GOOS=darwin go build -o ./bin/macos/dotnet-tpl