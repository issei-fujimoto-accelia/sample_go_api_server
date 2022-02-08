# readme

## start
```
docker-compose up -d
```


## server

```
$ curl localhost:8003/ok
{"status":200,"message":"ok"}
```

```
$ curl localhost:8003/server-error
{"status":500,"message":"server error"}
```

```
$ curl localhost:8003/not-found
{"status":404,"message":"not found"}
```

