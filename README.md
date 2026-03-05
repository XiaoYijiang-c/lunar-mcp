# lunar-mcp

MCP Server for Chinese Calendar

## 简介

基于 lunar-go 实现的 MCP 服务器，提供中国农历、八字、风水等传统日历能力给 AI Agent 使用。

## 功能

| 分类 | 工具 |
|------|------|
| 基础 | lunar_date, zodiac_bazi, solar_terms, festivals |
| 日历 | solar_calendar, month_calendar, year_calendar |
| 八字 | eight_char_full, advanced_bazi, destiny_analysis, fortune_periods |
| 命理 | date_selector, marriage_compat, name_generator |
| 高级 | iching_divination, nine_star_flying |
| 神煞 | shen_sha, pengzu_baiji |

## 快速开始

### Docker

```bash
docker run -d -p 8080:8080 ghcr.io/xiaoyijiang-c/lunar-mcp:latest
```

### 本地运行

```bash
go build -o lunar-mcp .
./lunar-mcp
```

## API

```bash
# 健康检查
curl http://localhost:8080/health

# 查询农历
curl -X POST http://localhost:8080/rpc \
  -d '{"method":"tools/call","params":{"name":"lunar_date","arguments":{"year":2026,"month":3,"day":5}}}'
```

## 技术栈

- Go
- MCP Protocol (JSON-RPC 2.0)
- Docker

## License

MIT
