Go HTTP Router Benchmark For Water
==================================

forked from [go-http-routing-benchmark](github.com/julienschmidt/go-http-routing-benchmark)

This benchmark suite aims to compare the performance of HTTP request routers for [Go](https://golang.org) by implementing the routing structure of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.

Of course the tested routers can be used for any kind of HTTP request → handler function routing, not only (REST) APIs.


#### Tested routers & frameworks:

 * [water](https://github.com/meilihao/water)

## Usage

If you'd like to run these benchmarks locally, you'll need to install the packge first:

```bash
go get github.com/meilihao/go-http-routing-benchmark-water
```
This may take a while due to the large number of dependencies that need to be downloaded. Once that command completes, you can run the full set of benchmarks like this:

```bash
cd $GOPATH/src/github.com/meilihao/go-http-routing-benchmark-water
go test -bench=.
```

> **Note:** If you run the tests and it SIGQUIT's make the go test timeout longer (#44)
>
>```
go test -timeout=2h -bench=.
```
