# req
```json
{
  "event": "/user/withdraw/set_pay_pass",
  "query": {
    "pass": "123456"
  }
}
```
# res  ok
```json
{
  "event": "/user/withdraw/set_pay_pass_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {}
  },
  "query": null,
  "tag": "",
  "_": 0
}
```
## res error
```json
{
  "event": "/user/withdraw/set_pay_pass_response",
  "body": {
    "code": 1,
    "msg": "pay pass has been set",
    "data": null
  },
  "query": null,
  "tag": "",
  "_": 0
}
```