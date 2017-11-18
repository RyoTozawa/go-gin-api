# go-gin-api
[gin-gonic](https://github.com/gin-gonic/gin)を使ったgoによるAPIのテスト実装

# Dockerfileを使う

```
$ docker pull ytakky2014/go-gin-api:v1
$ docker run -d -p 8080:8080 ytakky2014/go-gin-api:v1
```


リクエスト

```
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```

レスポンス

```
{"message":"Hello World!!"}
```
