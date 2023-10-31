# req

```json
{
  "event":"/user/wallet/listTransType"
}
```
# res
```json
{
  "event": "/user/wallet/listTransType_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "list": [
        {
          "title": "freeze",
          "code": -800,
          "remark": "冻结",
          "type": -1
        },
        {
          "title": "bet",
          "code": -700,
          "remark": "下注",
          "type": -1
        },
        {
          "title": "manual deduction",
          "code": -102,
          "remark": "人工扣除",
          "type": -1
        },
        {
          "title": "withdraw",
          "code": -100,
          "remark": "提现",
          "type": -1
        },
        {
          "title": "withdrawal failed",
          "code": 100,
          "remark": "提现失败",
          "type": -1
        },
        {
          "title": "income",
          "code": 201,
          "remark": "收益",
          "type": 1
        },
        {
          "title": "profit",
          "code": 400,
          "remark": "利润",
          "type": 1
        },
        {
          "title": "gift",
          "code": 401,
          "remark": "赠送",
          "type": 1
        },
        {
          "title": "deposit",
          "code": 500,
          "remark": "充值",
          "type": 1
        },
        {
          "title": "deposit bank",
          "code": 501,
          "remark": "银行卡充值",
          "type": 1
        },
        {
          "title": "unfreeze",
          "code": 800,
          "remark": "解冻",
          "type": 1
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```