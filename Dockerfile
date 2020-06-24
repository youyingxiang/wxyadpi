FROM golang:1.14 AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN  go mod download

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o wxyapi .

# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist

# 将二进制文件从 /build 目录复制到这里
RUN cp /build/wxyapi .
###################
# 接下来创建一个小镜像
###################
FROM scratch

#COPY ./templates /templates
#COPY ./static /static
COPY ./conf /conf
COPY .env .

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/wxyapi /
#RUN docker run -itd --name wxyapi   -p 9000:9000  --link mysql_local:mysql-dev --link redis-test:redis-dev  wxyapi /bin/bash
# 需要运行的命令
ENTRYPOINT ["/wxyapi"]