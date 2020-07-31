# Intro

This is a test project. Run different protocols on multiple ports to test the protocol forwarding in istio.


| Port | Protocols                                                                          |
|------|------------------------------------------------------------------------------------|
| 8001 | Plain http server. Supports http/1.0, http/1.1                                     |
| 8002 | Non-TLS http2 server, aka h2c. Supports upgrade from http/1.1 and prior knowledge. |
| 8003 | TLS-Enabled http2 server.                                                          |
| 8004 | Non-TLS Grpc server.                                                               |
| 8005 | TLS-Enabled Grpc server.                                                           |

# Test 
Before run tests. Upgrade your `curl` to the lastest version.

### Test 8001

```bash
# http 1.0
curl --http1.0 http://localhost:8001/path\?query\=1 -v
# => http 1.0 response

# http 1.1
curl --http1.1 http://localhost:8001/path\?query\=1 -v
# => http 1.1 response 

# http 2
curl --http2 http://localhost:8001/path\?query\=1 -v
# => http 1.1 response
```

### Test 8002

```bash
# http 1.0
curl --http1.0 http://localhost:8002/path\?query\=1 -v
# => http 1.0 response

# http 1.1
curl --http1.1 http://localhost:8002/path\?query\=1 -v
# => http 1.1 response 

# http 2~
curl --http2 http://localhost:8002/path\?query\=1 -v
# => switching protocol, connection upgrade
# => http 2 response

# http 2
curl --http2-prior-knowledge http://localhost:8002/path\?query\=1 -v
# => use h2 directly
# => http 2 response
```

### Test 8003

```bash
# http 1.1
curl --http1.1 https://localhost:8003/path\?query\=1 -v -k
# => http 1.1 response 

# http 2
curl --http2 https://localhost:8003/path\?query\=1 -v -k
# => ALPN h2
# => http 2 response
```

### Test 8004

Please make sure you have  [grpcurl](https://github.com/fullstorydev/grpcurl) installed.

```bash
grpcurl -plaintext localhost:8004 main.HelloWorld/Greeting
```

### Test 8005

Please make sure you have [grpcurl](https://github.com/fullstorydev/grpcurl) installed.

```bash
grpcurl -insecure localhost:8005 main.HelloWorld/Greeting
```