
# 安装

## Linux Alpine

- 安装zeromq

```bash
apk add --no-cache zeromq
go get github.com/xtech-cloud/omo-mod-net
```

## Windows Msys2

- 编译zeromq源码

在https://github.com/zeromq/libzmq/releases中下载4.2.5

```bash
tar -zxf zeromq-4.2.5.tar.gz
cd zeromq-4.2.5
configure --prefix=/usr/local
make
make install
```

- 编译zeromq的go语言库

```bash
go get github.com/pebbe/zmq4
CGO_CFLAGS="-I/usr/local/include -L/usr/local/lib" CGO_LDFLAGS="-L/usr/local/lib" go install github.com/pebbe/zmq4
go get github.com/xtech-cloud/omo-mod-net
```

# 启动测试服务

```bash
go test -v github.com/xtech-cloud/omo-mod-net
```

