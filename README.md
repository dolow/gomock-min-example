# gomock showcase

install gomock

```
go get github.com/golang/mock/mockgen@v1.4.4
```

generate mock

```
mockgen -source=dao/user.go -destination=mock/dao/user.go
```

run test

```
go test service/*
```
