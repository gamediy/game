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
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 1
        },
        {
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 2
        },
        {
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 3
        },
        {
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 4
        },
        {
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 5
        },
        {
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 6
        },
        {
          "rewardRule": "1",
          "taskId": 1,
          "taskItemId": 7
        }
      ]
    }
  },
  "query": null,
  "tag": "",
  "_": 0
}
```