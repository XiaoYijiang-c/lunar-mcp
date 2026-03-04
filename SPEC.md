# lunar-mcp 功能规范

> 当前版本: v1.0.0  
> 最后更新: 2026-03-05

## 1. 项目概述

**lunar-mcp** 是一个基于 [lunar-go](https://github.com/6tail/lunar-go) 实现的 MCP (Model Context Protocol) 服务器，为 AI Agent 提供中国农历、黄历、八字等传统日历能力。

**lunar-go 库共有 233+ 个函数**，本项目目前实现了部分核心功能。

---

## 2. 功能覆盖矩阵

### 2.1 已实现功能 ✅

| 工具名 | 功能描述 | 状态 | 优先级 |
|--------|---------|------|--------|
| `lunar_date` | 农历日期查询 | ✅ 已完成 | P0 |
| `zodiac_bazi` | 八字计算 | ✅ 已完成 | P0 |
| `solar_terms` | 节气查询 | ✅ 已完成 | P0 |
| `festivals` | 节日查询 | ✅ 已完成 | P0 |
| `auspicious_date` | 吉日查询 | ✅ 已完成 | P0 |
| `daily_omen` | 每日宜忌 | ✅ 已完成 | P0 |

### 2.2 待实现功能 📋

#### P1 - 高优先级

| 工具名 | 功能描述 | 说明 |
|--------|---------|------|
| `solar_calendar` | 公历信息 | 星座、星期、儒略日 |
| `lunar_detail` | 农历详细信息 | 年月日、闰月、法定节假日 |
| `month_calendar` | 月历 | 指定月份的完整日历 |
| `year_calendar` | 年历 | 指定年份的日历 |

#### P2 - 中优先级

| 工具名 | 功能描述 | 说明 |
|--------|---------|------|
| `eight_char_full` | 完整八字 | 包含五行、十神、纳音 |
| `destiny_analysis` | 命理分析 | 五行强弱、用神分析 |
| `daily_fortune` | 每日运势 | 喜神、福神、财神方位 |
| `time_bazi` | 时辰八字 | 指定时辰的八字 |

#### P3 - 低优先级

| 工具名 | 功能描述 | 说明 |
|--------|---------|------|
| `solar_terms_detail` | 节气详情 | 交接时间、物候 |
| `tao_holiday` | 道历 | 道教节日 |
| `buddhist_holiday` | 佛历 | 佛教节日 |
| `lunar_calendar` | 农历历书 | 完整的黄历信息 |
| `yi_jing` | 易经占卜 | 六十四卦 |

---

## 3. 详细功能说明

### 3.1 已实现功能详情

#### lunar_date - 农历日期查询

```json
{
  "name": "lunar_date",
  "description": "获取指定公历日期的农历信息",
  "params": {
    "year": "公历年份",
    "month": "公历月份", 
    "day": "公历日期"
  }
}
```

**返回值**:
- `lunar`: 完整农历日期字符串
- `solar`: 公历日期字符串
- `weekday`: 星期几 (0-6)
- `animal`: 生肖

---

#### zodiac_bazi - 八字计算

```json
{
  "name": "zodiac_bazi",
  "description": "获取八字信息",
  "params": {
    "year": "公历年份",
    "month": "公历月份",
    "day": "公历日期",
    "hour": "时辰 (0-23，可选)"
  }
}
```

**返回值**:
- `bazi`: 四柱八字 [年, 月, 日, 时]
- `nayin`: 纳音五行
- `animal`: 生肖
- `fullBazi`: 完整八字字符串

---

#### solar_terms - 节气查询

```json
{
  "name": "solar_terms",
  "description": "获取指定日期的节气信息",
  "params": {
    "year": "公历年份",
    "month": "公历月份",
    "day": "公历日期"
  }
}
```

**返回值**:
- `currentJie`: 当前节气
- `currentJieQi`: 当前节或气
- `jieQiList`: 节气列表

---

#### festivals - 节日查询

```json
{
  "name": "festivals",
  "description": "获取指定日期的节日信息",
  "params": {
    "year": "公历年份",
    "month": "公历月份",
    "day": "公历日期"
  }
}
```

**返回值**:
- `solarFestivals`: 公历节日
- `lunarFestivals`: 农历节日

---

#### auspicious_date - 吉日查询

```json
{
  "name": "auspicious_date",
  "description": "查询指定月份的吉日",
  "params": {
    "year": "公历年份",
    "month": "公历月份",
    "type": "类型 (嫁娶/搬家/开业/动土)"
  }
}
```

**返回值**:
- `results`: 吉日列表

---

#### daily_omen - 每日宜忌

```json
{
  "name": "daily_omen",
  "description": "获取每日宜忌",
  "params": {
    "year": "公历年份",
    "month": "公历月份",
    "day": "公历日期"
  }
}
```

**返回值**:
- `auspicious`: 宜做的事情
- `inauspicious`: 忌做的事情
- `chong`: 冲煞信息

---

## 4. lunar-go 函数映射

### 4.1 已映射

| MCP 工具 | 使用的 lunar-go 函数 |
|---------|---------------------|
| lunar_date | GetAnimal, ToFullString, GetWeek |
| zodiac_bazi | GetBaZi, GetBaZiNaYin |
| solar_terms | GetCurrentJie, GetCurrentJieQi |
| festivals | GetFestivals |
| auspicious_date | GetDayYi, GetDayJi |
| daily_omen | GetDayYi, GetDayJi, GetDayLu, GetChong |

### 4.2 未映射（可扩展）

**八字相关 (GetBaZi*)**:
- GetBaZiShiShenGan, GetBaZiShiShenZhi
- GetBaZiWuXing
- GetBaZiShiShenYearZhi/MonthZhi/DayZhi/TimeZhi

**五行纳音**:
- GetDayNaYin, GetTimeNaYin, GetYearNaYin, GetMonthNaYin

**神煞方位**:
- GetDayPositionXi, GetDayPositionFu, GetDayPositionCai
- GetDayPositionTaiSui
- GetDayTianShen, GetTimeTianShen

**其他**:
- GetShiChen (时辰)
- GetJieQiTable (节气表)
- GetNext/Prev 方法 (日期推算)

---

## 5. 扩展计划

### 5.1 v1.1.0 - 完善核心功能

- [ ] 添加更多吉日类型（订盟、纳采、入学、搬家等）
- [ ] 支持农历转公历
- [ ] 支持日期范围查询

### 5.2 v1.2.0 - 增强八字

- [ ] 时辰八字（含时辰名称）
- [ ] 十神分析
- [ ] 五行强弱分析

### 5.3 v1.3.0 - 完整历书

- [ ] 月历视图
- [ ] 年历视图
- [ ] 道历、佛历

### 5.4 v2.0.0 - 高级功能

- [ ] 命理分析
- [ ] 合婚查询
- [ ] 易经占卜

---

## 6. 技术规范

### 6.1 MCP 协议

- JSON-RPC 2.0
- 端口: 8080 (可配置)
- 端点: /rpc

### 6.2 依赖

- Go 1.21+
- lunar-go v1.4.6

### 6.3 部署

- Docker 支持
- 无状态服务

---

## 7. 验收标准

- [x] MCP 协议三板斧正常
- [x] 6 个工具函数返回正确结果
- [x] Docker 镜像可构建
- [ ] 单元测试覆盖 > 80%
- [ ] 文档完整度 100%

---

*此文档使用 OpenSpec 规范管理*
