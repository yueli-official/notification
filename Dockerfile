# --------------------------------------------------
# 1️⃣ 构建阶段：使用官方 Go 镜像编译二进制文件
# --------------------------------------------------
FROM golang:1.25-alpine AS builder

# 设置工作目录
WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git

# 复制 go.mod 和 go.sum 并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 构建二进制文件（关闭 CGO, 减小体积）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notification ./cmd/main.go

# --------------------------------------------------
# 2️⃣ 运行阶段
# --------------------------------------------------
FROM alpine:3.20

RUN adduser -D appuser
WORKDIR /app

COPY --from=builder /app/notification .

USER appuser

EXPOSE 9000

ENTRYPOINT ["./notification"]
