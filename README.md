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

```
https://hub.docker.com/repository/docker/drylm/ricardo
docker run --rm -it drylm/ricardo:v0.1
```

# RMQ

```
docker run --rm --hostname rabbit1 --name rabbit1 rabbitmq:3-management

Admin console: http://container-ip:15672
```

# Improvements:
  * support search for multiple occurences of each attributes
  * add details to events (clientIP, ...)
  * how to spoof ip address to bypass check of Swiss
  * docker files + docker compose
  * inject RMQ endpoint via environment variables
  * usage of pgsql instead of sqlite3
