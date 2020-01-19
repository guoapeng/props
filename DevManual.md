The file mock_propsReader.go is generated file by following the approach below 
```bash
cd <project_home>
go get github.com/golang/mock/mockgen
cd $GOPATH/src/github.com/golang/mock/mockgen
go build
cd <project_home>
echo $PATH
cp mockgen.exe <User.Home>\go\bin\   #e.g. User.Home = C:\Users\eagle\
mockgen -destination mocks/mock_propsReader.go -package mocks -source PropsReader.go

```




