# http2
research of http2 vulnarabilities

# run server
```
go run ./server
```

# generate certs
```
./generate.sh
```

# test connections
```
curl -k -v --http2 https://localhost:7002/hello/sayHello -d "Hello Go"
curl -v --http2 --cacert ./cert/localhost.crt https://localhost:7002/hello/sayHello -d "Hello Go!"
```