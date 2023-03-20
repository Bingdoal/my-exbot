# Ex-bot

**施工中...**

嘗試自製虛擬貨幣交易機器人，目前以國內 ACE 交易所為例

目標是能根據 K 線圖制定自己的策略，長期自動套利

## 使用

在目錄底下新增 `env.prod.yaml` 的檔案，補全下面資料

```yaml
exchange:
  ace:
    uid:
    apiKey:
    securityKey:
```

然後執行指令 `go run *.go prod`

## 說明

通過 `ExchangeApi` 介面可以替換不同的交易所 API 實作
