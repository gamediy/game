
# req

```json
{
  "event": "/user/withdraw/listWithdrawMethod"
}
```

# res
```json
{
  "event": "/user/withdraw/listWithdrawMethod_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "tips": "withdrawal prompt get the configuration file",
      "list": [
        {
          "id": 87,
          "title": "USDT",
          "detail": "USDT(TRC20) 1-1000",
          "logo": "https://pub-e700c1631e3a4dedaf546e73002fac2f.r2.dev/5T3wUb.jpeg",
          "currency": "USDT",
          "protocol": "Trc20",
          "address": "12******90",
          "exchangeRate": 1,
          "selectMoney": "[1, 10, 20, 50, 100, 200, 1000]",
          "bind": {
            "address": "address3",
            "protocol": "Trc20"
          }
        },
        {
          "id": 88,
          "title": "EURO ",
          "detail": "EURO 1-1000",
          "logo": "https://pub-e700c1631e3a4dedaf546e73002fac2f.r2.dev/5/2023/10/6PI3NA.png",
          "currency": "EURO",
          "protocol": "EURO",
          "address": "4123*********323",
          "exchangeRate": 1.06,
          "selectMoney": "[1, 5, 10, 50, 100, 200]",
          "bind": {}
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```