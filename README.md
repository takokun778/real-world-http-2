# real-world-http-2

[Real World HTTP 第2版](https://www.oreilly.co.jp/books/9784873119038/)

# 1章

```sh
curl --http1.0 http://localhost:18888/greeting
```

```sh
curl -v --http1.0 http://localhost:18888/greeting
```

# 2章

```sh
curl -v --http1.0 --digest -u user:pass http://localhost:18888/digest
```

# 3章
```sh
go run ./server/main.go
```

```sh
go run ./client/main.go
```
