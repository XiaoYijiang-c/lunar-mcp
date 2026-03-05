# lunar-mcp 测试方案

## 1. 测试目标

- 确保 28 个工具函数正常工作
- 验证 MCP 协议实现正确
- 保证边界情况和错误处理

---

## 2. 测试类型

### 2.1 单元测试 (Unit Tests)

| 模块 | 测试内容 | 覆盖工具 |
|------|---------|---------|
| protocol | JSON-RPC 解析、错误处理 | 全部 |
| tools/registry | 工具注册、查找 | 全部 |
| tools/lunar | 农历计算 | lunar_date |
| tools/bazi | 八字计算 | zodiac_bazi, advanced_bazi |
| tools/fortune | 命理分析 | fortune_periods |

### 2.2 集成测试 (Integration Tests)

| 测试项 | 说明 |
|--------|------|
| MCP 协议流程 | initialize → tools/list → tools/call |
| HTTP 端点 | /health, /rpc |
| 并发测试 | 多请求同时调用 |
| 错误处理 | 无效参数、非法日期 |

### 2.3 端到端测试 (E2E Tests)

| 测试场景 | 预期结果 |
|---------|---------|
| Docker 容器启动 | 服务正常运行 |
| API 响应格式 | 符合 JSON-RPC 2.0 |
| 28 个工具全部调用 | 正常返回 |

---

## 3. 测试用例

### 3.1 协议测试

```go
// protocol_test.go
func TestInitialize(t *testing.T) {
    // 测试 initialize 方法
    req := `{"jsonrpc":"2.0","method":"initialize","params":{},"id":1}`
    resp := callRPC(req)
    
    assert.Equal(t, "2.0", resp.JsonRPC)
    assert.NotNil(t, resp.Result["capabilities"])
    assert.Equal(t, "lunar-mcp", resp.Result["serverInfo"].(map[string]string)["name"])
}

func TestToolsList(t *testing.T) {
    // 测试 tools/list 返回 28 个工具
    resp := callRPC(`{"method":"tools/list"}`)
    tools := resp.Result["tools"].([]Tool)
    
    assert.Equal(t, 28, len(tools))
    assert.Contains(t, toolNames, "lunar_date")
    assert.Contains(t, toolNames, "zodiac_bazi")
}

func TestToolsCall(t *testing.T) {
    // 测试 tools/call
    req := `{"method":"tools/call","params":{"name":"lunar_date","arguments":{"year":2026,"month":3,"day":5}}}`
    resp := callRPC(req)
    
    assert.Contains(t, resp.Result["lunar"], "2026")
}

func TestErrorHandling(t *testing.T) {
    // 测试错误码
    resp := callRPC(`{"method":"invalid_method"}`)
    assert.Equal(t, -32601, resp.Error.Code) // Method not found
    
    resp = callRPC(`{"method":"tools/call","params":{"name":"not_exist"}}`)
    assert.Equal(t, -32601, resp.Error.Code)
}
```

### 3.2 工具函数测试

```go
// tools/lunar_test.go
func TestLunarDate(t *testing.T) {
    tests := []struct {
        name     string
        year     int
        month    int
        day      int
        wantLuna string
    }{
        {"2026-03-05", 2026, 3, 5, "二〇二六年正月十七"},
        {"2024-02-10", 2024, 2, 10, "二〇二四年正月初一"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resp := callRPC(buildReq("lunar_date", map[string]int{
                "year": tt.year, "month": tt.month, "day": tt.day,
            }))
            assert.Contains(t, resp.Result["lunar"], tt.wantLuna)
        })
    }
}

// tools/bazi_test.go
func TestZodiacBazi(t *testing.T) {
    resp := callRPC(buildReq("zodiac_bazi", map[string]int{
        "year": 1990, "month": 5, "day": 15,
    }))
    
    bazi := resp.Result["bazi"].(map[string]string)
    assert.Equal(t, "庚午", bazi["year"])
    assert.Equal(t, "辛巳", bazi["month"])
}

func TestAdvancedBazi(t *testing.T) {
    resp := callRPC(buildReq("advanced_bazi", map[string]int{
        "year": 1990, "month": 5, "day": 15,
    }))
    
    // 验证十神
    bazi := resp.Result["bazi"].([]map[string]interface{})
    assert.Contains(t, bazi[0]["shishenGan"], "比肩")
}
```

### 3.3 边界测试

```go
// boundary_test.go
func TestBoundaryDates(t *testing.T) {
    tests := []struct {
        name  string
        year  int
        month int
        day   int
    }{
        {"闰年", 2024, 2, 29},
        {"春节", 2026, 1, 29},
        {"除夕", 2025, 1, 28},
        {"2000年", 2000, 1, 1},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resp := callRPC(buildReq("lunar_date", map[string]int{
                "year": tt.year, "month": tt.month, "day": tt.day,
            }))
            assert.NotEmpty(t, resp.Result["lunar"])
        })
    }
}

func TestInvalidDates(t *testing.T) {
    tests := []struct {
        year  int
        month int
        day   int
    }{
        {2026, 13, 1},   // 无效月份
        {2026, 2, 30},   // 无效日期
        {2026, 6, 31},   // 6月无31天
    }
    
    for _, tt := range tests {
        resp := callRPC(buildReq("lunar_date", map[string]int{
            "year": tt.year, "month": tt.month, "day": tt.day,
        }))
        assert.NotNil(t, resp.Error) // 应返回错误
    }
}
```

### 3.4 并发测试

```go
// concurrency_test.go
func TestConcurrentRequests(t *testing.T) {
    var wg sync.WaitGroup
    errors := make(chan error, 100)
    
    // 100 个并发请求
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            resp := callRPC(buildReq("lunar_date", map[string]int{
                "year": 2026, "month": 3, "day": 5,
            }))
            if resp.Error != nil {
                errors <- resp.Error
            }
        }()
    }
    
    wg.Wait()
    close(errors)
    
    // 不应有错误
    for err := range errors {
        t.Errorf("concurrent request failed: %v", err)
    }
}
```

---

## 4. 测试数据

### 4.1 标准测试用例

| 日期 | 农历 | 八字 | 备注 |
|------|------|------|------|
| 2026-03-05 | 正月十七 | 丙午年辛卯月戊寅日 | 惊蛰 |
| 1990-05-15 | 四月廿一 | 庚午年辛巳月庚辰日 | - |
| 2000-01-01 | 腊月廿五 | 己卯年丙子月辛酉日 | 千禧年 |

### 4.2 边界测试数据

| 测试类型 | 数据 |
|---------|------|
| 闰年 | 2024-02-29 |
| 春节 | 2026-01-29 |
| 除夕 | 2025-01-28 |
| 最早日期 | 1900-01-01 |
| 最晚日期 | 2100-12-31 |

---

## 5. 执行计划

### Phase 1: 协议测试
- [ ] HTTP 端点测试
- [ ] JSON-RPC 解析
- [ ] 错误码验证

### Phase 2: 工具测试
- [ ] 28 个工具逐一测试
- [ ] 边界条件覆盖
- [ ] 错误处理

### Phase 3: 集成测试
- [ ] MCP 完整流程
- [ ] 并发测试
- [ ] Docker 容器测试

### Phase 4: 性能测试
- [ ] 响应时间
- [ ] 并发能力

---

## 6. 测试覆盖率目标

| 指标 | 目标 |
|------|------|
| 单元测试覆盖 | > 70% |
| 边界测试覆盖 | 100% |
| 工具函数测试 | 28/28 (100%) |

---

## 7. 测试工具

- Go testing
- testify/assert
- httptest
- docker-test