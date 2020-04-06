# HOW TO RUN

From sources:

```
go run main.go
```

From github:

```
# make sure $GOPATH/bin is in the path
go get -u github.com/Bhaal22/ricardo_user_creation
ricardo_user_creation
```


```
curl -XPOST localhost:8080/user -H "Content-Type: application/json" --data '{"first_name":"john", "email": "john@plop.io","password":"xyz"}'
```