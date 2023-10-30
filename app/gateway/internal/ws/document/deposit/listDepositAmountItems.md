## req

```json
{
  "event": "/user/wallet/depositAmountItems",
  "query": {
    "page": 1,
    "size": 10
  }
}
```

## res

```json
{
  "event": "/user/wallet/depositAmountItems_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "tips": "payment prompt get configuration file",
      "list": [
        {
          "title": "银行卡账转",
          "categoryId": 1,
          "item": [
            {
              "id": 3,
              "title": "USDT",
              "detail": "USDT(TRC20) 指定账号",
              "min": 1,
              "max": 1000,
              "logo": "https://pub-e700c1631e3a4dedaf546e73002fac2f.r2.dev/5T3wUb.jpeg",
              "currency": "USDT",
              "protocol": "Trc20",
              "address": "1234567890",
              "bankId": 3
            },
            {
              "id": 86,
              "title": "EURO",
              "detail": "EURO 指定账号",
              "min": 1,
              "max": 1000,
              "logo": "https://pub-e700c1631e3a4dedaf546e73002fac2f.r2.dev/5/2023/10/6PI3NA.png",
              "currency": "EURO",
              "protocol": "EURO",
              "address": "4123813103412323",
              "bankId": 1
            }
          ]
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```