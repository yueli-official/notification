# Notification Service

统一通知服务，支持邮件和短信发送，基于 Go + Gin 构建。

## 功能

- 单封邮件发送，支持 HTML / 纯文本
- 单条短信发送，支持腾讯云、阿里云
- 批量邮件 + 短信混合发送
- 多提供商配置，按 `name` 指定或自动选用第一个启用的
- API Key 认证（请求头或 Cookie）

## 快速开始

**1. 准备配置文件**

```bash
cp config.example.toml config.toml
```

编辑 `config.toml`，填入真实的 SMTP 账号、短信密钥和 API Key。

**2. 启动服务**

```bash
go run cmd/main.go
```

或使用 Docker：

```bash
docker compose up -d
```

服务默认监听 `:9000`，Swagger 文档地址：`http://localhost:9000/swagger/index.html`

## 认证

所有接口需在**请求头**或 **Cookie** 中携带 API Key：

| 方式 | 示例 |
|------|------|
| 请求头 | `X-API-Key: your-secret-api-key-here` |
| Cookie | `api_key=your-secret-api-key-here` |

缺失或错误时返回 `401 Unauthorized`。

## 接口

### 发送邮件

```
POST /api/v1/notify/email
```

```json
{
  "provider": "account",
  "to": ["user@example.com"],
  "subject": "欢迎使用",
  "body": "<h1>你好</h1>",
  "is_html": true
}
```

`provider` 可省略，省略时使用配置中第一个 `enabled = true` 的邮件账号。

### 发送短信

```
POST /api/v1/notify/sms
```

```json
{
  "provider": "tencent",
  "phone_number": "13800138000",
  "template_id": "SMS_123456",
  "params": {
    "code": "1234"
  }
}
```

### 批量发送

```
POST /api/v1/notify/batch
```

```json
{
  "emails": [
    {
      "to": ["a@example.com"],
      "subject": "通知",
      "body": "内容"
    }
  ],
  "sms": [
    {
      "phone_number": "13800138000",
      "template_id": "SMS_123456"
    }
  ]
}
```

响应：
- `200` — 全部成功
- `206` — 部分失败，`errors` 字段包含失败详情

## 配置说明

| 字段 | 说明 |
|------|------|
| `server.port` | 监听端口，默认 `9000` |
| `server.api_key` | 接口认证密钥 |
| `email[].name` | 提供商标识，发送时通过 `provider` 字段指定 |
| `email[].enabled` | 是否启用，`false` 时跳过 |
| `sms[].provider` | 短信厂商类型：`tencent` / `aliyun` |

> 支持通过环境变量覆盖配置，变量名规则：`SERVER_API_KEY`、`EMAIL_0_PASSWORD` 等（Viper AutomaticEnv）。

## 项目结构

```
.
├── cmd/main.go          # 入口，路由注册
├── config/config.go     # 配置结构与加载
├── handler/             # HTTP 处理器
├── middleware/auth.go   # API Key 认证中间件
├── model/               # 请求 / 响应结构体
├── service/             # 业务逻辑
├── docs/                # Swagger 文档（自动生成）
├── config.toml          # 本地配置（不提交）
└── config.example.toml  # 配置模板
```
