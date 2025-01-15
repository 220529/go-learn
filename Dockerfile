# 使用官方 Go 镜像作为基础镜像
FROM golang:1.23.4-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制项目代码到容器中
COPY . .

# 自动下载和清理依赖，更侧重于清理和修复文件
RUN go mod tidy

# 编译 Go 项目
RUN go build -o /go-app

# 暴露容器端口
EXPOSE 8080

FROM scratch

COPY --from=builder /go-app /

# 设置容器启动时运行的命令
CMD ["/go-app"]