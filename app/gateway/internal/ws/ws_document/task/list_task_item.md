req

```json
{
  "event": "/task/listTaskItem",
  "query": {
    "taskId": 1
  }
}
```

res
```json
{
  "event": "/task/listTaskItem_response",
  "body": {
    "code": 0,
    "msg": "",
    "data": {
      "list": [
        {
          "status": "0",
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 1
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```