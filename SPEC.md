# lunar-mcp 功能规范 (完整版)

> 当前版本: v1.1.0  
> 最后更新: 2026-03-05

## 项目概述

**lunar-mcp** 是基于 [lunar-go](https://github.com/6tail/lunar-go) 实现的 MCP 服务器。

**lunar-go 库共有 233+ 个函数**，本项目目前实现了 16 个工具，覆盖部分核心功能。

---

## 功能覆盖

### 已实现 ✅ (16个工具)

| # | 工具名 | 功能描述 | 状态 |
|---|--------|---------|------|
| 1 | `lunar_date` | 农历日期查询 | ✅ |
| 2 | `zodiac_bazi` | 八字计算 | ✅ |
| 3 | `solar_terms` | 节气查询 | ✅ |
| 4 | `festivals` | 节日查询 | ✅ |
| 5 | `auspicious_date` | 吉日查询 | ✅ |
| 6 | `daily_omen` | 每日宜忌 | ✅ |
| 7 | `solar_calendar` | 公历详情 | ✅ |
| 8 | `month_calendar` | 月历 | ✅ |
| 9 | `year_calendar` | 年历 | ✅ |
| 10 | `eight_char_full` | 完整八字 | ✅ |
| 11 | `destiny_analysis` | 命理分析 | ✅ |
| 12 | `daily_fortune` | 每日运势 | ✅ |
| 13 | `time_bazi` | 时辰八字 | ✅ |
| 14 | `tao_holiday` | 道历 | ✅ |
| 15 | `buddhist_holiday` | 佛历 | ✅ |
| 16 | `lunar_calendar` | 完整黄历 | ✅ |

### 待实现功能 📋

#### 高优先级

| 功能 | 说明 |
|------|------|
| 日期推算 | Next/Prev 方法，支持日期偏移 |
| 日期范围查询 | 批量查询多天 |
| 农历转公历 | 反向查询 |
| 吉时查询 | 每日吉时 |

#### 中优先级

| 功能 | 说明 |
|------|------|
| 择日 | 根据目的选择吉日 |
| 合婚 | 两人八字合婚 |
| 起名 | 根据八字起名 |
| 开业吉日 | 开业吉日查询 |

#### 低优先级

| 功能 | 说明 |
|------|------|
| 胎神方位 | 每日胎神位置 |
| 易经卦象 | 六十四卦 |
| 九星飞宫 | 玄空飞星 |
| 建除十二值星 | 每日星宿 |

---

## lunar-go 函数映射详情

### 已使用

| 工具 | 使用的函数 |
|------|-----------|
| lunar_date | GetAnimal, ToFullString, GetWeek |
| zodiac_bazi | GetBaZi, GetBaZiNaYin |
| solar_terms | GetCurrentJie, GetCurrentJieQi |
| festivals | GetFestivals |
| auspicious_date | GetDayYi, GetDayJi |
| daily_omen | GetDayYi, GetDayJi, GetDayLu, GetChong |
| solar_calendar | GetXingZuo, GetJulianDay, IsLeapYear |
| month_calendar | GetCurrentJie |
| eight_char_full | GetBaZi, GetBaZiNaYin, GetBaZiWuXing |
| destiny_analysis | GetBaZiWuXing |
| daily_fortune | GetDayPositionXi, GetDayPositionFu, GetDayPositionCai |
| time_bazi | GetTime, GetTimeGan, GetTimeZhi |
| tao_holiday | GetTao |
| buddhist_holiday | GetFoto |
| lunar_calendar | GetXiu, GetWuHou, GetZheng, GetYueXiang |

### 未使用（可扩展）

```go
// 八字增强
GetBaZiShiShenGan, GetBaZiShiShenZhi
GetBaZiShiShenYearZhi, GetBaZiShiShenMonthZhi

// 神煞
GetDayTianShenLuck, GetDayTianShenType
GetTimeTianShen, GetTimeTianShenLuck

// 方位
GetDayPositionTaiSui, GetDayPositionYangGui, GetDayPositionYinGui

// 彭祖百忌
GetPengZu, GetDayLu

// 九星
GetDayNineStar, GetYearNineStar

// 农历月份
GetMonth, GetMonthInGanZhi, GetMonthNaYin

// 其他
GetJieQiTable, GetShiChen
```

---

## 版本规划

### v1.0 (已完成)
- 6 个基础工具

### v1.1 (已完成)
- 10 个增强工具
- 共 16 个工具

### v1.2 (规划中)
- 日期推算
- 农历转公历
- 吉时查询

### v1.3 (规划中)
- 择日系统
- 合婚功能
- 起名功能

### v2.0 (规划中)
- 易经占卜
- 九星飞宫
- 完整命理

---

## 验收标准

- [x] MCP 协议正常
- [x] 16 个工具可用
- [x] Docker 支持
- [x] GitHub 开源
- [ ] 单元测试覆盖 > 80%
- [ ] 文档完整

---

*本文档使用 OpenSpec 管理*
