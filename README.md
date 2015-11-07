# UrlShortener

Url shortener RESTfull API developed in Go


## Installation

To run this program, you need to setup Go (https://golang.org/doc/install) and Redis (http://redis.io/topics/quickstart) in a Linux environment.

To get the code:

    go get github.com/Tip-Sy/UrlShortener

To resolve imports in the code *(should be automated with https://github.com/tools/godep)*:

    go get gopkg.in/redis.v3
    go get github.com/gorilla/mux

Installation:

    go install github.com/Tip-Sy/UrlShortener

Run:
```shell
cd $GOBIN
./UrlShortener
```


## Configuration

The config file ([config.go](config.go)) contains default Port binding, IP binding and Datastore configuration:
- Server IP: localhost
- API Port: 8080
- Db Port: 6379
- ...

*(Should be replaced by a yaml file)*


## Test case

1. Prerequisite: 'php-cli' package installed
2. Run Redis server
3. Run the Go code
4. Run the file test/[test.php](test/test.php) with the command:
```
    php test.php
```

.

***Tip-Sy***
