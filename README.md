# xii-go-http
一个轻量级、易用、功能丰富的Go HTTP客户端库，提供了优雅的链式API和中间件支持。

## 特性

- **链式API设计**：提供流畅的链式调用体验
- **多格式支持**：内置支持JSON、XML、YAML、表单等多种数据格式
- **中间件机制**：可扩展的中间件系统，方便添加自定义逻辑
- **灵活配置**：支持基础URL、超时、代理、默认头信息等配置
- **响应处理**：提供便捷的响应状态检查和内容解析方法

## 安装

```bash
go get -u github.com/gaussli/xii-go-http
```

## 快速开始

### 基本使用

```go
package main

import (
  "fmt"
  "github.com/gaussli/xii-go-http/pkg/http"
  "time"
)

func main() {
  // 创建客户端
  client := http.NewClient(
    http.WithBaseURL("https://api.example.com"),
    http.WithTimeout(30*time.Second),
  )

  // 发送请求
  req := http.NewXiiRequest().
    GET("/users").
    QueryParam("page", "1").
    QueryParam("limit", "10").
    Header("Authorization", "Bearer token123")

  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }

  // 检查响应
  if resp.IsSuccess() {
    fmt.Println("请求成功!")
    fmt.Println("响应内容:", resp.TextBody())
  } else {
    fmt.Printf("请求失败: %d", resp.StatusCode)
  }
}
```

### JSON请求和响应

```go
// JSON请求
type User struct {
  Name  string `json:"name"`
  Email string `json:"email"`
}

user := User{Name: "张三", Email: "zhangsan@example.com"}

req := http.NewXiiRequest().POST("/users").JSONBody(user)

resp, err := client.Do(req)

// 解析JSON响应
var result struct {
  ID   int    `json:"id"`
  Name string `json:"name"`
}
if err := resp.JSONBody(&result); err != nil {
  // 处理错误
}
```

### 使用中间件

```go
// 创建一个添加认证信息的中间件
authMiddleware := func(req *http.Request) error {
  req.Header.Set("X-API-Key", "api-key-123")
  return nil
}

// 使用中间件
client.Use(authMiddleware)
```

## API文档

### XiiClient

#### 创建客户端

```go
client := http.NewClient(opts ...Option)
```

#### 配置选项

- `WithBaseURL(baseURL string)`: 设置基础URL
- `WithTimeout(timeout time.Duration)`: 设置请求超时
- `WithProxy(proxyStr string)`: 设置代理服务器
- `WithHeader(key, value string)`: 添加默认请求头

#### 方法

- `Use(mw XiiMiddleware)`: 添加中间件
- `Do(req *XiiRequest) (*XiiResponse, error)`: 发送请求

### XiiRequest

#### 创建请求

```go
req := http.NewXiiRequest()
```

#### HTTP方法

- `Method(method string)`: 设置HTTP方法
- `GET(endpoint string)`: 设置为GET请求
- `POST(endpoint string)`: 设置为POST请求
- `PUT(endpoint string)`: 设置为PUT请求
- `DELETE(endpoint string)`: 设置为DELETE请求
- `PATCH(endpoint string)`: 设置为PATCH请求

#### 请求配置

- `Endpoint(endpoint string)`: 设置请求路径
- `Header(key, value string)`: 添加请求头
- `QueryParam(key, value string)`: 添加查询参数
- `Context(ctx context.Context)`: 设置请求上下文

#### 请求体

- `Body(body io.Reader)`: 设置原始请求体
- `TextBody(text string)`: 设置文本请求体
- `JSONBody(jsonData any)`: 设置JSON请求体
- `XMLBody(xmlData any)`: 设置XML请求体
- `YAMLBody(yamlData any)`: 设置YAML请求体
- `FormBody(formData url.Values)`: 设置表单请求体
- `MultipartFormBody(formData url.Values)`: 设置multipart表单请求体

### XiiResponse

#### 状态检查

- `IsSuccess() bool`: 检查是否成功响应(2xx)
- `IsRedirect() bool`: 检查是否重定向响应(3xx)
- `IsClientError() bool`: 检查是否客户端错误(4xx)
- `IsServerError() bool`: 检查是否服务器错误(5xx)
- `IsError() bool`: 检查是否错误响应(4xx或5xx)

#### 内容解析

- `TextBody() string`: 获取文本响应体
- `JSONBody(v any) error`: 解析JSON响应体
- `XMLBody(v any) error`: 解析XML响应体
- `YAMLBody(v any) error`: 解析YAML响应体

## 高级特性

### 中间件链

您可以添加多个中间件，它们将按照添加顺序依次执行：

```go
client.Use(loggingMiddleware)
client.Use(authMiddleware)
client.Use(metricsMiddleware)
```

### 自定义中间件示例

#### 日志中间件

```go
loggingMiddleware := func(req *http.Request) error {
  fmt.Printf("[INFO] Request: %s %s", req.Method, req.URL.String())
  return nil
}
```

#### 重试中间件

```go
retryMiddleware := func(req *http.Request) error {
  // 这里可以实现重试逻辑
  return nil
}
```

## 许可证

MIT License

## 贡献

欢迎提交Issue和Pull Request！