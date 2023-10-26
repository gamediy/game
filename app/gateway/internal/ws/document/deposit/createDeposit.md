## 创建充值订单

```json
{
  "event": "/user/wallet/createDeposit",
  "query": {
    "payId": 2,
    "amount": 20,
    "transferOrderNo": "test",
    "transferImg": "testimg"
  }
}
```
res
```json
{
  "event": "/user/wallet/createDeposit_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "orderNo": 1717395400610025472
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```