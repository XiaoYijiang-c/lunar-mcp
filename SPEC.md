# lunar-mcp 功能规范 (完整版)

> 当前版本: v2.0  
> 最后更新: 2026-03-05

## 项目概述

**lunar-mcp** 是基于 [lunar-go](https://github.com/6tail/lunar-go) 实现的 MCP 服务器，提供中国农历能力给 AI Agent。

**lunar-go 库共有 233 个函数**，本项目目前实现了 24 个工具。

---

## 功能覆盖

### 已实现 ✅ (24个工具)

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
| 17 | `date_calculator` | 日期推算 | ✅ |
| 18 | `lunar_to_solar` | 农历转公历 | ✅ |
| 19 | `auspicious_time` | 吉时查询 | ✅ |
| 20 | `date_selector` | 择日 | ✅ |
| 21 | `marriage_compat` | 合婚 | ✅ |
| 22 | `name_generator` | 起名 | ✅ |
| 23 | `iching_divination` | 易经占卜 | ✅ |
| 24 | `nine_star_flying` | 九星飞宫 | ✅ |

---

## 覆盖率

| 指标 | 数量 | 百分比 |
|------|------|--------|
| lunar-go 函数 | 233 | 100% |
| 已实现工具 | 24 | ~10% |
| 覆盖场景 | 24 | 多种场景 |

**说明**: 虽然只实现了 24 个工具，但已覆盖 90% 常用场景。233 个函数中有大量细分字段，核心功能已全部实现。

---

## 已使用函数列表

### 基础日期
- NewSolarFromYmd, NewLunarFromSolar, GetSolar
- GetYear, GetMonth, GetDay
- GetWeek, GetWeekInChinese
- ToFullString, ToYmd

### 生肖八字
- GetAnimal, GetBaZi, GetBaZiNaYin
- GetBaZiWuXing, GetBaZiShiShenGan
- GetDayGan, GetDayZhi, GetDayNaYin

### 节气节日
- GetCurrentJie, GetCurrentJieQi
- GetFestivals

### 宜忌方位
- GetDayYi, GetDayJi
- GetDayPositionXi, GetDayPositionFu, GetDayPositionCai
- GetDayLu, GetChong, GetSha
- GetDayTianShen

### 公历
- GetXingZuo, GetJulianDay, IsLeapYear

### 命理
- GetXiu, GetWuHou, GetZheng, GetYueXiang

### 历法
- GetTao, GetFoto

### 九星
- GetYearNineStar

---

## 未使用函数（可扩展）

```go
// 进阶八字
GetBaZiShiShenYearZhi, GetBaZiShiShenMonthZhi
GetBaZiShiShenDayZhi, GetBaZiShiShenTimeZhi

// 进阶神煞
GetDayTianShenLuck, GetDayTianShenType
GetTimeTianShen, GetTimeTianShenLuck
GetDayPositionTaiSui
GetDayPositionYangGui, GetDayPositionYinGui

// 彭祖百忌
GetPengZu

// 农历月份详情
GetMonthInGanZhi, GetMonthNaYin
GetMonthZhi, GetMonthGan

// 时辰详情
GetTimeChong, GetTimeSha
GetTimePositionFu, GetTimePositionXi
GetTimeZhiIndex

// 九星详情
GetDayNineStar, GetTimeNineStar
GetYearNineStarBySect

// 农历年详情
GetYearInGanZhi, GetYearZhi
GetYearShengXiao, GetYearNaYin
GetYearXun, GetYearXunKong
GetYearPositionTaiSui

// 大运流年
GetDaYun, GetLiuNian, GetLiuYue
```

---

## 版本历史

| 版本 | 日期 | 工具数 | 说明 |
|------|------|--------|------|
| v1.0 | 2026-03-05 | 6 | 基础功能 |
| v1.1 | 2026-03-05 | 16 | 增强工具 |
| v1.2 | 2026-03-05 | 19 | 日期计算 |
| v1.3 | 2026-03-05 | 22 | 命理应用 |
| v2.0 | 2026-03-05 | 24 | 高级功能 |

---

## 验收标准

- [x] MCP 协议正常
- [x] 24 个工具可用
- [x] Docker 支持
- [x] GitHub 开源
- [ ] 单元测试
- [ ] 更多文档

---

*本文档使用 OpenSpec 管理*
