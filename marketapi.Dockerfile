# 第一部分：构建阶段
# 使用golang:alpine作为基础镜像，用于构建Go应用
FROM golang:alpine AS builder

MAINTAINER runningriven@gmail.com

# 给当前构建阶段添加一个标签，便于在多阶段构建中识别
LABEL stage=gobuilder

# 禁用C语言绑定，有助于生成更轻量的静态链接二进制文件
ENV CGO_ENABLED 0
# 设置Go语言的代理，加速Go模块的下载
ENV GOPROXY https://goproxy.cn,direct
# 替换Alpine Linux的默认镜像源为阿里云镜像，加快软件包的下载速度
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 更新apk包为tzdata，该包包含时区数据
RUN apk update --no-cache && apk add --no-cache tzdata
# 设置工作目录，所有后续的命令都在这个目录下执行
WORKDIR /usr/src/stock

# 将当前目录下的所有文件(除了.dockerignore中排除的除外)拷贝到WORKDIR指定的目录下，即：/usr/src/stock目录下面
COPY . .
# 下载程序构建所需依赖库
RUN go mod tidy
# 构建Go应用程序，使用-ldflags="-s -w"标志以剥离符号信息，生成更小的二进制文件，构建结果输出到/app/stock/marketapi
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/stock/marketapi cmd/market/api/marketapi.go

# 第二部分：运行阶段
# 使用scratch作为基础镜像
FROM scratch
# 从builder阶段的镜像中复制CA证书文件和时区信息
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

# 从builder阶段复制构建好的应用程序和配置文件到镜像中
# 在构建阶段生成的可执行文件全路径为/app/stock/marketapi，这里拷贝到目的路径即可
COPY --from=builder /app/stock/marketapi /app/stock/marketapi
# 构建阶段的配置文件的全路径为/usr/src/stock/cmd/market/api/etc/marketapi.yaml，这里也拷贝到目的路径即可
COPY --from=builder /usr/src/stock/cmd/market/api/etc/marketapi.yaml /etc/stock/marketapi.yaml

# 设置容器启动时的默认命令，执行./marketapi命令，并传递-f /app/stock/etc/marketapi.yaml作为参数
CMD ["/app/stock/marketapi", "-f", "/etc/stock/marketapi.yaml"]
