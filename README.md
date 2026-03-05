# lunar-mcp

MCP Server for Chinese Calendar (农历日历 MCP 服务)

## 简介

基于 [lunar-go](https://github.com/6tail/lunar-go) 实现的 Model Context Protocol (MCP) 服务器，提供中国农历、黄历、八字等传统日历能力给 AI Agent 使用。

**28 工具函数**，支持日期计算、八字分析、命理、风水等功能。

## 功能概览

| 分类 | 工具 |
|------|------|
| 基础 | lunar_date, zodiac_bazi, solar_terms, festivals |
| 日历 | solar_calendar, month_calendar, year_calendar |
| 八字 | eight_char_full, advanced_bazi, destiny_analysis, fortune_periods |
| 命理 | date_selector, marriage_compat, name_generator |
| 高级 | iching_divination, nine_star_flying |
| 神煞 | shen_sha, pengzu_baiji |

## 快速开始

### Docker (推荐)

```bash
# 直接运行
docker run -d -p 8080:8080 ghcr.io/xiaoyijiang-c/lunar-mcp:latest

# docker-compose
docker-compose up -d
```

### 本地运行

```bash
# 编译
go build -o lunar-mcp .

# 运行
./lunar-mcp

# 默认端口 8080
```

## API 示例

```bash
# 健康检查
curl http://localhost:8080/health

# 查询农历
curl -X POST http://localhost:8080/rpc \
  -d '{"method":"tools/call","params":{"name":"lunar_date","arguments":{"year":2026,"month":3,"day":5}}}'

# 八字分析
curl -X POST http://localhost:8080/rpc \
  -d '{"method":"tools/call","params":{"name":"zodiac_bazi","arguments":{"year":1990,"month":5,"day":15}}}'

# 易经占卜
curl -X POST http://localhost:8080/rpc \
  -d '{"method":"tools/call","params":{"name":"iching_divination","arguments":{"question":"今日运势"}}}'
```

## 技术栈

- Go 1.21+
- MCP Protocol (JSON-RPC 2.0)
- lunar-go
- Docker

## 相关链接

- GitHub: https://github.com/XiaoYijiang-c/lunar-mcp
- awesome-mcp-servers: https://github.com/punkpeye/awesome-mcp-servers (PR #2747)
- Docker Hub: ghcr.io/xiaoyijiang-c/lunar-mcp

## 协议

MIT License - see [LICENSE](LICENSE)
