## req
```json
{
  "event": "/user/wallet/listDeposit",
  "query": {
    "page": 1,
    "size": 10
  }
}
```
res
```json
{
  "event": "/user/wallet/listDeposit_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "total": 4,
      "list": [
        {
          "logo": "https://pub-e700c1631e3a4dedaf546e73002fac2f.r2.dev/eydjs6.jpeg",
          "uid": 156,
          "orderNo": "1717394878939271168",
          "amount": 20,
          "currency": "PHP",
          "protocol": "BDO",
          "transferImg": "testimg",
          "createdAt": "2023-10-26 12:17:01",
          "amountItemId": 2
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```