# todo-app

# test
```shell


# 新增 TODO
curl -X POST localhost:8091/api/todo \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"title":"Learn Go"}'

# 查询 TODO
curl -H "Authorization: Bearer $TOKEN" localhost:8091/api/todos


```