# lunar-mcp

MCP Server for Chinese Calendar (农历日历 MCP 服务)

## 简介

基于 [lunar-go](https://github.com/6tail/lunar-go) 实现的 Model Context Protocol (MCP) 服务器，提供中国农历、黄历、八字等传统日历能力给 AI Agent 使用。

## 功能

| 工具 | 说明 |
|------|------|
| `lunar_date` | 获取指定公历日期的农历信息 |
| `zodiac_bazi` | 获取八字信息 |
| `solar_terms` | 获取节气信息 |
| `festivals` | 获取节日信息 |
| `auspicious_date` | 查询吉日 |
| `daily_omen` | 获取每日宜忌 |

## 快速开始

### 本地运行

```bash
# 编译
go build -o lunar-mcp .

# 运行
./lunar-mcp

# 默认端口 8080
```

### Docker 运行

```bash
docker run -d -p 8080:8080 lunar-mcp
```

## API

### MCP 方法

- `initialize` - 握手初始化
- `tools/list` - 获取工具列表
- `tools/call` - 调用工具
- `ping` - 健康检查

### 示例

```bash
# 健康检查
curl http://localhost:8080/health

# 初始化
curl -X POST http://localhost:8080/rpc \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"initialize","params":{},"id":1}'

# 获取工具列表
curl -X POST http://localhost:8080/rpc \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/list","params":{},"id":2}'

# 调用工具 - 查询农历
curl -X POST http://localhost:8080/rpc \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"tools/call","params":{"name":"lunar_date","arguments":{"year":2026,"month":3,"day":5}},"id":3}'
```

## 技术栈

- Go 1.21+
- MCP Protocol (JSON-RPC 2.0)
- lunar-go

## 协议

MIT License - see [LICENSE](LICENSE)
