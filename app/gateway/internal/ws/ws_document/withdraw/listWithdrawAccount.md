# req

```json
{
  "event": "/user/withdraw/listWithdrawAccount"
}
```

# res
```json
{
  "event": "/user/withdraw/listWithdrawAccount_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "total": 3,
      "list": [
        {
          "id": 1718886715747209216,
          "net": "TRON",
          "protocol": "Trc20",
          "address": "a****ss",
          "currency": "USDT",
          "status": "1",
          "title": "test"
        },
        {
          "id": 1718887644248674304,
          "net": "TRON",
          "protocol": "Trc20",
          "address": "ad*****2",
          "currency": "USDT",
          "status": "1",
          "title": "test"
        },
        {
          "id": 1718887673415864320,
          "net": "TRON",
          "protocol": "Trc20",
          "address": "ad*****3",
          "currency": "USDT",
          "status": "1",
          "title": "test3"
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```