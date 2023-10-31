# req

```json
{
  "event": "/user/withdraw/listWithdraw"
}
```
# res
```json
{
  "event": "/user/withdraw/listWithdraw_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "total": 5,
      "list": [
        {
          "orderNo": "1718952566189985792",
          "uid": "161",
          "status": "-1",
          "finishAt": "2023-10-31 11:23:00",
          "detail": "USDT USDT(TRC20) 指定账号",
          "amount": "1.00",
          "address": "address3",
          "fee": "0",
          "net": "TRON",
          "currency": "USDT",
          "protocol": "Trc20",
          "createdAt": "2023-10-30 19:26:48"
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```