# go-api-test

# 起動・テスト

```bash
# 以下のいずれか
cd api-basic
cd api-gin

# 実行
go run .
# 別ターミナルで:
curl 'http://localhost:8080/hello?name=wat'
# => {"message":"Hello, wat!"}
```
