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
curl -XPATCH localhost:8080/user/1 -H "Content-Type: application/json" --data '{"first_name":"john2"}'
curl "localhost:8080/user?first_name=j&email=plop"
```

# Improvements:
  * support search for multiple occurences of each attributes
  * add details to events (clientIP, ...)
  * how to spoof ip address to bypass check of Swiss
